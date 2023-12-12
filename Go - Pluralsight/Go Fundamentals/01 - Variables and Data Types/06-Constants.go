package main

import "fmt"

func main() {
	const a = "Israel"       // Implicitly typed (will be treated as a literal)
	const b string = "Mogom" // Explicitly typed (will be treated as a string)

	// Grouping constant declaration
	const (
		nameOne   = "Alex"
		nameTwo   = "Sandy"
		nameThree = "Luis"
	)

	// Constant expression (inheriting values)
	const (
		title    = "Support Engineer"
		position // "Support Engineer"
		role     // "Support Engineer"
	)

	const (
		zero  = iota
		one   // 1
		two   // 2
		three // 3
		four  // 4
	)

	const (
		_  = iota
		KB = float64(1 << (10 * iota))
		MB
		GB
		TB
	)

	fmt.Printf("KB: %v\nMB: %v\nGB: %v\nTB: %v\n", KB, MB, GB, TB)
	// KB: 1024
	// MB: 1.048576e+06
	// GB: 1.073741824e+09
	// TB: 1.099511627776e+12
}
