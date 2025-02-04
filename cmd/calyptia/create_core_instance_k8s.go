package main

import (
	"context"
	"fmt"

	"github.com/calyptia/core-images-index/go-index"

	"github.com/spf13/cobra"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	cloud "github.com/calyptia/api/types"
	"github.com/calyptia/cli/k8s"
)

const (
	//nolint: gosec // this is not a secret leak, it's just a format declaration.
	defaultCoreDockerImage = "ghcr.io/calyptia/core"
)

func newCmdCreateCoreInstanceOnK8s(config *config, testClientSet kubernetes.Interface) *cobra.Command {
	var coreInstanceName string
	var coreInstanceVersion string
	var noHealthCheckPipeline bool
	var environment string
	var tags []string

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}

	cmd := &cobra.Command{
		Use:     "kubernetes",
		Aliases: []string{"kube", "k8s"},
		Short:   "Setup a new core instance on Kubernetes",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()

			coreDockerImage := fmt.Sprintf("%s:%s", defaultCoreDockerImage, coreInstanceVersion)

			var environmentID string
			if environment != "" {
				var err error
				environmentID, err = config.loadEnvironmentID(environment)
				if err != nil {
					return err
				}
			}

			created, err := config.cloud.CreateAggregator(ctx, cloud.CreateAggregator{
				Name:                   coreInstanceName,
				AddHealthCheckPipeline: !noHealthCheckPipeline,
				EnvironmentID:          environmentID,
				Tags:                   tags,
			})
			if err != nil {
				return fmt.Errorf("could not create core instance at calyptia cloud: %w", err)
			}

			if configOverrides.Context.Namespace == "" {
				configOverrides.Context.Namespace = apiv1.NamespaceDefault
			}

			var clientSet kubernetes.Interface
			if testClientSet != nil {
				clientSet = testClientSet
			} else {
				kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
				kubeClientConfig, err := kubeConfig.ClientConfig()
				if err != nil {
					return err
				}

				clientSet, err = kubernetes.NewForConfig(kubeClientConfig)
				if err != nil {
					return err
				}

			}

			k8sClient := &k8s.Client{
				Interface:    clientSet,
				Namespace:    configOverrides.Context.Namespace,
				ProjectToken: config.projectToken,
				CloudBaseURL: config.baseURL,
				LabelsFunc: func() map[string]string {
					return map[string]string{
						k8s.LabelVersion:      version,
						k8s.LabelPartOf:       "calyptia",
						k8s.LabelManagedBy:    "calyptia-cli",
						k8s.LabelCreatedBy:    "calyptia-cli",
						k8s.LabelProjectID:    config.projectID,
						k8s.LabelAggregatorID: created.ID,
					}
				},
			}

			if err := k8sClient.EnsureOwnNamespace(ctx); err != nil {
				return fmt.Errorf("could not ensure kubernetes namespace exists: %w", err)
			}

			secret, err := k8sClient.CreateSecret(ctx, created)
			if err != nil {
				return fmt.Errorf("could not create kubernetes secret from private key: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "secret=%q\n", secret.Name)

			clusterRole, err := k8sClient.CreateClusterRole(ctx, created)
			if err != nil {
				return fmt.Errorf("could not create kubernetes cluster role: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "cluster_role=%q\n", clusterRole.Name)

			serviceAccount, err := k8sClient.CreateServiceAccount(ctx, created)
			if err != nil {
				return fmt.Errorf("could not create kubernetes service account: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "service_account=%q\n", serviceAccount.Name)

			binding, err := k8sClient.CreateClusterRoleBinding(ctx, created, clusterRole, serviceAccount)
			if err != nil {
				return fmt.Errorf("could not create kubernetes cluster role binding: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "cluster_role_binding=%q\n", binding.Name)

			deploy, err := k8sClient.CreateDeployment(ctx, coreDockerImage, created, serviceAccount)
			if err != nil {
				return fmt.Errorf("could not create kubernetes deployment: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "deployment=%q\n", deploy.Name)

			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&coreInstanceVersion, "version", "latest", "Core instance version")
	fs.StringVar(&coreInstanceName, "name", "", "Core instance name (autogenerated if empty)")
	fs.BoolVar(&noHealthCheckPipeline, "no-health-check-pipeline", false, "Disable health check pipeline creation alongside the core instance")
	fs.StringVar(&environment, "environment", "", "Calyptia environment name")
	fs.StringSliceVar(&tags, "tags", nil, "Tags to apply to the core instance")
	clientcmd.BindOverrideFlags(configOverrides, fs, clientcmd.RecommendedConfigOverrideFlags("kube-"))

	_ = cmd.RegisterFlagCompletionFunc("environment", config.completeEnvironments)
	_ = cmd.RegisterFlagCompletionFunc("version", config.completeCoreContainerVersion)

	return cmd
}

func (c *config) completeCoreContainerVersion(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	containerIndex, err := index.NewContainer()
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	vv, err := containerIndex.All(c.ctx)
	if err != nil {
		return nil, cobra.ShellCompDirectiveError
	}

	return vv, cobra.ShellCompDirectiveNoFileComp
}
