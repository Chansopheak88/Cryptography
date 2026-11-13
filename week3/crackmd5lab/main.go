package main

import (
	"fmt"

	"crackmd5lab/utils/crack"
)

func main() {
	wordlistPath := "wordlist/nord_vpn.txt"
	verboseLogPath := "verbose.txt"
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	fmt.Println("========MD5 Password Cracker========")
	fmt.Println("Target hash:", targetHash)
	fmt.Println("Wordlist:", wordlistPath)
	fmt.Println("Cracking... (this may take a moment)")

	password, found := crack.CrackMD5(wordlistPath, targetHash, verboseLogPath)

	if found {
		fmt.Println("\nSUCCESS! Password cracked:")
		fmt.Println("Password =", password)
	} else {
		fmt.Println("\nFAILED. Password not found in wordlist.")
	}

	fmt.Println("\nVerbose log saved to:", verboseLogPath)
}