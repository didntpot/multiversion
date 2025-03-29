package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CodeBuilderCategoryNone = iota
	CodeBuilderCategoryStatus
	CodeBuilderCategoryInstantiation
)

const (
	CodeBuilderOperationNone = iota
	CodeBuilderOperationGet
	CodeBuilderOperationSet
	CodeBuilderOperationReset
)

// CodeBuilderSource is an Education Edition packet sent by the client to the server to run an operation with a
// code builder.
type CodeBuilderSource struct {
	Operation byte
	Category  byte
	Value     []byte
}

// ID ...
func (pk *CodeBuilderSource) ID() uint32 {
	return IDCodeBuilderSource
}

// Marshal ...
func (pk *CodeBuilderSource) Marshal(io protocol.IO) {
	io.Uint8(&pk.Operation)
	io.Uint8(&pk.Category)
	io.ByteSlice(&pk.Value)
}
