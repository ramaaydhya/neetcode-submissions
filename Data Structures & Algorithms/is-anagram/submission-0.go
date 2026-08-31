import "slices"
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	tempS := []byte(s)
	tempT := []byte(t)
	slices.Sort(tempS)
	slices.Sort(tempT)
	for i := range tempS {
		if tempS[i] != tempT[i] {
			return false
		}
	}
	return true	
}
