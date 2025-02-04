package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"

	cloud "github.com/calyptia/api/types"
)

func newCmdCreatePipelineFile(config *config) *cobra.Command {
	var pipelineKey string
	var file string
	var encrypt bool
	var outputFormat string
	cmd := &cobra.Command{
		Use:   "pipeline_file",
		Short: "Create a new file within a pipeline",
		RunE: func(cmd *cobra.Command, args []string) error {
			name := filepath.Base(file)
			name = strings.TrimSuffix(name, filepath.Ext(name))

			contents, err := readFile(file)
			if err != nil {
				return err
			}

			pipelineID, err := config.loadPipelineID(pipelineKey)
			if err != nil {
				return err
			}

			out, err := config.cloud.CreatePipelineFile(config.ctx, pipelineID, cloud.CreatePipelineFile{
				Name:      name,
				Contents:  contents,
				Encrypted: encrypt,
			})
			if err != nil {
				return err
			}

			switch outputFormat {
			case "table":
				tw := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 4, 1, ' ', 0)
				fmt.Fprintln(tw, "ID\tAGE")
				fmt.Fprintf(tw, "%s\t%s\n", out.ID, fmtTime(out.CreatedAt))
				tw.Flush()

				return nil
			case "json":
				err := json.NewEncoder(cmd.OutOrStdout()).Encode(out)
				if err != nil {
					return fmt.Errorf("could not json encode your newly created file: %w", err)
				}
			default:
				return fmt.Errorf("unknown output format %q", outputFormat)
			}

			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&pipelineKey, "pipeline", "", "Pipeline ID or name")
	fs.StringVar(&file, "file", "", "File path. You will be able to reference the file from a fluentbit config using its base name without the extension. Ex: `some_dir/my_file.txt` will be referenced as `{{files.my_file}}`")
	fs.BoolVar(&encrypt, "encrypt", false, "Encrypt file contents")
	fs.StringVar(&outputFormat, "output-format", "table", "Output format. Allowed: table, json")

	_ = cmd.MarkFlagRequired("pipeline")
	_ = cmd.MarkFlagRequired("file")

	_ = cmd.RegisterFlagCompletionFunc("pipeline", config.completePipelines)
	_ = cmd.RegisterFlagCompletionFunc("output-format", config.completeOutputFormat)

	return cmd
}
