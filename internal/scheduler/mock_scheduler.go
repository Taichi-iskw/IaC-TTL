package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

// SchedulerClient is an interface that defines the methods we need from the AWS Event Scheduler client
type SchedulerClient interface {
	CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
}

// MockSchedulerClient is a mock implementation of SchedulerClient
type MockSchedulerClient struct {
	CreateScheduleFunc func(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
}

// CreateSchedule implements the SchedulerClient interface
func (m *MockSchedulerClient) CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error) {
	if m.CreateScheduleFunc != nil {
		return m.CreateScheduleFunc(ctx, params, optFns...)
	}
	return &scheduler.CreateScheduleOutput{}, nil
}
