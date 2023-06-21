package set

func dist(element int, pos uint, mask uint) uint {
	return (pos - uint(element) + mask + 1) & mask
}

func udist(element uint, pos uint, mask uint) uint {
	return (pos - element + mask + 1) & mask
}
