package isogram

import (
	"strings"
)

func IsIsogram(str string) bool {
	if str == "" || str == " " {
		return true
	}
	flags := [26]int{}
	for _, v := range strings.ToLower(str) {
		if v == '-' || v == ' ' {
			continue
		}
		if flags[v-'a'] == 0 {
			flags[v-'a']++
		} else {
			return false
		}

	}
	return true
}
