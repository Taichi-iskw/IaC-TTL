package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

func TestCreateSchedule(t *testing.T) {
	// Create a mock client
	mockClient := &MockSchedulerClient{
		CreateScheduleFunc: func(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error) {
			// Verify the input parameters
			if *params.Name != "ttl-delete-test-stack" {
				t.Errorf("Expected schedule name 'ttl-delete-test-stack', got '%s'", *params.Name)
			}
			if !contains(*params.ScheduleExpression, "at(") {
				t.Errorf("Expected schedule expression to contain 'at(', got '%s'", *params.ScheduleExpression)
			}
			if *params.Target.Arn == "" {
				t.Error("Expected target ARN to be set, but it was empty")
			}
			if *params.Target.Input != `{"StackName":"test-stack"}` {
				t.Errorf("Expected target input '{\"StackName\":\"test-stack\"}', got '%s'", *params.Target.Input)
			}
			return &scheduler.CreateScheduleOutput{}, nil
		},
	}

	// Test the function
	err := CreateScheduleWithClient(mockClient, "test-stack", 1*time.Hour)
	if err != nil {
		t.Fatalf("Failed to create schedule: %v", err)
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr
}
