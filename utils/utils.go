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
