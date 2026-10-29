//go:build ignore
package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

// ---------- Calculator functions ----------
func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul(a, b float64) float64 {
	return a * b
}

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func mod(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("modulo by zero")
	}
	return math.Mod(a, b), nil
}

// ---------- Input helper ----------
// Keeps asking until the user types a valid number
func readNumber(prompt string) float64 {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		num, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err == nil {
			return num
		}
		fmt.Println("Error: please enter a valid number")
	}
}

func main() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		if choice == "6" {
			fmt.Println("Goodbye!")
			break
		}

		// Check for invalid operator before asking for numbers
		if choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
			fmt.Println("Error: invalid option, choose 1-6")
			continue
		}

		a := readNumber("Enter a: ")
		b := readNumber("Enter b: ")

		switch choice {
		case "1":
			fmt.Printf("Result: %g + %g = %g\n", a, b, add(a, b))
		case "2":
			fmt.Printf("Result: %g - %g = %g\n", a, b, sub(a, b))
		case "3":
			fmt.Printf("Result: %g * %g = %g\n", a, b, mul(a, b))
		case "4":
			result, err := div(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %g / %g = %g\n", a, b, result)
			}
		case "5":
			result, err := mod(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %g %% %g = %g\n", a, b, result)
			}
		}
	}
}