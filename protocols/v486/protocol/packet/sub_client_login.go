package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// SubClientLogin is sent when a sub-client joins the server while another client is already connected to it.
// The packet is sent as a result of split-screen game play, and allows up to four players to play using the
// same network connection. After an initial Login packet from the 'main' client, each sub-client that
// connects sends a SubClientLogin to request their own login.
type SubClientLogin struct {
	ConnectionRequest []byte
}

// ID ...
func (*SubClientLogin) ID() uint32 {
	return IDSubClientLogin
}

// Marshal ...
func (pk *SubClientLogin) Marshal(io protocol.IO) {
	io.ByteSlice(&pk.ConnectionRequest)
}
