package comon

import "errors"

func Sumar(a int, b int) (error, int) {
	rs := a + b

	if rs > 5 {
		return errors.New("el resultado es mayor que 5"), 0
	}

	return nil, rs
}
