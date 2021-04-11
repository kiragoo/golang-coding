package math

func IsUgly(num int) bool {
	if num == 0 {
		return false
	}
	for {
		if num == 1 {
			break
		}
		if num%2 == 0 {
			num /= 2
			continue
		}
		if num%3 == 0 {
			num /= 3
			continue
		}
		if num%5 == 0 {
			num /= 5
			continue
		}
		return false
	}
	return true
}
