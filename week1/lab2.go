//go:build ignore

package main

import "fmt"

func main(){
	var x=10

	// && operator
	fmt.Print("&& Operator\n")
	fmt.Println(x>5 && x==10)	// true && true
	fmt.Println(x>10 && x==20)	// false && false
	fmt.Println(x>8 && x>10)	// true && false
	fmt.Println(x==12 && x>20)	// false && false

	// || operator
	fmt.Print("|| Operator\n")
	fmt.Println(x>5 || x==10)	// true || true
	fmt.Println(x>10 || x==20)	// false || false
	fmt.Println(x>8 || x>10)	// true || false
	fmt.Println(x==12 || x>20)	// false || false

	// ! operator 
	fmt.Print("! Operator with &&\n")
	fmt.Println(!(x>5 && x==10))	// true && true
	fmt.Println(!(x>10 && x==20))	// false && false
	fmt.Println(!(x>8 && x>10))	// true && false
	fmt.Println(!(x==12 && x>20))	// false && false

	fmt.Print("! Operator with || \n")
	fmt.Println(!(x>5 || x==10))	// true || true
	fmt.Println(!(x>10 || x==20))	// false || false
	fmt.Println(!(x>8 || x>10))		// true || false
	fmt.Println(!(x==12 || x>20))	// false || false
	
}