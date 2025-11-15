package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
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

	if readabilityPoints > 4 {
		return result
	} else {
		return nil
	}
}

func main() {
	file, _ := os.Open("set-1-basics\\single-character-xor\\4.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for k := 0; k < 256; k++ {
			decipheredString := ReverseHexXOR(scanner.Text(), k)
			if decipheredString != nil {
				fmt.Println(string(decipheredString))
			}
		}
	}
}
