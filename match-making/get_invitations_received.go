// Package protocol implements the Match Making protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetInvitationsReceived sets the GetInvitationsReceived handler function
func (protocol *Protocol) GetInvitationsReceived(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.getInvitationsReceivedHandler = handler
}

func (protocol *Protocol) handleGetInvitationsReceived(packet nex.PacketInterface) {
	if protocol.getInvitationsReceivedHandler == nil {
		globals.Logger.Warning("MatchMaking::GetInvitationsReceived not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.getInvitationsReceivedHandler(nil, client, callID)
}
