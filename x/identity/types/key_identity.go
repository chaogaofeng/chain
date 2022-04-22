package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// IdentityKeyPrefix is the prefix to retrieve all Identity
	IdentityKeyPrefix = "Identity/value/"
)

// IdentityKey returns the store key to retrieve a Identity from the index fields
func IdentityKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
