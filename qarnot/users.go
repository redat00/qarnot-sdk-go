package qarnot

import "encoding/json"

type ReservedQuotas struct {
	MachineKey   string `json:"machineKey"`
	MaxInstances int    `json:"maxInstances"`
	MaxCores     int    `json:"maxCores"`
}

type UserInfo struct {
	Email                           string           `json:"email"`
	MaxBucket                       int              `json:"maxBucket"`
	MaxTask                         int              `json:"maxTask"`
	TaskCount                       int              `json:"taskCount"`
	MaxJob                          int              `json:"maxJob"`
	JobCount                        int              `json:"jobCount"`
	MaxPool                         int              `json:"maxPool"`
	PoolCount                       int              `json:"poolCount"`
	MaxRunningTask                  int              `json:"maxRunningTask"`
	MaxRunningPool                  int              `json:"maxRunningPool"`
	RunningTaskCount                int              `json:"runningTaskCount"`
	RunningPoolCount                int              `json:"runningPoolCount"`
	RunningInstanceCount            int              `json:"runningInstanceCount"`
	RunningCoreCount                int              `json:"runningCoreCount"`
	MaxInstances                    int              `json:"maxInstances"`
	MaxCores                        int              `json:"maxCores"`
	MaxFlexInstances                int              `json:"maxFlexInstances"`
	MaxFlexCores                    int              `json:"maxFlexCores"`
	MaxOnDemandInstances            int              `json:"maxOnDemandInstances"`
	MaxOnDemandCores                int              `json:"maxOnDemandCores"`
	ReservedQuotas                  []ReservedQuotas `json:"reservedQuotas"`
	QuotaBytes                      int              `json:"quotaBytes"`
	UsedQuotaBytes                  int              `json:"usedQuotasBytes"`
	QuotaBytesBucket                int              `json:"quotaBytesBucket"`
	UsedQuotaBytesBucket            int              `json:"usedQuotaBytesBucket"`
	DefaultScheduling               string           `json:"defaultScheduling"`
	DefaultReservedSpecificationKey string           `json:"defaultReservedSpecificationKey"`
}

func (c *Client) GetUserInfo() UserInfo {
	// Send request and get back data
	data, _ := c.sendRequest("GET", []byte{}, make(map[string]string), "info")

	// Convert data to UserInfo struct
	var userInfo UserInfo
	err := json.Unmarshal(data, &userInfo)
	if err != nil {
		panic(err)
	}

	// Return UserInfo
	return userInfo
}
