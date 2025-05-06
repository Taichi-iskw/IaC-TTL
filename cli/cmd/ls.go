/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/Taichi-iskw/IaC-TTL/internal/scheduler"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all scheduled CloudFormation stack deletions",
	Long: `List all CloudFormation stacks that are scheduled for automatic deletion.

This command displays a table of all stacks that have been scheduled for deletion,
showing their names and scheduled deletion times.

Example:
  # List all scheduled stack deletions
  iac-ttl ls`,
	Run: func(cmd *cobra.Command, args []string) {
		schedules, err := scheduler.ListSchedules(cmd.Context())
		if err != nil {
			log.Fatalf("failed to list schedules: %v", err)
		}
		if len(schedules) == 0 {
			fmt.Println("No scheduled deletions found.")
			return
		}

		fmt.Printf("%-3s %-20s %s\n", "#", "Stack Name", "Scheduled Deletion Time")

		for i, s := range schedules {
			fmt.Printf("%-3d %-20s %s\n", i+1, s.Name, s.Time)
		}
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
}
