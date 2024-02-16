package matrix

import "fmt"

func ExamplePrintMatrix() {
	m1 := [][]float64{
		{1, 2},
		{3, 4},
	}

	m2 := [][]float64{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	PrintMatrix(m1)
	fmt.Println()
	PrintMatrix(m2)

	// Output:
	// | 1 2 |
	// | 3 4 |
	//
	// | 1 2 3 |
	// | 1 2 3 |
	// | 1 2 3 |
}

func ExampleSubMatrix() {
	m1 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	PrintMatrix(SubMatrix(m1, 1, 2))

	// Output:
	// | 1 3 |
	// | 4 6 |
}
