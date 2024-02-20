package qarnot

import (
	"encoding/json"
	"fmt"

	"github.com/redat00/qarnot-sdk-go/internal/helpers"
)

type Discriminator string

const (
	MinimumCoreHardware         Discriminator = "MinimumCoreHardwareConstraint"
	MaximumCoreHardawre         Discriminator = "MaximumCoreHardwareConstraint"
	MinimumRamHardware          Discriminator = "MinimumRamHardwareConstraint"
	MaximumRamHardware          Discriminator = "MaximumRamHardwareConstraint"
	SpecificHardware            Discriminator = "SpecificHardwareConstraint"
	MinimumRamCoreRatioHardware Discriminator = "MinimumRamCoreRatioHardwareConstraint"
	MaximumRamCoreRatioHardware Discriminator = "MaximumRamCoreRatioHardwareConstraint"
	SSDHardware                 Discriminator = "SSDHardwareConstraint"
	NoSSDHardware               Discriminator = "NoSSDHardwareConstraint"
	NoGpuHardware               Discriminator = "NoGpuHardwareConstraint"
	GpuHardware                 Discriminator = "GpuHardwareConstraint"
	CpuModelHardware            Discriminator = "CpuModelHardwareConstraint"
)

type HardwareConstraint struct {
	Discriminator            Discriminator `json:"discriminator"`
	CoreCount                int           `json:"coreCount"`
	MinimumMemoryMB          float64       `json:"minimumMemoryMB"`
	MaximumMemoryMB          float64       `json:"maximumMemoryMB"`
	SpecificationKey         string        `json:"specificationKey"`
	MinimumMemoryGBCoreRatio float64       `json:"minimumMemoryGBCoreRatio"`
	MaximumMemoryGBCoreRatio float64       `json:"maximumMemoryGBCoreRatio"`
	CpuModel                 string        `json:"cpuModel"`
}

type HardwareConstraintsResponse struct {
	Data   []HardwareConstraint `json:"data"`
	Offset int                  `json:"offset"`
	Limit  int                  `json:"limit"`
	Total  int                  `json:"total"`
}

func (c *Client) ListHardwareConstraints() (HardwareConstraintsResponse, error) {
	data, _, err := c.sendRequest("GET", []byte{}, nil, "hardware-constraints")
	if err != nil {
		return HardwareConstraintsResponse{}, fmt.Errorf("could not retrieve list of hardware constraints due to the following error : %v", err)
	}

	var response HardwareConstraintsResponse
	err = json.Unmarshal(data, &response)
	helpers.JsonUnmarshalCheckError(err)

	return response, nil
}
