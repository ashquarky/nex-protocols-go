// Package protocol implements the Subscriber protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetNumFollowers(packet nex.PacketInterface) {
	if protocol.GetNumFollowers == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::GetNumFollowers not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	globals.Logger.Warning("Subscriber::GetNumFollowers STUBBED")

	request := packet.RMCMessage()
	callID := request.CallID

	rmcMessage, rmcError := protocol.GetNumFollowers(nil, packet, callID, packet.Payload())
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
