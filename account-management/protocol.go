// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

const (
	// ProtocolID is the protocol ID for the Account Management protocol
	ProtocolID = 0x19

	// MethodCreateAccount is the method ID for method CreateAccount
	MethodCreateAccount = 0x1

	// MethodDeleteAccount is the method ID for method DeleteAccount
	MethodDeleteAccount = 0x2

	// MethodDisableAccount is the method ID for method DisableAccount
	MethodDisableAccount = 0x3

	// MethodChangePassword is the method ID for method ChangePassword
	MethodChangePassword = 0x4

	// MethodTestCapability is the method ID for method TestCapability
	MethodTestCapability = 0x5

	// MethodGetName is the method ID for method GetName
	MethodGetName = 0x6

	// MethodGetAccountData is the method ID for method GetAccountData
	MethodGetAccountData = 0x7

	// MethodGetPrivateData is the method ID for method GetPrivateData
	MethodGetPrivateData = 0x8

	// MethodGetPublicData is the method ID for method GetPublicData
	MethodGetPublicData = 0x9

	// MethodGetMultiplePublicData is the method ID for method GetMultiplePublicData
	MethodGetMultiplePublicData = 0xA

	// MethodUpdateAccountName is the method ID for method UpdateAccountName
	MethodUpdateAccountName = 0xB

	// MethodUpdateAccountEmail is the method ID for method UpdateAccountEmail
	MethodUpdateAccountEmail = 0xC

	// MethodUpdateCustomData is the method ID for method UpdateCustomData
	MethodUpdateCustomData = 0xD

	// MethodFindByNameRegex is the method ID for method FindByNameRegex
	MethodFindByNameRegex = 0xE

	// MethodUpdateAccountExpiryDate is the method ID for method UpdateAccountExpiryDate
	MethodUpdateAccountExpiryDate = 0xF

	// MethodUpdateAccountEffectiveDate is the method ID for method UpdateAccountEffectiveDate
	MethodUpdateAccountEffectiveDate = 0x10

	// MethodUpdateStatus is the method ID for method UpdateStatus
	MethodUpdateStatus = 0x11

	// MethodGetStatus is the method ID for method GetStatus
	MethodGetStatus = 0x12

	// MethodGetLastConnectionStats is the method ID for method GetLastConnectionStats
	MethodGetLastConnectionStats = 0x13

	// MethodResetPassword is the method ID for method ResetPassword
	MethodResetPassword = 0x14

	// MethodCreateAccountWithCustomData is the method ID for method CreateAccountWithCustomData
	MethodCreateAccountWithCustomData = 0x15

	// MethodRetrieveAccount is the method ID for method RetrieveAccount
	MethodRetrieveAccount = 0x16

	// MethodUpdateAccount is the method ID for method UpdateAccount
	MethodUpdateAccount = 0x17

	// MethodChangePasswordByGuest is the method ID for method ChangePasswordByGuest
	MethodChangePasswordByGuest = 0x18

	// MethodFindByNameLike is the method ID for method FindByNameLike
	MethodFindByNameLike = 0x19

	// MethodCustomCreateAccount is the method ID for method CustomCreateAccount
	MethodCustomCreateAccount = 0x1A

	// MethodNintendoCreateAccount is the method ID for method NintendoCreateAccount
	MethodNintendoCreateAccount = 0x1B

	// MethodLookupOrCreateAccount is the method ID for method LookupOrCreateAccount
	MethodLookupOrCreateAccount = 0x1C

	// MethodDisconnectPrincipal is the method ID for method DisconnectPrincipal
	MethodDisconnectPrincipal = 0x1D

	// MethodDisconnectAllPrincipals is the method ID for method DisconnectAllPrincipals
	MethodDisconnectAllPrincipals = 0x1E
)

// Protocol stores all the RMC method handlers for the Account Management protocol and listens for requests
type Protocol struct {
	server                      nex.ServerInterface
	CreateAccount               func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) (*nex.RMCMessage, uint32)
	DeleteAccount               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	DisableAccount              func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtUntil *nex.DateTime, strMessage string) (*nex.RMCMessage, uint32)
	ChangePassword              func(err error, packet nex.PacketInterface, callID uint32, strNewKey string) (*nex.RMCMessage, uint32)
	TestCapability              func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) (*nex.RMCMessage, uint32)
	GetName                     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetAccountData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetPrivateData              func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	GetPublicData               func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetMultiplePublicData       func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []*nex.PID) (*nex.RMCMessage, uint32)
	UpdateAccountName           func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)
	UpdateAccountEmail          func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)
	UpdateCustomData            func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	FindByNameRegex             func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	UpdateAccountExpiryDate     func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtExpiry *nex.DateTime, strExpiredMessage string) (*nex.RMCMessage, uint32)
	UpdateAccountEffectiveDate  func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) (*nex.RMCMessage, uint32)
	UpdateStatus                func(err error, packet nex.PacketInterface, callID uint32, strStatus string) (*nex.RMCMessage, uint32)
	GetStatus                   func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	GetLastConnectionStats      func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	ResetPassword               func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	CreateAccountWithCustomData func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	RetrieveAccount             func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
	UpdateAccount               func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)
	ChangePasswordByGuest       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) (*nex.RMCMessage, uint32)
	FindByNameLike              func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)
	CustomCreateAccount         func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	NintendoCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	LookupOrCreateAccount       func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)
	DisconnectPrincipal         func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)
	DisconnectAllPrincipals     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)
}

// Interface implements the methods present on the Account Management Protocol struct
type Interface interface {
	Server() nex.ServerInterface
	SetServer(server nex.ServerInterface)
	SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) (*nex.RMCMessage, uint32))
	SetHandlerDeleteAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerDisableAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtUntil *nex.DateTime, strMessage string) (*nex.RMCMessage, uint32))
	SetHandlerChangePassword(handler func(err error, packet nex.PacketInterface, callID uint32, strNewKey string) (*nex.RMCMessage, uint32))
	SetHandlerTestCapability(handler func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetAccountData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetPrivateData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerGetPublicData(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetMultiplePublicData(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []*nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerUpdateAccountName(handler func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateAccountEmail(handler func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerFindByNameRegex(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerUpdateAccountExpiryDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtExpiry *nex.DateTime, strExpiredMessage string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateAccountEffectiveDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) (*nex.RMCMessage, uint32))
	SetHandlerUpdateStatus(handler func(err error, packet nex.PacketInterface, callID uint32, strStatus string) (*nex.RMCMessage, uint32))
	SetHandlerGetStatus(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerGetLastConnectionStats(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerResetPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerCreateAccountWithCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerRetrieveAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
	SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerChangePasswordByGuest(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) (*nex.RMCMessage, uint32))
	SetHandlerFindByNameLike(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32))
	SetHandlerCustomCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerNintendoCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerLookupOrCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32))
	SetHandlerDisconnectPrincipal(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32))
	SetHandlerDisconnectAllPrincipals(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32))
}

// Server returns the server implementing the protocol
func (protocol *Protocol) Server() nex.ServerInterface {
	return protocol.server
}

// SetServer sets the server implementing the protocol
func (protocol *Protocol) SetServer(server nex.ServerInterface) {
	protocol.server = server
}

// SetHandlerCreateAccount sets the handler for the CreateAccount method
func (protocol *Protocol) SetHandlerCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string) (*nex.RMCMessage, uint32)) {
	protocol.CreateAccount = handler
}

// SetHandlerDeleteAccount sets the handler for the DeleteAccount method
func (protocol *Protocol) SetHandlerDeleteAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.DeleteAccount = handler
}

// SetHandlerDisableAccount sets the handler for the DisableAccount method
func (protocol *Protocol) SetHandlerDisableAccount(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtUntil *nex.DateTime, strMessage string) (*nex.RMCMessage, uint32)) {
	protocol.DisableAccount = handler
}

// SetHandlerChangePassword sets the handler for the ChangePassword method
func (protocol *Protocol) SetHandlerChangePassword(handler func(err error, packet nex.PacketInterface, callID uint32, strNewKey string) (*nex.RMCMessage, uint32)) {
	protocol.ChangePassword = handler
}

// SetHandlerTestCapability sets the handler for the TestCapability method
func (protocol *Protocol) SetHandlerTestCapability(handler func(err error, packet nex.PacketInterface, callID uint32, uiCapability uint32) (*nex.RMCMessage, uint32)) {
	protocol.TestCapability = handler
}

// SetHandlerGetName sets the handler for the GetName method
func (protocol *Protocol) SetHandlerGetName(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetName = handler
}

// SetHandlerGetAccountData sets the handler for the GetAccountData method
func (protocol *Protocol) SetHandlerGetAccountData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetAccountData = handler
}

// SetHandlerGetPrivateData sets the handler for the GetPrivateData method
func (protocol *Protocol) SetHandlerGetPrivateData(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.GetPrivateData = handler
}

// SetHandlerGetPublicData sets the handler for the GetPublicData method
func (protocol *Protocol) SetHandlerGetPublicData(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetPublicData = handler
}

// SetHandlerGetMultiplePublicData sets the handler for the GetMultiplePublicData method
func (protocol *Protocol) SetHandlerGetMultiplePublicData(handler func(err error, packet nex.PacketInterface, callID uint32, lstPrincipals []*nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetMultiplePublicData = handler
}

// SetHandlerUpdateAccountName sets the handler for the UpdateAccountName method
func (protocol *Protocol) SetHandlerUpdateAccountName(handler func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateAccountName = handler
}

// SetHandlerUpdateAccountEmail sets the handler for the UpdateAccountEmail method
func (protocol *Protocol) SetHandlerUpdateAccountEmail(handler func(err error, packet nex.PacketInterface, callID uint32, strName string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateAccountEmail = handler
}

// SetHandlerUpdateCustomData sets the handler for the UpdateCustomData method
func (protocol *Protocol) SetHandlerUpdateCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.UpdateCustomData = handler
}

// SetHandlerFindByNameRegex sets the handler for the FindByNameRegex method
func (protocol *Protocol) SetHandlerFindByNameRegex(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strRegex string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.FindByNameRegex = handler
}

// SetHandlerUpdateAccountExpiryDate sets the handler for the UpdateAccountExpiryDate method
func (protocol *Protocol) SetHandlerUpdateAccountExpiryDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtExpiry *nex.DateTime, strExpiredMessage string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateAccountExpiryDate = handler
}

// SetHandlerUpdateAccountEffectiveDate sets the handler for the UpdateAccountEffectiveDate method
func (protocol *Protocol) SetHandlerUpdateAccountEffectiveDate(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID, dtEffectiveFrom *nex.DateTime, strNotEffectiveMessage string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateAccountEffectiveDate = handler
}

// SetHandlerUpdateStatus sets the handler for the UpdateStatus method
func (protocol *Protocol) SetHandlerUpdateStatus(handler func(err error, packet nex.PacketInterface, callID uint32, strStatus string) (*nex.RMCMessage, uint32)) {
	protocol.UpdateStatus = handler
}

// SetHandlerGetStatus sets the handler for the GetStatus method
func (protocol *Protocol) SetHandlerGetStatus(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetStatus = handler
}

// SetHandlerGetLastConnectionStats sets the handler for the GetLastConnectionStats method
func (protocol *Protocol) SetHandlerGetLastConnectionStats(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.GetLastConnectionStats = handler
}

// SetHandlerResetPassword sets the handler for the ResetPassword method
func (protocol *Protocol) SetHandlerResetPassword(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.ResetPassword = handler
}

// SetHandlerCreateAccountWithCustomData sets the handler for the CreateAccountWithCustomData method
func (protocol *Protocol) SetHandlerCreateAccountWithCustomData(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.CreateAccountWithCustomData = handler
}

// SetHandlerRetrieveAccount sets the handler for the RetrieveAccount method
func (protocol *Protocol) SetHandlerRetrieveAccount(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.RetrieveAccount = handler
}

// SetHandlerUpdateAccount sets the handler for the UpdateAccount method
func (protocol *Protocol) SetHandlerUpdateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strKey string, strEmail string, oPublicData *nex.DataHolder, oPrivateData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.UpdateAccount = handler
}

// SetHandlerChangePasswordByGuest sets the handler for the ChangePasswordByGuest method
func (protocol *Protocol) SetHandlerChangePasswordByGuest(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, strEmail string) (*nex.RMCMessage, uint32)) {
	protocol.ChangePasswordByGuest = handler
}

// SetHandlerFindByNameLike sets the handler for the FindByNameLike method
func (protocol *Protocol) SetHandlerFindByNameLike(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroups uint32, strLike string, resultRange *nex.ResultRange) (*nex.RMCMessage, uint32)) {
	protocol.FindByNameLike = handler
}

// SetHandlerCustomCreateAccount sets the handler for the CustomCreateAccount method
func (protocol *Protocol) SetHandlerCustomCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.CustomCreateAccount = handler
}

// SetHandlerNintendoCreateAccount sets the handler for the NintendoCreateAccount method
func (protocol *Protocol) SetHandlerNintendoCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.NintendoCreateAccount = handler
}

// SetHandlerLookupOrCreateAccount sets the handler for the LookupOrCreateAccount method
func (protocol *Protocol) SetHandlerLookupOrCreateAccount(handler func(err error, packet nex.PacketInterface, callID uint32, strPrincipalName string, strKey string, uiGroups uint32, strEmail string, oAuthData *nex.DataHolder) (*nex.RMCMessage, uint32)) {
	protocol.LookupOrCreateAccount = handler
}

// SetHandlerDisconnectPrincipal sets the handler for the DisconnectPrincipal method
func (protocol *Protocol) SetHandlerDisconnectPrincipal(handler func(err error, packet nex.PacketInterface, callID uint32, idPrincipal *nex.PID) (*nex.RMCMessage, uint32)) {
	protocol.DisconnectPrincipal = handler
}

// SetHandlerDisconnectAllPrincipals sets the handler for the DisconnectAllPrincipals method
func (protocol *Protocol) SetHandlerDisconnectAllPrincipals(handler func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, uint32)) {
	protocol.DisconnectAllPrincipals = handler
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.server.OnData(func(packet nex.PacketInterface) {
		message := packet.RMCMessage()

		if message.IsRequest && message.ProtocolID == ProtocolID {
			protocol.HandlePacket(packet)
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCMessage()

	switch request.MethodID {
	case MethodCreateAccount:
		protocol.handleCreateAccount(packet)
	case MethodDeleteAccount:
		protocol.handleDeleteAccount(packet)
	case MethodDisableAccount:
		protocol.handleDisableAccount(packet)
	case MethodChangePassword:
		protocol.handleChangePassword(packet)
	case MethodTestCapability:
		protocol.handleTestCapability(packet)
	case MethodGetName:
		protocol.handleGetName(packet)
	case MethodGetAccountData:
		protocol.handleGetAccountData(packet)
	case MethodGetPrivateData:
		protocol.handleGetPrivateData(packet)
	case MethodGetPublicData:
		protocol.handleGetPublicData(packet)
	case MethodGetMultiplePublicData:
		protocol.handleGetMultiplePublicData(packet)
	case MethodUpdateAccountName:
		protocol.handleUpdateAccountName(packet)
	case MethodUpdateAccountEmail:
		protocol.handleUpdateAccountEmail(packet)
	case MethodUpdateCustomData:
		protocol.handleUpdateCustomData(packet)
	case MethodFindByNameRegex:
		protocol.handleFindByNameRegex(packet)
	case MethodUpdateAccountExpiryDate:
		protocol.handleUpdateAccountExpiryDate(packet)
	case MethodUpdateAccountEffectiveDate:
		protocol.handleUpdateAccountEffectiveDate(packet)
	case MethodUpdateStatus:
		protocol.handleUpdateStatus(packet)
	case MethodGetStatus:
		protocol.handleGetStatus(packet)
	case MethodGetLastConnectionStats:
		protocol.handleGetLastConnectionStats(packet)
	case MethodResetPassword:
		protocol.handleResetPassword(packet)
	case MethodCreateAccountWithCustomData:
		protocol.handleCreateAccountWithCustomData(packet)
	case MethodRetrieveAccount:
		protocol.handleRetrieveAccount(packet)
	case MethodUpdateAccount:
		protocol.handleUpdateAccount(packet)
	case MethodChangePasswordByGuest:
		protocol.handleChangePasswordByGuest(packet)
	case MethodFindByNameLike:
		protocol.handleFindByNameLike(packet)
	case MethodCustomCreateAccount:
		protocol.handleCustomCreateAccount(packet)
	case MethodNintendoCreateAccount:
		protocol.handleNintendoCreateAccount(packet)
	case MethodLookupOrCreateAccount:
		protocol.handleLookupOrCreateAccount(packet)
	case MethodDisconnectPrincipal:
		protocol.handleDisconnectPrincipal(packet)
	case MethodDisconnectAllPrincipals:
		protocol.handleDisconnectAllPrincipals(packet)
	default:
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported AccountManagement method ID: %#v\n", request.MethodID)
	}
}

// NewProtocol returns a new Account Management protocol
func NewProtocol(server nex.ServerInterface) *Protocol {
	protocol := &Protocol{server: server}

	protocol.Setup()

	return protocol
}
