package helpers

import "fmt"

func JsonUnmarshalCheckError(err error) {
	if err != nil {
		panic(fmt.Errorf("could not unmarshall json due to the following error: %v", err))
	}
}

func JsonMarshalCheckError(err error) {
	if err != nil {
		panic(fmt.Errorf("could not marshall json due to the following error: %v", err))
	}
}
