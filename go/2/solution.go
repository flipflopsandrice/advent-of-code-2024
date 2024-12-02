package two

func solution(arr1 []int, arr2 []int) int {
	var similarity int = 0
	for _, v := range arr1 {
		count := 0
		for _, v2 := range arr2 {
			if v2 == v {
				count += 1
			}
		}
		similarity += count * v
	}

	return similarity
}
