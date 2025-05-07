package scheduler

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/stretchr/testify/assert"
)

func TestRemoveSchedule(t *testing.T) {
	tests := []struct {
		name          string
		stackName     string
		mockClient    *MockSchedulerClient
		expectedError error
	}{
		{
			name:      "successfully delete schedule",
			stackName: "test-stack",
			mockClient: &MockSchedulerClient{
				DeleteScheduleFunc: func(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
					assert.Equal(t, "ttl-delete-test-stack", aws.ToString(params.Name))
					assert.Equal(t, "iac-ttl", aws.ToString(params.GroupName))
					return &scheduler.DeleteScheduleOutput{}, nil
				},
			},
			expectedError: nil,
		},
		{
			name:      "schedule not found",
			stackName: "non-existent-stack",
			mockClient: &MockSchedulerClient{
				DeleteScheduleFunc: func(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
					return nil, &types.ResourceNotFoundException{}
				},
			},
			expectedError: errors.New("no scheduled deletion found for stack 'non-existent-stack'"),
		},
		{
			name:      "error deleting schedule",
			stackName: "error-stack",
			mockClient: &MockSchedulerClient{
				DeleteScheduleFunc: func(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
					return nil, assert.AnError
				},
			},
			expectedError: errors.New("failed to remove schedule: assert.AnError general error for testing"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultMockClient = tt.mockClient
			defer func() {
				defaultMockClient = nil
			}()

			ctx := context.Background()
			err := RemoveSchedule(ctx, tt.stackName)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
