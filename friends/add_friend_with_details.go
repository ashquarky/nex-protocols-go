// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendByNameWithDetails(packet nex.PacketInterface) {
	if protocol.AddFriendByNameWithDetails == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::AddFriendByNameWithDetails not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	uiPlayer := types.NewPrimitiveU32(0)
	uiDetails := types.NewPrimitiveU32(0)
	strMessage := types.NewString("")

	var err error

	err = uiPlayer.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiDetails.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendByNameWithDetails(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddFriendByNameWithDetails(nil, packet, callID, uiPlayer, uiDetails, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
