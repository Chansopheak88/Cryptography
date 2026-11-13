package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

// hashSHA1 returns the hex SHA1 hash of a plain text string
func hashSHA1(text string) string {
	sum := sha1.Sum([]byte(text))
	return hex.EncodeToString(sum[:])
}

// CrackSHA1 reads the wordlist line by line, hashes each word with SHA1,
// and compares it to targetHash. It writes every attempt to the verbose
// log file as it goes. It returns the cracked password and true if found.
func CrackSHA1(wordlistPath string, targetHash string, verboseLogPath string) (string, bool) {
	wordlistFile, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist file:", err)
		return "", false
	}
	defer wordlistFile.Close()

	logFile, err := os.Create(verboseLogPath)
	if err != nil {
		fmt.Println("Error creating log file:", err)
		return "", false
	}
	defer logFile.Close()

	scanner := bufio.NewScanner(wordlistFile)

	attemptCount := 0
	found := false
	crackedPassword := ""

	for scanner.Scan() {
		word := scanner.Text()
		if word == "" {
			continue
		}

		attemptCount++
		wordHash := hashSHA1(word)

		logLine := fmt.Sprintf("Attempt #%d: word=%s  hash=%s\n", attemptCount, word, wordHash)
		fmt.Fprint(logFile, logLine)

		if wordHash == targetHash {
			found = true
			crackedPassword = word
			successLine := fmt.Sprintf("\n=== MATCH FOUND ===\nWord: %s\nHash: %s\nTotal attempts: %d\n", word, wordHash, attemptCount)
			fmt.Fprint(logFile, successLine)
			break
		}
	}

	return crackedPassword, found
}