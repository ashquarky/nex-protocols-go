// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAcceptFriendRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AcceptFriendRequest == nil {
		globals.Logger.Warning("FriendsWiiU::AcceptFriendRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	id := types.NewPrimitiveU64(0)
	err = id.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AcceptFriendRequest(fmt.Errorf("Failed to read id from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AcceptFriendRequest(nil, packet, callID, id)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
