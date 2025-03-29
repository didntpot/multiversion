package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ScriptCustomEvent is sent by both the client and the server. It is a way to let scripts communicate with
// the server, so that the client can let the server know it triggered an event, or the other way around.
// It is essentially an RPC kind of system.
type ScriptCustomEvent struct {
	EventName string
	EventData []byte
}

// ID ...
func (*ScriptCustomEvent) ID() uint32 {
	return IDScriptCustomEvent
}

// Marshal ...
func (pk *ScriptCustomEvent) Marshal(io protocol.IO) {
	io.String(&pk.EventName)
	io.ByteSlice(&pk.EventData)
}
