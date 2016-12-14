mkdir -p mock_math
mockgen -source=../formula/math.go -destination mock_math/math_mock.go
