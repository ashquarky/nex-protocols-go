// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleDisableAccount(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.DisableAccount == nil {
		globals.Logger.Warning("AccountManagement::DisableAccount not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idPrincipal := types.NewPID(0)
	err = idPrincipal.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DisableAccount(fmt.Errorf("Failed to read idPrincipal from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	dtUntil := types.NewDateTime(0)
	err = dtUntil.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DisableAccount(fmt.Errorf("Failed to read dtUntil from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage := types.NewString("")
	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.DisableAccount(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.DisableAccount(nil, packet, callID, idPrincipal, dtUntil, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
