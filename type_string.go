// Code generated by "stringer -type=Type"; DO NOT EDIT

package j

import "fmt"

const _Type_name = "InvalidTypeObjectTypeArrayTypeBoolTypeNumberTypeStringTypeNullType"

var _Type_index = [...]uint8{0, 11, 21, 30, 38, 48, 58, 66}

func (i Type) String() string {
	if i >= Type(len(_Type_index)-1) {
		return fmt.Sprintf("Type(%d)", i)
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
