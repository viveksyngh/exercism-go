package hamming

import "errors"

//Distance finds hamming distance between two DNA
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("Length is not equal")
	}

	distance := 0

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance += 1
		}
	}
	return distance, nil
}
