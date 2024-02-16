package matrix

import "fmt"

func ExampleDet() {
	m1 := [][]float64{
		{1, 2},
		{3, 4},
	}

	m2 := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	m4 := [][]float64{
		{0, 1, 2},
		{3, 2, 1},
		{1, 1, 0},
	}

	m3 := [][]float64{
		{2, 2, 0, 0},
		{4, 1, -3, 2},
		{0, -1, 2, 1},
		{1, -1, 0, 0},
	}

	fmt.Println(Det(m1))
	fmt.Println(Det(m2))
	fmt.Println(Det(m4))
	fmt.Println(Det(m3))

	// Output:
	// -2
	// 0
	// 3
	// 28
}
