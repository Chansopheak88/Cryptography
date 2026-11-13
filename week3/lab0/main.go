package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/sha3"
)

// compute the hash of a string using the chosen algorithm
func computeHash(algo string, text string) string {
	data := []byte(text)

	if algo == "MD5" {
		sum := md5.Sum(data)
		return hex.EncodeToString(sum[:])
	} else if algo == "SHA1" {
		sum := sha1.Sum(data)
		return hex.EncodeToString(sum[:])
	} else if algo == "SHA256" {
		sum := sha256.Sum256(data)
		return hex.EncodeToString(sum[:])
	} else if algo == "SHA512" {
		sum := sha512.Sum512(data)
		return hex.EncodeToString(sum[:])
	} else if algo == "SHA3" {
		sum := sha3.Sum256(data)
		return hex.EncodeToString(sum[:])
	}

	return ""
}

// hash both inputs with every algorithm and compare the results
func proofMe(txt1 string, txt2 string) {
	algos := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3"}

	for i := 0; i < len(algos); i++ {
		outputA := computeHash(algos[i], txt1)
		outputB := computeHash(algos[i], txt2)

		fmt.Printf("\nHash (%s)\n", algos[i])
		fmt.Println("  Output A =", outputA)
		fmt.Println("  Output B =", outputB)

		if outputA == outputB {
			fmt.Println("  => Match!")
		} else {
			fmt.Println("  => No Match!")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("========Princess + Hashing Program========")

	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)

	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)

	proofMe(input1, input2)
}