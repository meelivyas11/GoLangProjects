package main

import formula "GoUnitTesting/formula"

func main() {
	// "Sum()" function from mathCaller.go - Dependency Injection
	snums := formula.LoadNums(5, 7)
	mc := formula.MathCaller{MCaller: snums}
	mc.Sum()

	// CallImplementation - calls implementation directly
	formula.CallImplementation()
}
