package main

import "github.com/spf13/cobra"

func newCmdDelete(config *config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete core instances, pipelines, etc.",
	}

	cmd.AddCommand(
		newCmdDeleteAgents(config),
		newCmdDeleteAgent(config),
		newCmdDeletePipeline(config),
		newCmdDeleteEndpoint(config),
		newCmdDeletePipelineFile(config),
		newCmdDeleteCoreInstance(config, nil),
		newCmdDeleteCoreInstances(config),
		newCmdDeleteEnvironment(config),
		newCmdDeleteTraceSession(config),
	)

	return cmd
}
