// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateMii(packet nex.PacketInterface) {
	if protocol.UpdateMii == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::UpdateMii not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	miiV2 := friends_wiiu_types.NewMiiV2()

	err := miiV2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMii(fmt.Errorf("Failed to read miiV2 from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMii(nil, packet, callID, miiV2)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
