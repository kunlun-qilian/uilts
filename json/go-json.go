//go:build !amd64

package json

import (
	"github.com/goccy/go-json"
)

var (
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	Unmarshal     = json.Unmarshal
	Valid         = json.Valid
)
