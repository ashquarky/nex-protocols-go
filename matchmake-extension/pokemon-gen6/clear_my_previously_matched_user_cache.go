// Package protocol implements the Pokemon GEN 6 Matchmake Extension protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleClearMyPreviouslyMatchedUserCache(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ClearMyPreviouslyMatchedUserCache == nil {
		globals.Logger.Warning("MatchmakeExtension::ClearMyPreviouslyMatchedUserCache not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID

	rmcMessage, errorCode := protocol.ClearMyPreviouslyMatchedUserCache(nil, packet, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
