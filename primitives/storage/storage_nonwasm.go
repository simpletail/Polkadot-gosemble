//go:build nonwasmenv

package storage

import (
	"bytes"
	sc "github.com/LimeChain/goscale"
)

func ChangesRoot(parent_hash int64) int64 {
	panic("not implemented")
}

func Clear(key []byte) {
	panic("not implemented")
}

func ClearPrefix(key []byte, limit []byte) {
	panic("not implemented")
}

func Exists(key []byte) int32 {
	panic("not implemented")
}

func Get(key []byte) sc.Option[sc.Sequence[sc.U8]] {
	panic("not implemented")
}

func GetDecode[T sc.Encodable](key []byte, decodeFunc func(buffer *bytes.Buffer) T) T {
	panic("not implemented")
}

func NextKey(key int64) int64 {
	panic("not implemented")
}

func Read(key int64, value_out int64, offset int32) int64 {
	panic("not implemented")
}

func Root(key int32) []byte {
	panic("not implemented")
}

func Set(key []byte, value []byte) {
	panic("not implemented")
}

func TakeBytes(key []byte) []byte {
	panic("not implemented")
}

func TakeDecode[T sc.Encodable](key []byte, decodeFunc func(buffer *bytes.Buffer) T) T {
	panic("not implemented")
}
