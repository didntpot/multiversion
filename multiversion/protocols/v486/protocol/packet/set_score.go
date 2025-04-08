package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ScoreboardActionModify = iota
	ScoreboardActionRemove
)

// SetScore is sent by the server to send the contents of a scoreboard to the player. It may be used to either
// add, remove or edit entries on the scoreboard.
type SetScore struct {
	ActionType byte
	Entries    []protocol.ScoreboardEntry
}

// ID ...
func (*SetScore) ID() uint32 {
	return IDSetScore
}

// Marshal ...
func (pk *SetScore) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	switch pk.ActionType {
	case ScoreboardActionRemove:
		protocol.FuncIOSlice(io, &pk.Entries, protocol.ScoreRemoveEntry)
	case ScoreboardActionModify:
		protocol.Slice(io, &pk.Entries)
	default:
		io.UnknownEnumOption(pk.ActionType, "set score action type")
	}
}
