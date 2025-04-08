package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// AutomationClientConnect is used to make the client connect to a websocket server. This websocket server has
// the ability to execute commands on the behalf of the client and it can listen for certain events fired by
// the client.
type AutomationClientConnect struct {
	ServerURI string
}

// ID ...
func (*AutomationClientConnect) ID() uint32 {
	return IDAutomationClientConnect
}

// Marshal ...
func (pk *AutomationClientConnect) Marshal(io protocol.IO) {
	io.String(&pk.ServerURI)
}
