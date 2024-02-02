package beispiele

import "fmt"

func ExamplePrintList() {
	PrintList([]int{10, 20, 30, 40, 50})

	// Output:
	// 10
	// 20
	// 30
	// 40
	// 50
}

func ExampleAddElements() {
	fmt.Println(AddElements([]int{10, 20, 30, 40, 50}))

	// Output:
	// 150
}
