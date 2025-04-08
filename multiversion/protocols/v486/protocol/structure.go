package protocol

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StructureSettings is a struct holding settings of a structure block. Its fields may be changed using the
// in-game UI on the client-side.
type StructureSettings struct {
	protocol.StructureSettings
}

// Marshal ...
func (x *StructureSettings) Marshal(io protocol.IO) {
	io.String(&x.PaletteName)
	io.Bool(&x.IgnoreEntities)
	io.Bool(&x.IgnoreBlocks)
	io.UBlockPos(&x.Size)
	io.UBlockPos(&x.Offset)
	io.Varint64(&x.LastEditingPlayerUniqueID)
	io.Uint8(&x.Rotation)
	io.Uint8(&x.Mirror)
	io.Uint8(&x.AnimationMode)
	io.Float32(&x.AnimationDuration)
	io.Float32(&x.Integrity)
	io.Uint32(&x.Seed)
	io.Vec3(&x.Pivot)
}
