package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/multiversion/protocols/v486/protocol"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// StartGame is sent by the server to send information about the world the player will be spawned in. It
// contains information about the position the player spawns in, and information about the world in general
// such as its game rules.
type StartGame struct {
	EntityUniqueID                       int64
	EntityRuntimeID                      uint64
	PlayerGameMode                       int32
	PlayerPosition                       mgl32.Vec3
	Pitch                                float32
	Yaw                                  float32
	WorldSeed                            int32
	SpawnBiomeType                       int16
	UserDefinedBiomeName                 string
	Dimension                            int32
	Generator                            int32
	WorldGameMode                        int32
	Difficulty                           int32
	WorldSpawn                           protocol.BlockPos
	AchievementsDisabled                 bool
	DayCycleLockTime                     int32
	EducationEditionOffer                int32
	EducationFeaturesEnabled             bool
	EducationProductID                   string
	RainLevel                            float32
	LightningLevel                       float32
	ConfirmedPlatformLockedContent       bool
	MultiPlayerGame                      bool
	LANBroadcastEnabled                  bool
	XBLBroadcastMode                     int32
	PlatformBroadcastMode                int32
	CommandsEnabled                      bool
	TexturePackRequired                  bool
	GameRules                            []protocol.GameRule
	Experiments                          []protocol.ExperimentData
	ExperimentsPreviouslyToggled         bool
	BonusChestEnabled                    bool
	StartWithMapEnabled                  bool
	PlayerPermissions                    int32
	ServerChunkTickRadius                int32
	HasLockedBehaviourPack               bool
	HasLockedTexturePack                 bool
	FromLockedWorldTemplate              bool
	MSAGamerTagsOnly                     bool
	FromWorldTemplate                    bool
	WorldTemplateSettingsLocked          bool
	OnlySpawnV1Villagers                 bool
	BaseGameVersion                      string
	LimitedWorldWidth, LimitedWorldDepth int32
	NewNether                            bool
	EducationSharedResourceURI           protocol.EducationSharedResourceURI
	ForceExperimentalGameplay            bool
	LevelID                              string
	WorldName                            string
	TemplateContentIdentity              string
	Trial                                bool
	PlayerMovementSettings               protocol.PlayerMovementSettings
	Time                                 int64
	EnchantmentSeed                      int32
	Blocks                               []protocol.BlockEntry
	Items                                []legacyprotocol.ItemEntry
	MultiPlayerCorrelationID             string
	ServerAuthoritativeInventory         bool
	GameVersion                          string
	ServerBlockStateChecksum             uint64
}

// ID ...
func (*StartGame) ID() uint32 {
	return IDStartGame
}

// Marshal ...
func (pk *StartGame) Marshal(io protocol.IO) {
	io.Varint64(&pk.EntityUniqueID)
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.PlayerGameMode)
	io.Vec3(&pk.PlayerPosition)
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Varint32(&pk.WorldSeed)
	io.Int16(&pk.SpawnBiomeType)
	io.String(&pk.UserDefinedBiomeName)
	io.Varint32(&pk.Dimension)
	io.Varint32(&pk.Generator)
	io.Varint32(&pk.WorldGameMode)
	io.Varint32(&pk.Difficulty)
	io.UBlockPos(&pk.WorldSpawn)
	io.Bool(&pk.AchievementsDisabled)
	io.Varint32(&pk.DayCycleLockTime)
	io.Varint32(&pk.EducationEditionOffer)
	io.Bool(&pk.EducationFeaturesEnabled)
	io.String(&pk.EducationProductID)
	io.Float32(&pk.RainLevel)
	io.Float32(&pk.LightningLevel)
	io.Bool(&pk.ConfirmedPlatformLockedContent)
	io.Bool(&pk.MultiPlayerGame)
	io.Bool(&pk.LANBroadcastEnabled)
	io.Varint32(&pk.XBLBroadcastMode)
	io.Varint32(&pk.PlatformBroadcastMode)
	io.Bool(&pk.CommandsEnabled)
	io.Bool(&pk.TexturePackRequired)
	protocol.FuncSlice(io, &pk.GameRules, io.GameRule)
	protocol.SliceUint32Length(io, &pk.Experiments)
	io.Bool(&pk.ExperimentsPreviouslyToggled)
	io.Bool(&pk.BonusChestEnabled)
	io.Bool(&pk.StartWithMapEnabled)
	io.Varint32(&pk.PlayerPermissions)
	io.Int32(&pk.ServerChunkTickRadius)
	io.Bool(&pk.HasLockedBehaviourPack)
	io.Bool(&pk.HasLockedTexturePack)
	io.Bool(&pk.FromLockedWorldTemplate)
	io.Bool(&pk.MSAGamerTagsOnly)
	io.Bool(&pk.FromWorldTemplate)
	io.Bool(&pk.WorldTemplateSettingsLocked)
	io.Bool(&pk.OnlySpawnV1Villagers)
	io.String(&pk.BaseGameVersion)
	io.Int32(&pk.LimitedWorldWidth)
	io.Int32(&pk.LimitedWorldDepth)
	io.Bool(&pk.NewNether)
	protocol.Single(io, &pk.EducationSharedResourceURI)
	io.Bool(&pk.ForceExperimentalGameplay)
	if pk.ForceExperimentalGameplay {
		// This might look wrong, but it's correct: Mojank is reading/writing the same boolean twice if it's set to true
		io.Bool(&pk.ForceExperimentalGameplay)
	}
	io.String(&pk.LevelID)
	io.String(&pk.WorldName)
	io.String(&pk.TemplateContentIdentity)
	io.Bool(&pk.Trial)
	protocol.PlayerMoveSettings(io, &pk.PlayerMovementSettings)
	io.Int64(&pk.Time)
	io.Varint32(&pk.EnchantmentSeed)
	protocol.Slice(io, &pk.Blocks)
	protocol.Slice(io, &pk.Items)
	io.String(&pk.MultiPlayerCorrelationID)
	io.Bool(&pk.ServerAuthoritativeInventory)
	io.String(&pk.GameVersion)
	io.Uint64(&pk.ServerBlockStateChecksum)
}
