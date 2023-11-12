// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriend sets the AddFriend handler function
func (protocol *Protocol) AddFriend(handler func(err error, packet nex.PacketInterface, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string) uint32) {
	protocol.addFriendHandler = handler
}

func (protocol *Protocol) handleAddFriend(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.addFriendHandler == nil {
		globals.Logger.Warning("Friends::AddFriend not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.addFriendHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.addFriendHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, 0, 0, "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.addFriendHandler(nil, packet, callID, uiPlayer, uiDetails, strMessage)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
