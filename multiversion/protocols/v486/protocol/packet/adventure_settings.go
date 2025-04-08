package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	AdventureFlagWorldImmutable = 1 << iota
	AdventureSettingsFlagsNoPvM
	AdventureSettingsFlagsNoMvP
	AdventureSettingsFlagsUnused
	AdventureSettingsFlagsShowNameTags
	AdventureFlagAutoJump
	AdventureFlagAllowFlight
	AdventureFlagNoClip
	AdventureFlagWorldBuilder
	AdventureFlagFlying
	AdventureFlagMuted
)

const (
	CommandPermissionLevelNormal = iota
	CommandPermissionLevelGameDirectors
	CommandPermissionLevelAdmin
	CommandPermissionLevelHost
	CommandPermissionLevelOwner
	CommandPermissionLevelInternal
)

const (
	ActionPermissionMine = 1 << iota
	ActionPermissionDoorsAndSwitches
	ActionPermissionOpenContainers
	ActionPermissionAttackPlayers
	ActionPermissionAttackMobs
	ActionPermissionOperator
	ActionPermissionTeleport
	ActionPermissionBuild
	ActionPermissionDefault
)

const (
	PermissionLevelVisitor = iota
	PermissionLevelMember
	PermissionLevelOperator
	PermissionLevelCustom
)

// AdventureSettings is sent by the server to update game-play related features, in particular permissions to
// access these features for the client. It includes allowing the player to fly, build and mine, and attack
// entities. Most of these flags should be checked server-side instead of using this packet only.
// The client may also send this packet to the server when it updates one of these settings through the
// in-game settings interface. The server should verify if the player actually has permission to update those
// settings.
type AdventureSettings struct {
	Flags                   uint32
	CommandPermissionLevel  uint32
	ActionPermissions       uint32
	PermissionLevel         uint32
	CustomStoredPermissions uint32
	PlayerUniqueID          int64
}

// ID ...
func (*AdventureSettings) ID() uint32 {
	return IDAdventureSettings
}

// Marshal ...
func (pk *AdventureSettings) Marshal(io protocol.IO) {
	io.Varuint32(&pk.Flags)
	io.Varuint32(&pk.CommandPermissionLevel)
	io.Varuint32(&pk.ActionPermissions)
	io.Varuint32(&pk.PermissionLevel)
	io.Varuint32(&pk.CustomStoredPermissions)
	io.Int64(&pk.PlayerUniqueID)
}
