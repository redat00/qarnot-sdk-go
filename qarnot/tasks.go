package qarnot

import (
	"encoding/json"
	"fmt"
	"time"
)

// Struct representing the content of a prefix filtering field
type PrefixFiltering struct {
	Prefix string `json:"prefix"`
}

// Struct representing a prefix filtering field
type Filtering struct {
	PrefixFiltering PrefixFiltering `json:"prefixFiltering"`
}

// Struct representing the execution time by cpu model for a task status
type ExecutionTimeByCPUModel struct {
	Model string  `json:"model"`
	Time  float64 `json:"time"`
	Core  int     `json:"core"`
}

// Struct representing the execution time by machine specification for a task status
type ExecutionTimeByMachineSpecification struct {
	SpecificationKey string  `json:"specificationKey"`
	Time             float64 `json:"time"`
}

// Struct representing the execution time by instance id for a task status
type ExecutionTimeByInstanceID struct {
	InstanceId       int     `json:"instanceId"`
	SpecificationKey string  `json:"specificationKey"`
	SiteUuid         string  `json:"siteUuid"`
	Time             float64 `json:"time"`
	TimeGHz          float64 `json:"timeGHz"`
	ClockRatio       float64 `json:"clockRatio"`
}

// Struct representing the execution time by ghz and cpu model for a task status
type ExecutionTimeGhzByCPUModel struct {
	Model      string  `json:"model"`
	TimeGHz    float64 `json:"timeGHz"`
	ClockRatio float64 `json:"clockRatio"`
	Core       int     `json:"core"`
}

// Struct representing a task status for a task
type TaskStatus struct {
	Timestamp                           time.Time                             `json:"timestamp"`
	LastUpdateTimestamp                 time.Time                             `json:"lastUpdateTimestamp"`
	DownloadProgress                    float64                               `json:"downloadProgress"`
	ExecutionProgress                   float64                               `json:"executionProgress"`
	UploadProgress                      float64                               `json:"uploadProgress"`
	InstanceCount                       float64                               `json:"instanceCount"`
	DownloadTime                        string                                `json:"downloadTime"`
	DownloadTimeSec                     float64                               `json:"downloadTimeSec"`
	EnvironmentTime                     string                                `json:"environmentTime"`
	EnvironmentTimeSec                  float64                               `json:"environmentTimeSec"`
	ExecutionTime                       string                                `json:"executionTime"`
	ExecutionTimeSec                    float64                               `json:"executionTimeSec"`
	ExecutionTimeByCPUModel             []ExecutionTimeByCPUModel             `json:"executionTimeByCpuModel"`
	ExecutionTimeByMachineSpecification []ExecutionTimeByMachineSpecification `json:"executionTimeByMachineSpecification"`
	ExecutionTimeByInstanceID           []ExecutionTimeByInstanceID           `json:"executionTimeByInstanceId"`
	ExecutionTimeGhzByCPUModel          []ExecutionTimeGhzByCPUModel          `json:"executionTimeGhzByCpuModel"`
	UploadTime                          string                                `json:"uploadTime"`
	UploadTimeSec                       float64                               `json:"uploadTimeSec"`
	WallTime                            string                                `json:"wallTime"`
	WallTimeSec                         float64                               `json:"wallTimeSec"`
	SucceededRange                      string                                `json:"succeededRange"`
	ExecutedRange                       string                                `json:"executedRange"`
	FailedRange                         string                                `json:"failedRange"`
	CancelledRange                      string                                `json:"cancelledRange"`
	FailedOnlyRange                     string                                `json:"failedOnlyRange"`
	StartedOnceRange                    string                                `json:"startedOnceRange"`
}

// Struct representing a constant for a task
type Constant struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Struct representing a Qarnot public error
type QErrorPublic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Debug   string `json:"debug"`
}

// Struct representing a prefix field for the resource transformation struct
type StripPrefix struct {
	Prefix string `json:"prefix"`
}

// Struct representing the resource transformation settings for a task advanced resource bucket
type ResourceTransformation struct {
	StripPrefix StripPrefix `json:"stripPrefix"`
}

// Struct representing a task advanced resource bucket setting for a task
type TaskAdvancedResourceBucket struct {
	BucketName             string                 `json:"bucketName"`
	Filtering              Filtering              `json:"filtering"`
	ResourceTransformation ResourceTransformation `json:"resourceTransformation"`
	CacheTTLSec            int                    `jons:"cacheTTLSec"`
}

// Struct represetning a completed instance field for a task
type CompletedInstance struct {
	Results               []string     `json:"results"`
	InstanceId            int          `json:"instanceId"`
	WallTimeSec           float32      `json:"wallTimeSec"`
	ExecTimeSec           float32      `json:"execTimeSec"`
	ExecTimeSecGHz        float32      `json:"execTimeSecGHz"`
	PeakMemoryMB          int          `json:"peakMemoryMB"`
	State                 string       `json:"state"`
	Error                 QErrorPublic `json:"error"`
	SpecificationKey      string       `json:"specificationKey"`
	CpuModel              string       `json:"cpuModel"`
	CoreCount             int32        `json:"coreCount"`
	ClockRatio            float64      `json:"clockRatio"`
	AverageGHz            float64      `json:"averageGHz"`
	ExecutionAttemptCount int          `json:"executionAttemptCount"`
}

// Struct representing the by secret settings for secrets access rights
type BySecret struct {
	Key string `json:"key"`
}

// Struct representing the by prefix settings for secrets access rights
type ByPrefix struct {
	Prefix string `json:"prefix"`
}

// Struct representing the secrets and access rights for a task
type SecretsAccessRights struct {
	BySecret []BySecret `json:"bySecret"`
	ByPrefix []ByPrefix `json:"byPrefix"`
}

// Struct representing the dependencies for a task
type Dependencies struct {
	DependsOn []string `json:"dependsOn"`
}

// Struct representing the privileges for a task
type Privileges struct {
	ExportApiAndStorageCredentialsInEnvironment bool `json:"exportApiAndStorageCredentialsInEnvironment"`
}

// Struct representing the retry settings for a task
type RetrySettings struct {
	MaxTotalRetries       int `json:"maxTotalRetries"`
	MaxPerInstanceRetries int `json:"maxPerInstanceRetries"`
}

// Enum for the scheduling type
type SchedulingType string

const (
	Flex     SchedulingType = "flex"
	OnDemand SchedulingType = "onDemand"
	Reserved SchedulingType = "reserved"
)

// Struct representing a task with full details
type Task struct {
	Errors                              []QErrorPublic       `json:"errors"`
	ResourceBuckets                     []string             `json:"resourceBuckets"`
	AdvancedResourceBuckets             []string             `json:"advancedResourceBuckets"`
	ResultBucket                        string               `json:"resultBucket"`
	CompletedInstances                  []CompletedInstance  `json:"completedInstances"`
	Status                              TaskStatus           `json:"status"`
	SnapshotInterval                    int                  `json:"snapshotInterval"`
	ResultsCount                        int                  `json:"resultsCount"`
	Constants                           []Constant           `json:"constants"`
	SecretsAccessRights                 SecretsAccessRights  `json:"secretsAccessRights"`
	Tags                                []string             `json:"tags"`
	SnapshotWhitelist                   string               `json:"snapshotWhitelist"`
	SnapshotBlacklist                   string               `json:"snapshotBlacklist"`
	UploadResultsOnCancellation         bool                 `json:"uploadResultsOnCancellation"`
	Dependencies                        Dependencies         `json:"dependencies"`
	AutoDeleteOnCompletion              bool                 `json:"autoDeleteOnCompletion"`
	CompletionTimeToLive                string               `json:"completionTimeToLive"`
	HardwareConstraints                 []HardwareConstraint `json:"hardwareConstraints"`
	Labels                              map[string]string    `json:"labels"`
	SchedulingType                      SchedulingType       `json:"schedulingType"`
	Privileges                          Privileges           `json:"privileges"`
	RetrySettings                       RetrySettings        `json:"retrySettings"`
	UUID                                string               `json:"uuid"`
	Name                                string               `json:"name"`
	Shortname                           string               `json:"shortname"`
	Profile                             string               `json:"profile"`
	PoolUUID                            string               `json:"poolUuid"`
	JobUUID                             string               `json:"jobUuid"`
	Progress                            float64              `json:"progress"`
	RunningInstanceCount                int                  `json:"runningInstanceCount"`
	RunningCoreCount                    int                  `json:"runningCoreCount"`
	ExecutionTime                       string               `json:"executionTime"`
	WallTime                            string               `json:"wallTime"`
	State                               string               `json:"state"`
	PreviousState                       string               `json:"previousState"`
	InstanceCount                       int                  `json:"instanceCount"`
	MaxRetriesPerInstance               int                  `json:"maxRetriesPerInstance"`
	StateTransitionTime                 time.Time            `json:"stateTransitionTime"`
	PreviousStateTransitionTime         time.Time            `json:"previousStateTransitionTime"`
	LastModified                        time.Time            `json:"lastModified"`
	CreationDate                        time.Time            `json:"creationDate"`
	EndDate                             time.Time            `json:"endDate"`
	WaitForPoolResourcesSynchronization bool                 `json:"waitForPoolResourcesSynchronization"`
}

// Will list the tasks for the authenticated user
func (c *Client) ListTasks() []Task {
	data, _ := c.sendRequest(
		"GET",
		[]byte{},
		make(map[string]string),
		"tasks",
	)

	var tasks []Task
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks
}

// Will get the info for a task
func (c *Client) GetTaskInfo(uuid string) Task {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))

	var taskInfo Task
	err := json.Unmarshal(data, &taskInfo)
	if err != nil {
		panic(err)
	}

	return taskInfo
}

// Enum for the access constant string
type AccessConstant string

const (
	ReadOnly  AccessConstant = "readOnly"
	ReadWrite AccessConstant = "readWrite"
)

// Struct representing a ForcedConstant for the CreateTask payload
type ForcedConstant struct {
	ConstantName             string         `json:"constantName"`
	ForcedValue              string         `json:"forcedValue"`
	ForceExportInEnvironment bool           `json:"forceExportInEnvironment"`
	Access                   AccessConstant `json:"access"`
}

// Struct representing the payload for the CreateTask method
type CreateTaskPayload struct {
	Name                                string                        `json:"name"`
	Shortname                           string                        `json:"shortname,omitempty"`
	Profile                             string                        `json:"profile,omitempty"`
	PoolUUID                            string                        `json:"poolUuid,omitempty"`
	JobUUID                             string                        `json:"jobUuid,omitempty"`
	InstanceCount                       int                           `json:"instanceCount,omitempty"`
	AdvancedRanges                      string                        `json:"advancedRanges,omitempty"`
	ResourceBuckets                     string                        `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             *[]TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	ResultBucket                        string                        `json:"resultBucket,omitempty"`
	Constants                           *[]Constant                   `json:"constants,omitempty"`
	ForcedConstants                     *[]ForcedConstant             `json:"forcedConstants,omitempty"`
	Constraints                         *[]map[string]string          `json:"constraints,omitempty"`
	HardwareConstraints                 *[]HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	SecretsAccessRights                 *SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                      `json:"tags,omitempty"`
	SnapshotWhitelist                   string                        `json:"snapshotWhitelist,omitempty"`
	SnapshotBucket                      string                        `json:"snapshotBucket,omitempty"`
	SnapshotBucketPrefix                string                        `json:"snapshotBucketPrefix,omitempty"`
	ResultsWhitelist                    string                        `json:"resultsWhitelist,omitempty"`
	ResultsBlacklist                    string                        `json:"resultsBlacklist,omitempty"`
	ResultsBucket                       string                        `json:"resultsBucket,omitempty"`
	ResultsBucketPrefix                 string                        `json:"resultsBucketPrefix,omitempty"`
	Priority                            int                           `json:"priority,omitempty"`
	Dependencies                        *Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                          `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                int                           `json:"completionTimeToLive,omitempty"`
	WaitForPoolResourcesSynchronization bool                          `json:"waitForPoolResourcesSynchronization,omitempty"`
	UploadResultsOnCancellation         bool                          `json:"uploadResultsOnCancellation,omitempty"`
	Labels                              *map[string]string            `json:"labels,omitempty"`
	SchedulingType                      SchedulingType                `json:"schedulingType,omitempty"`
	TargetedReservedMachineKey          string                        `json:"targetedReservedMachineKey,omitempty"`
	DefaultResourcesCacheTTLSec         int                           `json:"defaultResourcesCacheTTLSec,omitempty"`
	Privileges                          *Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       *RetrySettings                `json:"retrySettings,omitempty"`
}

// Struct representing a UUID response
type UUIDResponse struct {
	Uuid string `json:"uuid"`
}

// Will create a task, based on a `CreateTaskPayload`
// Returns a `UUIDResponse` struct, containing a UUID for the newly created task
func (c *Client) CreateTask(payload CreateTaskPayload) UUIDResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, "tasks")
	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

// Struct representing a task summary
type TaskSummary struct {
	Uuid                                string
	Name                                string
	Shortname                           string
	Profile                             string
	PoolUuid                            string
	JobUuid                             string
	Progress                            float64
	RunningInstanceCount                int
	RunningCoreCount                    int
	ExecutionTime                       string
	WallTime                            string
	State                               string
	PreviousState                       string
	InstanceCount                       int
	AdvancedRanges                      string
	StateTransitionTime                 time.Time
	PreviousStateTransitionTime         time.Time
	LastModified                        time.Time
	CreationDate                        time.Time
	EndDate                             time.Time
	WaitForPoolResourcesSynchronization bool
}

// Will list task summaries for the authenticated user
func (c *Client) ListTaskSummaries() []TaskSummary {
	data, _ := c.sendRequest("GET", []byte{}, nil, "tasks/summaries")

	var summaries []TaskSummary
	err := json.Unmarshal(data, &summaries)
	if err != nil {
		panic(err)
	}

	return summaries
}

// Will delete a task
func (c *Client) DeleteTask(uuid string) {
	c.sendRequest("DELETE", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
}

// Will abort a task
func (c *Client) AbortTask(uuid string) {
	c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/abort", uuid))
}

// Will get the stdout for a task
func (c *Client) GetTaskStdout(uuid string) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

// Will get the previous stdout for a task
func (c *Client) GetLastTaskStdout(uuid string) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

// Will get the stdout for a task on a specific instance
func (c *Client) GetTaskInstanceStdout(uuid string, instanceId int) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

// Will get the previous stdout for a task on a specific instance
func (c *Client) GetLastTaskInstanceStdout(uuid string, instanceId int) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

// Will get the stderr for a task
func (c *Client) GetTaskStderr(uuid string) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

// Will get the previous stderr for a task
func (c *Client) GetLastTaskStderr(uuid string) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

// Will get the stderr for a task on a specific instance
func (c *Client) GetInstanceTaskStderr(uuid string, instanceId int) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

// Will get the previous stderr for a task on a specific instance
func (c *Client) GetInstanceLastTaskStderr(uuid string, instanceId int) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

// Struct representing the payload for creating unique and periodic snapshot
// Interval should only be filled for periodic snapshot, otherwise it will just get ignored
type CreateTaskSnapshotPayload struct {
	Interval     int    `json:"interval,omitempty"`
	Whitelist    string `json:"whitelist,omitempty"`
	Blacklist    string `json:"blacklist,omitempty"`
	Bucket       string `json:"bucket,omitempty"`
	BucketPrefix string `json:"bucketPrefix,omitempty"`
}

// Will create a periodic snapshot for a task using the UUID as string and a `CreateTaskSnapshotPayload` struct as arguments
func (c *Client) CreateTaskPeriodicSnapshot(uuid string, payload CreateTaskSnapshotPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot/periodic", uuid))
}

// Will create a unique snapshot for a task using the UUID as string and a `CreateTaskSnapshotPayload` struct as arguments
func (c *Client) CreateTaskUniqueSnapshot(uuid string, payload CreateTaskSnapshotPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot", uuid))
}

// Will retry a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly retried task
func (c *Client) RetryTask(uuid string, payload CreateTaskPayload) UUIDResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/retry", uuid))

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

// Will recover a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly recovered task
func (c *Client) RecoverTask(uuid string, payload CreateTaskPayload) UUIDResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/recover", uuid))

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

// Will resume a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly resumed task
func (c *Client) ResumeTask(uuid string, payload CreateTaskPayload) UUIDResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/resume", uuid))

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

// Will clone a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly cloned task
func (c *Client) CloneTask(uuid string, payload CreateTaskPayload) UUIDResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/clone", uuid))

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

// A struct representing the payload for the `UpdateTask` method
type UpdateTaskPayload struct {
	Constants []Constant `json:"constants,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

// Will update the fields of a task using the UUID as an argument, as well as a `UpdateTaskPayload` struct
func (c *Client) UpdateTask(uuid string, payload UpdateTaskPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c.sendRequest("PUT", payloadJson, nil, fmt.Sprintf("tasks/%v", uuid))
}

func (c *Client) UpdateTaskResources(uuid string) {
	c.sendRequest("PATCH", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
}
