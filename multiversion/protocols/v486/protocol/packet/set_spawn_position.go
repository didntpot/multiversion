package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	SpawnTypePlayer = iota
	SpawnTypeWorld
)

// SetSpawnPosition is sent by the server to update the spawn position of a player, for example when sleeping
// in a bed.
type SetSpawnPosition struct {
	SpawnType     int32
	Position      protocol.BlockPos
	Dimension     int32
	SpawnPosition protocol.BlockPos
}

// ID ...
func (*SetSpawnPosition) ID() uint32 {
	return IDSetSpawnPosition
}

// Marshal ...
func (pk *SetSpawnPosition) Marshal(io protocol.IO) {
	io.Varint32(&pk.SpawnType)
	io.UBlockPos(&pk.Position)
	io.Varint32(&pk.Dimension)
	io.UBlockPos(&pk.SpawnPosition)
}
