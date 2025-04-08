package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	StructureTemplateResponseExport = iota + 1
	StructureTemplateResponseQuery
)

// StructureTemplateDataResponse is sent by the server to send data of a structure to the client in response
// to a StructureTemplateDataRequest packet.
type StructureTemplateDataResponse struct {
	StructureName     string
	Success           bool
	ResponseType      byte
	StructureTemplate map[string]any
}

// ID ...
func (pk *StructureTemplateDataResponse) ID() uint32 {
	return IDStructureTemplateDataResponse
}

// Marshal ...
func (pk *StructureTemplateDataResponse) Marshal(io protocol.IO) {
	io.String(&pk.StructureName)
	io.Bool(&pk.Success)
	if pk.Success {
		io.NBT(&pk.StructureTemplate, nbt.NetworkLittleEndian)
	}
	io.Uint8(&pk.ResponseType)
}
