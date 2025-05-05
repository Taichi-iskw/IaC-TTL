package scheduler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
)

// ScheduleInfo holds minimal information about a schedule
type ScheduleInfo struct {
	Name string
	Time string
}

// TimeZone represents the timezone for schedule display
var TimeZone = "Asia/Tokyo"

func ListSchedules(ctx context.Context) ([]ScheduleInfo, error) {
	client, err := NewSchedulerClient(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]ScheduleInfo, 0)
	var nextToken *string

	loc, err := time.LoadLocation(TimeZone)
	if err != nil {
		return nil, fmt.Errorf("failed to load timezone %s: %w", TimeZone, err)
	}

	for {
		output, err := client.ListSchedules(ctx, &scheduler.ListSchedulesInput{
			NamePrefix: aws.String("ttl-delete-"),
			NextToken:  nextToken,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list schedules: %w", err)
		}

		for _, s := range output.Schedules {
			scheduleDetail, err := client.GetSchedule(ctx, &scheduler.GetScheduleInput{
				Name:      s.Name,
				GroupName: aws.String("iac-ttl"),
			})
			if err != nil {
				fmt.Println("failed to get schedule details: %w", err)
				continue
			}
			expr := aws.ToString(scheduleDetail.ScheduleExpression)
			raw := strings.TrimPrefix(strings.TrimSuffix(expr, ")"), "at(")

			t, err := time.Parse("2006-01-02T15:04:05", raw)
			if err != nil {
				fmt.Println("invalid time:", expr)
				continue
			}

			result = append(result, ScheduleInfo{
				Name: strings.TrimPrefix(aws.ToString(s.Name), "ttl-delete-"),
				Time: t.In(loc).Format("2006-01-02 15:04"),
			})
		}

		if output.NextToken == nil {
			break
		}
		nextToken = output.NextToken
	}

	return result, nil
}
