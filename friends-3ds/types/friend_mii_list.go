// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendMiiList is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendMiiList struct {
	types.Structure
	*types.Data
	Unknown1 *types.PrimitiveU32
	MiiList  *MiiList
	Unknown2 *types.DateTime
}

// WriteTo writes the Mii to the given writable
func (friendMiiList *FriendMiiList) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendMiiList.Unknown1.WriteTo(contentWritable)
	friendMiiList.MiiList.WriteTo(contentWritable)
	friendMiiList.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendMiiList.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendMiiList
func (friendMiiList *FriendMiiList) Copy() types.RVType {
	copied := NewFriendMiiList()

	copied.StructureVersion = friendMiiList.StructureVersion

	copied.Data = friendMiiList.Data.Copy().(*types.Data)

	copied.Unknown1 = friendMiiList.Unknown1
	copied.MiiList = friendMiiList.MiiList.Copy().(*MiiList)
	copied.Unknown2 = friendMiiList.Unknown2.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendMiiList *FriendMiiList) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendMiiList); !ok {
		return false
	}

	other := o.(*FriendMiiList)

	if friendMiiList.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendMiiList.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendMiiList.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !friendMiiList.MiiList.Equals(other.MiiList) {
		return false
	}

	if !friendMiiList.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendMiiList *FriendMiiList) String() string {
	return friendMiiList.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendMiiList *FriendMiiList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendMiiList{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendMiiList.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, friendMiiList.Unknown1))

	if friendMiiList.MiiList != nil {
		b.WriteString(fmt.Sprintf("%sMiiList: %s,\n", indentationValues, friendMiiList.MiiList.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMiiList: nil,\n", indentationValues))
	}

	if friendMiiList.Unknown2 != nil {
		b.WriteString(fmt.Sprintf("%sUnknown2: %s\n", indentationValues, friendMiiList.Unknown2.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown2: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendMiiList returns a new FriendMiiList
func NewFriendMiiList() *FriendMiiList {
	return &FriendMiiList{}
}
