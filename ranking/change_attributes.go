// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

// ChangeAttributes sets the ChangeAttributes handler function
func (protocol *Protocol) ChangeAttributes(handler func(err error, client *nex.Client, callID uint32, category uint32, changeParam *ranking_types.RankingChangeAttributesParam, uniqueID uint64) uint32) {
	protocol.changeAttributesHandler = handler
}

func (protocol *Protocol) handleChangeAttributes(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.changeAttributesHandler == nil {
		globals.Logger.Warning("Ranking::ChangeAttributes not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	category, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.changeAttributesHandler(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	changeParam, err := parametersStream.ReadStructure(ranking_types.NewRankingChangeAttributesParam())
	if err != nil {
		errorCode = protocol.changeAttributesHandler(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.changeAttributesHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.changeAttributesHandler(nil, client, callID, category, changeParam.(*ranking_types.RankingChangeAttributesParam), uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
