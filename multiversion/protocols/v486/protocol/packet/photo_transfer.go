package packet

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol"
)

const (
	PhotoTypePortfolio uint8 = iota
	PhotoTypePhotoItem
	PhotoTypeBook
)

// PhotoTransfer is sent by the server to transfer a photo (image) file to the client. It is typically used
// to transfer photos so that the client can display it in a portfolio in Education Edition.
// While previously usable in the default Bedrock Edition, the displaying of photos in books was disabled and
// the packet now has little use anymore.
type PhotoTransfer struct {
	PhotoName           string
	PhotoData           []byte
	BookID              string
	PhotoType           byte
	SourceType          byte
	OwnerEntityUniqueID int64
	NewPhotoName        string
}

// ID ...
func (*PhotoTransfer) ID() uint32 {
	return IDPhotoTransfer
}

// Marshal ...
func (pk *PhotoTransfer) Marshal(io protocol.IO) {
	io.String(&pk.PhotoName)
	io.ByteSlice(&pk.PhotoData)
	io.String(&pk.BookID)
	io.Uint8(&pk.PhotoType)
	io.Uint8(&pk.SourceType)
	io.Int64(&pk.OwnerEntityUniqueID)
	io.String(&pk.NewPhotoName)
}
