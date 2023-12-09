package main

import "fmt"

func main() {
	// You can declare multiple variables at once with var, and they can be of the same type:
	var a, b int = 10, 20
	// …all zero values of the same type:
	var c, d int
	// …or of different types:
	var e, f = 10, "hello"

	fmt.Println(a, b, c, d, e, f)
	// 10 20 0 0 10 hello

	// There’s one more way to use var.
	// If you are declaring multiple variables at once,
	// you can wrap them in a declaration list:
	var (
		x    int
		y        = 20
		z    int = 30
		g, h     = 40, "hello"
		i, j string
	)
	fmt.Println(x, y, z, g, h, i, j)
	// 0 20 30 40 hello

	// The following two declarations do exactly the same thing:
	var k = 10
	l := 10
	fmt.Println(k, l)
	// 10 10

	// Like var, you can declare multiple variables at once using :=.
	var m, n = 10, "hello"
	o, p := 10, "hello"
	fmt.Println(m, n, o, p)
	// 10 hello 10 hello

	// := also allows us to assign values to existing variables:
	q := 10
	r, s := 30, "hello"
	fmt.Println(q, r, s)
	// 10 30 hello

}
