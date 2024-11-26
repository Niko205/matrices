package matrix

import "fmt"

func ExampleMatrix_Gauss_inverse() {
	m := Matrix{
		{1, 2, 1, 0},
		{2, 3, 0, 1},
	}

	m.Gauss()
	fmt.Println(m[0])
	fmt.Println(m[1])

	// Output:
	// [1 0 -3 2]
	// [0 1 2 -1]
}
