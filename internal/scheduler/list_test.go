package scheduler

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/stretchr/testify/assert"
)

func TestListSchedules(t *testing.T) {
	tests := []struct {
		name          string
		mockClient    *MockSchedulerClient
		expected      []ScheduleInfo
		expectedError error
	}{
		{
			name: "successfully list schedules",
			mockClient: &MockSchedulerClient{
				ListSchedulesFunc: func(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
					return &scheduler.ListSchedulesOutput{
						Schedules: []types.ScheduleSummary{
							{
								Name: aws.String("ttl-delete-test-resource"),
							},
						},
					}, nil
				},
				GetScheduleFunc: func(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error) {
					return &scheduler.GetScheduleOutput{
						ScheduleExpression: aws.String("at(2024-03-20T15:30:00)"),
					}, nil
				},
			},
			expected: []ScheduleInfo{
				{
					Name: "test-resource",
					Time: "2024-03-21 00:30", // UTC+9 (Asia/Tokyo)
				},
			},
			expectedError: nil,
		},
		{
			name: "empty schedule list",
			mockClient: &MockSchedulerClient{
				ListSchedulesFunc: func(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
					return &scheduler.ListSchedulesOutput{
						Schedules: []types.ScheduleSummary{},
					}, nil
				},
			},
			expected:      []ScheduleInfo{},
			expectedError: nil,
		},
		{
			name: "error listing schedules",
			mockClient: &MockSchedulerClient{
				ListSchedulesFunc: func(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
					return nil, assert.AnError
				},
			},
			expected:      nil,
			expectedError: assert.AnError,
		},
		{
			name: "invalid timezone",
			mockClient: &MockSchedulerClient{
				ListSchedulesFunc: func(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
					return &scheduler.ListSchedulesOutput{
						Schedules: []types.ScheduleSummary{},
					}, nil
				},
			},
			expected:      nil,
			expectedError: fmt.Errorf("failed to load timezone"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultMockClient = tt.mockClient
			defer func() {
				defaultMockClient = nil
			}()

			// For invalid timezone test
			if tt.name == "invalid timezone" {
				originalTimeZone := TimeZone
				TimeZone = "Invalid/Timezone"
				defer func() {
					TimeZone = originalTimeZone
				}()
			}

			ctx := context.Background()
			result, err := ListSchedules(ctx)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
