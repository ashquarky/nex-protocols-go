// Package protocol implements the Matchmake Referee protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetNotSummarizedRound sets the GetNotSummarizedRound handler function
func (protocol *Protocol) GetNotSummarizedRound(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getNotSummarizedRoundHandler = handler
}

func (protocol *Protocol) handleGetNotSummarizedRound(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getNotSummarizedRoundHandler == nil {
		globals.Logger.Warning("MatchmakeReferee::GetNotSummarizedRound not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getNotSummarizedRoundHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
