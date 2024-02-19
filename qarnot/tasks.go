package qarnot

import (
	"encoding/json"
	"fmt"
	"time"
)

type PrefixFiltering struct {
	Prefix string `json:"prefix"`
}

type Filtering struct {
	PrefixFiltering PrefixFiltering `json:"prefixFiltering"`
}

type PerRunningInstanceInfo struct {
	ActiveForwards        any    `json:"activeForwards"`
	Phase                 string `json:"phase"`
	InstanceID            int    `json:"instanceId"`
	MaxFrequencyGHz       int    `json:"maxFrequencyGHz"`
	CurrentFrequencyGHz   int    `json:"currentFrequencyGHz"`
	CPUUsage              int    `json:"cpuUsage"`
	MaxMemoryMB           int    `json:"maxMemoryMB"`
	CurrentMemoryMB       int    `json:"currentMemoryMB"`
	MemoryUsage           int    `json:"memoryUsage"`
	NetworkInKbps         int    `json:"networkInKbps"`
	NetworkOutKbps        int    `json:"networkOutKbps"`
	Progress              int    `json:"progress"`
	ExecutionTimeSec      int    `json:"executionTimeSec"`
	ExecutionTimeGHz      int    `json:"executionTimeGHz"`
	SpecificationKey      any    `json:"specificationKey"`
	CPUModel              any    `json:"cpuModel"`
	CoreCount             int    `json:"coreCount"`
	ExecutionAttemptCount int    `json:"executionAttemptCount"`
	ClockRatio            int    `json:"clockRatio"`
}

type RunningInstancesInfo struct {
	PerRunningInstanceInfo     []PerRunningInstanceInfo `json:"perRunningInstanceInfo"`
	SnapshotResults            []any                    `json:"snapshotResults"`
	Timestamp                  time.Time                `json:"timestamp"`
	AverageFrequencyGHz        int                      `json:"averageFrequencyGHz"`
	MaxFrequencyGHz            int                      `json:"maxFrequencyGHz"`
	MinFrequencyGHz            int                      `json:"minFrequencyGHz"`
	AverageMaxFrequencyGHz     int                      `json:"averageMaxFrequencyGHz"`
	AverageCPUUsage            int                      `json:"averageCpuUsage"`
	ClusterPowerIndicator      int                      `json:"clusterPowerIndicator"`
	AverageMemoryUsage         int                      `json:"averageMemoryUsage"`
	AverageNetworkInKbps       int                      `json:"averageNetworkInKbps"`
	AverageNetworkOutKbps      int                      `json:"averageNetworkOutKbps"`
	TotalNetworkInKbps         int                      `json:"totalNetworkInKbps"`
	TotalNetworkOutKbps        int                      `json:"totalNetworkOutKbps"`
	RunningCoreCountByCPUModel any                      `json:"runningCoreCountByCpuModel"`
}

type ExecutionTimeByCPUModel struct {
	Model string  `json:"model"`
	Time  float64 `json:"time"`
	Core  int     `json:"core"`
}

type ExecutionTimeByMachineSpecification struct {
	SpecificationKey string  `json:"specificationKey"`
	Time             float64 `json:"time"`
}

type ExecutionTimeByInstanceID struct {
	InstanceId       int     `json:"instanceId"`
	SpecificationKey string  `json:"specificationKey"`
	SiteUuid         string  `json:"siteUuid"`
	Time             float64 `json:"time"`
	TimeGHz          float64 `json:"timeGHz"`
	ClockRatio       float64 `json:"clockRatio"`
}

type ExecutionTimeGhzByCPUModel struct {
	Model      string  `json:"model"`
	TimeGHz    float64 `json:"timeGHz"`
	ClockRatio float64 `json:"clockRatio"`
	Core       int     `json:"core"`
}

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

type Constant struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type QErrorPublic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Debug   string `json:"debug"`
}

type StripPrefix struct {
	Prefix string `json:"prefix"`
}

type ResourceTransformation struct {
	StripPrefix StripPrefix `json:"stripPrefix"`
}

type TaskAdvancedResourceBucket struct {
	BucketName             string                 `json:"bucketName"`
	Filtering              Filtering              `json:"filtering"`
	ResourceTransformation ResourceTransformation `json:"resourceTransformation"`
	CacheTTLSec            int                    `jons:"cacheTTLSec"`
}

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

type BySecret struct {
	Key string `json:"key"`
}

type ByPrefix struct {
	Prefix string `json:"prefix"`
}

type SecretsAccessRights struct {
	BySecret []BySecret `json:"bySecret"`
	ByPrefix []ByPrefix `json:"byPrefix"`
}

type Dependencies struct {
	DependsOn []string `json:"dependsOn"`
}

type Privileges struct {
	ExportApiAndStorageCredentialsInEnvironment bool `json:"exportApiAndStorageCredentialsInEnvironment"`
}

type RetrySettings struct {
	MaxTotalRetries       int `json:"maxTotalRetries"`
	MaxPerInstanceRetries int `json:"maxPerInstanceRetries"`
}

type SchedulingType string

const (
	Flex     SchedulingType = "flex"
	OnDemand SchedulingType = "onDemand"
	Reserved SchedulingType = "reserved"
)

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

func (c *Client) GetTaskInfo(uuid string) Task {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))

	var taskInfo Task
	err := json.Unmarshal(data, &taskInfo)
	if err != nil {
		panic(err)
	}

	return taskInfo
}

type AccessConstant string

const (
	ReadOnly  AccessConstant = "readOnly"
	ReadWrite AccessConstant = "readWrite"
)

type ForcedConstant struct {
	ConstantName             string         `json:"constantName"`
	ForcedValue              string         `json:"forcedValue"`
	ForceExportInEnvironment bool           `json:"forceExportInEnvironment"`
	Access                   AccessConstant `json:"access"`
}

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

type UUIDResponse struct {
	Uuid string `json:"uuid"`
}

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

func (c *Client) ListTaskSummaries() []TaskSummary {
	data, _ := c.sendRequest("GET", []byte{}, nil, "tasks/summaries")

	var summaries []TaskSummary
	err := json.Unmarshal(data, &summaries)
	if err != nil {
		panic(err)
	}

	return summaries
}

func (c *Client) DeleteTask(uuid string) {
	c.sendRequest("DELETE", []byte{}, nil, fmt.Sprintf("tasks/%v", uuid))
}

func (c *Client) AbortTask(uuid string) {
	c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/abort", uuid))
}

func (c *Client) GetTaskStdout(uuid string) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

func (c *Client) GetLastTaskStdout(uuid string) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout", uuid))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

func (c *Client) GetTaskInstanceStdout(uuid string, instanceId int) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

func (c *Client) GetLastTaskInstanceStdout(uuid string, instanceId int) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stdout/%v", uuid, instanceId))

	var stdout string
	err := json.Unmarshal(data, &stdout)
	if err != nil {
		panic(err)
	}

	return stdout
}

func (c *Client) GetTaskStderr(uuid string) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

func (c *Client) GetLastTaskStderr(uuid string) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

func (c *Client) GetInstanceTaskStderr(uuid string) string {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

func (c *Client) GetInstanceLastTaskStderr(uuid string) string {
	data, _ := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr", uuid))

	var stderr string
	err := json.Unmarshal(data, &stderr)
	if err != nil {
		panic(err)
	}

	return stderr
}

type CreateTaskSnapshotPayload struct {
	Interval     int    `json:"interval,omitempty"`
	Whitelist    string `json:"whitelist,omitempty"`
	Blacklist    string `json:"blacklist,omitempty"`
	Bucket       string `json:"bucket,omitempty"`
	BucketPrefix string `json:"bucketPrefix,omitempty"`
}

func (c *Client) CreateTaskPeriodicSnapshot(uuid string, payload CreateTaskSnapshotPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot/periodic", uuid))
}

func (c *Client) CreateTaskUniqueSnapshot(uuid string, payload CreateTaskSnapshotPayload) {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot", uuid))
}

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

type UpdateTaskPayload struct {
	Constants []Constant `json:"constants,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

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
