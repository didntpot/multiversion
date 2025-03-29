package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	SimpleEventCommandsEnabled = iota + 1
	SimpleEventCommandsDisabled
	SimpleEventUnlockWorldTemplateSettings
)

// SimpleEvent is sent by the server to send a 'simple event' to the client, meaning an event without any
// additional event data. The event is typically used by the client for telemetry.
type SimpleEvent struct {
	EventType int16
}

// ID ...
func (*SimpleEvent) ID() uint32 {
	return IDSimpleEvent
}

// Marshal ...
func (pk *SimpleEvent) Marshal(io protocol.IO) {
	io.Int16(&pk.EventType)
}
