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

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm <stack-name>",
	Short: "Remove a scheduled CloudFormation stack deletion",
	Long: `Remove a scheduled deletion for a CloudFormation stack.

This command cancels a previously scheduled deletion for a specified stack.
The stack will no longer be automatically deleted at the scheduled time.

Example:
  # Remove a scheduled deletion for a stack
  iac-ttl rm my-stack`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stackName := args[0]
		err := scheduler.RemoveSchedule(cmd.Context(), stackName)
		if err != nil {
			log.Fatalf("failed to remove schedule: %v", err)
		}
		fmt.Printf("Successfully removed scheduled deletion for stack '%s'\n", stackName)
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
