package util

// Max returns max int value
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns min int value
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
