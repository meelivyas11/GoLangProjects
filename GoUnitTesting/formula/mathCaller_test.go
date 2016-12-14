package formula

import (
	math_mock "GoUnitTesting/mocks/mock_math"
	"testing"

	"github.com/golang/mock/gomock"
)

// contract validation
// Make sure MathCaller.Sum() return what ever is returned by MathImpl.DoSum()
func TestSum(t *testing.T) {
	ctr := gomock.NewController(t)
	defer ctr.Finish()

	snums := math_mock.NewMockMath(ctr)

	// It is expected that DoSum() method return 12 as output
	snums.EXPECT().DoSum().Return(12)
	tc := MathCaller{MCaller: snums}
	result := tc.Sum()

	// Check if the output from DoSum() == to that from Sum()
	if result != 12 {
		t.Fail()
	}
}
