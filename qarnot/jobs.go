package qarnot

import (
	"encoding/json"
	"fmt"
)

type Job struct {
	Uuid                        string   `json:"uuid"`
	Name                        string   `json:"name"`
	Shortname                   string   `json:"shortname"`
	PoolUuid                    string   `json:"poolUuid"`
	State                       string   `json:"state"`
	PreviousState               string   `json:"previousState"`
	UseDependencies             bool     `json:"useDependencies"`
	StateTransitionTime         string   `json:"stateTransitionTime"`
	PreviousStateTransitionTime string   `json:"previousStateTransitionTime"`
	CreationDate                string   `json:"creationDate"`
	LastModified                string   `json:"lastModified"`
	MaxWallTime                 string   `json:"maxWallTime"`
	Tags                        []string `json:"tags"`
	AutoDeleteOnCompletion      bool     `json:"autoDeleteOnCompletion"`
	CompletionTimeToLive        string   `json:"completionTimeToLive"`
}

func (c *Client) ListJobs() []Job {
	resp, _ := c.sendRequest("GET", []byte{}, nil, "jobs")

	var jobs []Job
	json.Unmarshal(resp, &jobs)

	return jobs
}

type CreateJobPayload struct {
	Name                   string   `json:"name"`
	ShortName              string   `json:"shortName,omitempty"`
	PoolUuid               string   `json:"poolUuid,omitempty"`
	UseDependencies        bool     `json:"useDependencies,omitempty"`
	Tags                   []string `json:"tags,omitempty"`
	MaxWallTime            string   `json:"maxWallTime,omitempty"`
	AutoDeleteOnCompletion bool     `json:"autoDeleteOnCompletion,omitempty"`
	CompletionTimeToLive   string   `json:"completionTimeToLive,omitempty"`
}

type CreateJobResponse struct {
	Uuid string `json:"uuid"`
}

func (c *Client) CreateJob(payload CreateJobPayload) CreateJobResponse {
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	data, _ := c.sendRequest("POST", payloadJson, nil, "jobs")

	var response CreateJobResponse
	err = json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	return response
}

func (c *Client) DeleteJob(uuid string, force bool) {
	var endpoint string
	if force {
		endpoint = fmt.Sprintf("jobs/%v?force=true", uuid)
	} else {
		endpoint = fmt.Sprintf("jobs/%v", uuid)
	}
	c.sendRequest("DELETE", []byte{}, nil, endpoint)
}

func (c *Client) TerminateJob(uuid string) {
	c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("jobs/%v/terminate", uuid))
}

func (c *Client) GetJobInfo(uuid string) Job {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("jobs/%v", uuid))

	var job Job
	err := json.Unmarshal(data, &job)
	if err != nil {
		panic(err)
	}

	return job
}

func (c *Client) ListJobTasks(uuid string) []Task {
	data, _ := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("jobs/%v/tasks", uuid))

	var tasks []Task
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		panic(err)
	}

	return tasks
}
