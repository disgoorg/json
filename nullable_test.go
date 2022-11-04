package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullBool(t *testing.T) {
	type v struct {
		NullBool Nullable[bool] `json:"null_bool"`
	}
	cases := []struct {
		input Nullable[bool]
		data  string
	}{
		{Null[bool](), `{"null_bool":null}`},
		{NewNullable(true), `{"null_bool":true}`},
		{NewNullable(false), `{"null_bool":false}`},
	}

	for _, c := range cases {
		data, err := Marshal(v{
			NullBool: c.input,
		})
		assert.NoError(t, err)
		assert.Equal(t, c.data, string(data))
	}
}

func TestNullBoolPtr(t *testing.T) {
	type v struct {
		NullBool *Nullable[bool] `json:"null_bool,omitempty"`
	}
	cases := []struct {
		input *Nullable[bool]
		data  string
	}{
		{nil, `{}`},
		{NullPtr[bool](), `{"null_bool":null}`},
		{NewNullablePtr(true), `{"null_bool":true}`},
		{NewNullablePtr(false), `{"null_bool":false}`},
	}

	for _, c := range cases {
		data, err := Marshal(v{
			NullBool: c.input,
		})
		assert.NoError(t, err)
		assert.Equal(t, c.data, string(data))
	}
}
