package main

import "fmt" //displays messages and texts

func main() { // -> entrpoint of the program
	var bln bool
	fmt.Println(bln) // single quotes '' doesnt work, single quotations only work for charecters
}

/*

-> go is a compiled language

	so before we actually execute the code in order to get the output, the code needs to be translated
	to some lower level language, closer to what the computer can understand in order to process and output the
	result as something which hoomans can read

-> go build tutorial.go - tutorial.exe

	go code -> compiled executable/.exec file for computer to understand -> run it to get output for hoomans

	So if you want to have a intermediate exec file you need to fire the build command
	this exec file can be shared on any machine of any OS and it will still give you the output without having go installed

-> go is statically typed

	which means once you assign a value to a variable of a particular type it can't be changed. Basically it is not like Python
	which is dynamically typed and so you can type case values there.

-> implicit vs explicit type assignment to variables

	for example: 
		var name_of_variable type = value_of_variable
		var name_of_variable = value_of_variable // here go s implicitly guessing the type of the variable
		~ easiest way to declare variable              name_of_the_variable := value_of_the variable


*/
