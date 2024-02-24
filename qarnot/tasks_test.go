package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/r3labs/diff"
)

func TestListTasksOK(t *testing.T) {
	expected := `[
		{
			"errors": [],
			"resourceBuckets": [],
			"advancedResourceBuckets": [],
			"resultBucket": null,
			"completedInstances": [
			  {
				"results": null,
				"instanceId": 0,
				"wallTimeSec": 181.25049,
				"execTimeSec": 1.0,
				"execTimeSecGHz": 1.8559625,
				"peakMemoryMB": 0,
				"state": "Success",
				"error": null,
				"specificationKey": "32c-128g-amd-tr2990wx-ssd",
				"cpuModel": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				"coreCount": 32,
				"clockRatio": 0.619,
				"averageGHz": 1.8559625,
				"executionAttemptCount": 1
			  }
			],
			"status": {
			  "timestamp": "0001-01-01T00:00:00Z",
			  "lastUpdateTimestamp": "0001-01-01T00:00:00Z",
			  "downloadProgress": 0.0,
			  "executionProgress": 100.0,
			  "uploadProgress": 100.0,
			  "instanceCount": 0,
			  "downloadTime": "00:00:00",
			  "downloadTimeSec": 0.0,
			  "environmentTime": "00:02:50",
			  "environmentTimeSec": 170.0,
			  "executionTime": "00:00:01",
			  "executionTimeSec": 1.0,
			  "executionTimeByCpuModel": [
				{
				  "model": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				  "time": 1.0,
				  "core": 32
				}
			  ],
			  "executionTimeByMachineSpecification": [
				{
				  "specificationKey": "32c-128g-amd-tr2990wx-ssd",
				  "time": 1.0
				}
			  ],
			  "executionTimeByInstanceId": null,
			  "executionTimeGhzByCpuModel": [
				{
				  "model": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				  "timeGhz": 1.8559625148773193,
				  "clockRatio": 0.619,
				  "core": 32
				}
			  ],
			  "uploadTime": "00:00:00",
			  "uploadTimeSec": 0.0,
			  "wallTime": "00:03:07",
			  "wallTimeSec": 187.0,
			  "succeededRange": "0",
			  "executedRange": "0",
			  "failedRange": "",
			  "cancelledRange": "",
			  "failedOnlyRange": "",
			  "startedOnceRange": "0",
			  "runningInstancesInfo": {
				"perRunningInstanceInfo": [],
				"snapshotResults": [],
				"timestamp": "0001-01-01T00:00:00Z",
				"averageFrequencyGHz": 0.0,
				"maxFrequencyGHz": 0.0,
				"minFrequencyGHz": 0.0,
				"averageMaxFrequencyGHz": 0.0,
				"averageCpuUsage": 0.0,
				"clusterPowerIndicator": 1.0,
				"averageMemoryUsage": 0.0,
				"averageNetworkInKbps": 0.0,
				"averageNetworkOutKbps": 0.0,
				"totalNetworkInKbps": 0.0,
				"totalNetworkOutKbps": 0.0,
				"runningCoreCountByCpuModel": []
			  }
			},
			"snapshotInterval": 0,
			"resultsCount": 0,
			"constants": [
			  {
				"key": "DOCKER_CMD",
				"value": "echo \"Hello world\""
			  }
			],
			"secretsAccessRights": {
			  "bySecret": [],
			  "byPrefix": []
			},
			"tags": [],
			"uploadResultsOnCancellation": null,
			"dependencies": null,
			"autoDeleteOnCompletion": false,
			"completionTimeToLive": "00:00:00",
			"hardwareConstraints": [],
			"labels": {},
			"schedulingType": "flex",
			"privileges": {
			  "exportApiAndStorageCredentialsInEnvironment": false
			},
			"retrySettings": {
			  "maxTotalRetries": null,
			  "maxPerInstanceRetries": null
			},
			"uuid": "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			"name": "test-hello-world",
			"shortname": "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			"profile": "docker-batch",
			"poolUuid": null,
			"jobUuid": null,
			"progress": 100.0,
			"runningInstanceCount": 0,
			"runningCoreCount": 0,
			"executionTime": "00:00:01",
			"wallTime": "00:03:07",
			"state": "Success",
			"previousState": "UploadingResults",
			"instanceCount": 1,
			"maxRetriesPerInstance": 0,
			"stateTransitionTime": "2024-02-20T22:09:35Z",
			"previousStateTransitionTime": "2024-02-20T22:09:26Z",
			"lastModified": "2024-02-20T22:09:36Z",
			"creationDate": "2024-02-20T22:06:24Z",
			"endDate": "2024-02-20T22:09:35Z",
			"waitForPoolResourcesSynchronization": null
		  }
	]
	`
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks" {
				fmt.Fprint(w, expected)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	time_status_timestamp, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_status_lastupdate, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-20T22:09:35Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_previous_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-20T22:09:26Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_last_modified, err := time.Parse(time.RFC3339, "2024-02-20T22:09:36Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_creation_date, err := time.Parse(time.RFC3339, "2024-02-20T22:06:24Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_end_date, err := time.Parse(time.RFC3339, "2024-02-20T22:09:35Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	expectedData := []Task{
		{
			Errors:                  []QErrorPublic{},
			ResourceBuckets:         []string{},
			AdvancedResourceBuckets: []TaskAdvancedResourceBucket{},
			ResultBucket:            "",
			CompletedInstances: []CompletedInstance{
				{
					Results:        nil,
					InstanceId:     0,
					WallTimeSec:    181.25049,
					ExecTimeSec:    1,
					ExecTimeSecGHz: 1.8559625,
					PeakMemoryMB:   0,
					State:          "Success",
					Error: QErrorPublic{
						Code:    "",
						Message: "",
						Debug:   "",
					},
					SpecificationKey:      "32c-128g-amd-tr2990wx-ssd",
					CpuModel:              "AMD Ryzen Threadripper 2990WX 32-Core Processor",
					CoreCount:             32,
					ClockRatio:            0.619,
					AverageGHz:            1.8559625,
					ExecutionAttemptCount: 1,
				},
			},
			Status: TaskStatus{
				Timestamp:           time_status_timestamp,
				LastUpdateTimestamp: time_status_lastupdate,
				DownloadProgress:    0,
				ExecutionProgress:   100,
				UploadProgress:      100,
				InstanceCount:       0,
				DownloadTime:        "00:00:00",
				DownloadTimeSec:     0,
				EnvironmentTime:     "00:02:50",
				EnvironmentTimeSec:  170,
				ExecutionTime:       "00:00:01",
				ExecutionTimeSec:    1,
				ExecutionTimeByCPUModel: []ExecutionTimeByCPUModel{
					{
						Model: "AMD Ryzen Threadripper 2990WX 32-Core Processor",
						Time:  1,
						Core:  32,
					},
				},
				ExecutionTimeByMachineSpecification: []ExecutionTimeByMachineSpecification{
					{
						SpecificationKey: "32c-128g-amd-tr2990wx-ssd",
						Time:             1,
					},
				},
				ExecutionTimeByInstanceID: nil,
				ExecutionTimeGhzByCPUModel: []ExecutionTimeGhzByCPUModel{
					{
						Model:      "AMD Ryzen Threadripper 2990WX 32-Core Processor",
						TimeGHz:    1.8559625148773193,
						ClockRatio: 0.619,
						Core:       32,
					},
				},
				UploadTime:       "00:00:00",
				UploadTimeSec:    0,
				WallTime:         "00:03:07",
				WallTimeSec:      187,
				SucceededRange:   "0",
				ExecutedRange:    "0",
				FailedRange:      "",
				CancelledRange:   "",
				FailedOnlyRange:  "",
				StartedOnceRange: "0",
			},
			SnapshotInterval: 0,
			ResultsCount:     0,
			Constants: []Constant{
				{
					Key:   "DOCKER_CMD",
					Value: "echo \"Hello world\"",
				},
			},
			SecretsAccessRights: SecretsAccessRights{
				BySecret: []BySecret{},
				ByPrefix: []ByPrefix{},
			},
			Tags:                        []string{},
			SnapshotWhitelist:           "",
			SnapshotBlacklist:           "",
			UploadResultsOnCancellation: false,
			Dependencies: Dependencies{
				DependsOn: nil,
			},
			AutoDeleteOnCompletion: false,
			CompletionTimeToLive:   "00:00:00",
			HardwareConstraints:    []HardwareConstraint{},
			Labels:                 map[string]string{},
			SchedulingType:         SchedulingType(Flex),
			Privileges: Privileges{
				ExportApiAndStorageCredentialsInEnvironment: false,
			},
			RetrySettings: RetrySettings{
				MaxTotalRetries:       0,
				MaxPerInstanceRetries: 0,
			},
			UUID:                                "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			Name:                                "test-hello-world",
			Shortname:                           "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			Profile:                             "docker-batch",
			PoolUUID:                            "",
			JobUUID:                             "",
			Progress:                            100,
			RunningInstanceCount:                0,
			RunningCoreCount:                    0,
			ExecutionTime:                       "00:00:01",
			WallTime:                            "00:03:07",
			State:                               "Success",
			PreviousState:                       "UploadingResults",
			InstanceCount:                       1,
			MaxRetriesPerInstance:               0,
			StateTransitionTime:                 time_state_transition_time,
			PreviousStateTransitionTime:         time_previous_state_transition_time,
			LastModified:                        time_last_modified,
			CreationDate:                        time_creation_date,
			EndDate:                             time_end_date,
			WaitForPoolResourcesSynchronization: false,
		},
	}

	tasks, err := client.ListTasks()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(tasks, expectedData) {
		t.Errorf("values are not equal:")
		changelog, _ := diff.Diff(expectedData, tasks)
		for i := range changelog {
			fmt.Println(changelog[i].Path)
			fmt.Println(changelog[i].From)
			fmt.Println(changelog[i].To)
			fmt.Println("---")
		}
	}
}

func TestListTaskKO(t *testing.T) {
	expected := "{\"error\":{\"message\":\"Bad authentication token\",\"code\":401},\"data\":{}}"

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(401)
			fmt.Fprint(w, expected)
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	_, err = client.ListTasks()
	expectedData := "could not list tasks due to the following error : [HTTP 401] Bad authentication token"

	if err.Error() != expectedData {
		t.Error("different values.")
		t.Errorf("expected: %v", expectedData)
		t.Errorf("found: %v", err.Error())
	}
}

func TestGetTaskInfo(t *testing.T) {
	expected := `
		{
			"errors": [],
			"resourceBuckets": [],
			"advancedResourceBuckets": [],
			"resultBucket": null,
			"completedInstances": [
			  {
				"results": null,
				"instanceId": 0,
				"wallTimeSec": 181.25049,
				"execTimeSec": 1.0,
				"execTimeSecGHz": 1.8559625,
				"peakMemoryMB": 0,
				"state": "Success",
				"error": null,
				"specificationKey": "32c-128g-amd-tr2990wx-ssd",
				"cpuModel": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				"coreCount": 32,
				"clockRatio": 0.619,
				"averageGHz": 1.8559625,
				"executionAttemptCount": 1
			  }
			],
			"status": {
			  "timestamp": "0001-01-01T00:00:00Z",
			  "lastUpdateTimestamp": "0001-01-01T00:00:00Z",
			  "downloadProgress": 0.0,
			  "executionProgress": 100.0,
			  "uploadProgress": 100.0,
			  "instanceCount": 0,
			  "downloadTime": "00:00:00",
			  "downloadTimeSec": 0.0,
			  "environmentTime": "00:02:50",
			  "environmentTimeSec": 170.0,
			  "executionTime": "00:00:01",
			  "executionTimeSec": 1.0,
			  "executionTimeByCpuModel": [
				{
				  "model": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				  "time": 1.0,
				  "core": 32
				}
			  ],
			  "executionTimeByMachineSpecification": [
				{
				  "specificationKey": "32c-128g-amd-tr2990wx-ssd",
				  "time": 1.0
				}
			  ],
			  "executionTimeByInstanceId": null,
			  "executionTimeGhzByCpuModel": [
				{
				  "model": "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				  "timeGhz": 1.8559625148773193,
				  "clockRatio": 0.619,
				  "core": 32
				}
			  ],
			  "uploadTime": "00:00:00",
			  "uploadTimeSec": 0.0,
			  "wallTime": "00:03:07",
			  "wallTimeSec": 187.0,
			  "succeededRange": "0",
			  "executedRange": "0",
			  "failedRange": "",
			  "cancelledRange": "",
			  "failedOnlyRange": "",
			  "startedOnceRange": "0",
			  "runningInstancesInfo": {
				"perRunningInstanceInfo": [],
				"snapshotResults": [],
				"timestamp": "0001-01-01T00:00:00Z",
				"averageFrequencyGHz": 0.0,
				"maxFrequencyGHz": 0.0,
				"minFrequencyGHz": 0.0,
				"averageMaxFrequencyGHz": 0.0,
				"averageCpuUsage": 0.0,
				"clusterPowerIndicator": 1.0,
				"averageMemoryUsage": 0.0,
				"averageNetworkInKbps": 0.0,
				"averageNetworkOutKbps": 0.0,
				"totalNetworkInKbps": 0.0,
				"totalNetworkOutKbps": 0.0,
				"runningCoreCountByCpuModel": []
			  }
			},
			"snapshotInterval": 0,
			"resultsCount": 0,
			"constants": [
			  {
				"key": "DOCKER_CMD",
				"value": "echo \"Hello world\""
			  }
			],
			"secretsAccessRights": {
			  "bySecret": [],
			  "byPrefix": []
			},
			"tags": [],
			"uploadResultsOnCancellation": null,
			"dependencies": null,
			"autoDeleteOnCompletion": false,
			"completionTimeToLive": "00:00:00",
			"hardwareConstraints": [],
			"labels": {},
			"schedulingType": "flex",
			"privileges": {
			  "exportApiAndStorageCredentialsInEnvironment": false
			},
			"retrySettings": {
			  "maxTotalRetries": null,
			  "maxPerInstanceRetries": null
			},
			"uuid": "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			"name": "test-hello-world",
			"shortname": "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
			"profile": "docker-batch",
			"poolUuid": null,
			"jobUuid": null,
			"progress": 100.0,
			"runningInstanceCount": 0,
			"runningCoreCount": 0,
			"executionTime": "00:00:01",
			"wallTime": "00:03:07",
			"state": "Success",
			"previousState": "UploadingResults",
			"instanceCount": 1,
			"maxRetriesPerInstance": 0,
			"stateTransitionTime": "2024-02-20T22:09:35Z",
			"previousStateTransitionTime": "2024-02-20T22:09:26Z",
			"lastModified": "2024-02-20T22:09:36Z",
			"creationDate": "2024-02-20T22:06:24Z",
			"endDate": "2024-02-20T22:09:35Z",
			"waitForPoolResourcesSynchronization": null
		  }
	`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	time_status_timestamp, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_status_lastupdate, err := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-20T22:09:35Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_previous_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-20T22:09:26Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_last_modified, err := time.Parse(time.RFC3339, "2024-02-20T22:09:36Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_creation_date, err := time.Parse(time.RFC3339, "2024-02-20T22:06:24Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	time_end_date, err := time.Parse(time.RFC3339, "2024-02-20T22:09:35Z")
	if err != nil {
		t.Errorf("could not parse time: %v", err)
	}

	expectedData := Task{
		Errors:                  []QErrorPublic{},
		ResourceBuckets:         []string{},
		AdvancedResourceBuckets: []TaskAdvancedResourceBucket{},
		ResultBucket:            "",
		CompletedInstances: []CompletedInstance{
			{
				Results:        nil,
				InstanceId:     0,
				WallTimeSec:    181.25049,
				ExecTimeSec:    1,
				ExecTimeSecGHz: 1.8559625,
				PeakMemoryMB:   0,
				State:          "Success",
				Error: QErrorPublic{
					Code:    "",
					Message: "",
					Debug:   "",
				},
				SpecificationKey:      "32c-128g-amd-tr2990wx-ssd",
				CpuModel:              "AMD Ryzen Threadripper 2990WX 32-Core Processor",
				CoreCount:             32,
				ClockRatio:            0.619,
				AverageGHz:            1.8559625,
				ExecutionAttemptCount: 1,
			},
		},
		Status: TaskStatus{
			Timestamp:           time_status_timestamp,
			LastUpdateTimestamp: time_status_lastupdate,
			DownloadProgress:    0,
			ExecutionProgress:   100,
			UploadProgress:      100,
			InstanceCount:       0,
			DownloadTime:        "00:00:00",
			DownloadTimeSec:     0,
			EnvironmentTime:     "00:02:50",
			EnvironmentTimeSec:  170,
			ExecutionTime:       "00:00:01",
			ExecutionTimeSec:    1,
			ExecutionTimeByCPUModel: []ExecutionTimeByCPUModel{
				{
					Model: "AMD Ryzen Threadripper 2990WX 32-Core Processor",
					Time:  1,
					Core:  32,
				},
			},
			ExecutionTimeByMachineSpecification: []ExecutionTimeByMachineSpecification{
				{
					SpecificationKey: "32c-128g-amd-tr2990wx-ssd",
					Time:             1,
				},
			},
			ExecutionTimeByInstanceID: nil,
			ExecutionTimeGhzByCPUModel: []ExecutionTimeGhzByCPUModel{
				{
					Model:      "AMD Ryzen Threadripper 2990WX 32-Core Processor",
					TimeGHz:    1.8559625148773193,
					ClockRatio: 0.619,
					Core:       32,
				},
			},
			UploadTime:       "00:00:00",
			UploadTimeSec:    0,
			WallTime:         "00:03:07",
			WallTimeSec:      187,
			SucceededRange:   "0",
			ExecutedRange:    "0",
			FailedRange:      "",
			CancelledRange:   "",
			FailedOnlyRange:  "",
			StartedOnceRange: "0",
		},
		SnapshotInterval: 0,
		ResultsCount:     0,
		Constants: []Constant{
			{
				Key:   "DOCKER_CMD",
				Value: "echo \"Hello world\"",
			},
		},
		SecretsAccessRights: SecretsAccessRights{
			BySecret: []BySecret{},
			ByPrefix: []ByPrefix{},
		},
		Tags:                        []string{},
		SnapshotWhitelist:           "",
		SnapshotBlacklist:           "",
		UploadResultsOnCancellation: false,
		Dependencies: Dependencies{
			DependsOn: nil,
		},
		AutoDeleteOnCompletion: false,
		CompletionTimeToLive:   "00:00:00",
		HardwareConstraints:    []HardwareConstraint{},
		Labels:                 map[string]string{},
		SchedulingType:         SchedulingType(Flex),
		Privileges: Privileges{
			ExportApiAndStorageCredentialsInEnvironment: false,
		},
		RetrySettings: RetrySettings{
			MaxTotalRetries:       0,
			MaxPerInstanceRetries: 0,
		},
		UUID:                                "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
		Name:                                "test-hello-world",
		Shortname:                           "e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3",
		Profile:                             "docker-batch",
		PoolUUID:                            "",
		JobUUID:                             "",
		Progress:                            100,
		RunningInstanceCount:                0,
		RunningCoreCount:                    0,
		ExecutionTime:                       "00:00:01",
		WallTime:                            "00:03:07",
		State:                               "Success",
		PreviousState:                       "UploadingResults",
		InstanceCount:                       1,
		MaxRetriesPerInstance:               0,
		StateTransitionTime:                 time_state_transition_time,
		PreviousStateTransitionTime:         time_previous_state_transition_time,
		LastModified:                        time_last_modified,
		CreationDate:                        time_creation_date,
		EndDate:                             time_end_date,
		WaitForPoolResourcesSynchronization: false,
	}

	task, err := client.GetTaskInfo("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Errorf("error should be equal to nil.")
	}

	if !reflect.DeepEqual(task, expectedData) {
		t.Error("different values.")
		t.Errorf("expected : %v", expected)
		t.Errorf("found    : %v", task)
	}

	_, err = client.GetTaskInfo("test")
	expectedErrorString := "could not get task info due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetTaskStdout(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stdout" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	stdout, err := client.GetTaskStdout("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != stdout {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", stdout)
	}

	_, err = client.GetTaskStdout("test")
	expectedErrorString := "could not get task stdout due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetLastTaskStdout(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stdout" && r.Method == "POST" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	stdout, err := client.GetLastTaskStdout("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != stdout {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", stdout)
	}

	_, err = client.GetLastTaskStdout("test")
	expectedErrorString := "could not get last task stdout due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetTaskInstanceStdout(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stdout/1" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	stdout, err := client.GetTaskInstanceStdout("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3", 1)
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != stdout {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", stdout)
	}

	_, err = client.GetTaskInstanceStdout("test", 1)
	expectedErrorString := "could not get task instance stdout due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetLastTaskInstanceStdout(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stdout/1" && r.Method == "POST" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	stdout, err := client.GetLastTaskInstanceStdout("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3", 1)
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != stdout {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", stdout)
	}

	_, err = client.GetLastTaskInstanceStdout("test", 1)
	expectedErrorString := "could not get last task instance stdout due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetTaskStderr(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stderr" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	stderr, err := client.GetTaskStderr("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != stderr {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", stderr)
	}

	_, err = client.GetTaskStderr("test")
	expectedErrorString := "could not get task stderr due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetTaskLastStderr(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stderr" && r.Method == "POST" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	lastStderr, err := client.GetLastTaskStderr("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != lastStderr {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", lastStderr)
	}

	_, err = client.GetLastTaskStderr("test")
	expectedErrorString := "could not get last task stderr due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetTaskInstanceStderr(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stderr/1" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	lastStderr, err := client.GetTaskInstanceStderr("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3", 1)
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != lastStderr {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", lastStderr)
	}

	_, err = client.GetTaskInstanceStderr("test", 1)
	expectedErrorString := "could not get task instance stderr due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestGetLastTaskInstanceStderr(t *testing.T) {
	expected := `" 0> Hello world"`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/stderr/1" && r.Method == "POST" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	lastStderr, err := client.GetLastTaskInstanceStderr("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3", 1)
	if err != nil {
		t.Error("error shout be nil")
	}

	expectedData := ` 0> Hello world`
	if expectedData != lastStderr {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", lastStderr)
	}

	_, err = client.GetLastTaskInstanceStderr("test", 1)
	expectedErrorString := "could not get last task instance stderr due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestDeleteTask(t *testing.T) {
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3" && r.Method == "DELETE" {
				fmt.Fprint(w, nil)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	err = client.DeleteTask("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Errorf("err should be equal to nil: %v", err)
	}

	err = client.DeleteTask("test")
	expectedErrorString := "could not delete task due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestAbortTask(t *testing.T) {
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3/abort" && r.Method == "POST" {
				fmt.Fprint(w, nil)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	err = client.AbortTask("e1d8e5fc-b28f-4ed8-9b72-7ea991d9cfc3")
	if err != nil {
		t.Errorf("err should be equal to nil: %v", err)
	}

	err = client.AbortTask("test")
	expectedErrorString := "could not abort task due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}

func TestListTaskSummaries(t *testing.T) {
	expected := `[
			{
				"uuid": "effce24b-a15e-4b1b-b81b-93b7379ac2f4",
				"name": "test2",
				"shortname": "test2-shortname",
				"profile": "docker-batch",
				"poolUuid": null,
				"jobUuid": null,
				"progress": 100.0,
				"runningInstanceCount": 0,
				"runningCoreCount": 0,
				"executionTime": "00:00:01",
				"wallTime": "00:02:40",
				"state": "Success",
				"previousState": "UploadingResults",
				"instanceCount": 1,
				"maxRetriesPerInstance": 0,
				"stateTransitionTime": "2024-02-17T22:35:40Z",
				"previousStateTransitionTime": "2024-02-17T22:35:37Z",
				"lastModified": "2024-02-17T22:35:40Z",
				"creationDate": "2024-02-17T22:32:57Z",
				"endDate": "2024-02-17T22:35:40Z",
				"waitForPoolResourcesSynchronization": null
			}
	]
	`
	expectedNotFound := `{
		"message": "No such task: test"
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/tasks/summaries" && r.Method == "GET" {
				fmt.Fprint(w, expected)
			} else {
				w.WriteHeader(404)
				fmt.Fprint(w, expectedNotFound)
			}
		}),
	)
	defer srv.Close()

	qarnotConfig := QarnotConfig{
		ApiUrl:     srv.URL,
		ApiKey:     "xxx",
		Email:      "test@example.org",
		Version:    "v1",
		StorageUrl: "http://fake.storage.qarnope.com",
	}

	client, err := NewClient(&qarnotConfig)
	if err != nil {
		t.Errorf("could not create a new client: %v", err)
	}

	summaries, err := client.ListTaskSummaries()
	if err != nil {
		t.Errorf("err should be equal to nil: %v", err)
	}

	time_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-17T22:35:40Z")
	if err != nil {
		t.Error("could not parse time")
	}

	time_previous_state_transition_time, err := time.Parse(time.RFC3339, "2024-02-17T22:35:37Z")
	if err != nil {
		t.Error("could not parse time")
	}

	time_last_modified, err := time.Parse(time.RFC3339, "2024-02-17T22:35:40Z")
	if err != nil {
		t.Error("could not parse time")
	}

	time_creation_date, err := time.Parse(time.RFC3339, "2024-02-17T22:32:57Z")
	if err != nil {
		t.Error("could not parse time")
	}

	time_end_date, err := time.Parse(time.RFC3339, "2024-02-17T22:35:40Z")
	if err != nil {
		t.Error("could not parse time")
	}

	expectedData := []TaskSummary{
		{
			Uuid:                                "effce24b-a15e-4b1b-b81b-93b7379ac2f4",
			Name:                                "test2",
			Shortname:                           "test2-shortname",
			Profile:                             "docker-batch",
			PoolUuid:                            "",
			JobUuid:                             "",
			Progress:                            100.0,
			RunningInstanceCount:                0,
			RunningCoreCount:                    0,
			ExecutionTime:                       "00:00:01",
			WallTime:                            "00:02:40",
			State:                               "Success",
			PreviousState:                       "UploadingResults",
			InstanceCount:                       1,
			MaxRetriesPerInstance:               0,
			StateTransitionTime:                 time_state_transition_time,
			PreviousStateTransitionTime:         time_previous_state_transition_time,
			LastModified:                        time_last_modified,
			CreationDate:                        time_creation_date,
			EndDate:                             time_end_date,
			WaitForPoolResourcesSynchronization: false,
		},
	}

	if !reflect.DeepEqual(summaries, expectedData) {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", summaries)
	}

	err = client.AbortTask("test")
	expectedErrorString := "could not abort task due to the following error : [HTTP 404] No such task: test"

	if err.Error() != expectedErrorString {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedErrorString)
		t.Errorf("found    : %v", err.Error())
	}
}
