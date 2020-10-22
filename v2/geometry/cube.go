package geometry

import "errors"

// Cubevolueme function is used to Find the Volume of a Ccube
func Cubevolueme(n int) (int, error) {

	if n != 0 {
		return n * n * n, nil
	}
	return 0, errors.New("Cube with side size 0 is not allowed")

}
