package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	ActorEventJump = iota + 1
	ActorEventHurt
	ActorEventDeath
	ActorEventStartAttacking
	ActorEventStopAttacking
	ActorEventTamingFailed
	ActorEventTamingSucceeded
	ActorEventShakeWetness
	ActorEventUseItem
	ActorEventEatGrass
	ActorEventFishhookBubble
	ActorEventFishhookFishPosition
	ActorEventFishhookHookTime
	ActorEventFishhookTease
	ActorEventSquidFleeing
	ActorEventZombieConverting
	ActorEventPlayAmbient
	ActorEventSpawnAlive
	ActorEventStartOfferFlower
	ActorEventStopOfferFlower
	ActorEventLoveHearts
	ActorEventVillagerAngry
	ActorEventVillagerHappy
	ActorEventWitchHatMagic
	ActorEventFireworksExplode
	ActorEventInLoveHearts
	ActorEventSilverfishMergeAnimation
	ActorEventGuardianAttackSound
	ActorEventDrinkPotion
	ActorEventThrowPotion
	ActorEventCartWithPrimeTNT
	ActorEventPrimeCreeper
	ActorEventAirSupply
	ActorEventAddPlayerLevels
	ActorEventGuardianMiningFatigue
	ActorEventAgentSwingArm
	ActorEventDragonStartDeathAnim
	ActorEventGroundDust
	ActorEventShake
)

const (
	ActorEventFeed = iota + 57
	ActorEventBabyEat
	ActorEventInstantDeath
	ActorEventNotifyTrade
	ActorEventLeashDestroyed
	ActorEventCaravanUpdated
	ActorEventTalismanActivate
	ActorEventUpdateStructureFeature
	ActorEventPlayerSpawnedMob
	ActorEventPuke
	ActorEventUpdateStackSize
	ActorEventStartSwimming
	ActorEventBalloonPop
	ActorEventTreasureHunt
	ActorEventSummonAgent
	ActorEventFinishedChargingCrossbow
	ActorEventLandedOnGround
	ActorEventActorGrowUp
)

// ActorEvent is sent by the server when a particular event happens that has to do with an entity. Some of
// these events are entity-specific, for example a wolf shaking itself dry, but others are used for each
// entity, such as dying.
type ActorEvent struct {
	EntityRuntimeID uint64
	EventType       byte
	EventData       int32
}

// ID ...
func (*ActorEvent) ID() uint32 {
	return IDActorEvent
}

// Marshal ...
func (pk *ActorEvent) Marshal(io protocol.IO) {
	io.Varuint64(&pk.EntityRuntimeID)
	io.Uint8(&pk.EventType)
	io.Varint32(&pk.EventData)
}
