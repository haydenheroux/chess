package main

func flip(s0 [][]string) {
	for l, r := 0, len(s0)-1; l < r; l, r = l+1, r-1 {
		s0[l], s0[r] = s0[r], s0[l]
	}

	for _, s1 := range s0 {
		reverse(s1)
	}
}

func reverse(s []string) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		s[l], s[r] = s[r], s[l]
	}
}
