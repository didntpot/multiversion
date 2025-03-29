package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

const (
	InputFlagAscend = 1 << iota
	InputFlagDescend
	InputFlagNorthJump
	InputFlagJumpDown
	InputFlagSprintDown
	InputFlagChangeHeight
	InputFlagJumping
	InputFlagAutoJumpingInWater
	InputFlagSneaking
	InputFlagSneakDown
	InputFlagUp
	InputFlagDown
	InputFlagLeft
	InputFlagRight
	InputFlagUpLeft
	InputFlagUpRight
	InputFlagWantUp
	InputFlagWantDown
	InputFlagWantDownSlow
	InputFlagWantUpSlow
	InputFlagSprinting
	InputFlagAscendBlock
	InputFlagDescendBlock
	InputFlagSneakToggleDown
	InputFlagPersistSneak
	InputFlagStartSprinting
	InputFlagStopSprinting
	InputFlagStartSneaking
	InputFlagStopSneaking
	InputFlagStartSwimming
	InputFlagStopSwimming
	InputFlagStartJumping
	InputFlagStartGliding
	InputFlagStopGliding
	InputFlagPerformItemInteraction
	InputFlagPerformBlockActions
	InputFlagPerformItemStackRequest
	InputFlagHandledTeleport
	InputFlagEmoting
	InputFlagMissedSwing
	InputFlagStartCrawling
	InputFlagStopCrawling
	InputFlagStartFlying
	InputFlagStopFlying
	InputFlagClientAckServerData
	InputFlagClientPredictedVehicle
	InputFlagPaddlingLeft
	InputFlagPaddlingRight
	InputFlagBlockBreakingDelayEnabled
	InputFlagHorizontalCollision
	InputFlagVerticalCollision
	InputFlagDownLeft
	InputFlagDownRight
	InputFlagCameraRelativeMovementEnabled
	InputFlagRotControlledByMoveDirection
	InputFlagStartSpinAttack
	InputFlagStopSpinAttack
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to synchronise
// the player input with the position server-side.
// The client sends this packet when the ServerAuthoritativeMovementMode field in the StartGame packet is set
// to true, instead of the MovePlayer packet. The client will send this packet once every tick.
type PlayerAuthInput struct {
	Pitch, Yaw          float32
	Position            mgl32.Vec3
	MoveVector          mgl32.Vec2
	HeadYaw             float32
	InputData           uint64
	InputMode           uint32
	PlayMode            uint32
	GazeDirection       mgl32.Vec3
	Tick                uint64
	Delta               mgl32.Vec3
	ItemInteractionData protocol.UseItemTransactionData
	ItemStackRequest    legacyprotocol.ItemStackRequest
	BlockActions        []protocol.PlayerBlockAction
}

// ID ...
func (pk *PlayerAuthInput) ID() uint32 {
	return IDPlayerAuthInput
}

// Marshal ...
func (pk *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	io.Varuint64(&pk.InputData)
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)
	if pk.PlayMode == packet.PlayModeReality {
		io.Vec3(&pk.GazeDirection)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)
	if pk.InputData&InputFlagPerformItemInteraction != 0 {
		io.PlayerInventoryAction(&pk.ItemInteractionData)
	}
	if pk.InputData&InputFlagPerformItemStackRequest != 0 {
		protocol.Single(io, &pk.ItemStackRequest)
	}
	if pk.InputData&InputFlagPerformBlockActions != 0 {
		protocol.SliceVarint32Length(io, &pk.BlockActions)
	}
}
