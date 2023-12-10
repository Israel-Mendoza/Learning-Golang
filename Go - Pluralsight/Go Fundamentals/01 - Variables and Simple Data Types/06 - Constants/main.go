package main

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
}
