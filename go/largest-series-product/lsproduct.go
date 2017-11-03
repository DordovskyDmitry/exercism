package lsproduct

import "errors"

const testVersion = 5

func LargestSeriesProduct(digits string, span int) (result int, ok error) {
	if span > len(digits) || span < 0 {
		return -1, errors.New("Invalid length")
	}
	if span == 0 && digits == "" {
		return 1, nil
	}
	prod := 1
	count := 0
	i := 0
	for i < len(digits) {
		if digits[i] < '0' || digits[i] > '9' {
			return -1, errors.New("Not a digit")
		}
		if digits[i] == '0' {
			prod = 1
			count = 0
			i++
			continue
		}
		if count < span {
			count++
			prod *= runeToInt(digits[i])
		} else {
			prod = prod / runeToInt(digits[i-span]) * runeToInt(digits[i])
		}
		if count == span && prod > result {
			result = prod
		}
		i++
	}
	return result, nil
}

func runeToInt(r byte) int {
	return int(r - '0')
}
