package qarnot

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
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

// Struct representing the payload for creating unique and periodic snapshot
// Interval should only be filled for periodic snapshot, otherwise it will just get ignored
type CreateTaskSnapshotPayload struct {
	Interval     int    `json:"interval,omitempty"`
	Whitelist    string `json:"whitelist,omitempty"`
	Blacklist    string `json:"blacklist,omitempty"`
	Bucket       string `json:"bucket,omitempty"`
	BucketPrefix string `json:"bucketPrefix,omitempty"`
}

// A struct representing the payload for the `UpdateTask` method
type UpdateTaskPayload struct {
	Constants []Constant `json:"constants,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

// Will list the tasks for the authenticated user
func (c *Client) ListTasks() ([]Task, error) {
	data, statusCode, err := c.sendRequest(
		"GET",
		[]byte{},
		make(map[string]string),
		"tasks",
	)
	if err != nil {
		return []Task{}, fmt.Errorf("could not list tasks due to the following error (HTTP %v) : %v", statusCode, err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	helpers.JsonUnmarshalCheckError(err)

	return tasks, nil
}

// Will get the info for a task
func (c *Client) GetTaskInfo(uuid string) (Task, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
	if err != nil {
		return Task{}, fmt.Errorf("could not get task info due to the following error : %v", err)
	}

	var taskInfo Task
	err = json.Unmarshal(data, &taskInfo)
	helpers.JsonUnmarshalCheckError(err)

	return taskInfo, nil
}

// Will create a task, based on a `CreateTaskPayload`
// Returns a `UUIDResponse` struct, containing a UUID for the newly created task
func (c *Client) CreateTask(payload CreateTaskPayload) (UUIDResponse, error) {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	data, _, err := c.sendRequest("POST", payloadJson, nil, "tasks")
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not create task due to the following error : %v", err)
	}

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)

	return response, nil
}

// Will list task summaries for the authenticated user
func (c *Client) ListTaskSummaries() ([]TaskSummary, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, "tasks/summaries")
	if err != nil {
		return []TaskSummary{}, fmt.Errorf("could not list task summaries due to the following error : %v", err)
	}

	var summaries []TaskSummary
	err = json.Unmarshal(data, &summaries)
	helpers.JsonUnmarshalCheckError(err)

	return summaries, nil
}

// Will delete a task
func (c *Client) DeleteTask(uuid string) error {
	_, _, err := c.sendRequest("DELETE", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
	if err != nil {
		return fmt.Errorf("could not delete task due to the following error : %v", err)
	}
	return nil
}

// Will abort a task
func (c *Client) AbortTask(uuid string) error {
	_, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/abort", uuid))
	if err != nil {
		return fmt.Errorf("could not abort task due to the following error : %v", err)
	}
	return nil
}

// Will get the stdout for a task
func (c *Client) GetTaskStdout(uuid string) (string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))
	if err != nil {
		return "", fmt.Errorf("could not get task stdout due to the following error : %v", err)
	}

	var stdout string
	err = json.Unmarshal(data, &stdout)
	helpers.JsonUnmarshalCheckError(err)

	return stdout, nil
}

// Will get the previous stdout for a task
func (c *Client) GetLastTaskStdout(uuid string) (string, error) {
	data, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))
	if err != nil {
		return "", fmt.Errorf("could not get last task stdout due to the following error : %v", err)
	}

	var stdout string
	err = json.Unmarshal(data, &stdout)
	helpers.JsonUnmarshalCheckError(err)

	return stdout, nil
}

// Will get the stdout for a task on a specific instance
func (c *Client) GetTaskInstanceStdout(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get task instance stdout due to the following error : %v", err)
	}

	var stdout string
	err = json.Unmarshal(data, &stdout)
	helpers.JsonUnmarshalCheckError(err)

	return stdout, nil
}

// Will get the previous stdout for a task on a specific instance
func (c *Client) GetLastTaskInstanceStdout(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get last task instance stdout due to the following error : %v", err)
	}

	var stdout string
	err = json.Unmarshal(data, &stdout)
	helpers.JsonUnmarshalCheckError(err)

	return stdout, nil
}

// Will get the stderr for a task
func (c *Client) GetTaskStderr(uuid string) (string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))
	if err != nil {
		return "", fmt.Errorf("could not get task stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	helpers.JsonUnmarshalCheckError(err)

	return stderr, nil
}

// Will get the previous stderr for a task
func (c *Client) GetLastTaskStderr(uuid string) (string, error) {
	data, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))
	if err != nil {
		return "", fmt.Errorf("could not get last task stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	helpers.JsonUnmarshalCheckError(err)

	return stderr, nil
}

// Will get the stderr for a task on a specific instance
func (c *Client) GetInstanceTaskStderr(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get task instance stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	helpers.JsonUnmarshalCheckError(err)

	return stderr, nil
}

// Will get the previous stderr for a task on a specific instance
func (c *Client) GetInstanceLastTaskStderr(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get last task instance stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	helpers.JsonUnmarshalCheckError(err)

	return stderr, nil
}

// Will create a periodic snapshot for a task using the UUID as string and a `CreateTaskSnapshotPayload` struct as arguments
func (c *Client) CreateTaskPeriodicSnapshot(uuid string, payload CreateTaskSnapshotPayload) error {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	_, _, err = c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot/periodic", uuid))
	if err != nil {
		return fmt.Errorf("could not create a task periodic snapshot due to the following error : %v", err)
	}

	return nil
}

// Will create a unique snapshot for a task using the UUID as string and a `CreateTaskSnapshotPayload` struct as arguments
func (c *Client) CreateTaskUniqueSnapshot(uuid string, payload CreateTaskSnapshotPayload) error {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	_, _, err = c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot", uuid))
	if err != nil {
		return fmt.Errorf("could not create a task unique snapshot due to the following error : %v", err)
	}

	return nil
}

// Will retry a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly retried task
func (c *Client) RetryTask(uuid string, payload CreateTaskPayload) (UUIDResponse, error) {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/retry", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not retry task due to the following error : %v", err)
	}

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)

	return response, nil
}

// Will recover a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly recovered task
func (c *Client) RecoverTask(uuid string, payload CreateTaskPayload) (UUIDResponse, error) {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/recover", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not recover task due to the following error : %v", err)
	}

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)
	return response, nil
}

// Will resume a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly resumed task
func (c *Client) ResumeTask(uuid string, payload CreateTaskPayload) (UUIDResponse, error) {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/resume", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not resume task due to the following error : %v", err)
	}

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)

	return response, nil
}

// Will clone a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly cloned task
func (c *Client) CloneTask(uuid string, payload CreateTaskPayload) (UUIDResponse, error) {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/clone", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not clone task due to the following error : %v", err)
	}

	var response UUIDResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)
	return response, nil
}

// Will update the fields of a task using the UUID as an argument, as well as a `UpdateTaskPayload` struct
func (c *Client) UpdateTask(uuid string, payload UpdateTaskPayload) error {
	payloadJson, err := json.Marshal(payload)
	helpers.JsonMarshalCheckError(err)

	_, _, err = c.sendRequest("PUT", payloadJson, nil, fmt.Sprintf("tasks/%v", uuid))
	if err != nil {
		return fmt.Errorf("could not update task due to the following error : %v", err)
	}

	return nil
}

func (c *Client) UpdateTaskResources(uuid string) error {
	_, _, err := c.sendRequest("PATCH", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
	if err != nil {
		return fmt.Errorf("could not update task resources due to the following error : %v", err)
	}

	return nil
}
