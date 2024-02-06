// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteGathering(packet nex.PacketInterface) {
	if protocol.DeleteGathering == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::DeleteGathering not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)

	err := idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DeleteGathering(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DeleteGathering(nil, packet, callID, idGathering)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
