package main

import (
	"fmt"

	"crackSha512lab/utils/crack"
)

func main() {
	wordlistPath := "wordlist/nord_vpn.txt"
	verboseLogPath := "verbose.txt"
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	fmt.Println("========SHA512 Password Cracker========")
	fmt.Println("Target hash:", targetHash)
	fmt.Println("Wordlist:", wordlistPath)
	fmt.Println("Cracking... (this may take a moment)")

	password, found := crack.CrackSHA512(wordlistPath, targetHash, verboseLogPath)

	if found {
		fmt.Println("\nSUCCESS! Password cracked:")
		fmt.Println("Password =", password)
	} else {
		fmt.Println("\nFAILED. Password not found in wordlist.")
	}

	fmt.Println("\nVerbose log saved to:", verboseLogPath)
}