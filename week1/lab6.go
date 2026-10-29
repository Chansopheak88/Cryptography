//go:build ignore
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Single-byte key (the function from the lab)
func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}

// Repeating key: the key repeats along the message
func xorRepeat(text string, key string) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key[i%len(key)]
	}
	return string(result)
}

func readLine(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	message := readLine(reader, "Enter plaintext message: ")
	mode := readLine(reader, "Key type (1 = single byte, 2 = repeating key): ")

	var cipher, decrypted string

	if mode == "2" {
		key := readLine(reader, "Enter key (word): ")
		if key == "" {
			fmt.Println("Error: key cannot be empty")
			return
		}
		cipher = xorRepeat(message, key)
		decrypted = xorRepeat(cipher, key) // same function decrypts
	} else {
		key := readLine(reader, "Enter key (one character): ")
		if key == "" {
			fmt.Println("Error: key cannot be empty")
			return
		}
		cipher = xorEncrypt(message, key[0])
		decrypted = xorEncrypt(cipher, key[0]) // same function decrypts
	}

	fmt.Println("\n===== Results =====")
	fmt.Println("Plaintext  :", message)
	fmt.Println("Ciphertext :", hex.EncodeToString([]byte(cipher)), "(hex)")
	fmt.Println("Decrypted  :", decrypted)
}