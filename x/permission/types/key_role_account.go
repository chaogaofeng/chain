package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RoleAccountKeyPrefix is the prefix to retrieve all RoleAccount
	RoleAccountKeyPrefix = "RoleAccount/value/"
)

// RoleAccountKey returns the store key to retrieve a RoleAccount from the index fields
func RoleAccountKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
