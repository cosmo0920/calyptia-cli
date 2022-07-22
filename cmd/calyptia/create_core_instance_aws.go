package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/calyptia/api/types"
	awsClient "github.com/calyptia/cli/aws"
	"github.com/pkg/errors"
	"github.com/sethvargo/go-retry"
	"github.com/spf13/cobra"
)

const (
	coreInstanceUpCheckTimeout = 10 * time.Minute
	coreInstanceUpCheckBackoff = 5 * time.Second
)

var (
	errCoreInstanceNotFound        = errors.New("core instance not found")
	errCoreInstanceNotRunning      = errors.New("core instance not in running status")
	coreInstanceUpCheckMaxDuration = func() retry.Backoff {
		return retry.WithMaxDuration(coreInstanceUpCheckTimeout, retry.NewConstant(coreInstanceUpCheckBackoff))
	}
)

func newCmdCreateCoreInstanceOnAWS(config *config, client awsClient.Client) *cobra.Command {
	var (
		tags                  []string
		noHealthCheckPipeline bool
		coreInstanceVersion   string
		coreInstanceName      string
		environmentKey        string
		credentials           string
		profileFile           string
		profileName           string
		region                string
		subnetID              string
		keyName               string
		instanceTypeName      string
		securityGroupName     string
	)

	cmd := &cobra.Command{
		Use:     "aws",
		Aliases: []string{"ec2", "amazon"},
		Short:   "Setup a new core instance on Amazon EC2",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var params awsClient.CreateInstanceParams

			ctx := context.Background()

			if client == nil {
				client, err = awsClient.New(ctx, region, credentials, profileFile, profileName)
				if err != nil {
					return errors.Wrap(err, "could not initialize client")
				}
			}

			err = ensureCoreInstanceName(ctx, config, coreInstanceName)
			if err != nil {
				return errors.Wrap(err, "core instance name already exists, choose a different one")
			}

			fmt.Fprintln(cmd.OutOrStdout(), "Booting calyptia core instance on AWS")

			imageID, err := client.FindMatchingAMI(ctx, coreInstanceVersion)
			if err != nil {
				return errors.Wrap(err, "could not find a matching AMI for version: "+coreInstanceVersion)
			}

			keyPairName, err := client.EnsureKeyPair(ctx, keyName)
			if err != nil {
				return errors.Wrap(err, "could not find a suitable key pair for a key")
			}

			instanceType, err := client.EnsureInstanceType(ctx, instanceTypeName)
			if err != nil {
				return errors.Wrap(err, "could not find a suitable instance type")
			}

			vpcID, err := client.EnsureSubnet(ctx, subnetID)
			if err != nil {
				return errors.Wrap(err, "could not find a suitable subnet")
			}

			params.SubnetID = subnetID

			securityGroupID, err := client.EnsureSecurityGroup(ctx, securityGroupName, vpcID)
			if err != nil {
				return errors.Wrap(err, "could not find a suitable security group")
			}

			err = client.EnsureSecurityGroupIngressRules(ctx, securityGroupID)
			if err != nil {
				return errors.Wrap(err, "could not apply ingress security rules")
			}

			userData, err := client.CreateUserdata(ctx, &awsClient.CreateUserDataParams{
				ProjectToken:     config.projectToken,
				CoreInstanceName: coreInstanceName,
			})
			if err != nil {
				return errors.Wrap(err, "could not generate instance user data")
			}

			params.ImageID = imageID
			params.InstanceType = instanceType
			params.UserData = userData
			params.SecurityGroupID = securityGroupID
			params.KeyPairName = keyPairName

			createdInstance, err := client.CreateInstance(ctx, &params)
			if err != nil {
				return errors.Wrap(err, "could not create instance")
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Booted AWS instance as: %s\n", createdInstance.String())

			var instance *types.Aggregator

			err = retry.Do(ctx, coreInstanceUpCheckMaxDuration(), func(ctx context.Context) error {
				instance, err = getCoreInstanceByName(ctx, config, coreInstanceName)
				if err != nil {
					return retry.RetryableError(err)
				}

				fmt.Fprintf(cmd.OutOrStdout(), "core instance: %s not yet running, current status: %s\n", instance.Name, instance.Status)

				if instance.Status != types.AggregatorStatusRunning {
					return retry.RetryableError(errCoreInstanceNotRunning)
				}

				return nil
			})

			if err != nil {
				return errors.Wrap(err, "core instance didn't reach running status")
			}

			metadata, err := json.Marshal(createdInstance)
			if err != nil {
				return errors.Wrap(err, "could not encode core instance metadata")
			}

			awsMetadata := json.RawMessage(metadata)
			err = config.cloud.UpdateAggregator(ctx, instance.ID, types.UpdateAggregator{
				Metadata: &awsMetadata,
			})

			if err != nil {
				return errors.Wrap(err, "could not update core instance metadata on cloud API")
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Calyptia core instance: %s, is ready. Happy logs, metrics and traces on AWS :-)\n", instance.Name)
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&coreInstanceVersion, "version", "", "Core instance version (latest is the default)")
	fs.StringVar(&coreInstanceName, "name", "", "Core instance name (autogenerated if empty)")
	fs.BoolVar(&noHealthCheckPipeline, "no-health-check-pipeline", false, "Disable health check pipeline creation alongside the core instance")
	fs.StringVar(&environmentKey, "environment", "", "Calyptia environment name or ID")
	fs.StringSliceVar(&tags, "tags", nil, "Tags to apply to the core instance.")
	fs.StringVar(&credentials, "credentials", "", "Path to the AWS credentials file. If not specified the default credential loader will be used.")
	fs.StringVar(&profileFile, "profile-file", "", "Path to the AWS profile file. If not specified the default credential loader will be used.")
	fs.StringVar(&profileName, "profile", "", "Name of the AWS profile to use, if not specified, the default profileFile will be used.")
	// Set of parameters that map into https://docs.aws.amazon.com/sdk-for-go/api/service/ec2/#RunInstancesInput
	fs.StringVar(&keyName, "key", awsClient.DefaultSecurityGroupName, "AWS Key to use for SSH into the core instance.")
	fs.StringVar(&region, "region", awsClient.DefaultRegionName, "AWS region name to use in the instance.")
	fs.StringVar(&instanceTypeName, "instance-type", awsClient.DefaultInstanceTypeName, "AWS Instance type to use (see https://aws.amazon.com/es/ec2/instance-types/) for details.")
	fs.StringVar(&securityGroupName, "security-group", awsClient.DefaultSecurityGroupName, "AWS Security group name to use.")
	fs.StringVar(&subnetID, "subnet-id", "", "AWS subnet name to use.If you don't specify a subnet ID, we choose a default subnet from your default VPC for you. If you don't have a default VPC, you MUST specify a subnet.")

	// TODO: pass the environment name to the virtual machines.
	_ = cmd.RegisterFlagCompletionFunc("environment", config.completeEnvironments)

	return cmd
}

func ensureCoreInstanceName(ctx context.Context, config *config, name string) error {
	_, err := getCoreInstanceByName(ctx, config, name)
	if errors.Is(err, errCoreInstanceNotFound) {
		return nil
	}
	return err
}

func getCoreInstanceByName(ctx context.Context, config *config, name string) (*types.Aggregator, error) {
	coreInstances, _ := config.cloud.Aggregators(ctx, config.projectID, types.AggregatorsParams{
		Name: &name,
	})

	if len(coreInstances.Items) > 0 {
		return &coreInstances.Items[0], nil
	}

	return nil, errCoreInstanceNotFound
}
