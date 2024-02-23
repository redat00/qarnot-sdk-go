package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
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

func (c *Client) ListJobs() ([]Job, error) {
	resp, _, err := c.sendRequest("GET", []byte{}, nil, "jobs")
	if err != nil {
		return []Job{}, err
	}

	var jobs []Job
	err = json.Unmarshal(resp, &jobs)
	if err != nil {
		return nil, helpers.FormatJsonUnmarshalError(err)
	}

	return jobs, nil
}

func (c *Client) CreateJob(payload CreateJobPayload) (CreateJobResponse, error) {
	var response CreateJobResponse

	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return response, helpers.FormatJsonMarshalError(err)
	}

	data, _, err := c.sendRequest("POST", payloadJson, nil, "jobs")
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return response, helpers.FormatJsonUnmarshalError(err)
	}

	return response, nil
}

func (c *Client) DeleteJob(uuid string, force bool) error {
	var endpoint string
	if force {
		endpoint = fmt.Sprintf("jobs/%v?force=true", uuid)
	} else {
		endpoint = fmt.Sprintf("jobs/%v", uuid)
	}

	_, _, err := c.sendRequest("DELETE", []byte{}, nil, endpoint)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) TerminateJob(uuid string) error {
	_, _, err := c.sendRequest("POST", []byte{}, nil, fmt.Sprintf("jobs/%v/terminate", uuid))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetJobInfo(uuid string) (Job, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("jobs/%v", uuid))
	if err != nil {
		return Job{}, err
	}

	var job Job
	err = json.Unmarshal(data, &job)
	if err != nil {
		return Job{}, helpers.FormatJsonUnmarshalError(err)
	}

	return job, nil
}

func (c *Client) ListJobTasks(uuid string) ([]Task, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, fmt.Sprintf("jobs/%v/tasks", uuid))
	if err != nil {
		return []Task{}, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, helpers.FormatJsonUnmarshalError(err)
	}

	return tasks, nil
}
