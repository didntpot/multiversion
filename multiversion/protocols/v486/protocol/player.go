package protocol

import "github.com/sandertv/gophertunnel/minecraft/protocol"

const (
	PlayerMovementModeClient = iota
	PlayerMovementModeServer
	PlayerMovementModeServerWithRewind
)

// PlayerMovementSettings represents the different server authoritative movement settings. These control how
// the client will provide input to the server.
type PlayerMovementSettings struct {
	// MovementType specifies the way the server handles player movement. Available options are
	// protocol.PlayerMovementModeClient, protocol.PlayerMovementModeServer and
	// protocol.PlayerMovementModeServerWithRewind, where the server authoritative types result
	// in the client sending PlayerAuthInput packets instead of MovePlayer packets and the rewind mode
	// requires sending the tick of movement and several actions.
	MovementType int32
	// RewindHistorySize is the amount of history to keep at maximum.
	RewindHistorySize int32
	// ServerAuthoritativeBlockBreaking specifies if block breaking should be sent through
	// packet.PlayerAuthInput or not.
	ServerAuthoritativeBlockBreaking bool
}

// PlayerMoveSettings reads/writes PlayerMovementSettings x to/from IO r.
func PlayerMoveSettings(r protocol.IO, x *PlayerMovementSettings) {
	r.Varint32(&x.MovementType)
	r.Varint32(&x.RewindHistorySize)
	r.Bool(&x.ServerAuthoritativeBlockBreaking)
}
