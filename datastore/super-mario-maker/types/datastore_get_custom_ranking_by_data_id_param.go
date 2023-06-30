// Package datastore_super_mario_maker_types implements all the types used by the DataStore Super Mario Maker protocol
package datastore_super_mario_maker_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreGetCustomRankingByDataIDParam is sent in the GetCustomRankingByDataID method
type DataStoreGetCustomRankingByDataIDParam struct {
	nex.Structure
	ApplicationID uint32
	DataIDList    []uint64
	ResultOption  uint8
}

// ExtractFromStream extracts a DataStoreGetCustomRankingByDataIDParam structure from a stream
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreGetCustomRankingByDataIDParam.ApplicationID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ApplicationID. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIDParam.DataIDList, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.DataIDList. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIDParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIDParam.ResultOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCustomRankingByDataIDParam
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCustomRankingByDataIDParam()

	copied.ApplicationID = dataStoreGetCustomRankingByDataIDParam.ApplicationID
	copied.DataIDList = make([]uint64, len(dataStoreGetCustomRankingByDataIDParam.DataIDList))

	copy(copied.DataIDList, dataStoreGetCustomRankingByDataIDParam.DataIDList)

	copied.ResultOption = dataStoreGetCustomRankingByDataIDParam.ResultOption

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCustomRankingByDataIDParam)

	if dataStoreGetCustomRankingByDataIDParam.ApplicationID != other.ApplicationID {
		return false
	}

	if len(dataStoreGetCustomRankingByDataIDParam.DataIDList) != len(other.DataIDList) {
		return false
	}

	for i := 0; i < len(dataStoreGetCustomRankingByDataIDParam.DataIDList); i++ {
		if dataStoreGetCustomRankingByDataIDParam.DataIDList[i] != other.DataIDList[i] {
			return false
		}
	}

	if dataStoreGetCustomRankingByDataIDParam.ResultOption != other.ResultOption {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) String() string {
	return dataStoreGetCustomRankingByDataIDParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreGetCustomRankingByDataIDParam *DataStoreGetCustomRankingByDataIDParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreGetCustomRankingByDataIDParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sApplicationID: %d,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.ApplicationID))
	b.WriteString(fmt.Sprintf("%sDataIDList: %v,\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.DataIDList))
	b.WriteString(fmt.Sprintf("%sResultOption: %d\n", indentationValues, dataStoreGetCustomRankingByDataIDParam.ResultOption))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreGetCustomRankingByDataIDParam returns a new DataStoreGetCustomRankingByDataIDParam
func NewDataStoreGetCustomRankingByDataIDParam() *DataStoreGetCustomRankingByDataIDParam {
	return &DataStoreGetCustomRankingByDataIDParam{}
}
