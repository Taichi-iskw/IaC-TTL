package scheduler

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSchedule(t *testing.T) {
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
			expectedError: &types.ResourceNotFoundException{},
		},
		{
			name:      "error deleting schedule",
			stackName: "error-stack",
			mockClient: &MockSchedulerClient{
				DeleteScheduleFunc: func(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
					return nil, assert.AnError
				},
			},
			expectedError: assert.AnError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultMockClient = tt.mockClient
			defer func() {
				defaultMockClient = nil
			}()

			ctx := context.Background()
			err := DeleteSchedule(ctx, tt.stackName)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.IsType(t, tt.expectedError, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
