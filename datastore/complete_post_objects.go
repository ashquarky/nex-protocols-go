// Package datastore implements the DataStore NEX protocol
package datastore

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CompletePostObjects sets the CompletePostObjects handler function
func (protocol *DataStoreProtocol) CompletePostObjects(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64)) {
	protocol.CompletePostObjectsHandler = handler
}

func (protocol *DataStoreProtocol) handleCompletePostObjects(packet nex.PacketInterface) {
	if protocol.CompletePostObjectsHandler == nil {
		globals.Logger.Warning("DataStore::CompletePostObjects not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIDs, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.CompletePostObjectsHandler(fmt.Errorf("Failed to read dataIDs from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.CompletePostObjectsHandler(nil, client, callID, dataIDs)
}
