package internal_test

import (
	"fmt"
	"testing"

	"github.com/michimani/evidentlylocal/internal"
	"github.com/michimani/evidentlylocal/types"
	"github.com/stretchr/testify/assert"
)

type badMarshaler struct{}

func (b badMarshaler) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("an error occurred")
}

func Test_GenerateResponseBody(t *testing.T) {

	cases := []struct {
		name    string
		data    any
		wantErr bool
		expect  []byte
	}{
		{
			"data is nil",
			nil,
			true,
			nil,
		},
		{
			"json marshal error",
			badMarshaler{},
			true,
			nil,
		},
		{
			"success",
			types.EvaluateFeatureResponse{
				Details:   "{}",
				Reason:    "test-reason",
				Value:     types.VariableValue{types.VariableValueTypeBool: true},
				Variation: "test-variation",
			},
			false,
			[]byte(`{"details":"{}","reason":"test-reason","value":{"boolValue":true},"variation":"test-variation"}`),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			asst := assert.New(tt)
			got, requestID, err := internal.GenerateResponseBody(c.data)
			if c.wantErr {
				asst.Nil(got, string(got))
				asst.Empty(requestID, requestID)
				asst.Error(err)
				return
			}

			asst.Equal(c.expect, got, string(got))
			asst.Len(requestID, 36, requestID)
		})
	}
}
