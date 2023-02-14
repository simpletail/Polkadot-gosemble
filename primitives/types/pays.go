package types

import (
	"bytes"

	sc "github.com/LimeChain/goscale"
)

const (
	// Transactor will pay related fees.
	PaysYes = Pays(iota)

	// Transactor will NOT pay related fees.
	PaysNo
)

type Pays sc.U8

func (p Pays) Encode(buffer *bytes.Buffer) {
	switch p {
	case PaysYes:
		sc.U8(0).Encode(buffer)
	case PaysNo:
		sc.U8(1).Encode(buffer)
	default:
		panic("invalid Pays type")
	}
}

func DecodePays(buffer *bytes.Buffer) Pays {
	b := sc.DecodeU8(buffer)

	switch b {
	case 0:
		return PaysYes
	case 1:
		return PaysNo
	default:
		panic("invalid Pays type")
	}
}

func (p Pays) Bytes() []byte {
	return sc.EncodedBytes(p)
}
