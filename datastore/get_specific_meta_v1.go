// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetSpecificMetaV1 sets the GetSpecificMetaV1 handler function
func (protocol *Protocol) GetSpecificMetaV1(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreGetSpecificMetaParamV1)) {
	protocol.GetSpecificMetaV1Handler = handler
}

func (protocol *Protocol) handleGetSpecificMetaV1(packet nex.PacketInterface) {
	if protocol.GetSpecificMetaV1Handler == nil {
		globals.Logger.Warning("DataStore::GetSpecificMetaV1 not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreGetSpecificMetaParamV1())
	if err != nil {
		go protocol.GetSpecificMetaV1Handler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetSpecificMetaV1Handler(nil, client, callID, param.(*datastore_types.DataStoreGetSpecificMetaParamV1))
}
