package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	mockData := `{
		"email": "test@example.org",
		"maxBucket": 100,
		"maxTask": 100,
		"taskCount": 2,
		"maxJob": 100,
		"jobCount": 4,
		"maxPool": 50,
		"poolCount": 0,
		"maxRunningTask": 10,
		"maxRunningPool": 10,
		"runningTaskCount": 0,
		"runningPoolCount": 0,
		"runningInstanceCount": 0,
		"runningCoreCount": 0,
		"maxInstances": 128,
		"maxCores": 1024,
		"maxFlexInstances": 128,
		"maxFlexCores": 1024,
		"maxOnDemandInstances": 0,
		"maxOnDemandCores": 0,
		"reservedQuotas": [],
		"quotaBytes": 0,
		"quotaBytesBucket": 10737418240,
		"usedQuotaBytesBucket": 0,
		"usedQuotaBytes": 0,
		"defaultScheduling": "Flex",
		"defaultReservedSpecificationKey": ""
	}`
	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/info" {
				fmt.Fprint(w, mockData)
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

	userInfo, _ := client.GetUserInfo()

	expectedData := UserInfo{
		Email:                           "test@example.org",
		MaxBucket:                       100,
		MaxTask:                         100,
		TaskCount:                       2,
		MaxJob:                          100,
		JobCount:                        4,
		MaxPool:                         50,
		PoolCount:                       0,
		MaxRunningTask:                  10,
		MaxRunningPool:                  10,
		RunningTaskCount:                0,
		RunningPoolCount:                0,
		RunningInstanceCount:            0,
		RunningCoreCount:                0,
		MaxInstances:                    128,
		MaxCores:                        1024,
		MaxFlexInstances:                128,
		MaxFlexCores:                    1024,
		MaxOnDemandInstances:            0,
		MaxOnDemandCores:                0,
		ReservedQuotas:                  []ReservedQuotas{},
		QuotaBytes:                      0,
		QuotaBytesBucket:                10737418240,
		UsedQuotaBytesBucket:            0,
		UsedQuotaBytes:                  0,
		DefaultScheduling:               "Flex",
		DefaultReservedSpecificationKey: "",
	}

	if !reflect.DeepEqual(expectedData, userInfo) {
		t.Errorf("error in values. Expected %+v, found %+v", expectedData, userInfo)
	}
}

func TestGetUserInfoBadToken(t *testing.T) {
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

	_, err = client.GetUserInfo()
	expectedData := "could not get user info due to the following error : [HTTP 401] Bad authentication token"

	if err.Error() != expectedData {
		t.Errorf("different values.")
		t.Errorf("expected : %v", expectedData)
		t.Errorf("found    : %v", err.Error())
	}
}
