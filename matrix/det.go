package matrix

// Berechnet die Determinante einer nichtleeren quadratischen Matrix.
func Det(m [][]float64) float64 {
	if len(m) == 0 || len(m) != len(m[0]) {
		return 0
	}

	if len(m) == 1 {
		return m[0][0]
	}

	// if len(m) == 2 {
	// 	return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	// }

	// if len(m) == 3 {
	// 	u0 := SubMatrix(m, 0, 0)
	// 	u1 := SubMatrix(m, 0, 1)
	// 	u2 := SubMatrix(m, 0, 2)
	// 	return m[0][0]*Det(u0) - m[1][0]*Det(u1) + m[2][0]*Det(u2)
	// }

	result := 0.0
	sign := 1.0
	for row := 0; row < len(m); row++ {
		u := SubMatrix(m, 0, row)
		coeff := m[row][0]
		result += sign * coeff * Det(u)
		sign = sign * -1
	}

	return result
}
