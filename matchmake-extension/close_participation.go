// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCloseParticipation(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.CloseParticipation == nil {
		globals.Logger.Warning("MatchmakeExtension::CloseParticipation not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		_, errorCode = protocol.CloseParticipation(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CloseParticipation(nil, packet, callID, gid)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
