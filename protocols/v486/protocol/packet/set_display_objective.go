package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ScoreboardSortOrderAscending = iota
	ScoreboardSortOrderDescending
)

// noinspection SpellCheckingInspection
const (
	ScoreboardSlotList      = "list"
	ScoreboardSlotSidebar   = "sidebar"
	ScoreboardSlotBelowName = "belowname"
)

// SetDisplayObjective is sent by the server to display an object as a scoreboard to the player. Once sent,
// it should be followed up by a SetScore packet to set the lines of the packet.
type SetDisplayObjective struct {
	DisplaySlot   string
	ObjectiveName string
	DisplayName   string
	CriteriaName  string
	SortOrder     int32
}

// ID ...
func (*SetDisplayObjective) ID() uint32 {
	return IDSetDisplayObjective
}

// Marshal ...
func (pk *SetDisplayObjective) Marshal(io protocol.IO) {
	io.String(&pk.DisplaySlot)
	io.String(&pk.ObjectiveName)
	io.String(&pk.DisplayName)
	io.String(&pk.CriteriaName)
	io.Varint32(&pk.SortOrder)
}
