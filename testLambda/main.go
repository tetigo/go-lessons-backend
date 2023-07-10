package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber int
}

func getMonth(month int) (string, error) {
	if month < 1 || month > 12 {
		return "", fmt.Errorf("month should be between 1 and 12 received wrong month: %d", month)
	}
	return time.Month(month).String(), nil
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	return getMonth(event.MonthNumber)
}

func main() {
	lambda.Start(HandleRequest)
}
