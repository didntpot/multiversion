package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AvailableCommands is sent by the server to send a list of all commands that
// the player is able to use on the server. This packet holds all the arguments
// of each commands as well, making it possible for the client to provide
// auto-completion and command usages. AvailableCommands packets can be resent,
// but the packet is often very big, so doing this very often should be avoided.
type AvailableCommands struct {
	EnumValues   []string
	Suffixes     []string
	Enums        []protocol.CommandEnum
	Commands     []legacyprotocol.Command
	DynamicEnums []protocol.DynamicEnum
	Constraints  []protocol.CommandEnumConstraint
}

// ID ...
func (*AvailableCommands) ID() uint32 {
	return IDAvailableCommands
}

// Marshal ...
func (pk *AvailableCommands) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &pk.EnumValues, io.String)
	protocol.FuncSlice(io, &pk.Suffixes, io.String)
	protocol.FuncIOSlice(io, &pk.Enums, protocol.CommandEnumContext{EnumValues: pk.EnumValues}.Marshal)
	protocol.Slice(io, &pk.Commands)
	protocol.Slice(io, &pk.DynamicEnums)
	protocol.Slice(io, &pk.Constraints)
}
