// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_super_mario_maker_types "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddToBufferQueues(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.AddToBufferQueues == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::AddToBufferQueues not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	params, err := nex.StreamReadListStructure(parametersStream, datastore_super_mario_maker_types.NewBufferQueueParam())
	if err != nil {
		_, errorCode = protocol.AddToBufferQueues(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	buffers, err := parametersStream.ReadListQBuffer()
	if err != nil {
		_, errorCode = protocol.AddToBufferQueues(fmt.Errorf("Failed to read buffers from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddToBufferQueues(nil, packet, callID, params, buffers)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
