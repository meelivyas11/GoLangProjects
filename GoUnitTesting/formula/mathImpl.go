package formula

type nums struct {
	num1 int
	num2 int
}

func LoadNums(n1 int, n2 int) *nums {
	return &nums{
		num1: n1,
		num2: n2,
	}
}
func (n *nums) DoSum() int {
	return n.num1 + n.num2
}

func (n *nums) DoArea() int {
	return n.num1 * n.num2
}
