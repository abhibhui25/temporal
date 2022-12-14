package workflows

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func SampleGreetingsWorkflow(ctx workflow.Context) error {

	ao := workflow.ActivityOptions{

		StartToCloseTimeout: 10 * time.Second,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)

	var greetResult string

	err := workflow.ExecuteActivity(ctx, "GetGreeting").Get(ctx, &greetResult)

	if err != nil {

		logger.Error("Get greeting failed.", "Error", err)

		return err

	}

	var nameResult string

	err = workflow.ExecuteActivity(ctx, "GetName").Get(ctx, &nameResult)

	if err != nil {

		logger.Error("Get name failed.", "Error", err)

		return err

	}

	var sayResult string

	err = workflow.ExecuteActivity(ctx, "SayGreeting", greetResult, nameResult).Get(ctx, &sayResult)

	if err != nil {

		logger.Error("Marshalling with errors.", "Error", err)

		return err

	}

	logger.Info("Workflow completed.", "Result", sayResult)

	return nil

}
