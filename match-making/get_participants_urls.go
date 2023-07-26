// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetParticipantsURLs sets the GetParticipantsURLs handler function
func (protocol *Protocol) GetParticipantsURLs(handler func(err error, client *nex.Client, callID uint32, idGathering uint32)) {
	protocol.getParticipantsURLsHandler = handler
}

func (protocol *Protocol) handleGetParticipantsURLs(packet nex.PacketInterface) {
	if protocol.getParticipantsURLsHandler == nil {
		globals.Logger.Warning("MatchMaking::GetParticipantsURLs not implemented")
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
		go protocol.getParticipantsURLsHandler(fmt.Errorf("Failed to read gatheringID from parameters. %s", err.Error()), client, callID, 0)
	}

	go protocol.getParticipantsURLsHandler(nil, client, callID, idGathering)
}
