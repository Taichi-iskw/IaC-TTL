package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

var defaultMockClient *MockSchedulerClient

type SchedulerClient interface {
	CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
	DeleteSchedule(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error)
	ListSchedules(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)
	GetSchedule(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)
}

type MockSchedulerClient struct {
	CreateScheduleFunc func(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
	DeleteScheduleFunc func(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error)
	ListSchedulesFunc  func(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)
	GetScheduleFunc    func(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)
}

func (m *MockSchedulerClient) CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error) {
	if m.CreateScheduleFunc != nil {
		return m.CreateScheduleFunc(ctx, params, optFns...)
	}
	return &scheduler.CreateScheduleOutput{}, nil
}

func (m *MockSchedulerClient) DeleteSchedule(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error) {
	if m.DeleteScheduleFunc != nil {
		return m.DeleteScheduleFunc(ctx, params, optFns...)
	}
	return &scheduler.DeleteScheduleOutput{}, nil
}

func (m *MockSchedulerClient) ListSchedules(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error) {
	if m.ListSchedulesFunc != nil {
		return m.ListSchedulesFunc(ctx, params, optFns...)
	}
	return &scheduler.ListSchedulesOutput{
		Schedules: []types.ScheduleSummary{},
	}, nil
}

func (m *MockSchedulerClient) GetSchedule(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error) {
	if m.GetScheduleFunc != nil {
		return m.GetScheduleFunc(ctx, params, optFns...)
	}
	return &scheduler.GetScheduleOutput{}, nil
}

// NewSchedulerClient creates a new AWS Event Scheduler client
func NewSchedulerClient(ctx context.Context) (SchedulerClient, error) {
	if defaultMockClient != nil {
		return defaultMockClient, nil
	}
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return scheduler.NewFromConfig(cfg), nil
}
