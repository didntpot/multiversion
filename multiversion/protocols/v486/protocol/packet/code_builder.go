package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// CodeBuilder is an Education Edition packet sent by the server to the client to open the URL to a Code
// Builder (websocket) server.
type CodeBuilder struct {
	URL                   string
	ShouldOpenCodeBuilder bool
}

// ID ...
func (*CodeBuilder) ID() uint32 {
	return IDCodeBuilder
}

// Marshal ...
func (pk *CodeBuilder) Marshal(io protocol.IO) {
	io.String(&pk.URL)
	io.Bool(&pk.ShouldOpenCodeBuilder)
}
