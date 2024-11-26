package matrix

import "fmt"

func ExampleMatrix_Row() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(m.Row(1))

	// Output: [4 5 6]
}

func ExampleGetCol() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(m.Col(1))

	// Output: [2 5]
}

func ExampleMatrix_AddRows() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	m.AddRows(0, 1)
	fmt.Println(m)

	// Output: [[5 7 9] [4 5 6]]
}

func ExampleMatrix_ScalarMultRow() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}
	m.ScalarMultRow(1, 2)
	fmt.Println(m)

	// Output: [[1 2 3] [8 10 12]]
}

func ExampleMatrix_Transposed() {
	m := Matrix{
		{1, 2, 3},
		{4, 5, 6},
	}

	transposed := m.Transposed()
	fmt.Println(transposed)

	// Output: [[1 4] [2 5] [3 6]]
}
