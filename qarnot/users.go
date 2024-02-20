package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

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

func (c *Client) GetUserInfo() (UserInfo, error) {
	// Send request and get back data
	data, _, err := c.sendRequest("GET", []byte{}, make(map[string]string), "info")
	if err != nil {
		return UserInfo{}, fmt.Errorf("could not get user info due to the following error : %v", err)
	}

	// Convert data to UserInfo struct
	var userInfo UserInfo
	err = json.Unmarshal(data, &userInfo)
	helpers.JsonUnmarshalCheckError(err)

	// Return UserInfo
	return userInfo, nil
}
