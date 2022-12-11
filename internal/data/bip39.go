package data

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func Hex(bytes []byte) (string, error) {
	blength := len(bytes)
	if blength < 1 {
		return "", errors.New("bip39.Hex - input should have at least one byte")
	}

	hexBytes := make([]byte, hex.EncodedLen(len(bytes)))
	hex.Encode(hexBytes, bytes)

	result := fmt.Sprintf("%s", hexBytes)
	log.Printf("bip39.Hex %s\n", result)
	return result, nil
}

func ToBinaryString(bytess []byte) (string, error) {
	blength := len(bytess)
	if blength < 1 {
		return "", errors.New("bip39.ToBinaryString - input should have at least one byte")
	}

	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.LittleEndian, bytess)
	if err != nil {
		return "", errors.New("bip39.ToBinaryString - error on write binary: " + err.Error())
	}

	var sbinary string
	for _, byte := range buffer.Bytes() {
		sbinary += fmt.Sprintf("%08b", byte)
	}

	log.Printf("bip39.ToBinaryString %d - %s\n", len(sbinary), sbinary)
	return sbinary, nil
}
