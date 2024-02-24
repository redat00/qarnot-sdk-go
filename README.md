# Qarnot Computing Go SDK

![CI](https://github.com/redat00/qarnot-sdk-go/actions/workflows/go.yml/badge.svg)

This Go library allows you to create jobs, tasks and buckets onto the Qarnot API.

1. [Installation](#installation)
2. [Basic usage](#basic-usage)
	1. [Creating a client](#creating-a-client)
	2. [Creating a task](#create-a-task)
	3. [Creating a bucket](#creating-a-bucket)
3. [Status of the project](#status-of-the-project)
	1. [TODO](#todo)
	2. [Endpoints implementation](#endpoint-implementation)
	3. [Bucket implementation](#bucket-implementation)
4. [Contributing](#contributing)
5. [License](#license)

## Installation

```
go get -u github.com/redat00/qarnot-sdk-go
```

## Basic usage

### Creating a client

The base of everything in this library is the client. Once you've created one, you have access to all the methods to do your work.

```go
package main

import (
	"github.com/redat00/qarnot-sdk-go/qarnot"
)

func main() {
	// Creating a client
	client, err := qarnot.NewClient(
		&qarnot.QarnotConfig{
			ApiUrl:     "https://api.qarnot.com",
			ApiKey:     "MY_SUPER_TOKEN",
			Email:      "test@example.org",
			Version:    "v1",
			StorageUrl: "https://storage.qarnot.com",
		},
	)
	if err != nil {
		panic(err)
	}
```
### Create a task

Creating a task is also fairly easy : You just have to use the `CreateTask` method of the client, and pass it a `CreateTaskPayload` with the parameters you wish your task to use.

```go
package main

import (
	"github.com/redat00/qarnot-sdk-go/qarnot"
)

func main() {
	// Creating a client
	client, err := qarnot.NewClient(
		&qarnot.QarnotConfig{
			ApiUrl:     "https://api.qarnot.com",
			ApiKey:     "MY_SUPER_TOKEN",
			Email:      "test@example.org",
			Version:    "v1",
			StorageUrl: "https://storage.qarnot.com",
		},
	)
	if err != nil {
		panic(err)
	}

	// Create payload describing the task
	newTaskPayload := qarnot.CreateTaskPayload{
		Name:          "hello-world",
		Profile:       "docker-batch",
		InstanceCount: 1,
		Constants: &[]qarnot.Constant{
			{Key: "DOCKER_CMD", Value: "echo \"Hello world\""},
		},
		SchedulingType: qarnot.SchedulingType("flex"),
	}
	
	// Creating the task
	uuid, err := client.CreateTask(newTaskPayload)
	if err != nil {
		panic(err)
	}

	// Get task status
	status, err := client.GetTaskInfo(uuid.Uuid)
	if err != nil {
		panic(err)
	}
}
```

### Creating a bucket

Once again, the creation of the bucket is also very easy. It's done through the use of the `CreateBucket` method, which only takes a string as an argument for the bucket name.

```go
package main

import (
	"github.com/redat00/qarnot-sdk-go/qarnot"
)

func main() {
	client, err := qarnot.NewClient(
		&qarnot.QarnotConfig{
			ApiUrl:     "https://api.qarnot.com",
			ApiKey:     "MY_SUPER_TOKEN",
			Email:      "test@example.org",
			Version:    "v1",
			StorageUrl: "https://storage.qarnot.com",
		},
	)
	if err != nil {
		panic(err)
	}

	// Create the bucket
	bucket, err := client.CreateBucket("my_big_bucket")
	if err != nil {
		panic(err)
	}

	// Upload object to the created bucket
	err = client.UploadObject(&qarnot.ObjectToUpload{
		Bucket:    "my_big_bucket",
		LocalPath: "/tmp/file.txt",
		Key:       "file.txt",
	})
	if err != nil {
		panic(err)
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

## Contributing

Contributions are more than welcome. There is no special rules to contribute to this project. Feel free to open issues and pull requests if you deem it necessary. 

Asking for help is always ok, and reporting bug even more. 

## License

This library is distributed with the [MIT License](https://opensource.org/license/mit).