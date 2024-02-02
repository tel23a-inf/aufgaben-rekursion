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

func ExampleAddListElements() {
	fmt.Println(AddListElements([]int{10, 20, 30, 40, 50}))

	// Output:
	// 150
}
