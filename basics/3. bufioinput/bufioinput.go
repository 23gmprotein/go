package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {

	soumyascanner := bufio.NewScanner(os.Stdin)

	/*
	soumyascanner is a variable of type Scanner pointing to an object of type Scanner

	Scanner object is getting created as a result of bufio.NewScanner()

	this scanner object is configured to read from (os.Stdin)

	When you create a Scanner object with bufio.NewScanner(os.Stdin), you're essentially setting up a scanner 
	that's configured to read from the standard input (os.Stdin). However, just creating the Scanner object doesn't 
	automatically start reading input. It's just prepared to do so.
	*/

	fmt.Print("Enter Something:")
	soumyascanner.Scan()

	/* The Scan() method is what actually initiates the reading process. When you call soumyascanner.Scan(), 
	it waits for input from the user (in this case, from the standard input).
	Once the user enters something and presses Enter, the Scan() method reads that input and prepares it to be retrieved. */


	input := soumyascanner.Text() /* Text() function is used to retrieve the data that was scanned and store it in the input variable */

	fmt.Printf("You entered %q\n", input)

}
