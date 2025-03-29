package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StructureBlockUpdate is sent by the client when it updates a structure block using the in-game UI. The
// data it contains depends on the type of structure block that it is. In Minecraft Bedrock Edition v1.11,
// there is only the Export structure block type, but in v1.13 the ones present in Java Edition will,
// according to the wiki, be added too.
type StructureBlockUpdate struct {
	Position           protocol.BlockPos
	StructureName      string
	DataField          string
	IncludePlayers     bool
	ShowBoundingBox    bool
	StructureBlockType int32
	Settings           legacyprotocol.StructureSettings
	RedstoneSaveMode   int32
	ShouldTrigger      bool
	Waterlogged        bool
}

// ID ...
func (*StructureBlockUpdate) ID() uint32 {
	return IDStructureBlockUpdate
}

// Marshal ...
func (pk *StructureBlockUpdate) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.String(&pk.StructureName)
	io.String(&pk.DataField)
	io.Bool(&pk.IncludePlayers)
	io.Bool(&pk.ShowBoundingBox)
	io.Varint32(&pk.StructureBlockType)
	protocol.Single(io, &pk.Settings)
	io.Varint32(&pk.RedstoneSaveMode)
	io.Bool(&pk.ShouldTrigger)
	io.Bool(&pk.Waterlogged)
}
