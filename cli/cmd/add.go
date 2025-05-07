/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Taichi-iskw/IaC-TTL/internal/manifest"
	"github.com/Taichi-iskw/IaC-TTL/internal/scheduler"
	"github.com/spf13/cobra"
)

var (
	hours   int
	minutes int
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [stack-name]",
	Short: "Schedule a CloudFormation stack for deletion after a specified time",
	Long: `Schedule a CloudFormation stack for automatic deletion after a specified time period.

This command creates a schedule in AWS EventBridge Scheduler to delete the specified
CloudFormation stack after the given time period (TTL - Time To Live).

If stack-name is not provided, it will be automatically detected from the manifest.json
in the current directory's cdk.out folder.

Examples:
  # Schedule a stack for deletion after 24 hours
  iac-ttl add my-stack -H 24

  # Schedule a stack for deletion after 30 minutes
  iac-ttl add my-stack -m 30

  # Schedule a stack for deletion after 1 hour and 30 minutes
  iac-ttl add my-stack -H 1 -m 30

  # Schedule the current stack for deletion after 24 hours
  iac-ttl add -H 24`,
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var stack string
		var err error

		if len(args) == 0 {
			// Get stack name from manifest.json if not provided
			stack, err = manifest.GetStackNameFromManifest()
			if err != nil {
				return fmt.Errorf("failed to get stack name from manifest: %v", err)
			}
		} else {
			stack = args[0]
		}

		ttl := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute

		if ttl <= 0 {
			return fmt.Errorf("TTL must be greater than 0")
		}

		// Create schedule
		err = scheduler.AddSchedule(cmd.Context(), stack, ttl)
		if err != nil {
			return fmt.Errorf("failed to create schedule: %v", err)
		}

		fmt.Printf("Stack '%s' will be automatically deleted in %v\n", stack, ttl)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVarP(&hours, "hours", "H", 0, "TTL in hours")
	addCmd.Flags().IntVarP(&minutes, "minutes", "m", 0, "TTL in minutes")
}
