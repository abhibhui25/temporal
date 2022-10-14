package main

import (
	"log"
	"wire_poc/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{

		HostPort: client.DefaultHostPort,

		Namespace: "Demo",
	})

	if err != nil {

		log.Fatalln("Unable to create client", err)

	}

	defer c.Close()

	w := worker.New(c, "dynamic", worker.Options{})

	w.RegisterWorkflow(workflows.SampleGreetingsWorkflow)

	w.RegisterActivity(&workflows.Activities{})

	err = w.Run(worker.InterruptCh())

	if err != nil {

		log.Fatalln("Unable to start worker")

	}
}
