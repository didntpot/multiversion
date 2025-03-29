package packet

import (
	legacyprotocol "github.com/didntpot/multiversion/protocols/v486/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	EventTypeAchievementAwarded = iota
	EventTypeEntityInteract
	EventTypePortalBuilt
	EventTypePortalUsed
	EventTypeMobKilled
	EventTypeCauldronUsed
	EventTypePlayerDied
	EventTypeBossKilled
	EventTypeAgentCommand
	EventTypeAgentCreated // Unused for whatever reason?
	EventTypePatternRemoved
	EventTypeSlashCommandExecuted
	EventTypeFishBucketed
	EventTypeMobBorn
	EventTypePetDied
	EventTypeCauldronInteract
	EventTypeComposterInteract
	EventTypeBellUsed
	EventTypeEntityDefinitionTrigger
	EventTypeRaidUpdate
	EventTypeMovementAnomaly
	EventTypeMovementCorrected
	EventTypeExtractHoney
	EventTypeTargetBlockHit
	EventTypePiglinBarter
	EventTypePlayerWaxedOrUnwaxedCopper
	EventTypeCodeBuilderRuntimeAction
	EventTypeCodeBuilderScoreboard
	EventTypeStriderRiddenInLavaInOverworld
	EventTypeSneakCloseToSculkSensor
)

// Event is sent by the server to send an event with additional data. It is typically sent to the client for
// telemetry reasons, much like the SimpleEvent packet.
type Event struct {
	EntityRuntimeID uint64
	EventType       int32
	UsePlayerID     byte
	EventData       legacyprotocol.EventData
}

// ID ...
func (*Event) ID() uint32 {
	return IDEvent
}

// Marshal ...
func (pk *Event) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Varint32(&pk.EventType)
	io.Uint8(&pk.UsePlayerID)

	pk.EventData.Marshal(io)
}
