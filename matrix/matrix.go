package matrix

import "fmt"

func PrintMatrix(m [][]float64) {
	for row := 0; row < len(m); row++ {
		fmt.Print("|")
		fmt.Print(" ")
		for column := 0; column < len(m[row]); column++ {
			fmt.Print(m[row][column])
			fmt.Print(" ")
		}
		fmt.Println("|")
	}

}

// Berechnet die Untermatrix von m,
// bei der Spalte s und Zeile z fehlt.
func SubMatrix(m [][]float64, s, z int) [][]float64 {
	a := [][]float64{}

	for row := 0; row < len(m); row++ {
		if row != z {
			r := []float64{}
			for column := 0; column < len(m[row]); column++ {
				if column != s {
					r = append(r, m[row][column])
				}
			}
			a = append(a, r)
		}
	}

	return a
}
