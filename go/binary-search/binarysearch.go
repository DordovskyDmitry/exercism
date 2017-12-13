package binarysearch

func SearchInts(array []int, key int) int {
	if len(array) == 0 {
		return -1
	}
	return find(array, key, 0)
}

func find(array []int, key, shift int) int {
	size := len(array)
	if size == 1 {
		if key == array[0] {
			return shift
		}
		return -1
	}
	if array[size/2] > key {
		return find(array[:size/2], key, shift)
	}
	return find(array[size/2:], key, shift+size/2)
}
