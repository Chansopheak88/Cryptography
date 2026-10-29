//go:build ignore

package main

import "fmt"

func main(){
	var a int 
	var b int 
	fmt.Print("Please enter integer 1:")
	fmt.Scanln(&a)
	fmt.Print("Please enter integer 2:")
	fmt.Scanln(&b)

	// = operator
	var x int = a
	var y int = b
	fmt.Printf("The value of x and y are %v and %v",x,y)

	// += operator
	x+=a
	y+=b
	fmt.Printf("\nThe value of x and y are %v and %v",x,y)

	// -= operator
	x-=a
	y-=b
	fmt.Printf("\nThe value of x and y are %v and %v",x,y)

	// *= operator
	x*=a
	y*=b
	fmt.Printf("\nThe value of x and y are %v and %v",x,y)

	// /= operator
	x/=a
	y/=b
	fmt.Printf("\nThe value of x and y are %v and %v",x,y)

	// %= operator
	x%=2
	y%=2
	fmt.Printf("\nThe value of x and y are %v and %v",x,y)

	
}