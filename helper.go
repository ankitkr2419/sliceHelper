package slicehelper

//Internal Helpers
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//TO CHECK IF AN ELEMENT IS PRESENT IN SLICE OR NOT
func FindElement(elem string, s []string) bool {
	return Contains(s, elem)
}

//TO FIND THE DIFFERNECE ELEMENT IN TWO SLICES
func DiffElements(s1 []string, s2 []string) (ans []string) {
	unique := make(map[string]bool)
	for _, i := range s1 {
		if !Contains(s2, i) {
			unique[i] = true
		}
	}
	for k := range unique {
		ans = append(ans, k)
	}
	return
}

//TO GET UNIQUE ELEMENTS IN SLICE
func UniqueElements(s []string) []string {
	visited := make(map[string]int)
	unique := []string{}
	for _, ss := range s {
		visited[ss]++
	}
	for k, v := range visited {
		if v <= 1 {
			unique = append(unique, k)
		}
	}
	return unique
}
