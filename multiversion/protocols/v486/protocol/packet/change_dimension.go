package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	DimensionOverworld = iota
	DimensionNether
	DimensionEnd
)

// ChangeDimension is sent by the server to the client to send a dimension change screen client-side. Once the
// screen is cleared client-side, the client will send a PlayerAction packet with
// PlayerActionDimensionChangeDone.
type ChangeDimension struct {
	Dimension       int32
	Position        mgl32.Vec3
	Respawn         bool
	LoadingScreenID protocol.Optional[uint32]
}

// ID ...
func (*ChangeDimension) ID() uint32 {
	return IDChangeDimension
}

// Marshal ...
func (pk *ChangeDimension) Marshal(io protocol.IO) {
	io.Varint32(&pk.Dimension)
	io.Vec3(&pk.Position)
	io.Bool(&pk.Respawn)
	protocol.OptionalFunc(io, &pk.LoadingScreenID, io.Uint32)
}
