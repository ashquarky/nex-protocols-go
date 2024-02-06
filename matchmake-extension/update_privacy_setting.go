// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdatePrivacySetting(packet nex.PacketInterface) {
	if protocol.UpdatePrivacySetting == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdatePrivacySetting not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	onlineStatus := types.NewPrimitiveBool(false)
	participationCommunity := types.NewPrimitiveBool(false)

	var err error

	err = onlineStatus.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePrivacySetting(fmt.Errorf("Failed to read onlineStatus from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = participationCommunity.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdatePrivacySetting(fmt.Errorf("Failed to read participationCommunity from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdatePrivacySetting(nil, packet, callID, onlineStatus, participationCommunity)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
