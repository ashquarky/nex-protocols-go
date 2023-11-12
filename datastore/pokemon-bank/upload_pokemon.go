// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UploadPokemon sets the UploadPokemon handler function
func (protocol *Protocol) UploadPokemon(handler func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam) uint32) {
	protocol.uploadPokemonHandler = handler
}

func (protocol *Protocol) handleUploadPokemon(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.uploadPokemonHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::UploadPokemon not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(datastore_pokemon_bank_types.NewGlobalTradeStationUploadPokemonParam())
	if err != nil {
		errorCode = protocol.uploadPokemonHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.uploadPokemonHandler(nil, packet, callID, param.(*datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
