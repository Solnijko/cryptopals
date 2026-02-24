package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func ReverseHexXOR(hexString string, k int) (string, int) {
	hexBytes, _ := hex.DecodeString(hexString)
	var result []byte
	key := byte(k)
	for i := 0; i < len(hexBytes); i++ {
		result = append(result, hexBytes[i]^key)
	}

	mostCommonRunes := []string{"e", "t", "a", "o", "i", "n"}
	otherCommonRunes := []string{"s", "h", "r", "d", "l", "u"}
	readabilityPoints := 0

	byteStr := string(result)
	if strings.Count(byteStr, " ") > 1 {
		readabilityPoints += 3
	}
	for i := range mostCommonRunes {
		if strings.Count(byteStr, mostCommonRunes[i]) > 1 {
			readabilityPoints += 2
		}
		if strings.Count(byteStr, otherCommonRunes[i]) > 1 {
			readabilityPoints += 1
		}
	}
	return string(result), readabilityPoints

}

func main() {
	file, _ := os.Open("set-1-basics/single-character-xor/4.txt")
	scanner := bufio.NewScanner(file)
	bestScore := 0
	bestText := ""
	for scanner.Scan() {
		for k := 0; k < 256; k++ {
			decipheredString, score := ReverseHexXOR(scanner.Text(), k)
			if score > bestScore {
				bestScore = score
				bestText = decipheredString
			}
		}
	}
	fmt.Println(bestText)
}
