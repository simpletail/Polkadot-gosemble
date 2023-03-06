package types

import (
	"bytes"
	"testing"

	sc "github.com/LimeChain/goscale"
	"github.com/stretchr/testify/assert"
)

func Test_EncodeExtrinsicSignature(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       ExtrinsicSignature
		expectation []byte
	}{
		{
			label: "Encode(ExtrinsicSignature())",
			input: ExtrinsicSignature{
				Signer:    NewMultiAddress(AccountId{U64: 1}),
				Signature: NewMultiSignature(NewEd25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
				Extra: SignedExtra{
					Era:   NewImmortalEra(),
					Nonce: 0,
					Fee:   0,
				},
			},
			expectation: []byte{0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x0, 0x0, 0x0},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assert.Equal(t, testExample.expectation, buffer.Bytes())
		})
	}
}

func Test_DecodeExtrinsicSignature(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation ExtrinsicSignature
	}{
		{
			label: "Decode(ExtrinsicSignature())",
			input: []byte{0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x0, 0x0, 0x0},
			expectation: ExtrinsicSignature{
				Signer:    NewMultiAddress(AccountId{U64: 1}),
				Signature: NewMultiSignature(NewEd25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
				Extra: SignedExtra{
					Era:   NewImmortalEra(),
					Nonce: 0,
					Fee:   0,
				},
			},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			s := DecodeExtrinsicSignature(buffer)

			assert.Equal(t, testExample.expectation.Signer, s.Signer)
		})
	}
}