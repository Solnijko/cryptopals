package main

import (
	"encoding/hex"
	"fmt"
)

func FixedXOR(hexStr0 string, hexStr1 string) string {
	hexBytes0, err := hex.DecodeString(hexStr0)
	if err != nil {
		return ""
	}

	hexBytes1, err := hex.DecodeString(hexStr1)
	if err != nil {
		return ""
	}

	var result []byte
	var xor byte
	for i := 0; i < len(hexBytes0); i++ {
		xor = hexBytes0[i] ^ hexBytes1[i]
		result = append(result, xor)
	}
	return hex.EncodeToString(result)
}

func main() {
	fmt.Println(FixedXOR("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965"))
}
