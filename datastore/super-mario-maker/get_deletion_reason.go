package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetDeletionReason sets the GetDeletionReason handler function
func (protocol *DataStoreSuperMarioMakerProtocol) GetDeletionReason(handler func(err error, client *nex.Client, callID uint32, dataIdLst []uint64)) {
	protocol.GetDeletionReasonHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) HandleGetDeletionReason(packet nex.PacketInterface) {
	if protocol.GetDeletionReasonHandler == nil {
		globals.Logger.Warning("DataStoreSMM::GetDeletionReason not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataIdLst, err := parametersStream.ReadListUInt64LE()
	if err != nil {
		go protocol.GetDeletionReasonHandler(fmt.Errorf("Failed to read dataIdLst from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.GetDeletionReasonHandler(nil, client, callID, dataIdLst)
}
