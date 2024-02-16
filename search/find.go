package search

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {
	if len(list) == 0 {
		return -1
	}

	if list[0] == x {
		return 0
	}

	h := Find(list[1:], x) + 1
	if h >= 0 {
		return h + 1
	}
	return -1
}
