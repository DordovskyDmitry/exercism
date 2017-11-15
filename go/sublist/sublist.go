package sublist

type Relation string

func Sublist(listOne []int, listTwo []int) Relation {
	if isEqual(listOne, listTwo) {
		return Relation("equal")
	}
	if isSublist(listOne, listTwo) {
		return Relation("sublist")
	}
	if isSublist(listTwo, listOne) {
		return Relation("superlist")
	}
	return Relation("unequal")
}

func isSublist(listOne []int, listTwo []int) bool {
	if len(listOne) >= len(listTwo) {
		return false
	}
	if len(listOne) == 0 {
		return true
	}
	for i := 0; i < len(listTwo); i++ {
		if listTwo[i] == listOne[0] {
			j := 1
			for j < len(listOne) {
				if i+j >= len(listTwo) {
					return false
				}
				if listTwo[i+j] != listOne[j] {
					break
				}
				j++
			}
			if j == len(listOne) {
				return true
			}
		}
	}
	return false
}

func isEqual(listOne []int, listTwo []int) bool {
	if len(listOne) != len(listTwo) {
		return false
	}
	for i, v := range listOne {
		if listTwo[i] != v {
			return false
		}
	}
	return true
}
