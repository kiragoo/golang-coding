package hamming

import "errors"

func Distance(a, b string) (int, error) {
	r := 0
	if (a == "" && b != "") || (a != "" && b == "") || len(a) != len(b) {
		return r, errors.New("false")
	}
	length := len(a)
	i := 0
	for {
		if i == length {
			break
		}
		if a[i] != b[i] {
			r++
		}
		i++
	}
	return r, nil
}
