/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Taichi-iskw/IaC-TTL/internal/scheduler"
	"github.com/spf13/cobra"
)

var (
	hours   int
	minutes int
)

// scheduleCmd represents the schedule command
var scheduleCmd = &cobra.Command{
	Use:   "schedule <stack-name>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		stack := args[0]
		ttl := time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute

		if ttl <= 0 {
			return fmt.Errorf("TTL must be greater than 0")
		}

		// create schedule
		err := scheduler.CreateSchedule(stack, ttl)
		if err != nil {
			return fmt.Errorf("failed to create schedule: %v", err)
		}

		fmt.Printf("[DEBUG] scheduling stack '%s' to expire in %v\n", stack, ttl)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(scheduleCmd)
	scheduleCmd.Flags().IntVarP(&hours, "hours", "H", 0, "TTL in hours")
	scheduleCmd.Flags().IntVarP(&minutes, "minutes", "m", 0, "TTL in minutes")
}
