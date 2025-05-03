package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

// NewSchedulerClient creates a new AWS Event Scheduler client
func NewSchedulerClient(ctx context.Context) (SchedulerClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return scheduler.NewFromConfig(cfg), nil
}

func CreateSchedule(stackName string, ttl time.Duration) error {
	// create client
	client, err := NewSchedulerClient(context.TODO())
	if err != nil {
		return err
	}

	return CreateScheduleWithClient(client, stackName, ttl)
}

// CreateScheduleWithClient is a helper function that accepts a client for testing purposes
func CreateScheduleWithClient(client SchedulerClient, stackName string, ttl time.Duration) error {
	// create schedule input
	execTime := time.Now().Add(ttl).UTC().Format(time.RFC3339)
	scheduleName := fmt.Sprintf("ttl-delete-%s", stackName)

	input := &scheduler.CreateScheduleInput{
		Name:               aws.String(scheduleName),
		ScheduleExpression: aws.String(fmt.Sprintf("at(%s)", execTime)),
		Target: &types.Target{
			Arn:     aws.String("arn:aws:scheduler:::target/CFN-DeleteStack"), // TODO: mock
			RoleArn: aws.String("arn:aws:iam::<account>:role/<scheduler-role>"),
			Input:   aws.String(fmt.Sprintf(`{"StackName":"%s"}`, stackName)),
		},
		FlexibleTimeWindow: &types.FlexibleTimeWindow{
			Mode: types.FlexibleTimeWindowModeOff,
		},
	}

	// call api
	_, err := client.CreateSchedule(context.TODO(), input)
	return err
}
