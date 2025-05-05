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
	Use:   "rm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		stackName := args[0]
		fmt.Printf("[INFO] cancelling schedule for stack: %s\n", stackName)

		if err := scheduler.RemoveSchedule(cmd.Context(), stackName); err != nil {
			log.Fatalf("Error: failed to delete schedule: %v\n", err)
		}

		fmt.Println("[OK] schedule deleted successfully")
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}
