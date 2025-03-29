package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// PlayerList is sent by the server to update the client-side player list in the in-game menu screen. It shows
// the icon of each player if the correct XUID is written in the packet.
// Sending the PlayerList packet is obligatory when sending an AddPlayer packet. The added player will not
// show up to a client if it has not been added to the player list, because several properties of the player
// are obtained from the player list, such as the skin.
type PlayerList struct {
	ActionType byte
	Entries    []legacyprotocol.PlayerListEntry
}

// ID ...
func (*PlayerList) ID() uint32 {
	return IDPlayerList
}

// Marshal ...
func (pk *PlayerList) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	switch pk.ActionType {
	case packet.PlayerListActionAdd:
		protocol.Slice(io, &pk.Entries)
	case packet.PlayerListActionRemove:
		protocol.FuncIOSlice(io, &pk.Entries, func(r protocol.IO, x *legacyprotocol.PlayerListEntry) {
			r.UUID(&x.UUID)
		})
	default:
		io.UnknownEnumOption(pk.ActionType, "player list action type")
	}
	if pk.ActionType == packet.PlayerListActionAdd {
		for i := 0; i < len(pk.Entries); i++ {
			io.Bool(&pk.Entries[i].Skin.Trusted)
		}
	}
}
