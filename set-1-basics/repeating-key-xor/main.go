package main

import (
	"encoding/hex"
	"fmt"
)

func RepeatingKeyXOR(plaintextStr string, key string) string {
	byteStr := []byte(plaintextStr)
	key_i := 0
	var result []byte
	for i := range byteStr {
		result = append(result, byteStr[i]^key[i%len(key)])
		key_i += 1
	}
	return hex.EncodeToString(result)
}

func main() {
	stanza := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	fmt.Println(RepeatingKeyXOR(stanza, key))
}
