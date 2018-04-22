package tHelper

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

// ReadFile provide file path and get file content back
// fail test on any error during process with additional description
// of what was tested.
func ReadFile(t *testing.T, filepath, description string) []byte {
	payload, err := ioutil.ReadFile(filepath)
	if err != nil {
		t.Fatalf("Description: %s\nfailed to read file: %s, with err: %s", description, filepath, err)
	}

	return payload
}

// UnmarshalObject try to unmarshal given slice of bytes into provided object.
// fail test on any error during process with additional description
// of what was tested.
func UnmarshalObject(t *testing.T, data []byte, object interface{}, description string) {
	err := json.Unmarshal(data, object)
	if err != nil {
		t.Fatalf("Description: %s\nfailed to unmarshal with err: %s", description, err)
	}
}

// ReadFileAndUnmarshal read file at given filepath and unmarshal into object
// fail test on any error during process with additional description
// of what was tested.
func ReadFileAndUnmarshal(t *testing.T, filepath string, object interface{}, description string) {
	data := ReadFile(t, filepath, description)

	UnmarshalObject(t, data, object, description)
}
