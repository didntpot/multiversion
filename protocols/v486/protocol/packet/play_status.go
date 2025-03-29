package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PlayStatusLoginSuccess int32 = iota
	PlayStatusLoginFailedClient
	PlayStatusLoginFailedServer
	PlayStatusPlayerSpawn
	PlayStatusLoginFailedInvalidTenant
	PlayStatusLoginFailedVanillaEdu
	PlayStatusLoginFailedEduVanilla
	PlayStatusLoginFailedServerFull
)

// PlayStatus is sent by the server to update a player on the play status. This includes failed statuses due
// to a mismatched version, but also success statuses.
type PlayStatus struct {
	Status int32
}

// ID ...
func (*PlayStatus) ID() uint32 {
	return IDPlayStatus
}

// Marshal ...
func (pk *PlayStatus) Marshal(io protocol.IO) {
	io.BEInt32(&pk.Status)
}
