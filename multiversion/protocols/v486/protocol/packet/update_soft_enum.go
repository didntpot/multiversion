package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	SoftEnumActionAdd = iota
	SoftEnumActionRemove
	SoftEnumActionSet
)

// UpdateSoftEnum is sent by the server to update a soft enum, also known as a dynamic enum, previously sent
// in the AvailableCommands packet. It is sent whenever the enum should get new options or when some of its
// options should be removed.
// The UpdateSoftEnum packet will apply for enums that have been set in the AvailableCommands packet with the
// 'Dynamic' field of the CommandEnum set to true.
type UpdateSoftEnum struct {
	EnumType   string
	Options    []string
	ActionType byte
}

// ID ...
func (*UpdateSoftEnum) ID() uint32 {
	return IDUpdateSoftEnum
}

// Marshal ...
func (pk *UpdateSoftEnum) Marshal(io protocol.IO) {
	io.String(&pk.EnumType)
	protocol.FuncSlice(io, &pk.Options, io.String)
	io.Uint8(&pk.ActionType)
}
