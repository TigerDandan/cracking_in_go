package src

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	backup := x
	y := 0

	for x/10 > 0 {
		y = 10*y + x%10
		x = x / 10
	}
	y = y*10 + x

	if y == backup {
		return true
	}
	return false
}
