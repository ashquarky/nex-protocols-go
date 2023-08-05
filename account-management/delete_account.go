// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteAccount sets the DeleteAccount handler function
func (protocol *Protocol) DeleteAccount(handler func(err error, client *nex.Client, callID uint32, idPrincipal uint32) uint32) {
	protocol.deleteAccountHandler = handler
}

func (protocol *Protocol) handleDeleteAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteAccountHandler == nil {
		globals.Logger.Warning("AccountManagement::DeleteAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	idPrincipal, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.deleteAccountHandler(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteAccountHandler(nil, client, callID, idPrincipal)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
