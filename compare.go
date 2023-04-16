package template

func attributesAreEqual(attrs1, attrs2 []string) bool {
	if len(attrs1) != len(attrs2) {
		return false
	}
	for i := range attrs1 {
		if attrs1[i] != attrs2[i] {
			return false
		}
	}
	return true
}
