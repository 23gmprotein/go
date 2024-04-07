package main

import "fmt"

type life struct {
	love string
	wlb string
}

func main() {

	var soumyas life = life {"not good", "good"}
	fmt.Printf("Soumya's work life balance is %s and love life is %s", soumyas.wlb, soumyas.love )

}