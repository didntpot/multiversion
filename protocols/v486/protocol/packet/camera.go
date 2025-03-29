package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Camera is sent by the server to use an Education Edition camera on a player. It produces an image
// client-side.
type Camera struct {
	CameraEntityUniqueID int64
	TargetPlayerUniqueID int64
}

// ID ...
func (*Camera) ID() uint32 {
	return IDCamera
}

// Marshal ...
func (pk *Camera) Marshal(io protocol.IO) {
	io.Varint64(&pk.CameraEntityUniqueID)
	io.Varint64(&pk.TargetPlayerUniqueID)
}
