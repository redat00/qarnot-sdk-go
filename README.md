# Qarnot Computing Go SDK

THIS IS STILL A WIP, PREPARE YOURSELF FOR BREAKING CHANGE, DON'T USE IN PRODUCTION

This library allows you to interact with Qarnot API through your Go code.

## Basic usage

After creating a `client`, with valid credentials, it's as simple as it gets. Simply create a new payload for your task (specifying the name of the task, the profile...) and create it using the client. And voila! You launched a task on Qarnot Computing's platform !


```go
package main

import (
	"github.com/redat00/qarnot-sdk-go/qarnot"
)

func main() {
	client, err := qarnot.NewClient(
		&qarnot.QarnotConfig{
			ApiUrl:     "https://api.qarnot.com",
			ApiKey:     "fdd906d352ebd46c77e822f8b0de15fc6cb2bdafc65bd1cb6ddef7240e8635db5574613ab8d645f74b94368a0d1598ea2a5f65bb55ca5ce097fa51fcfab2e5df",
			Email:      "contact@renaudduret.fr",
			Version:    "v1",
			StorageUrl: "https://storage.qarnot.com",
		},
	)
	if err != nil {
		panic(err)
	}

	newTaskPayload := qarnot.CreateTaskPayload{
		Name:          "hello-world",
		Profile:       "docker-batch",
		InstanceCount: 1,
		Constants: &[]qarnot.Constant{
			{Key: "DOCKER_CMD", Value: "echo \"Hello world\""},
		},
		SchedulingType: qarnot.SchedulingType("flex"),
	}
	err := client.CreateTask(newTaskPayload)
	if err != nil {
		panic(err)
	}
}
```

## Status of the project

This section aims at keeping track of the project, see where we're at and give you an idea of what you can expect.

### TODO

- Write a CI/CD
- Clearly needs more tests. E2E tests would be great (with Minio for bucket for example). Coverage should at least be at 70% to seem acceptable
- Needs to rework the way decoding errors are actually handled. `panic` is not acceptable in a library : The end user should be responsible for handling error.
- Needs to rework the differents methods, especially for tasks.
- Document a little bit more, add examples.

### Endpoint implementation

In this section you can find the list of all the endpoints, and how we implement them in the SDK.

#### Hardware Constraints

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /hardware-constraints` | `Client.ListHardwareConstraints` | ✅ | - |
| `GET /hardware-constraints/cpu-model-constraints/search` | - | ❌ | - |

#### Jobs

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /jobs` | `Client.ListJobs` | ✅ | - |
| `POST /jobs` | `Client.CreateJob` | ✅ | - |
| `POST /jobs/search` | - | ❌ | - |
| `POST /jobs/paginate` | - | ❌ | - |
| `POST /jobs/{uuid}/terminate` | `Client.TerminateJob` | ✅ | - |
| `DELETE /jobs/{uuid}` | `Client.DeleteJob` | ✅ | - |
| `GET /jobs/{uuid}` | `Client.GetJobInfo` | ✅ | - |
| `GET /jobs/{uuid}/tasks` | `Client.ListJobTasks` | ✅ | - |

#### Pools

None (for now).

#### Pool Scaling Sanity Check

None (for now).

#### Profiles

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /profiles` | `Client.ListProfiles` | ✅ | - |
| `GET /profiles/{profile}` | `Client.GetProfileDetails` | ✅ | - |

#### Settings

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /settings` | `Client.GetSettings` | ✅ | - |

#### Tasks

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /tasks` | `Client.ListTasks` | ✅ | - |
| `POST /tasks` | `Client.CreateTask` | ✅ | - |
| `GET /tasks/summaries` | `Client.ListTasksSummaries` | ✅ | - |
| `POST /tasks/summaries/paginate` | - | ❌ | - |
| `POST /tasks/search` | - | ❌ | - |
| `POST /tasks/paginate` | - | ❌ | - |
| `POST /tasks/{uuid}/snapshot/periodic` | `Client.CreateTaskPeriodicSnapshot` | ✅ | - |
| `POST /tasks/{uuid}/snapshot/unique` | `Client.CreateTaskUniqueSnapshot` | ✅ | - |
| `POST /tasks/{uuid}/retry` | `Client.RetryTask` | ✅ | - |
| `POST /tasks/{uuid}/recover` | `Client.RecoverTask` | ✅ | - |
| `POST /tasks/{uuid}/resume` | `Client.ResumeTask` | ✅ | - |
| `POST /tasks/{uuid}/clone` | `Client.CloneTask` | ✅ | - |
| `PUT /tasks/{uuid}` | `Client.UpdateTask` | ✅ | - |
| `PATCH /tasks/{uuid}` | `Client.UpdateTaskResources` | ✅ | - |
| `DELETE /tasks/{uuid}` | `Client.DeleteTask` | ✅ | - |
| `GET /tasks/{uuid}` | `Client.GetTaskInformation` | ✅ | - |
| `POST /tasks/{uuid}/abort` | `Client.AbortTask` | ✅ | - |
| `GET /tasks/{uuid}/stdout` | `Client.GetTaskStdout` | ✅ | - |
| `POST /tasks/{uuid}/stdout` | `Client.GetLastTaskStdout` | ✅ | - |
| `GET /tasks/{uuid}/stdout/{instance}` | `Client.GetTaskInstanceStdout` | ✅ | - |
| `POST /tasks/{uuid}/stdout/{instance}` | `Client.GetLastTaskInstanceStdout` | ✅ | - |
| `GET /tasks/{uuid}/stderr` | `Client.GetTaskStderr` | ✅ | - |
| `POST /tasks/{uuid}/stderr` | `Client.GetLastTaskStderr` | ✅ | - |
| `GET /tasks/{uuid}/stderr/{instance}` | `Client.GetTaskInstanceStderr` | ✅ | - |
| `POST /tasks/{uuid}/stderr/{instance}` | `Client.GetLastTaskInstanceStderr` | ✅ | - |

#### Users

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /info` | `Client.GetUserInfo` | ✅ | - |


#### Versions

| Endpoint | SDK Equivalent | Status | Comment |
| --- | --- | --- | --- |
| `GET /versions` | `Client.GetVersions` | ✅ | - |

### Bucket Implementation

Bucket are managed through classic S3 protocol. This SDK handles directly the S3 part so you don't have to do it on your side.

For now some options are not really used/exposed, such as the multiparts upload. More to come.



| Action | SDK Equivalent | Status | Comment |
| ------ | -------------- | ------ | ------- |
| Create bucket | `client.CreateBucket` | ✅ | - |
| Delete bucket | `client.DeleteBucket` | ✅ | - |
| List buckets | `client.ListBuckets` | ✅ | - |
| List bucket objects | `client.ListObjects` | ✅ | - |
| Upload object | `client.UploadObject` | ✅ | - |
| Delete object | `client.DeleteObject` | ✅ | - |
| Get object head | `client.GetObjectHead` | ✅ | - |