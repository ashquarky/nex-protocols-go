package friends_wiiu

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdatePresence sets the UpdatePresence handler function
func (protocol *FriendsWiiUProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresenceV2)) {
	protocol.UpdatePresenceHandler = handler
}

func (protocol *FriendsWiiUProtocol) HandleUpdatePresence(packet nex.PacketInterface) {
	if protocol.UpdatePresenceHandler == nil {
		globals.Logger.Warning("FriendsWiiU::UpdatePresence not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nintendoPresenceV2, err := parametersStream.ReadStructure(NewNintendoPresenceV2())
	if err != nil {
		go protocol.UpdatePresenceHandler(fmt.Errorf("Failed to read nintendoPresenceV2 from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.UpdatePresenceHandler(nil, client, callID, nintendoPresenceV2.(*NintendoPresenceV2))
}
