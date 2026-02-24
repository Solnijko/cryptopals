package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"sort"
	"strings"
)

func hamming(a, b []byte) int {
	// Hamming distance is strictly defined only for strings of the same length.
	if len(a) != len(b) {
		panic("inputs must have equal length")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		b1 := a[i]
		b2 := b[i]
		for j := 0; j < 8; j++ {
			mask := byte(1 << uint(j))
			if (b1 & mask) != (b2 & mask) {
				diff++
			}
		}
	}
	return diff
}

type candidate struct {
	keysize int
	score   float64
}

func findKeysize(ciphertext []byte) []candidate {
	var result []candidate
	for keysize := 2; keysize <= 40; keysize++ {
		if len(ciphertext) < keysize*4 {
			continue
		}

		b0 := ciphertext[0:keysize]
		b1 := ciphertext[keysize : keysize*2]
		b2 := ciphertext[keysize*2 : keysize*3]
		b3 := ciphertext[keysize*3 : keysize*4]

		blocks := [][]byte{b0, b1, b2, b3}
		h := 0.0
		count := 0
		for i := 0; i < len(blocks); i++ {
			for j := i + 1; j < len(blocks); j++ {
				h += float64(hamming(blocks[i], blocks[j])) / float64(keysize)
				count++
			}
		}
		avg := h / float64(count)
		candidate := candidate{keysize, avg}
		result = append(result, candidate)
	}
	return result
}

func breakCiphertext(ciphertext []byte, keysize int) [][]byte {
	var result [][]byte
	for sliceRange := 0; sliceRange+keysize <= len(ciphertext); sliceRange += keysize {
		result = append(result, ciphertext[sliceRange:sliceRange+keysize])
	}
	return result
}

func transpose(brokenCiphertext [][]byte) [][]byte {
	var result [][]byte
	for column := 0; column < len(brokenCiphertext[0]); column += 1 {
		var block []byte
		for row := 0; row < len(brokenCiphertext); row += 1 {
			block = append(block, brokenCiphertext[row][column])
		}
		result = append(result, block)
	}
	return result
}

func calculateKeyReadability(block []byte, k int) int {
	var result []byte
	key := byte(k)
	for i := 0; i < len(block); i++ {
		result = append(result, block[i]^key)
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

	return readabilityPoints
}

func repeatingKeyXOR(ciphertext []byte, key string) string {
	var result []byte
	for i := range ciphertext {
		result = append(result, ciphertext[i]^key[i%len(key)])
	}
	return string(result)
}

func main() {
	b64File, _ := os.ReadFile("set-1-basics/break-repeating-key-xor/6.txt")
	ciphertext, _ := base64.StdEncoding.DecodeString(string(b64File))

	candidates := findKeysize(ciphertext)
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score < candidates[j].score
	})
	keysize := candidates[0].keysize

	brokenCiphertext := breakCiphertext(ciphertext, keysize)
	transposedCiphertext := transpose(brokenCiphertext)

	var fullkey []byte
	for block := range transposedCiphertext {
		bestScore := 0
		bestKey := 0
		for k := 0; k < 256; k++ {
			score := calculateKeyReadability(transposedCiphertext[block], k)
			if score > bestScore {
				bestScore = score
				bestKey = k
			}
		}
		fullkey = append(fullkey, byte(bestKey))
	}
	fmt.Println("---", string(fullkey), "---")
	fmt.Println(repeatingKeyXOR(ciphertext, string(fullkey)))
}
