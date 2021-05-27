package function

// RemoveSliceDuplication intのスライスの中から重複した数字を削除する
func RemoveSliceDuplication(slice []int) []int {
	m := make(map[int]struct{})
	newSlice := make([]int, 0)

	for _, element := range slice {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newSlice = append(newSlice, element)
		}
	}

	return newSlice
}
