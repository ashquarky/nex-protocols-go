// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleClearRelationship(packet nex.PacketInterface) {
	if protocol.ClearRelationship == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::ClearRelationship not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiPlayer := types.NewPrimitiveU32(0)

	err := uiPlayer.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ClearRelationship(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ClearRelationship(nil, packet, callID, uiPlayer)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
