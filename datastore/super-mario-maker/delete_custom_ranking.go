// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDeleteCustomRanking(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DeleteCustomRanking == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteCustomRanking not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataIDList := types.NewList[*types.PrimitiveU64]()
	dataIDList.Type = types.NewPrimitiveU64(0)
	err = dataIDList.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DeleteCustomRanking(fmt.Errorf("Failed to read dataIDList from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DeleteCustomRanking(nil, packet, callID, dataIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
