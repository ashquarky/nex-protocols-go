// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetObject sets the PrepareGetObject handler function
func (protocol *Protocol) PrepareGetObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStorePrepareGetParam) uint32) {
	protocol.prepareGetObjectHandler = handler
}

func (protocol *Protocol) handlePrepareGetObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.prepareGetObjectHandler == nil {
		globals.Logger.Warning("DataStore::PrepareGetObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStorePrepareGetParam())
	if err != nil {
		errorCode = protocol.prepareGetObjectHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.prepareGetObjectHandler(nil, client, callID, param.(*datastore_types.DataStorePrepareGetParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
