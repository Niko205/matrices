package matrix

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {
	// TODO
}

// HINT
// Überprüfen Sie, ob das Diagonalelement ungleich 0 ist.
// Falls ja, berechnen Sie den Kehrwert des Diagonalelements und
// multiplizieren Sie die gesamte Zeile mit diesem Faktor.
// Dazu können Sie die Funktion `ScalarMultRow` verwenden.

// EliminateBelow erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(row int) {
	// TODO
}

// HINT
// Iterieren Sie über alle Zeilen unter der Zeile `row`, also von `row+1` bis `len(m)`.
// Skalieren Sie jeder dieser Zeilen, so dass sie mit `-1` beginnen.
// Addieren Sie dann die Zeile `row` zu jeder dieser Zeilen.

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {
	// TODO
}

// HINT
// Gehen Sie analog zu `EliminateBelow` vor, aber iterieren Sie über alle Zeilen über der Zeile `row`, also von `0` bis `row`.

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {
	// TODO
}

// HINT
// Iterieren Sie über alle Zeilen der Matrix.
// Normieren Sie jede Zeile und eliminieren Sie alle Elemente unter dem Diagonalelement.

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	// TODO
}

// HINT
// Gehen Sie analog zu `UpperTriangular` vor, aber iterieren Sie über die Zeilen in umgekehrter Reihenfolge.

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	// TODO
}

// HINT
// Die Gauß-Elimination muss lediglich zuerst eine obere Dreiecksmatrix und dann eine untere Dreiecksmatrix erzeugen.

// Allgemeiner Hinweis:
// Die beiden folgenden Funktionen können nützlich sein, falls Ihre Ergebnisse
// am Ende nicht ganz wie erwartet sind:

// almostEqual ist eine Hilfsfunktion, die zwei float64-Werte auf Gleichheit prüft.
// Da float64-Werte nicht exakt vergleichbar sind, wird ein Toleranzwert von 1e-10 verwendet.
func almostEqual(a, b float64) bool {
	return a-b < 1e-10 && b-a < 1e-10
}

// fixZeros erwartet eine Matrix und setzt alle Elemente, die kleiner als 1e-10 sind, auf 0.
func fixZeros(m Matrix) {
	for i := range m {
		for j := range m[i] {
			if almostEqual(m[i][j], 0) {
				m[i][j] = 0
			}
		}
	}
}
