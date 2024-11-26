package matrixtools

import "fmt"

func ExampleGetRow() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	row := GetRow(m, 1)
	fmt.Println(row)

	// Output: [4 5 6]
}

func ExampleGetCol() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	col := GetCol(m, 1)
	fmt.Println(col)

	// Output: [2 5]
}

func ExampleAddRows() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	AddRows(m, 0, 1)
	fmt.Println(m[0])
	fmt.Println(m[1])

	// Output:
	// [5 7 9]
	// [4 5 6]
}

func ExampleScalarMultRow() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}
	ScalarMultRow(m, 1, 2)
	fmt.Println(m[0])
	fmt.Println(m[1])

	// Output:
	// [1 2 3]
	// [8 10 12]
}

func ExampleTransposed() {
	m := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
	}

	transposed := Transposed(m)
	fmt.Println(transposed[0])
	fmt.Println(transposed[1])
	fmt.Println(transposed[2])

	// Output:
	// [1 4]
	// [2 5]
	// [3 6]
}
