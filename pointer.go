package main

func change(n *int) {
	*n += 10
}

func modifySlice(s *[]int) {
	for idx := range *s {
		(*s)[idx] *= 2
	}
}
