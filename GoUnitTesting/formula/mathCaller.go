package formula

import (
	"fmt"
	"strconv"
)

type MathCaller struct {
	MCaller Math
}

// Dependency (Math interface) is injected
func (mc *MathCaller) Sum() int {
	res := mc.MCaller.DoSum()
	fmt.Println("The result from Sum() method is: " + strconv.Itoa(res))
	return res
}

func CallImplementation() {
	fmt.Print("Interface Sample\n")
	nums := LoadNums(5, 7)
	result := Math.DoSum(nums)
	fmt.Println("The result CallInterface() is: " + strconv.Itoa(result))
}
