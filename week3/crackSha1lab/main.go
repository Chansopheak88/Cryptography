package main

import (
	"fmt"

	"crackSha1lab/utils/crack"
)

func main() {
	wordlistPath := "wordlist/nord_vpn.txt"
	verboseLogPath := "verbose.txt"
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	fmt.Println("========SHA1 Password Cracker========")
	fmt.Println("Target hash:", targetHash)
	fmt.Println("Wordlist:", wordlistPath)
	fmt.Println("Cracking... (this may take a moment)")

	password, found := crack.CrackSHA1(wordlistPath, targetHash, verboseLogPath)

	if found {
		fmt.Println("\nSUCCESS! Password cracked:")
		fmt.Println("Password =", password)
	} else {
		fmt.Println("\nFAILED. Password not found in wordlist.")
	}

	fmt.Println("\nVerbose log saved to:", verboseLogPath)
}