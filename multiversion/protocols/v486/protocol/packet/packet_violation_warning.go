package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ViolationTypeMalformed = iota
)

const (
	ViolationSeverityWarning = iota
	ViolationSeverityFinalWarning
	ViolationSeverityTerminatingConnection
)

// PacketViolationWarning is sent by the client when it receives an invalid packet from the server. It holds
// some information on the error that occurred.
// noinspection GoNameStartsWithPackageName
type PacketViolationWarning struct {
	Type             int32
	Severity         int32
	PacketID         int32
	ViolationContext string
}

// ID ...
func (*PacketViolationWarning) ID() uint32 {
	return IDPacketViolationWarning
}

// Marshal ...
func (pk *PacketViolationWarning) Marshal(io protocol.IO) {
	io.Varint32(&pk.Type)
	io.Varint32(&pk.Severity)
	io.Varint32(&pk.PacketID)
	io.String(&pk.ViolationContext)
}
