package protocol

import (
	"strings"

	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

// EventData ...
type EventData interface {
	Marshal(io protocol.IO)
}

// AchievementAwardedEventData is the event data sent for achievements.
type AchievementAwardedEventData struct {
	AchievementID int32
}

// Marshal ...
func (x *AchievementAwardedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.AchievementID)
}

// EntityInteractEventData is the event data sent for entity interactions.
type EntityInteractEventData struct {
	InteractionType       int32
	InteractionEntityType int32
	EntityVariant         int32
	EntityColour          uint8
}

// Marshal ...
func (x *EntityInteractEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.InteractionType)
	io.Varint32(&x.InteractionEntityType)
	io.Varint32(&x.EntityVariant)
	io.Uint8(&x.EntityColour)
}

// PortalBuiltEventData is the event data sent when a portal is built.
type PortalBuiltEventData struct {
	DimensionID int32
}

// Marshal ...
func (x *PortalBuiltEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.DimensionID)
}

// PortalUsedEventData is the event data sent when a portal is used.
type PortalUsedEventData struct {
	FromDimensionID int32
	ToDimensionID   int32
}

// Marshal ...
func (x *PortalUsedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.FromDimensionID)
	io.Varint32(&x.ToDimensionID)
}

// MobKilledEventData is the event data sent when a mob is killed.
type MobKilledEventData struct {
	KillerEntityUniqueID int64
	VictimEntityUniqueID int64
	KillerEntityType     int32
	EntityDamageCause    int32
	VillagerTradeTier    int32
	VillagerDisplayName  string
}

// Marshal ...
func (x *MobKilledEventData) Marshal(io protocol.IO) {
	io.Varint64(&x.KillerEntityUniqueID)
	io.Varint64(&x.VictimEntityUniqueID)
	io.Varint32(&x.KillerEntityType)
	io.Varint32(&x.EntityDamageCause)
	io.Varint32(&x.VillagerTradeTier)
	io.String(&x.VillagerDisplayName)
}

// CauldronUsedEventData is the event data sent when a cauldron is used.
type CauldronUsedEventData struct {
	PotionID  int32
	Colour    int32
	FillLevel int32
}

// Marshal ...
func (x *CauldronUsedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.PotionID)
	io.Varint32(&x.Colour)
	io.Varint32(&x.FillLevel)
}

// PlayerDiedEventData is the event data sent when a player dies.
type PlayerDiedEventData struct {
	AttackerEntityID  int32
	EntityDamageCause int32
}

// Marshal ...
func (x *PlayerDiedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.AttackerEntityID)
	io.Varint32(&x.EntityDamageCause)
}

// BossKilledEventData is the event data sent when a boss dies.
type BossKilledEventData struct {
	BossEntityUniqueID    int64
	PlayerPartySize       int32
	InteractionEntityType int32
}

// Marshal ...
func (x *BossKilledEventData) Marshal(io protocol.IO) {
	io.Varint64(&x.BossEntityUniqueID)
	io.Varint32(&x.PlayerPartySize)
	io.Varint32(&x.InteractionEntityType)
}

// AgentCommandEventData is an event used in Education Edition.
type AgentCommandEventData struct {
	AgentResult int32
	DataValue   int32
	Command     string
	DataKey     string
	Output      string
}

// Marshal ...
func (x *AgentCommandEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.AgentResult)
	io.Varint32(&x.DataValue)
	io.String(&x.Command)
	io.String(&x.DataKey)
	io.String(&x.Output)
}

// PatternRemovedEventData is the event data sent when a pattern is removed.
type PatternRemovedEventData struct {
	ItemID        int32
	AuxValue      int32
	PatternsSize  int32
	PatternIndex  int32
	PatternColour int32
}

// Marshal ...
func (x *PatternRemovedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.ItemID)
	io.Varint32(&x.AuxValue)
	io.Varint32(&x.PatternsSize)
	io.Varint32(&x.PatternIndex)
	io.Varint32(&x.PatternColour)
}

// SlashCommandExecutedEventData is the event data sent when a slash command is executed.
type SlashCommandExecutedEventData struct {
	CommandName    string
	SuccessCount   int32
	OutputMessages []string
}

// Marshal ...
func (x *SlashCommandExecutedEventData) Marshal(io protocol.IO) {
	outputMessagesSize := int32(len(x.OutputMessages))
	outputMessagesJoined := strings.Join(x.OutputMessages, ";")

	io.Varint32(&x.SuccessCount)
	io.Varint32(&outputMessagesSize)
	io.String(&x.CommandName)
	io.String(&outputMessagesJoined)
}

// FishBucketedEventData is the event data sent when a fish is bucketed.
type FishBucketedEventData struct {
	Pattern            int32
	Preset             int32
	BucketedEntityType int32
	Release            bool
}

// Marshal ...
func (x *FishBucketedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.Pattern)
	io.Varint32(&x.Preset)
	io.Varint32(&x.BucketedEntityType)
	io.Bool(&x.Release)
}

// MobBornEventData is the event data sent when a mob is born.
type MobBornEventData struct {
	EntityType int32
	Variant    int32
	Colour     uint8
}

// Marshal ...
func (x *MobBornEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.EntityType)
	io.Varint32(&x.Variant)
	io.Uint8(&x.Colour)
}

// PetDiedEventData is the event data sent when a pet dies.
type PetDiedEventData struct {
	KilledByOwner        bool
	KillerEntityUniqueID int64
	PetEntityUniqueID    int64
	EntityDamageCause    int32
	PetEntityType        int32
}

// Marshal ...
func (x *PetDiedEventData) Marshal(io protocol.IO) {
	io.Bool(&x.KilledByOwner)
	io.Varint64(&x.KillerEntityUniqueID)
	io.Varint64(&x.PetEntityUniqueID)
	io.Varint32(&x.EntityDamageCause)
	io.Varint32(&x.PetEntityType)
}

// CauldronInteractEventData is the event data sent when a cauldron is interacted with.
type CauldronInteractEventData struct {
	BlockInteractionType int32
	ItemID               int32
}

// Marshal ...
func (x *CauldronInteractEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.BlockInteractionType)
	io.Varint32(&x.ItemID)
}

// ComposterInteractEventData is the event data sent when a composter is interacted with.
type ComposterInteractEventData struct {
	BlockInteractionType int32
	ItemID               int32
}

// Marshal ...
func (x *ComposterInteractEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.BlockInteractionType)
	io.Varint32(&x.ItemID)
}

// BellUsedEventData is the event data sent when a bell is used.
type BellUsedEventData struct {
	ItemID int32
}

// Marshal ...
func (x *BellUsedEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.ItemID)
}

// EntityDefinitionTriggerEventData is an event used for an unknown purpose.
type EntityDefinitionTriggerEventData struct {
	EventName string
}

// Marshal ...
func (x *EntityDefinitionTriggerEventData) Marshal(io protocol.IO) {
	io.String(&x.EventName)
}

// RaidUpdateEventData is an event used to update a raids progress client side.
type RaidUpdateEventData struct {
	CurrentRaidWave int32
	TotalRaidWaves  int32
	WonRaid         bool
}

// Marshal ...
func (x *RaidUpdateEventData) Marshal(io protocol.IO) {
	io.Varint32(&x.CurrentRaidWave)
	io.Varint32(&x.TotalRaidWaves)
	io.Bool(&x.WonRaid)
}

// MovementAnomalyEventData is an event used for updating the other party on movement data.
type MovementAnomalyEventData struct {
	EventType            uint8
	CheatingScore        float32
	AveragePositionDelta float32
	TotalPositionDelta   float32
	MinPositionDelta     float32
	MaxPositionDelta     float32
}

// Marshal ...
func (x *MovementAnomalyEventData) Marshal(io protocol.IO) {
	io.Uint8(&x.EventType)
	io.Float32(&x.CheatingScore)
	io.Float32(&x.AveragePositionDelta)
	io.Float32(&x.TotalPositionDelta)
	io.Float32(&x.MinPositionDelta)
	io.Float32(&x.MaxPositionDelta)
}

// MovementCorrectedEventData is an event sent by the server to correct movement client side.
type MovementCorrectedEventData struct {
	PositionDelta     float32
	CheatingScore     float32
	ScoreThreshold    float32
	DistanceThreshold float32
	DurationThreshold int32
}

// Marshal ...
func (x *MovementCorrectedEventData) Marshal(io protocol.IO) {
	io.Float32(&x.PositionDelta)
	io.Float32(&x.CheatingScore)
	io.Float32(&x.ScoreThreshold)
	io.Float32(&x.DistanceThreshold)
	io.Varint32(&x.DurationThreshold)
}

// ExtractHoneyEventData is an event with no purpose.
type ExtractHoneyEventData struct{}

// Marshal ...
func (*ExtractHoneyEventData) Marshal(protocol.IO) {}
