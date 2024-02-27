package qarnot

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestListHardwareConstraints(t *testing.T) {
	expectedOk := `{
		"data": [
		  {
			"discriminator": "MinimumRamHardwareConstraint",
			"minimumMemoryMB": 32000.0
		  },
		  {
			"discriminator": "MinimumRamHardwareConstraint",
			"minimumMemoryMB": 128000.0
		  },
		  {
			"discriminator": "MinimumCoreHardwareConstraint",
			"coreCount": 8
		  },
		  {
			"discriminator": "MinimumCoreHardwareConstraint",
			"coreCount": 16
		  },
		  {
			"discriminator": "MinimumCoreHardwareConstraint",
			"coreCount": 32
		  },
		  {
			"discriminator": "SSDHardwareConstraint"
		  }
		],
		"offset": 0,
		"limit": 6,
		"total": 6
	  }`

	srv := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/v1/hardware-constraints" && r.Method == "GET" {
				fmt.Fprint(w, expectedOk)
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

	hardwareConstraints, err := client.ListHardwareConstraints()
	if err != nil {
		t.Errorf("err should be equal to nil: %v", err)
	}

	expectedHardwareConstraintsResponse := HardwareConstraintsResponse{
		Data: []HardwareConstraint{
			{
				Discriminator:   Discriminator(MinimumRamHardware),
				MinimumMemoryMB: 32000.0,
			},
			{
				Discriminator:   Discriminator(MinimumRamHardware),
				MinimumMemoryMB: 128000.0,
			},
			{
				Discriminator: Discriminator(MinimumCoreHardware),
				CoreCount:     8,
			},
			{
				Discriminator: Discriminator(MinimumCoreHardware),
				CoreCount:     16,
			},
			{
				Discriminator: Discriminator(MinimumCoreHardware),
				CoreCount:     32,
			},
			{
				Discriminator: Discriminator(SSDHardware),
			},
		},
		Offset: 0,
		Limit:  6,
		Total:  6,
	}

	if !reflect.DeepEqual(hardwareConstraints, expectedHardwareConstraintsResponse) {
		t.Error("different values.")
		t.Errorf("expected : %v", expectedHardwareConstraintsResponse)
		t.Errorf("found    : %v", hardwareConstraints)
	}
}
