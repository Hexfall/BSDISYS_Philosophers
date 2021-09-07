package main

// PHILCOUNT Number of philosophers in program.
const PHILCOUNT int = 5
// PHILSWITCHCHANCE Likelihood that a philosopher decides to switch at any given moment.
const PHILSWITCHCHANCE float64 = 0.000000125
// ANNOUNCEEATING Debug tool. When true, philosophers write to the console, when they start and stop eating.
const ANNOUNCEEATING bool = true

// Returns the modulo of the number by PHILCOUNT, so an index will be in the range [0, PHILCOUNT - 1]
func limit(philId int) int {
	return philId % PHILCOUNT
}

// Boolean to integer function.
func btoi(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}