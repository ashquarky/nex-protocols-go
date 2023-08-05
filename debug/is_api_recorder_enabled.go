// Package protocol implements the Debug protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// IsAPIRecorderEnabled sets the IsAPIRecorderEnabled handler function
func (protocol *Protocol) IsAPIRecorderEnabled(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.isAPIRecorderEnabledHandler = handler
}

func (protocol *Protocol) handleIsAPIRecorderEnabled(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.isAPIRecorderEnabledHandler == nil {
		globals.Logger.Warning("Debug::IsAPIRecorderEnabled not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.isAPIRecorderEnabledHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
