//go:build amd64

package json

import (
	"github.com/bytedance/sonic"
)

var (
	Marshal       = sonic.ConfigDefault.Marshal
	MarshalIndent = sonic.ConfigDefault.MarshalIndent
	NewDecoder    = sonic.ConfigDefault.NewDecoder
	NewEncoder    = sonic.ConfigDefault.NewEncoder
	Unmarshal     = sonic.ConfigDefault.Unmarshal
	Valid         = sonic.ConfigDefault.Valid
)
