package packet

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// AddPlayer is sent by the server to the client to make a player entity show up client-side. It is one of the
// few entities that cannot be sent using the AddActor packet.
type AddPlayer struct {
	UUID              uuid.UUID
	Username          string
	EntityUniqueID    int64
	EntityRuntimeID   uint64
	PlatformChatID    string
	Position          mgl32.Vec3
	Velocity          mgl32.Vec3
	Pitch             float32
	Yaw               float32
	HeadYaw           float32
	HeldItem          protocol.ItemInstance
	EntityMetadata    map[uint32]any
	AdventureSettings packet.AdventureSettings
	EntityLinks       []protocol.EntityLink
	DeviceID          string
	BuildPlatform     int32
}

// ID ...
func (*AddPlayer) ID() uint32 {
	return IDAddPlayer
}

// Marshal ...
func (pk *AddPlayer) Marshal(io protocol.IO) {
	io.UUID(&pk.UUID)
	io.String(&pk.Username)
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.String(&pk.PlatformChatID)
	io.Vec3(&pk.Position)
	io.Vec3(&pk.Velocity)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Float32(&pk.HeadYaw)
	io.ItemInstance(&pk.HeldItem)
	io.EntityMetadata(&pk.EntityMetadata)
	protocol.Single(io, &pk.AdventureSettings)
	protocol.Slice(io, &pk.EntityLinks)
	io.String(&pk.DeviceID)
	io.Int32(&pk.BuildPlatform)
}
