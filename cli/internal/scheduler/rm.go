package scheduler

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	"github.com/aws/aws-sdk-go-v2/service/scheduler/types"
)

func RemoveSchedule(ctx context.Context, stackName string) error {
	client, err := NewSchedulerClient(ctx)
	if err != nil {
		return err
	}

	scheduleName := fmt.Sprintf("ttl-delete-%s", stackName)

	_, err = client.DeleteSchedule(ctx, &scheduler.DeleteScheduleInput{
		Name:      aws.String(scheduleName),
		GroupName: aws.String("iac-ttl"),
	})

	if err != nil {
		var resourceNotFound *types.ResourceNotFoundException
		if errors.As(err, &resourceNotFound) {
			return fmt.Errorf("no scheduled deletion found for stack '%s'", stackName)
		}
		return fmt.Errorf("failed to remove schedule: %w", err)
	}

	return nil
}
