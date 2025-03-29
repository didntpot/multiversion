package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// GameRulesChanged is sent by the server to the client to update client-side game rules, such as game rules
// like the 'showCoordinates' game rule.
type GameRulesChanged struct {
	GameRules []protocol.GameRule
}

// ID ...
func (*GameRulesChanged) ID() uint32 {
	return IDGameRulesChanged
}

// Marshal ...
func (pk *GameRulesChanged) Marshal(io protocol.IO) {
	protocol.FuncSlice(io, &pk.GameRules, io.GameRule)
}
