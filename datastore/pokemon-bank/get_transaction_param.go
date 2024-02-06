// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetTransactionParam(packet nex.PacketInterface) {
	if protocol.GetTransactionParam == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::GetTransactionParam not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	slotID := types.NewPrimitiveU16(0)

	err := slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTransactionParam(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetTransactionParam(nil, packet, callID, slotID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
