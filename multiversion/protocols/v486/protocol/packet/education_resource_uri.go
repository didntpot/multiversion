package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// EducationResourceURI is a packet that transmits education resource settings to all clients.
type EducationResourceURI struct {
	Resource legacyprotocol.EducationSharedResourceURI
}

// ID ...
func (*EducationResourceURI) ID() uint32 {
	return IDEducationResourceURI
}

// Marshal ...
func (pk *EducationResourceURI) Marshal(w *protocol.Writer) {
	legacyprotocol.EducationResourceURI(w, &pk.Resource)
}
