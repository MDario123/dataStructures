package set

func dist(element int, pos uint, mask uint) uint {
	return (pos - uint(element) + mask + 1) & mask
}
