package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// FilterText is sent by the both the client and the server. The client sends the packet to the server to
// allow the server to filter the text server-side. The server then responds with the same packet and the
// safer version of the text.
type FilterText struct {
	Text       string
	FromServer bool
}

// ID ...
func (*FilterText) ID() uint32 {
	return IDFilterText
}

// Marshal ...
func (pk *FilterText) Marshal(io protocol.IO) {
	io.String(&pk.Text)
	io.Bool(&pk.FromServer)
}
