package main

import "fmt"

func main() {

	var mp map[string]int = map[string]int{
//				^key   ^value
		"apple":  1,
		"pear":   8,
		"Orange": 10,
	}

	fmt.Println(mp["apple"])

}
