type sortableSomething []*Something

func (s sortableSomething) Len() int { return len(s) }
func (s sortableSomething) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortableSomething) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
