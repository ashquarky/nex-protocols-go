// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (protocol *DataStoreProtocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParams []*datastore_types.DataStoreGetMetaParam)) {
	protocol.GetMetasMultipleParamHandler = handler
}

func (protocol *DataStoreProtocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	if protocol.GetMetasMultipleParamHandler == nil {
		globals.Logger.Warning("DataStore::GetMetasMultipleParam not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	params, err := parametersStream.ReadListStructure(datastore_types.NewDataStoreGetMetaParam())
	if err != nil {
		go protocol.GetMetasMultipleParamHandler(fmt.Errorf("Failed to read params from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetMetasMultipleParamHandler(nil, client, callID, params.([]*datastore_types.DataStoreGetMetaParam))
}
