package main
import "fmt"

func main() {

	arr := [5]int{1,2,3,4,5}
	slizes := arr[1:3]

	fmt.Println(slizes[1:3])
	fmt.Println(len(slizes))
	fmt.Println(cap(slizes))

	// -> making slices with make function

	make_slice_example := make([]int, 5)
	fmt.Print(make_slice_example)

}