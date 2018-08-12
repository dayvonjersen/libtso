import "sort"

func unique(slice []string) []string {
	for i := len(slice) - 1; i >= 0; i-- {
		s := slice[i]
		j := sort.SearchStrings(slice, s)
		if i != j {
			slice = remove(slice, j)
		}
	}
	return slice
}

func remove(slice []string, i int) []string {
	return append(slice[0:i], slice[i+1:]...)
}

func insert(slice []string, s string, i int) []string {
	return append(slice[:i], append([]string{s}, slice[i:]...)...)
}
