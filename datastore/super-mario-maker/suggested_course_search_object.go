package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SuggestedCourseSearchObject sets the SuggestedCourseSearchObject handler function
func (protocol *DataStoreSuperMarioMakerProtocol) SuggestedCourseSearchObject(handler func(err error, client *nex.Client, callID uint32, param *datastore_types.DataStoreSearchParam, extraData []string)) {
	protocol.SuggestedCourseSearchObjectHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleSuggestedCourseSearchObject(packet nex.PacketInterface) {
	if protocol.SuggestedCourseSearchObjectHandler == nil {
		globals.Logger.Warning("DataStoreSMM::SuggestedCourseSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_types.NewDataStoreSearchParam())
	if err != nil {
		go protocol.SuggestedCourseSearchObjectHandler(err, client, callID, nil, []string{})
		return
	}

	extraData, err := parametersStream.ReadListString()
	if err != nil {
		go protocol.SuggestedCourseSearchObjectHandler(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.SuggestedCourseSearchObjectHandler(nil, client, callID, param.(*datastore_types.DataStoreSearchParam), extraData)
}
