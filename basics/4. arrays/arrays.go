package main
import "fmt"

func main(){

	var arr=[5]int{1,2,3,4,5} //-> here the syntax is important
	//-> alternatively we can do something like arr := [5]int{1,2,3,4,5}, both ways work

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		println(arr[i])
	}

}