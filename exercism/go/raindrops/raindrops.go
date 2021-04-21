package raindrops

import (
	"bytes"
	"strconv"
)

func Convert(input int) string {
	dict := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}
	var r bytes.Buffer
	for k, v := range dict {
		if input%k == 0 {
			r.WriteString(v)
		}
	}
	if r.String() == "" {
		return strconv.Itoa(input)
	} else {
		return r.String()
	}
}
