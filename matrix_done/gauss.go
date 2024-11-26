package matrix

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	if m[col][col] != 0 {
		m.ScalarMultRow(col, 1/m[col][col])
	}
}

// EliminateBelow erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(row int) {
	x := 0.0

	for i := row + 1; i < len(m); i++ {
		x = -m[row][row] / m[i][row]
		m.ScalarMultRow(i, x)
		m.AddRows(i, row)
	}
}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	x := 0.0
	for i := row - 1; i >= 0; i-- {
		x = -m[row][row] / m[i][row]
		m.ScalarMultRow(i, x)
		m.AddRows(i, row)
	}
}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {

	for i := 0; i < len(m)-1; i++ {
		m.Normalize(i)
		m.EliminateBelow(i)
		m.Normalize(i + 1)
	}
	m.CheckNull()
}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	for i := len(m) - 1; i >= 0; i-- {
		m.Normalize(i)
		m.EliminateAbove(i)
		m.Normalize(i)
	}
	m.CheckNull()
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	m.UpperTriangular()
	m.LowerTriangular()
}

// CheckNul geht die Matrix durch und erstetzt alle -0 mit 0
func (m Matrix) CheckNull() {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == -0 {
				m[i][j] = 0
			}
		}
	}
}
