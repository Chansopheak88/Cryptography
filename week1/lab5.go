//go:build ignore
package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Each byte becomes 8 bits, e.g. 'A' -> 01000001
func toBinary(text string) string {
	var parts []string
	for i := 0; i < len(text); i++ {
		parts = append(parts, fmt.Sprintf("%08b", text[i]))
	}
	return strings.Join(parts, " ")
}

func toHex(text string) string {
	return hex.EncodeToString([]byte(text))
}

func toBase64(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	fmt.Println("\n===== Results =====")
	fmt.Println("Text   :", input)
	fmt.Println("Binary :", toBinary(input))
	fmt.Println("Hex    :", toHex(input))
	fmt.Println("Base64 :", toBase64(input))
}