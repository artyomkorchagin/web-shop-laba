package mssqlutils

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func ReverseUUID(rawUUID string) string {
	// Convert the UUID string to a byte slice
	bytes, err := hex.DecodeString(strings.ReplaceAll(rawUUID, "-", ""))
	if err != nil {
		return rawUUID // Return as-is if decoding fails
	}

	// Reverse the byte order for the first three segments
	reverseBytes := make([]byte, len(bytes))
	copy(reverseBytes[0:4], reverseBytesInPlace(bytes[0:4])) // Time-low
	copy(reverseBytes[4:6], reverseBytesInPlace(bytes[4:6])) // Time-mid
	copy(reverseBytes[6:8], reverseBytesInPlace(bytes[6:8])) // Time-high-and-version
	copy(reverseBytes[8:], bytes[8:])                        // Clock sequence and node (big-endian)

	// Convert the reversed byte slice back to a UUID string
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		hex.EncodeToString(reverseBytes[0:4]),
		hex.EncodeToString(reverseBytes[4:6]),
		hex.EncodeToString(reverseBytes[6:8]),
		hex.EncodeToString(reverseBytes[8:10]),
		hex.EncodeToString(reverseBytes[10:]),
	)
}

func reverseBytesInPlace(data []byte) []byte {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}
