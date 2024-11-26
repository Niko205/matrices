package arraytools

import "fmt"

func ExampleScalarMult() {
	a := []float64{1, 2, 3}
	ScalarMult(a, 2)
	fmt.Println(a)

	// Output: [2 4 6]
}

func ExampleAdd() {
	a := []float64{1, 2, 3}
	b := []float64{4, -5, 6}
	Add(a, b)
	fmt.Println(a)

	// Output: [5 -3 9]
}
