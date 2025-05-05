package scheduler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

func DeleteSchedule(ctx context.Context, stackName string) error {
	client, err := NewSchedulerClient(ctx)
	if err != nil {
		return err
	}

	scheduleName := fmt.Sprintf("ttl-delete-%s", stackName)

	_, err = client.DeleteSchedule(ctx, &scheduler.DeleteScheduleInput{
		Name:      aws.String(scheduleName),
		GroupName: aws.String("iac-ttl"),
	})

	return err
}
