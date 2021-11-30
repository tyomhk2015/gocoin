package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// TIP: To receive any type of data, use interface{} as the argument.
func ToBytes(i interface{}) []byte {
	// Turn the data of block into bytes.
	// 'gob' package is used. https://pkg.go.dev/encoding/gob
	// 'gob' encode/decode bytes.
	var anyBuffer bytes.Buffer            // Set writer
	encoder := gob.NewEncoder(&anyBuffer) // Set the encoder with the writer.
	HandleErr(encoder.Encode(i))          // Encode the block with the encoder and return the bytes to blockBuffer.
	return anyBuffer.Bytes()
}

// Find data that matches the 'Hash'(arg: data).
// Decode the returned data and put it in the given interface, a Block(struct) or Blockchain(struct).
func FromBytes(i interface{}, data []byte) {
	// The decoder, with a target to decode, initiates decoding
	// and stores the decoded data to 'i'.
	decoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(decoder.Decode(i))
}
