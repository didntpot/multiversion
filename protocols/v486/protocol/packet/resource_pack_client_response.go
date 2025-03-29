package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PackResponseRefused = iota + 1
	PackResponseSendPacks
	PackResponseAllPacksDownloaded
	PackResponseCompleted
)

// ResourcePackClientResponse is sent by the client in response to resource packets sent by the server. It is
// used to let the server know what action needs to be taken for the client to have all resource packs ready
// and set.
type ResourcePackClientResponse struct {
	Response        byte
	PacksToDownload []string
}

// ID ...
func (*ResourcePackClientResponse) ID() uint32 {
	return IDResourcePackClientResponse
}

// Marshal ...
func (pk *ResourcePackClientResponse) Marshal(io protocol.IO) {
	io.Uint8(&pk.Response)
	protocol.FuncSliceUint16Length(io, &pk.PacksToDownload, io.String)
}
