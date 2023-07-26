// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetExcludeJoinedMatchmakeSession sets the GetExcludeJoinedMatchmakeSession handler function
func (protocol *Protocol) GetExcludeJoinedMatchmakeSession(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.GetExcludeJoinedMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleGetExcludeJoinedMatchmakeSession(packet nex.PacketInterface) {
	globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession STUBBED")

	if protocol.GetExcludeJoinedMatchmakeSessionHandler == nil {
		globals.Logger.Warning("Debug::GetExcludeJoinedMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
}
