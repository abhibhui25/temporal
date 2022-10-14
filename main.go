package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"wire_poc/workflows"

	"github.com/pborman/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
}

type Event struct {
	Greeter Greeter
}

func NewMessage() Message {
	return Message("Hi there!")
}

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

func (g Greeter) Greet() Message {
	return g.Message
}

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func StartWorker() {
	c, err := client.Dial(client.Options{

		HostPort: client.DefaultHostPort,

		Namespace: "Demo",
	})

	if err != nil {

		log.Fatalln("Unable to create client", err)

	}

	defer c.Close()

	w := worker.New(c, "main", worker.Options{})

	w.RegisterWorkflow(workflows.SampleGreetingsWorkflow)

	w.RegisterActivity(&workflows.Activities{})

	err = w.Run(worker.InterruptCh())

	if err != nil {

		log.Fatalln("Unable to start worker")

	}
}

func StartWorkflow() {
	c, err := client.Dial(client.Options{

		HostPort: client.DefaultHostPort,

		Namespace: "Demo",
	})

	if err != nil {

		log.Fatalln("Unable to create client", err)

	}

	defer c.Close()

	workflowID := "dynamic_" + uuid.New()

	workflowOptions := client.StartWorkflowOptions{

		ID: workflowID,

		TaskQueue: "dynamic",
	}

	we, err := c.ExecuteWorkflow(context.Background(),

		workflowOptions, "SampleGreetingsWorkflow")

	if err != nil {

		log.Fatalln("Unable to execute workflow", err)

	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID",

		we.GetRunID())
}

func main() {

	// msg := NewMessage()
	// g := NewGreeter(msg)
	// e := NewEvent(g)
	// e.Start()
	// go StartWorker()
	e := InitializeEvent()
	e.Start()
	go StartWorkflow()
	time.Sleep(60 * time.Second)
}
