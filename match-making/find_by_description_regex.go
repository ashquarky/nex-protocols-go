// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleFindByDescriptionRegex(packet nex.PacketInterface) {
	if protocol.FindByDescriptionRegex == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::FindByDescriptionRegex not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	strDescriptionRegex := types.NewString("")
	resultRange := types.NewResultRange()

	var err error

	err = strDescriptionRegex.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByDescriptionRegex(fmt.Errorf("Failed to read strDescriptionRegex from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByDescriptionRegex(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByDescriptionRegex(nil, packet, callID, strDescriptionRegex, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
