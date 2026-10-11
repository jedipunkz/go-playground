package calc

import "errors"

var ErrDivideByZero = errors.New("divide by zero")

// Div は a を b で割った結果を返す
func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}
