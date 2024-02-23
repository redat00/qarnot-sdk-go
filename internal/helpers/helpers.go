package helpers

import "fmt"

func FormatJsonUnmarshalError(err error) error {
	return fmt.Errorf("could not unmarshall json due to the following error: %v", err)
}

func FormatJsonMarshalError(err error) error {
	return fmt.Errorf("could not marshall json due to the following error: %v", err)
}
