package infra

import (
	"hash/crc32"
	"hash/crc64"
)

var crc64Table = crc64.MakeTable(crc64.ECMA)
var crc32Table = crc32.MakeTable(crc32.IEEE)

// func GetPasswordHash(password string) uint64 {
// 	return crc64.Checksum([]byte(password), crc64Table)
// }

func GetPasswordHash(password string) uint32 {
	return crc32.ChecksumIEEE([]byte(password))
}
