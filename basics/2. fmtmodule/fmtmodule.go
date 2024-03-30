package main

import (
	"fmt"
	//"strconv"
)

func main() {

	var sjm string = fmt.Sprintf("This is the power of") // -> here sprintf is actually returning "wassup" as a new string
	fmt.Printf(" \n%s, Sprintf \n\n", sjm)

	var sjm1 int = 6969
	fmt.Printf("Hey %d is the number \n\n", sjm1) // -> here we have a format specifier

	var sjm2 string = fmt.Sprintf("Just tring out") // -> you cannot print a modified string with println, you have to use  printf
	var sjm3 string = "Sprintf"
	fmt.Println(sjm2 + sjm3)
	fmt.Printf("\n")


}
