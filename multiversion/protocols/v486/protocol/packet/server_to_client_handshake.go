package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// ServerToClientHandshake is sent by the server to the client to complete the key exchange in order to
// initialise encryption on client and server side. It is followed up by a ClientToServerHandshake packet
// from the client.
type ServerToClientHandshake struct {
	JWT []byte
}

// ID ...
func (*ServerToClientHandshake) ID() uint32 {
	return IDServerToClientHandshake
}

// Marshal ...
func (pk *ServerToClientHandshake) Marshal(io protocol.IO) {
	io.ByteSlice(&pk.JWT)
}
