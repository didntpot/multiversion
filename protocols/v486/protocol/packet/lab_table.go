package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	LabTableActionCombine = iota
	LabTableActionReact
	LabTableActionReset
)

// LabTable is sent by the client to let the server know it started a chemical reaction in Education Edition,
// and is sent by the server to other clients to show the effects.
// The packet is only functional if Education features are enabled.
type LabTable struct {
	ActionType   byte
	Position     protocol.BlockPos
	ReactionType byte
}

// ID ...
func (*LabTable) ID() uint32 {
	return IDLabTable
}

// Marshal ...
func (pk *LabTable) Marshal(io protocol.IO) {
	io.Uint8(&pk.ActionType)
	io.BlockPos(&pk.Position)
	io.Uint8(&pk.ReactionType)
}
