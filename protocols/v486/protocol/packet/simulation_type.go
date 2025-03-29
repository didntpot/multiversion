package packet

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	SimulationTypeGame byte = iota
	SimulationTypeEditor
	SimulationTypeTest
	SimulationTypeInvalid
)

// SimulationType is an in-progress packet. We currently do not know the use case.
type SimulationType struct {
	SimulationType byte
}

// ID ...
func (*SimulationType) ID() uint32 {
	return IDSimulationType
}

// Marshal ...
func (pk *SimulationType) Marshal(io protocol.IO) {
	io.Uint8(&pk.SimulationType)
}
