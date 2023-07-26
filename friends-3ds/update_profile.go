// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_3ds_types "github.com/PretendoNetwork/nex-protocols-go/friends-3ds/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProfile sets the UpdateProfile handler function
func (protocol *Protocol) UpdateProfile(handler func(err error, client *nex.Client, callID uint32, profileData *friends_3ds_types.MyProfile)) {
	protocol.updateProfileHandler = handler
}

func (protocol *Protocol) handleUpdateProfile(packet nex.PacketInterface) {
	if protocol.updateProfileHandler == nil {
		globals.Logger.Warning("Friends3DS::UpdateProfile not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	profileData, err := parametersStream.ReadStructure(friends_3ds_types.NewMyProfile())
	if err != nil {
		go protocol.updateProfileHandler(fmt.Errorf("Failed to read showGame from profileData. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.updateProfileHandler(nil, client, callID, profileData.(*friends_3ds_types.MyProfile))
}
