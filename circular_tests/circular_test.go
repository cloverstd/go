package circular_test

import (
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/require"
)

func init() {
	jsoniter.MaxDepth = 10
}

func Test_map(t *testing.T) {
	m := map[string]interface{}{
		"foo": "foo",
	}
	m["bar"] = m

	should := require.New(t)
	_, err := jsoniter.MarshalToString(m)
	should.IsType(jsoniter.MaxDepthError{}, err)
}

func Test_struct(t *testing.T) {
	type Node struct {
		Next *Node
		Name string
	}
	m := Node{Name: "foo"}
	m.Next = &m

	should := require.New(t)
	_, err := jsoniter.MarshalToString(m)
	should.IsType(jsoniter.MaxDepthError{}, err)
}

func Test_slice(t *testing.T) {
	m := []interface{}{
		"123",
		nil,
	}
	m[1] = m

	should := require.New(t)
	_, err := jsoniter.MarshalToString(m)
	should.IsType(jsoniter.MaxDepthError{}, err)
}

func Test_map_struct(t *testing.T) {
	type Node struct {
		M    interface{}
		Name string
	}
	m := Node{Name: "foo"}
	m.M = map[string]interface{}{
		"m": &m,
	}

	should := require.New(t)
	_, err := jsoniter.MarshalToString(m)
	should.IsType(jsoniter.MaxDepthError{}, err)
}

func Test_depth(t *testing.T) {
	m := map[string]interface{}{
		"foo": "foo",
	}
	m["bar"] = m

	should := require.New(t)
	_, err := jsoniter.MarshalToString(m)
	should.IsType(jsoniter.MaxDepthError{}, err)
	should.Equal("exceeding maximum depth 11", err.Error())
}
