// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddFriendByPrincipalID sets the AddFriendByPrincipalID handler function
func (protocol *Protocol) AddFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)) {
	protocol.addFriendByPrincipalIDHandler = handler
}

func (protocol *Protocol) handleAddFriendByPrincipalID(packet nex.PacketInterface) {
	if protocol.addFriendByPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::AddFriendByPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lfc, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.addFriendByPrincipalIDHandler(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.addFriendByPrincipalIDHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.addFriendByPrincipalIDHandler(nil, client, callID, lfc, pid)
}
