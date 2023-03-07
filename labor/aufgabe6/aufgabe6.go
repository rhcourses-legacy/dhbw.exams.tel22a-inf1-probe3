package aufgabe6

/*
AUFGABENSTELLUNG:
Hier ist ein Struct Matrix vorgegeben.
Vervollständigen Sie die Funktion Mult.
PUNKTE: 15
BEWERTUNG:
*/

type Matrix struct {
	Rows [][]int
}

// Mult erwartet eine weitere Matrix m2, berechnet die Matrixmultiplikation der beiden und liefert diese zurück.
// Gehen Sie davon aus, dass die Matrizen wohlgeformt sind und von den Dimensionen her passen,
// so dass eine Multiplikation möglich ist.
func (m1 Matrix) Mult(m2 Matrix) Matrix {
	result := Matrix{Rows: [][]int{}}
	// BEGIN-SOLUTION
	resultRows := len(m1.Rows)
	resultColumns := len(m2.Rows[0])
	m1Columns := len(m1.Rows[0]) // auch == m2Rows
	for range m1.Rows {
		result.Rows = append(result.Rows, make([]int, resultColumns))
	}

	for i := 0; i < resultRows; i++ {
		for j := 0; j < resultColumns; j++ {
			for n := 0; n < m1Columns; n++ {
				result.Rows[i][j] += m1.Rows[i][n] * m2.Rows[n][j]
			}
		}
	}
	// END-SOLUTION
	return result
}
