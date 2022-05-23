package tools

import "hash/crc64"

var crc64Table *crc64.Table

func init() {
	crc64Table = crc64.MakeTable(crc64.ECMA)
}
func GetPasswordHash(password string) uint64 {
	return crc64.Checksum([]byte(password), crc64Table)
}
