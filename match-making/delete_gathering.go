// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteGathering sets the DeleteGathering handler function
func (protocol *Protocol) DeleteGathering(handler func(err error, client *nex.Client, callID uint32, idGathering uint32)) {
	protocol.deleteGatheringHandler = handler
}

func (protocol *Protocol) handleDeleteGathering(packet nex.PacketInterface) {
	if protocol.deleteGatheringHandler == nil {
		globals.Logger.Warning("MatchMaking::DeleteGathering not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idGathering, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.deleteGatheringHandler(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.deleteGatheringHandler(nil, client, callID, idGathering)
}
