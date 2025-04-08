package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// LevelEventGeneric is sent by the server to send a 'generic' level event to the client. This packet sends an
// NBT serialised object and may for that reason be used for any event holding additional data.
type LevelEventGeneric struct {
	EventID             int32
	SerialisedEventData []byte
}

// ID ...
func (pk *LevelEventGeneric) ID() uint32 {
	return IDLevelEventGeneric
}

// Marshal ...
func (pk *LevelEventGeneric) Marshal(io protocol.IO) {
	io.Varint32(&pk.EventID)
	io.Bytes(&pk.SerialisedEventData)
}
