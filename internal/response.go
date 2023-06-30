package internal

import (
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
)

const (
	dummyRequestID = "xxxxxxxx-0000-0000-0000-xxxxxxxxxxxx"
)

// GenerateResponseBody generates response body
//
// 1st return value: response body
// 2nd return value: response request ID
// 3rd return value: error
func GenerateResponseBody(data any) ([]byte, string, error) {
	if data == nil {
		return nil, "", errors.New("data is nil")
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, "", err
	}

	requestID := dummyRequestID
	uuid, err := uuid.NewV4()
	if err == nil {
		requestID = uuid.String()
	}

	return bytes, requestID, nil
}
