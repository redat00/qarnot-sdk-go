package qarnot

import (
	"encoding/json"
	"fmt"
	"net/http"
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

// Struct representing an active formward settings
type ActiveForward struct {
	ApplicationPort int    `json:"applicationPort"`
	ForwarderPort   int    `json:"forwarderPort"`
	ForwarderHost   string `json:"forwarderHost"`
	BindAddress     string `json:"bindAddress"`
}

// Struct representing a VPN Connection
type VpnConnection struct {
	VpnName           string `json:"vpnName"`
	NodeIPAddressCidr string `json:"nodeIPAddressCidr"`
}

type PerRunningInstanceInfo struct {
	ActiveForwards        []ActiveForward `json:"activeForwards"`
	VpnConnections        []VpnConnection `json:"vpnConnections"`
	Phase                 string          `json:"phase"`
	InstanceID            int             `json:"instanceId"`
	MaxFrequencyGHz       float64         `json:"maxFrequencyGHz"`
	CurrentFrequencyGHz   float64         `json:"currentFrequencyGHz"`
	CpuUsage              float64         `json:"cpuUsage"`
	MaxMemoryMB           int             `json:"maxMemoryMB"`
	CurrentMemoryMB       int             `json:"currentMemoryMB"`
	MemoryUsage           float64         `json:"memoryUsage"`
	NetworkInKbps         float64         `json:"networkInKbps"`
	NetworkOutKbps        float64         `json:"networkOutKbps"`
	Progress              float64         `json:"progress"`
	ExecutionTimeSec      float64         `json:"executionTimeSec"`
	ExecutionTimeGHz      float64         `json:"executionTimeGHz"`
	SpecificationKey      string          `json:"specificationKey"`
	CpuModel              string          `json:"cpuModel"`
	CoreCount             int             `json:"coreCount"`
	ExecutionAttemptCount int             `json:"executionAttemptCount"`
	ClockRatio            float64         `json:"clockRatio"`
}

type RunningCoreCountByCpuModel struct {
	Model            string `json:"model"`
	Core             int    `json:"core"`
	RunningCoreCount int    `json:"runningCoreCount"`
}

type RunningInstancesInfo struct {
	PerRunningInstanceInfo     []PerRunningInstanceInfo     `json:"perRunningInstanceInfo"`
	Timestamp                  string                       `json:"timestamp"`
	AverageFrequencyGHz        float64                      `json:"averageFrequencyGHz"`
	MaxFrequencyGHz            float64                      `json:"maxFrequencyGHz"`
	MinFrequencyGHz            float64                      `json:"minFrequencyGHz"`
	AverageMaxFrequencyGHz     float64                      `json:"averageMaxFrequencyGHz"`
	AverageCpuUsage            float64                      `json:"averageCpuUsage"`
	ClusterPowerIndicator      float64                      `json:"clusterPowerIndicator"`
	AverageMemoryUsage         float64                      `json:"averageMemoryUsage"`
	AverageNetworkInKbps       float64                      `json:"averageNetworkInKbps"`
	AverageNetworkOutKbps      float64                      `json:"averageNetworkOutKbps"`
	TotalNetworkInKbps         float64                      `json:"totalNetworkInKbps"`
	TotalNetworkOutKbps        float64                      `json:"totalNetworkOutKbps"`
	RunningCoreCountByCpuModel []RunningCoreCountByCpuModel `json:"runningCoreCountByCpuModel"`
}

// Struct representing a task status for a task
type TaskStatus struct {
	Timestamp                           time.Time                             `json:"timestamp,omitempty"`
	LastUpdateTimestamp                 time.Time                             `json:"lastUpdateTimestamp,omitempty"`
	DownloadProgress                    float64                               `json:"downloadProgress,omitempty"`
	ExecutionProgress                   float64                               `json:"executionProgress,omitempty"`
	UploadProgress                      float64                               `json:"uploadProgress,omitempty"`
	InstanceCount                       float64                               `json:"instanceCount,omitempty"`
	DownloadTime                        string                                `json:"downloadTime,omitempty"`
	DownloadTimeSec                     float64                               `json:"downloadTimeSec,omitempty"`
	EnvironmentTime                     string                                `json:"environmentTime,omitempty"`
	EnvironmentTimeSec                  float64                               `json:"environmentTimeSec,omitempty"`
	ExecutionTime                       string                                `json:"executionTime,omitempty"`
	ExecutionTimeSec                    float64                               `json:"executionTimeSec,omitempty"`
	ExecutionTimeByCPUModel             []ExecutionTimeByCPUModel             `json:"executionTimeByCpuModel,omitempty"`
	ExecutionTimeByMachineSpecification []ExecutionTimeByMachineSpecification `json:"executionTimeByMachineSpecification,omitempty"`
	ExecutionTimeByInstanceID           []ExecutionTimeByInstanceID           `json:"executionTimeByInstanceId,omitempty"`
	ExecutionTimeGhzByCPUModel          []ExecutionTimeGhzByCPUModel          `json:"executionTimeGhzByCpuModel,omitempty"`
	UploadTime                          string                                `json:"uploadTime,omitempty"`
	UploadTimeSec                       float64                               `json:"uploadTimeSec,omitempty"`
	WallTime                            string                                `json:"wallTime,omitempty"`
	WallTimeSec                         float64                               `json:"wallTimeSec,omitempty"`
	SucceededRange                      string                                `json:"succeededRange,omitempty"`
	ExecutedRange                       string                                `json:"executedRange,omitempty"`
	FailedRange                         string                                `json:"failedRange,omitempty"`
	CancelledRange                      string                                `json:"cancelledRange,omitempty"`
	FailedOnlyRange                     string                                `json:"failedOnlyRange,omitempty"`
	StartedOnceRange                    string                                `json:"startedOnceRange,omitempty"`
	RunningInstancesInfo                RunningInstancesInfo                  `json:"runningInstancesInfo,omitempty"`
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

type Proto string

const (
	TCP   Proto = "tcp"
	UDP   Proto = "udp"
	HTTPS Proto = "https"
)

// Struct representing the network rules for a task
type ForcedNetworkRule struct {
	Inbound         bool   `json:"inbound"`
	Proto           Proto  `json:"proto"`
	Port            string `json:"port"`
	To              string `json:"to,omitempty"`
	PublicHost      string `json:"publicHost,omitempty"`
	PublicPort      string `json:"publicPort,omitempty"`
	Forwarder       string `json:"forwarder,omitempty"`
	Priority        string `json:"priority,omitempty"`
	Description     string `json:"description,omitempty"`
	ToQbox          bool   `json:"toQbox,omitempty"`
	ToPayload       bool   `json:"toPayload,omitempty"`
	Name            string `json:"name,omitempty"`
	ApplicationType string `json:"applicationType,omitempty"`
}

// Struct representing a task with full details
type Task struct {
	Errors                              []QErrorPublic               `json:"errors,omitempty"`
	ResourceBuckets                     []string                     `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             []TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	ResultBucket                        string                       `json:"resultBucket,omitempty"`
	CompletedInstances                  []CompletedInstance          `json:"completedInstances,omitempty"`
	Status                              TaskStatus                   `json:"status,omitempty"`
	SnapshotInterval                    int                          `json:"snapshotInterval,omitempty"`
	ResultsCount                        int                          `json:"resultsCount,omitempty"`
	Constants                           []Constant                   `json:"constants,omitempty"`
	SecretsAccessRights                 SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                     `json:"tags,omitempty"`
	SnapshotWhitelist                   string                       `json:"snapshotWhitelist,omitempty"`
	SnapshotBlacklist                   string                       `json:"snapshotBlacklist,omitempty"`
	UploadResultsOnCancellation         bool                         `json:"uploadResultsOnCancellation,omitempty"`
	Dependencies                        Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                         `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                string                       `json:"completionTimeToLive,omitempty"`
	HardwareConstraints                 []HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	Labels                              map[string]string            `json:"labels,omitempty"`
	SchedulingType                      SchedulingType               `json:"schedulingType,omitempty"`
	Privileges                          Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       RetrySettings                `json:"retrySettings,omitempty"`
	UUID                                string                       `json:"uuid,omitempty"`
	Name                                string                       `json:"name,omitempty"`
	Shortname                           string                       `json:"shortname,omitempty"`
	Profile                             string                       `json:"profile,omitempty"`
	PoolUUID                            string                       `json:"poolUuid,omitempty"`
	JobUUID                             string                       `json:"jobUuid,omitempty"`
	Progress                            float64                      `json:"progress,omitempty"`
	RunningInstanceCount                int                          `json:"runningInstanceCount,omitempty"`
	RunningCoreCount                    int                          `json:"runningCoreCount,omitempty"`
	ExecutionTime                       string                       `json:"executionTime,omitempty"`
	WallTime                            string                       `json:"wallTime,omitempty"`
	State                               string                       `json:"state,omitempty"`
	PreviousState                       string                       `json:"previousState,omitempty"`
	InstanceCount                       int                          `json:"instanceCount,omitempty"`
	MaxRetriesPerInstance               int                          `json:"maxRetriesPerInstance,omitempty"`
	StateTransitionTime                 time.Time                    `json:"stateTransitionTime,omitempty"`
	PreviousStateTransitionTime         time.Time                    `json:"previousStateTransitionTime,omitempty"`
	LastModified                        time.Time                    `json:"lastModified,omitempty"`
	CreationDate                        time.Time                    `json:"creationDate,omitempty"`
	EndDate                             time.Time                    `json:"endDate,omitempty"`
	WaitForPoolResourcesSynchronization bool                         `json:"waitForPoolResourcesSynchronization,omitempty"`
	ForcedNetworkRules                  []ForcedNetworkRule          `json:"forcedNetworkRules,omitempty"`
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
	ForcedNetworkRules                  *[]ForcedNetworkRule          `json:"forcedNetworkRules,omitempty"`
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
	MaxRetriesPerInstance               int
	AdvancedRanges                      string
	StateTransitionTime                 time.Time
	PreviousStateTransitionTime         time.Time
	LastModified                        time.Time
	CreationDate                        time.Time
	EndDate                             time.Time
	WaitForPoolResourcesSynchronization bool
}

// Will list the tasks for the authenticated user
// Optionally filter the results if any tags are provided
func (c *Client) ListTasks(tags ...string) ([]Task, error) {
	addQuery := func(req *http.Request) error {
		query := req.URL.Query()
		for _, tag := range tags {
			query.Add("tag", tag)
		}
		req.URL.RawQuery = query.Encode()
		return nil
	}

	data, _, err := c.sendRequest(
		"GET",
		[]byte{},
		make(map[string]string),
		"tasks",
		addQuery,
	)
	if err != nil {
		return []Task{}, fmt.Errorf("could not list tasks due to the following error : %v", err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return tasks, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return Task{}, helpers.FormatJsonUnmarshalError(err)
	}

	return taskInfo, nil
}

// Will create a task, based on a `CreateTaskPayload`
// Returns a `UUIDResponse` struct, containing a UUID for the newly created task
func (c *Client) CreateTask(payload *CreateTaskPayload) (UUIDResponse, error) {
	var response UUIDResponse

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return UUIDResponse{}, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, "tasks")
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not create task due to the following error : %v", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return summaries, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stdout, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stdout, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stdout, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stdout, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stderr, helpers.FormatJsonUnmarshalError(err)
	}

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
	if err != nil {
		return stderr, helpers.FormatJsonUnmarshalError(err)
	}

	return stderr, nil
}

// Will get the stderr for a task on a specific instance
func (c *Client) GetTaskInstanceStderr(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get task instance stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	if err != nil {
		return stderr, helpers.FormatJsonUnmarshalError(err)
	}

	return stderr, nil
}

// Will get the previous stderr for a task on a specific instance
func (c *Client) GetLastTaskInstanceStderr(uuid string, instanceId int) (string, error) {
	data, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("tasks/%v/stderr/%v", uuid, instanceId))
	if err != nil {
		return "", fmt.Errorf("could not get last task instance stderr due to the following error : %v", err)
	}

	var stderr string
	err = json.Unmarshal(data, &stderr)
	if err != nil {
		return stderr, helpers.FormatJsonUnmarshalError(err)
	}

	return stderr, nil
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
func (c *Client) CreateTaskPeriodicSnapshot(uuid string, payload *CreateTaskSnapshotPayload) error {
	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return helpers.FormatJsonMarshalError(err)
	}

	if _, _, err = c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot/periodic", uuid)); err != nil {
		return fmt.Errorf("could not create a task periodic snapshot due to the following error : %v", err)
	}

	return nil
}

// Will create a unique snapshot for a task using the UUID as string and a `CreateTaskSnapshotPayload` struct as arguments
func (c *Client) CreateTaskUniqueSnapshot(uuid string, payload *CreateTaskSnapshotPayload) error {
	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return helpers.FormatJsonMarshalError(err)
	}

	_, _, err = c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/snapshot", uuid))
	if err != nil {
		return fmt.Errorf("could not create a task unique snapshot due to the following error : %v", err)
	}

	return nil
}

// Struct representing the payload the `RetryTask` method
// When using it, the field that been filled will be used to update parameters of the task you're retrying from
// A new task will be created with all the parameters from the old task you're retrying, and the updated settings you've set here
type RetryTaskPayload struct {
	Name                                string                       `json:"name,omitempty"`
	ResourceBuckets                     []string                     `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             []TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	Shortname                           string                       `json:"shortname,omitempty"`
	Profile                             string                       `json:"profile,omitempty"`
	JobUuid                             string                       `json:"jobUuid,omitempty"`
	ResultBucket                        string                       `json:"resultBucket,omitempty"`
	Constants                           []Constant                   `json:"constants,omitempty"`
	ForcedConstants                     []ForcedConstant             `json:"forcedConstants,omitempty"`
	Constraints                         []map[string]string          `json:"constraints,omitempty"`
	HardwareConstraints                 []HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	SecretsAccessRights                 SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                     `json:"tags,omitempty"`
	SnapshotWhitelist                   string                       `json:"snapshotWhitelist,omitempty"`
	SnapshotBlacklist                   string                       `json:"snapshotBlacklist,omitempty"`
	SnapshotBucket                      string                       `json:"snapshotBucket,omitempty"`
	SnapshotBucketPrefix                string                       `json:"snapshotBucketPrefix,omitempty"`
	ResultsWhitelist                    string                       `json:"resultsWhitelist,omitempty"`
	ResultsBlacklist                    string                       `json:"resultsBlacklist,omitempty"`
	ResultsBucket                       string                       `json:"resultsBucket,omitempty"`
	ResultsBucketPrefix                 string                       `json:"resultsBucketPrefix,omitempty"`
	Priority                            int                          `json:"priority,omitempty"`
	Dependencies                        Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                         `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                string                       `json:"completionTimeToLive,omitempty"`
	WaitForPoolResourcesSynchronization bool                         `json:"waitForPoolResourcesSynchronization,omitempty"`
	UploadResultsOnCancellation         bool                         `json:"uploadResultsOnCancellation,omitempty"`
	Labels                              []map[string]string          `json:"labels,omitempty"`
	SchedulingType                      SchedulingType               `json:"schedulingType,omitempty"`
	TargetedReservedMachineKey          string                       `json:"targetedReservedMachineKey,omitempty"`
	DefaultResourcesCacheTTLSec         int                          `json:"defaultResourcesCacheTTLSec,omitempty"`
	Privileges                          Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       RetrySettings                `json:"retrySettings,omitempty"`
}

// Will retry a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly retried task
func (c *Client) RetryTask(uuid string, payload *RetryTaskPayload) (UUIDResponse, error) {
	var response UUIDResponse

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return response, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/retry", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not retry task due to the following error : %v", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}

	return response, nil
}

// Struct representing the payload the `RecoverTask` method
// When using it, the field that been filled will be used to update parameters of the task you're recovering from
// A new task will be created with all the parameters from the old task you're recovering, and the updated settings you've set here
type RecoverTaskPayload struct {
	Name                                string                       `json:"name,omitempty"`
	ResourceBuckets                     []string                     `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             []TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	Shortname                           string                       `json:"shortname,omitempty"`
	Profile                             string                       `json:"profile,omitempty"`
	JobUuid                             string                       `json:"jobUuid,omitempty"`
	ResultBucket                        string                       `json:"resultBucket,omitempty"`
	Constants                           []Constant                   `json:"constants,omitempty"`
	ForcedConstants                     []ForcedConstant             `json:"forcedConstants,omitempty"`
	Constraints                         []map[string]string          `json:"constraints,omitempty"`
	HardwareConstraints                 []HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	SecretsAccessRights                 SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                     `json:"tags,omitempty"`
	SnapshotWhitelist                   string                       `json:"snapshotWhitelist,omitempty"`
	SnapshotBlacklist                   string                       `json:"snapshotBlacklist,omitempty"`
	SnapshotBucket                      string                       `json:"snapshotBucket,omitempty"`
	SnapshotBucketPrefix                string                       `json:"snapshotBucketPrefix,omitempty"`
	ResultsWhitelist                    string                       `json:"resultsWhitelist,omitempty"`
	ResultsBlacklist                    string                       `json:"resultsBlacklist,omitempty"`
	ResultsBucket                       string                       `json:"resultsBucket,omitempty"`
	ResultsBucketPrefix                 string                       `json:"resultsBucketPrefix,omitempty"`
	Priority                            int                          `json:"priority,omitempty"`
	Dependencies                        Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                         `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                string                       `json:"completionTimeToLive,omitempty"`
	WaitForPoolResourcesSynchronization bool                         `json:"waitForPoolResourcesSynchronization,omitempty"`
	UploadResultsOnCancellation         bool                         `json:"uploadResultsOnCancellation,omitempty"`
	Labels                              []map[string]string          `json:"labels,omitempty"`
	SchedulingType                      SchedulingType               `json:"schedulingType,omitempty"`
	TargetedReservedMachineKey          string                       `json:"targetedReservedMachineKey,omitempty"`
	DefaultResourcesCacheTTLSec         int                          `json:"defaultResourcesCacheTTLSec,omitempty"`
	Privileges                          Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       RetrySettings                `json:"retrySettings,omitempty"`
}

// Will recover a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly recovered task
func (c *Client) RecoverTask(uuid string, payload *RecoverTaskPayload) (UUIDResponse, error) {
	var response UUIDResponse

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return response, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/recover", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not recover task due to the following error : %v", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}
	return response, nil
}

// Struct representing the payload the `ResumeTask` method
// When using it, the field that been filled will be used to update parameters of the task you're resuming from
// A new task will be created with all the parameters from the old task you're resuming, and the updated settings you've set here
type ResumeTaskPayload struct {
	Name                                string                       `json:"name,omitempty"`
	ResourceBuckets                     []string                     `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             []TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	Shortname                           string                       `json:"shortname,omitempty"`
	Profile                             string                       `json:"profile,omitempty"`
	JobUuid                             string                       `json:"jobUuid,omitempty"`
	ResultBucket                        string                       `json:"resultBucket,omitempty"`
	Constants                           []Constant                   `json:"constants,omitempty"`
	ForcedConstants                     []ForcedConstant             `json:"forcedConstants,omitempty"`
	Constraints                         []map[string]string          `json:"constraints,omitempty"`
	HardwareConstraints                 []HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	SecretsAccessRights                 SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                     `json:"tags,omitempty"`
	SnapshotWhitelist                   string                       `json:"snapshotWhitelist,omitempty"`
	SnapshotBlacklist                   string                       `json:"snapshotBlacklist,omitempty"`
	SnapshotBucket                      string                       `json:"snapshotBucket,omitempty"`
	SnapshotBucketPrefix                string                       `json:"snapshotBucketPrefix,omitempty"`
	ResultsWhitelist                    string                       `json:"resultsWhitelist,omitempty"`
	ResultsBlacklist                    string                       `json:"resultsBlacklist,omitempty"`
	ResultsBucket                       string                       `json:"resultsBucket,omitempty"`
	ResultsBucketPrefix                 string                       `json:"resultsBucketPrefix,omitempty"`
	Priority                            int                          `json:"priority,omitempty"`
	Dependencies                        Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                         `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                string                       `json:"completionTimeToLive,omitempty"`
	WaitForPoolResourcesSynchronization bool                         `json:"waitForPoolResourcesSynchronization,omitempty"`
	UploadResultsOnCancellation         bool                         `json:"uploadResultsOnCancellation,omitempty"`
	Labels                              []map[string]string          `json:"labels,omitempty"`
	SchedulingType                      SchedulingType               `json:"schedulingType,omitempty"`
	TargetedReservedMachineKey          string                       `json:"targetedReservedMachineKey,omitempty"`
	DefaultResourcesCacheTTLSec         int                          `json:"defaultResourcesCacheTTLSec,omitempty"`
	Privileges                          Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       RetrySettings                `json:"retrySettings,omitempty"`
}

// Will resume a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly resumed task
func (c *Client) ResumeTask(uuid string, payload *ResumeTaskPayload) (UUIDResponse, error) {
	var response UUIDResponse

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return response, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/resume", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not resume task due to the following error : %v", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}

	return response, nil
}

// Struct representing the payload the `ResumeTask` method
// When using it, the field that been filled will be used to update parameters of the task you're cloning from
// A new task will be created with all the parameters from the old task you're cloning, and the updated settings you've set here
type CloneTaskPayload struct {
	Name                                string                       `json:"name,omitempty"`
	ResourceBuckets                     []string                     `json:"resourceBuckets,omitempty"`
	AdvancedResourceBuckets             []TaskAdvancedResourceBucket `json:"advancedResourceBuckets,omitempty"`
	Shortname                           string                       `json:"shortname,omitempty"`
	Profile                             string                       `json:"profile,omitempty"`
	JobUuid                             string                       `json:"jobUuid,omitempty"`
	ResultBucket                        string                       `json:"resultBucket,omitempty"`
	Constants                           []Constant                   `json:"constants,omitempty"`
	ForcedConstants                     []ForcedConstant             `json:"forcedConstants,omitempty"`
	Constraints                         []map[string]string          `json:"constraints,omitempty"`
	HardwareConstraints                 []HardwareConstraint         `json:"hardwareConstraints,omitempty"`
	SecretsAccessRights                 SecretsAccessRights          `json:"secretsAccessRights,omitempty"`
	Tags                                []string                     `json:"tags,omitempty"`
	SnapshotWhitelist                   string                       `json:"snapshotWhitelist,omitempty"`
	SnapshotBlacklist                   string                       `json:"snapshotBlacklist,omitempty"`
	SnapshotBucket                      string                       `json:"snapshotBucket,omitempty"`
	SnapshotBucketPrefix                string                       `json:"snapshotBucketPrefix,omitempty"`
	ResultsWhitelist                    string                       `json:"resultsWhitelist,omitempty"`
	ResultsBlacklist                    string                       `json:"resultsBlacklist,omitempty"`
	ResultsBucket                       string                       `json:"resultsBucket,omitempty"`
	ResultsBucketPrefix                 string                       `json:"resultsBucketPrefix,omitempty"`
	Priority                            int                          `json:"priority,omitempty"`
	Dependencies                        Dependencies                 `json:"dependencies,omitempty"`
	AutoDeleteOnCompletion              bool                         `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive                string                       `json:"completionTimeToLive,omitempty"`
	WaitForPoolResourcesSynchronization bool                         `json:"waitForPoolResourcesSynchronization,omitempty"`
	UploadResultsOnCancellation         bool                         `json:"uploadResultsOnCancellation,omitempty"`
	Labels                              []map[string]string          `json:"labels,omitempty"`
	SchedulingType                      SchedulingType               `json:"schedulingType,omitempty"`
	TargetedReservedMachineKey          string                       `json:"targetedReservedMachineKey,omitempty"`
	DefaultResourcesCacheTTLSec         int                          `json:"defaultResourcesCacheTTLSec,omitempty"`
	Privileges                          Privileges                   `json:"privileges,omitempty"`
	RetrySettings                       RetrySettings                `json:"retrySettings,omitempty"`
}

// Will clone a task using the UUID as string, and a `CreateTaskPayload` struct as arguments
// Return a `UUIDResponse` containing the UUID of the newly cloned task
func (c *Client) CloneTask(uuid string, payload *CloneTaskPayload) (UUIDResponse, error) {
	var response UUIDResponse

	payloadJson, err := json.Marshal(&payload)
	if err != nil {
		return response, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, fmt.Sprintf("tasks/%v/clone", uuid))
	if err != nil {
		return UUIDResponse{}, fmt.Errorf("could not clone task due to the following error : %v", err)
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}
	return response, nil
}

// A struct representing the payload for the `UpdateTask` method
type UpdateTaskPayload struct {
	Constants []Constant `json:"constants,omitempty"`
	Tags      []string   `json:"tags,omitempty"`
}

// Will update the fields of a task using the UUID as an argument, as well as a `UpdateTaskPayload` struct
func (c *Client) UpdateTask(uuid string, payload UpdateTaskPayload) error {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return helpers.FormatJsonMarshalError(err)
	}

	_, _, err = c.sendRequest("PUT", payloadJson, nil, fmt.Sprintf("tasks/%v", uuid))
	if err != nil {
		return fmt.Errorf("could not update task due to the following error : %v", err)
	}

	return nil
}
