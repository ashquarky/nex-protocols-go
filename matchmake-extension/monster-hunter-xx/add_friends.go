// Package protocol implements the MatchmakeExtensionMonsterHunterXX protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriends(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AddFriends == nil {
		globals.Logger.Warning("MatchmakeExtensionMonsterHunterXX::AddFriends not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)
	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AddFriends(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AddFriends(nil, packet, callID, pids)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
