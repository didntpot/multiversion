package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	CommandBlockImpulse = iota
	CommandBlockRepeating
	CommandBlockChain
)

// CommandBlockUpdate is sent by the client to update a command block at a specific position. The command
// block may be either a physical block or an entity.
type CommandBlockUpdate struct {
	Block                   bool
	Position                protocol.BlockPos
	Mode                    uint32
	NeedsRedstone           bool
	Conditional             bool
	MinecartEntityRuntimeID uint64
	Command                 string
	LastOutput              string
	Name                    string
	ShouldTrackOutput       bool
	TickDelay               int32
	ExecuteOnFirstTick      bool
}

// ID ...
func (*CommandBlockUpdate) ID() uint32 {
	return IDCommandBlockUpdate
}

// Marshal ...
func (pk *CommandBlockUpdate) Marshal(io protocol.IO) {
	io.Bool(&pk.Block)
	if pk.Block {
		io.UBlockPos(&pk.Position)
		io.Varuint32(&pk.Mode)
		io.Bool(&pk.NeedsRedstone)
		io.Bool(&pk.Conditional)
	} else {
		io.Varuint64(&pk.MinecartEntityRuntimeID)
	}
	io.String(&pk.Command)
	io.String(&pk.LastOutput)
	io.String(&pk.Name)
	io.Bool(&pk.ShouldTrackOutput)
	io.Int32(&pk.TickDelay)
	io.Bool(&pk.ExecuteOnFirstTick)
}
