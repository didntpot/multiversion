package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	BossEventShow = iota
	BossEventRegisterPlayer
	BossEventHide
	BossEventUnregisterPlayer
	BossEventHealthPercentage
	BossEventTitle
	BossEventAppearanceProperties
	BossEventTexture
	BossEventRequest
)

const (
	BossEventColourGrey = iota
	BossEventColourBlue
	BossEventColourRed
	BossEventColourGreen
	BossEventColourYellow
	BossEventColourPurple
	BossEventColourWhite
)

// BossEvent is sent by the server to make a specific 'boss event' occur in the world. It includes features
// such as showing a boss bar to the player and turning the sky dark.
type BossEvent struct {
	BossEntityUniqueID int64
	EventType          uint32
	PlayerUniqueID     int64
	BossBarTitle       string
	HealthPercentage   float32
	ScreenDarkening    int16
	Colour             uint32
	Overlay            uint32
}

// ID ...
func (*BossEvent) ID() uint32 {
	return IDBossEvent
}

// Marshal ...
func (pk *BossEvent) Marshal(io protocol.IO) {
	io.Varint64(&pk.BossEntityUniqueID)
	io.Varuint32(&pk.EventType)
	switch pk.EventType {
	case BossEventShow:
		io.String(&pk.BossBarTitle)
		io.Float32(&pk.HealthPercentage)
		io.Int16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	case BossEventRegisterPlayer, BossEventUnregisterPlayer, BossEventRequest:
		io.Varint64(&pk.PlayerUniqueID)
	case BossEventHide:
		// No extra payload for this boss event type.
	case BossEventHealthPercentage:
		io.Float32(&pk.HealthPercentage)
	case BossEventTitle:
		io.String(&pk.BossBarTitle)
	case BossEventAppearanceProperties:
		io.Int16(&pk.ScreenDarkening)
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	case BossEventTexture:
		io.Varuint32(&pk.Colour)
		io.Varuint32(&pk.Overlay)
	default:
		io.UnknownEnumOption(pk.EventType, "boss event type")
	}
}
