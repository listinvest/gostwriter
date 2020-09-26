package pca9685

func min(a, b uint16) uint16 {
	if a > b {
		return b
	}
	return a
}
