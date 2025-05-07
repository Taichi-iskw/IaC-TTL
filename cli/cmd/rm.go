/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Taichi-iskw/IaC-TTL/internal/manifest"
	"github.com/Taichi-iskw/IaC-TTL/internal/scheduler"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [stack-name]",
	Short: "Remove a scheduled stack deletion",
	Long: `Remove a scheduled stack deletion.
If stack-name is not provided, it will be automatically detected from the manifest.json file in the cdk.out folder.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var stack string
		var err error

		if len(args) > 0 {
			stack = args[0]
		} else {
			stack, err = manifest.GetStackNameFromManifest()
			if err != nil {
				log.Fatalf("Failed to get stack name from manifest: %v", err)
			}
		}

		err = scheduler.RemoveSchedule(cmd.Context(), stack)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Successfully removed scheduled deletion for stack '%s'\n", stack)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
