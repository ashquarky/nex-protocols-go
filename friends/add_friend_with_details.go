// Package friends implements the Friends QRV protocol
package friends

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByNameWithDetails sets the AddFriendByNameWithDetails handler function
func (protocol *FriendsProtocol) AddFriendByNameWithDetails(handler func(err error, client *nex.Client, callID uint32, uiPlayer uint32, uiDetails uint32, strMessage string)) {
	protocol.addFriendByNameWithDetailsHandler = handler
}

func (protocol *FriendsProtocol) handleAddFriendByNameWithDetails(packet nex.PacketInterface) {
	if protocol.addFriendByNameWithDetailsHandler == nil {
		globals.Logger.Warning("Friends::AddFriendByNameWithDetails not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uiPlayer, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.addFriendByNameWithDetailsHandler(fmt.Errorf("Failed to read uiPlayer from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	uiDetails, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.addFriendByNameWithDetailsHandler(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.addFriendByNameWithDetailsHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, 0, "")
		return
	}

	go protocol.addFriendByNameWithDetailsHandler(nil, client, callID, uiPlayer, uiDetails, strMessage)
}
