//go:build ignore
package main

import "fmt"

// ---------- Bitwise functions ----------
func myAND(a, b int) int {
	return a & b
}

func myOR(a, b int) int {
	return a | b
}

func myXor(a, b int) int {
	return a ^ b
}

func myNOT(a, b int) (int, int) {
	return ^a, ^b
}

func myLeftShift(a int, n uint) int {
	return a << n
}

func myRightShift(a int, n uint) int {
	return a >> n
}

func main() {
	var a, b int

	// Get two integers from the user
	fmt.Print("Please enter integer 1: ")
	fmt.Scanln(&a)
	fmt.Print("Please enter integer 2: ")
	fmt.Scanln(&b)

	// ---------- Bitwise operations ----------
	fmt.Println("\n===== Bitwise Operations =====")
	fmt.Printf("a = %d (binary: %b)\n", a, a)
	fmt.Printf("b = %d (binary: %b)\n\n", b, b)

	fmt.Printf("a & b  = %d (binary: %b)\n", myAND(a, b), myAND(a, b))
	fmt.Printf("a | b  = %d (binary: %b)\n", myOR(a, b), myOR(a, b))
	fmt.Printf("a ^ b  = %d (binary: %b)\n", myXor(a, b), myXor(a, b))

	notA, notB := myNOT(a, b)
	fmt.Printf("^a     = %d\n", notA)
	fmt.Printf("^b     = %d\n", notB)

	// Shift both numbers by 1 position
	fmt.Printf("a << 1 = %d (binary: %b)\n", myLeftShift(a, 1), myLeftShift(a, 1))
	fmt.Printf("a >> 1 = %d (binary: %b)\n", myRightShift(a, 1), myRightShift(a, 1))
	fmt.Printf("b << 1 = %d (binary: %b)\n", myLeftShift(b, 1), myLeftShift(b, 1))
	fmt.Printf("b >> 1 = %d (binary: %b)\n", myRightShift(b, 1), myRightShift(b, 1))

	// ---------- Assignment operators ----------
	fmt.Println("\n===== Assignment Operators =====")
	x := a
	fmt.Printf("x = a   -> x = %d\n", x)

	x += b
	fmt.Printf("x += b  -> x = %d\n", x)

	x -= b
	fmt.Printf("x -= b  -> x = %d\n", x)

	x *= b
	fmt.Printf("x *= b  -> x = %d\n", x)

	if b != 0 { // avoid division by zero
		x /= b
		fmt.Printf("x /= b  -> x = %d\n", x)
		x %= b
		fmt.Printf("x %%= b  -> x = %d\n", x)
	} else {
		fmt.Println("x /= b and x %= b skipped (b is 0)")
	}

	// Bitwise assignment operators
	y := a
	y &= b
	fmt.Printf("y &= b  -> y = %d\n", y)
	y = a
	y |= b
	fmt.Printf("y |= b  -> y = %d\n", y)
	y = a
	y ^= b
	fmt.Printf("y ^= b  -> y = %d\n", y)
	y = a
	y <<= 1
	fmt.Printf("y <<= 1 -> y = %d\n", y)
	y >>= 1
	fmt.Printf("y >>= 1 -> y = %d\n", y)
}