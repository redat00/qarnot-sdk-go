# Qarnot Computing Go SDK

# THIS IS STILL A WIP, PREPARE YOURSELF FOR BREAKING CHANGE, DON'T USE IN PRODUCTION

This library allows you to interact with Qarnot API through your Go code.

## Basic usage

After creating a `client`, with valid credentials, it's as simple as it gets. Simply create a new payload for your task (specifying the name of the task, the profile...) and create it using the client. And voila! You launched a task on Qarnot Computing's platform !


```go
package main

import (
	"github.com/redat00/qarnot-sdk-go/qarnot"
)

func main() {
	client, err := qarnot.NewClient("https://api.qarnot.com", "MY_SUPER_TOKEN", "v1")
	if err != nil {
		panic(err)
	}

	newTaskPayload := qarnot.CreateTaskPayload{
		Name:          "hello-world",
		Profile:       "docker-batch",
		InstanceCount: 1,
		Constants: &[]qarnot.Constant{
			qarnot.Constant{Key: "DOCKER_CMD", Value: "echo \"Hello world\""},
		},
		SchedulingType: qarnot.SchedulingType("flex"),
	}
	client.CreateTask(newTaskPayload)
}
```

