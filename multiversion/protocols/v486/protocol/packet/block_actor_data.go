package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// BlockActorData is sent by the server to update data of a block entity client-side, for example the data of
// a chest.
type BlockActorData struct {
	Position protocol.BlockPos
	NBTData  map[string]any
}

// ID ...
func (*BlockActorData) ID() uint32 {
	return IDBlockActorData
}

// Marshal ...
func (pk *BlockActorData) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.NBT(&pk.NBTData, nbt.NetworkLittleEndian)
}
