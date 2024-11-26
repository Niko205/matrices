package matrix

import matrixtools "matrices/matrixtools_done"

type Row = []float64
type Col = []float64
type Matrix []Row

// Row liefert die i-te Zeile der Matrix m.
func (m Matrix) Row(i int) Row {
	return matrixtools.GetRow(m, i)
}

// Col liefert die j-te Spalte der Matrix m.
func (m Matrix) Col(j int) Col {
	return matrixtools.GetCol(m, j)
}

// AddRows erwartet zwei Zeilennummern i und j.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in Zeile i.
func (m Matrix) AddRows(i, j int) {
	matrixtools.AddRows(m, i, j)
}

// ScalarMultRow erwartet eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func (m Matrix) ScalarMultRow(i int, factor float64) {
	matrixtools.ScalarMultRow(m, i, factor)
}

// Transposed liefert die Transponierte der Matrix.
func (m Matrix) Transposed() Matrix {
	return matrixtools.Transposed(m)
}
