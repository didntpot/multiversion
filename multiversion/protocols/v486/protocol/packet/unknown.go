package packet

import (
	"fmt"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// Unknown is an implementation of the Packet interface for unknown/unimplemented packets. It holds the packet
// ID and the raw payload. It serves as a way to read raw unknown packets and forward them to another
// connection, without necessarily implementing them.
type Unknown struct {
	PacketID uint32
	Payload  []byte
}

// ID ...
func (pk *Unknown) ID() uint32 {
	return pk.PacketID
}

// Marshal ...
func (pk *Unknown) Marshal(io protocol.IO) {
	io.Bytes(&pk.Payload)
}

// String ...
func (pk *Unknown) String() string {
	return fmt.Sprintf("{ID:0x%x Payload:0x%x}", pk.PacketID, pk.Payload)
}
