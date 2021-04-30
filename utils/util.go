package utils

import "encoding/binary"

func GetEndian(isBigEndian bool) binary.ByteOrder {
	if isBigEndian {
		return binary.BigEndian
	} else {
		return binary.LittleEndian
	}
}
