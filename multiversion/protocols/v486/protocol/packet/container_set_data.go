package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ContainerDataFurnaceTickCount = iota
	ContainerDataFurnaceLitTime
	ContainerDataFurnaceLitDuration
	_
	ContainerDataFurnaceFuelAux
)

const (
	ContainerDataBrewingStandBrewTime = iota
	ContainerDataBrewingStandFuelAmount
	ContainerDataBrewingStandFuelTotal
)

// ContainerSetData is sent by the server to update specific data of a single container, meaning a block such
// as a furnace or a brewing stand. This data is usually used by the client to display certain features
// client-side.
type ContainerSetData struct {
	WindowID byte
	Key      int32
	Value    int32
}

// ID ...
func (*ContainerSetData) ID() uint32 {
	return IDContainerSetData
}

// Marshal ...
func (pk *ContainerSetData) Marshal(io protocol.IO) {
	io.Uint8(&pk.WindowID)
	io.Varint32(&pk.Key)
	io.Varint32(&pk.Value)
}
