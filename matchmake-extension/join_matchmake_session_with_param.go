// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/match-making/types"
)

func (protocol *Protocol) handleJoinMatchmakeSessionWithParam(packet nex.PacketInterface) {
	if protocol.JoinMatchmakeSessionWithParam == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::JoinMatchmakeSessionWithParam not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	joinMatchmakeSessionParam := match_making_types.NewJoinMatchmakeSessionParam()

	err := joinMatchmakeSessionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionWithParam(fmt.Errorf("Failed to read joinMatchmakeSessionParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.JoinMatchmakeSessionWithParam(nil, packet, callID, joinMatchmakeSessionParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
