package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StructureTemplateDataRequest is sent by the client to request data of a structure.
type StructureTemplateDataRequest struct {
	StructureName string
	Position      protocol.BlockPos
	Settings      legacyprotocol.StructureSettings
	RequestType   byte
}

// ID ...
func (*StructureTemplateDataRequest) ID() uint32 {
	return IDStructureTemplateDataRequest
}

// Marshal ...
func (pk *StructureTemplateDataRequest) Marshal(io protocol.IO) {
	io.String(&pk.StructureName)
	io.UBlockPos(&pk.Position)
	protocol.Single(io, &pk.Settings)
	io.Uint8(&pk.RequestType)
}
