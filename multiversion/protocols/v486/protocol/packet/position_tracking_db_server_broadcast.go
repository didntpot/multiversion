package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PositionTrackingDBBroadcastActionUpdate = iota
	PositionTrackingDBBroadcastActionDestroy
	PositionTrackingDBBroadcastActionNotFound
)

// PositionTrackingDBServerBroadcast is sent by the server in response to the
// PositionTrackingDBClientRequest packet. This packet is, as of 1.16, currently only used for lodestones. The
// server maintains a database with tracking IDs and their position and dimension. The client will request
// these tracking IDs, (NBT tag set on the lodestone compass with the tracking ID?) and the server will
// respond with the status of those tracking IDs.
// What is actually done with the data sent depends on what the client chooses to do with it. For the
// lodestone compass, it is used to make the compass point towards lodestones and to make it spin if the
// lodestone at a position is no longer there.
type PositionTrackingDBServerBroadcast struct {
	BroadcastAction byte
	TrackingID      int32
	SerialisedData  []byte
}

// ID ...
func (*PositionTrackingDBServerBroadcast) ID() uint32 {
	return IDPositionTrackingDBServerBroadcast
}

// Marshal ...
func (pk *PositionTrackingDBServerBroadcast) Marshal(io protocol.IO) {
	io.Uint8(&pk.BroadcastAction)
	io.Varint32(&pk.TrackingID)
	io.Bytes(&pk.SerialisedData)
}
