package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

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
	execTime := time.Now().Add(ttl).UTC().Format("2006-01-02T15:04:05")
	scheduleName := fmt.Sprintf("ttl-delete-%s", stackName)

	Arn, err := getParameter("/iac-ttl/destroy-fn-arn")
	if err != nil {
		return err
	}
	RoleArn, err := getParameter("/iac-ttl/scheduler-role-arn")
	if err != nil {
		return err
	}

	input := &scheduler.CreateScheduleInput{
		Name:               aws.String(scheduleName),
		ScheduleExpression: aws.String(fmt.Sprintf("at(%s)", execTime)),
		GroupName:          aws.String("iac-ttl"),
		Target: &types.Target{
			Arn:     aws.String(Arn),
			RoleArn: aws.String(RoleArn),
			Input:   aws.String(fmt.Sprintf(`{"StackName":"%s"}`, stackName)),
		},
		FlexibleTimeWindow: &types.FlexibleTimeWindow{
			Mode: types.FlexibleTimeWindowModeOff,
		},
	}

	// call api
	_, err = client.CreateSchedule(context.TODO(), input)
	return err
}

func getParameter(name string) (string, error) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}
	client := ssm.NewFromConfig(cfg)

	resp, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name: aws.String(name),
	})
	if err != nil {
		return "", err
	}

	return *resp.Parameter.Value, nil
}
