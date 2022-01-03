package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

func newCmdGetMembers(config *config) *cobra.Command {
	var projectKey string
	var last uint64
	var format string
	var showIDs bool
	cmd := &cobra.Command{
		Use:   "members",
		Short: "Display latest members from a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			projectID, err := config.loadProjectID(projectKey)
			if err != nil {
				return err
			}

			mm, err := config.cloud.ProjectMembers(config.ctx, projectID, last)
			if err != nil {
				return fmt.Errorf("could not fetch your project members: %w", err)
			}

			switch format {
			case "table":
				tw := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
				if showIDs {
					fmt.Fprintf(tw, "ID\t")
				}
				fmt.Fprint(tw, "EMAIL\tNAME\tROLES\t")
				if showIDs {
					fmt.Fprint(tw, "MEMBER-ID\t")
				}
				fmt.Fprintln(tw, "AGE")
				for _, m := range mm {
					if showIDs {
						fmt.Fprintf(tw, "%s\t", m.ID)
					}
					roles := make([]string, len(m.Roles))
					for i, r := range m.Roles {
						roles[i] = string(r)
					}
					fmt.Fprintf(tw, "%s\t%s\t%s\t", m.User.Email, m.User.Name, strings.Join(roles, ", "))
					if showIDs {
						fmt.Fprintf(tw, "%s\t", m.ID)
					}
					fmt.Fprintln(tw, fmtAgo(m.CreatedAt))
				}
				tw.Flush()
			case "json":
				err := json.NewEncoder(os.Stdout).Encode(mm)
				if err != nil {
					return fmt.Errorf("could not json encode your project members: %w", err)
				}
			default:
				return fmt.Errorf("unknown output format %q", format)
			}
			return nil
		},
	}

	fs := cmd.Flags()
	fs.StringVar(&projectKey, "project", config.defaultProject, "Parent project ID or name")
	fs.Uint64VarP(&last, "last", "l", 0, "Last `N` members. 0 means no limit")
	fs.StringVarP(&format, "output-format", "o", "table", "Output format. Allowed: table, json")
	fs.BoolVar(&showIDs, "show-ids", false, "Include member IDs in table output")

	_ = cmd.RegisterFlagCompletionFunc("output-format", config.completeOutputFormat)
	_ = cmd.RegisterFlagCompletionFunc("project", config.completeProjects)

	return cmd
}
