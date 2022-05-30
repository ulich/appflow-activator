package main

import (
	"context"
	"flag"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/appflow"
)

func main() {
	appflowName := flag.String("name", "", "The flow to (de)activate (required)")
	shouldDeactivateFlow := flag.Bool("deactivate", false, "If provided, the flow will be deactivated instead")

	flag.Parse()

	if *appflowName == "" {
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	client := appflow.NewFromConfig(cfg)

	if *shouldDeactivateFlow {
		deactivateFlow(client, *appflowName)
	} else {
		activateFlow(client, *appflowName)
	}
}

func activateFlow(client *appflow.Client, flowName string) {
	_, err := client.StartFlow(context.TODO(), &appflow.StartFlowInput{
		FlowName: aws.String(flowName),
	})
	if err != nil {
		panic(err)
	}
}

func deactivateFlow(client *appflow.Client, flowName string) {
	_, err := client.StopFlow(context.TODO(), &appflow.StopFlowInput{
		FlowName: aws.String(flowName),
	})
	if err != nil {
		panic(err)
	}
}
