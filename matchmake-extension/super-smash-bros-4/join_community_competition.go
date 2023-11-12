// Package protocol implements the MatchmakeExtensionSuperSmashBros4 protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinCommunityCompetition sets the JoinCommunityCompetition handler function
func (protocol *Protocol) JoinCommunityCompetition(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) uint32) {
	protocol.joinCommunityCompetitionHandler = handler
}

func (protocol *Protocol) handleJoinCommunityCompetition(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinCommunityCompetitionHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetition not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	globals.Logger.Warning("MatchmakeExtensionSuperSmashBros4::JoinCommunityCompetition STUBBED")

	request := packet.RMCMessage()

	callID := request.CallID

	errorCode = protocol.joinCommunityCompetitionHandler(nil, packet, callID, packet.Payload())
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
