// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetInvitationsSent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetInvitationsSent == nil {
		globals.Logger.Warning("MatchMaking::GetInvitationsSent not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.GetInvitationsSent(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetInvitationsSent(nil, packet, callID, idGathering)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
