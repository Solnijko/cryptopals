package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func ReverseHexXOR(hexString string, k int) []byte {
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		return nil
	}
	var result []byte
	key := byte(k)
	for i := 0; i < len(hexBytes); i++ {
		result = append(result, hexBytes[i]^key)
	}

	commonRunes := []string{"e", "t", "a", "o", "i", "n", " ", "s", "h", "r", "d", "l", "u"}
	readabilityPoints := 0

	for i := range commonRunes {
		if strings.Count(string(result), commonRunes[i]) > 1 {
			readabilityPoints++
		}
	}

	if readabilityPoints > 1 {
		return result
	} else {
		return nil
	}
}

func main() {
	for k := 0; k < 256; k++ {
		decipheredString := ReverseHexXOR("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736", k)
		if decipheredString != nil {
			fmt.Println(string(decipheredString))
		}
	}
}
