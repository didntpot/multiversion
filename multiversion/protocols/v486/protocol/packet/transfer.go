package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Transfer is sent by the server to transfer a player from the current server to another. Doing so will
// fully disconnect the client, bring it back to the main menu and make it connect to the next server.
type Transfer struct {
	Address string
	Port    uint16
}

// ID ...
func (*Transfer) ID() uint32 {
	return IDTransfer
}

// Marshal ...
func (pk *Transfer) Marshal(io protocol.IO) {
	io.String(&pk.Address)
	io.Uint16(&pk.Port)
}
