package mapping

import (
	"bytes"
	"sort"

	"github.com/df-mc/worldupgrader/blockupgrader"
	"github.com/didntpot/multiversion/multiversion/internal"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/nbt"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/segmentio/fasthash/fnv1"
)

// Block ...
type Block interface {
	// HashToRuntimeID converts a hashed block ID to a runtime ID.
	HashToRuntimeID(hash uint32) (uint32, bool)
	// StateToRuntimeID converts a block state to a runtime ID.
	StateToRuntimeID(blockupgrader.BlockState) (uint32, bool)
	// RuntimeIDToState converts a runtime ID to a name and its state properties.
	RuntimeIDToState(uint32) (blockupgrader.BlockState, bool)
	// DowngradeBlockActorData downgrades the input sub chunk to a legacy block actor.
	DowngradeBlockActorData(map[string]any)
	// UpgradeBlockActorData upgrades the input sub chunk to the latest block actor.
	UpgradeBlockActorData(map[string]any)
	// Adjust adjusts the latest mappings to account for custom states.
	Adjust([]protocol.BlockEntry)
	// Air returns the runtime ID of air.
	Air() uint32
	// UnknownBlock returns the runtime ID of unknown block.
	UnknownBlock() uint32
}

// DefaultBlockMapping ...
type DefaultBlockMapping struct {
	// states holds a list of all possible vanilla block states.
	states []blockupgrader.BlockState
	// stateRuntimeIDs holds a map for looking up the runtime ID of a block by the stateHash it produces.
	stateRuntimeIDs map[internal.StateHash]uint32
	// hashRuntimeIDs holds a map for looking up the runtime ID of a block by the hashed block ID.
	hashRuntimeIDs map[uint32]uint32
	// runtimeIDToState holds a map for looking up the blockState of a block by its runtime ID.
	runtimeIDToState map[uint32]blockupgrader.BlockState
	// airRID is the runtime ID of the air block.
	airRID uint32
	// unknownBlockRID is the runtime ID of the unknown block.
	unknownBlockRID uint32

	upgrader, downgrader func(map[string]any) map[string]any
}

// NewBlockMapping ...
func NewBlockMapping(raw []byte) *DefaultBlockMapping {
	dec := nbt.NewDecoder(bytes.NewBuffer(raw))

	var states []blockupgrader.BlockState
	stateRuntimeIDs := make(map[internal.StateHash]uint32)
	runtimeIDToState := make(map[uint32]blockupgrader.BlockState)
	hashRuntimeIDs := make(map[uint32]uint32)
	var airRID *uint32
	var unknownBlockRID *uint32

	minecraft.DefaultProtocol.ID()
	var s blockupgrader.BlockState
	for {
		if err := dec.Decode(&s); err != nil {
			break
		}

		rid := uint32(len(states))
		states = append(states, s)
		switch s.Name {
		case "minecraft:air":
			airRID = &rid
		case "minecraft:unknown", "minecraft:info_update":
			unknownBlockRID = &rid
		}

		upgraded := blockupgrader.Upgrade(s)
		stateRuntimeIDs[internal.HashState(upgraded)] = rid
		runtimeIDToState[rid] = s
		hashRuntimeIDs[networkBlockHash(upgraded.Name, upgraded.Properties)] = rid
	}
	if airRID == nil {
		panic("couldn't find air")
	}
	if unknownBlockRID == nil {
		panic("couldn't find unknown block")
	}

	return &DefaultBlockMapping{
		states:           states,
		stateRuntimeIDs:  stateRuntimeIDs,
		hashRuntimeIDs:   hashRuntimeIDs,
		runtimeIDToState: runtimeIDToState,
		airRID:           *airRID,
		unknownBlockRID:  *unknownBlockRID,
	}
}

func (m *DefaultBlockMapping) HashToRuntimeID(hash uint32) (uint32, bool) {
	rid, found := m.hashRuntimeIDs[hash]
	return rid, found
}

func (m *DefaultBlockMapping) WithBlockActorRemapper(downgrader, upgrader func(map[string]any) map[string]any) *DefaultBlockMapping {
	m.downgrader = downgrader
	m.upgrader = upgrader
	return m
}

func (m *DefaultBlockMapping) StateToRuntimeID(state blockupgrader.BlockState) (uint32, bool) {
	rid, ok := m.stateRuntimeIDs[internal.HashState(blockupgrader.Upgrade(state))]
	return rid, ok
}

func (m *DefaultBlockMapping) RuntimeIDToState(runtimeId uint32) (blockupgrader.BlockState, bool) {
	state, found := m.runtimeIDToState[runtimeId]
	return state, found
}

func (m *DefaultBlockMapping) DowngradeBlockActorData(actorData map[string]any) {
	if m.downgrader != nil {
		m.downgrader(actorData)
	}
}

func (m *DefaultBlockMapping) UpgradeBlockActorData(actorData map[string]any) {
	if m.upgrader != nil {
		m.upgrader(actorData)
	}
}

func (m *DefaultBlockMapping) Adjust(entries []protocol.BlockEntry) {
	if len(entries) == 0 {
		return
	}

	customStates := convert(entries)
	var newStates []blockupgrader.BlockState
	for _, state := range customStates {
		if _, ok := m.StateToRuntimeID(state); !ok {
			newStates = append(newStates, state)
		}
	}
	if len(newStates) == 0 {
		return
	}

	adjustedStates := append(m.states, customStates...)
	sort.SliceStable(adjustedStates, func(i, j int) bool {
		stateOne, stateTwo := adjustedStates[i], adjustedStates[j]
		return stateOne.Name != stateTwo.Name && fnv1.HashString64(stateOne.Name) < fnv1.HashString64(stateTwo.Name)
	})

	m.stateRuntimeIDs = make(map[internal.StateHash]uint32, len(adjustedStates))
	m.runtimeIDToState = make(map[uint32]blockupgrader.BlockState, len(adjustedStates))
	for rid, state := range adjustedStates {
		m.stateRuntimeIDs[internal.HashState(blockupgrader.Upgrade(state))] = uint32(rid)
		m.runtimeIDToState[uint32(rid)] = state
	}
}

func (m *DefaultBlockMapping) Air() uint32 {
	return m.airRID
}

func (m *DefaultBlockMapping) UnknownBlock() uint32 {
	return m.unknownBlockRID
}
