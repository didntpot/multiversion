package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AddBehaviourTree is sent by the server to the client. The packet is currently unused by both client and
// server.
type AddBehaviourTree struct {
	BehaviourTree string
}

// ID ...
func (*AddBehaviourTree) ID() uint32 {
	return IDAddBehaviourTree
}

// Marshal ...
func (pk *AddBehaviourTree) Marshal(io protocol.IO) {
	io.String(&pk.BehaviourTree)
}
