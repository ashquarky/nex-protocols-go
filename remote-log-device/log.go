// Package remote_log_device implements the Remote Log Device NEX protocol
package remote_log_device

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Log sets the Log handler function
func (protocol *RemoteLogDeviceProtocol) Log(handler func(err error, client *nex.Client, callID uint32, strLine string)) {
	protocol.LogHandler = handler
}

func (protocol *RemoteLogDeviceProtocol) handleLog(packet nex.PacketInterface) {
	if protocol.LogHandler == nil {
		globals.Logger.Warning("RemoteLogDevice::Log not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	strLine, err := parametersStream.ReadString()
	if err != nil {
		go protocol.LogHandler(fmt.Errorf("Failed to read strLine from parameters. %s", err.Error()), client, callID, "")
	}

	go protocol.LogHandler(nil, client, callID, strLine)
}
