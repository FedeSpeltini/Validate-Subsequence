package main

func main() {

	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{22, 25, 6}
	IsValidSubsequence(array, sequence)
}

func IsValidSubsequence(array []int, sequence []int) bool {

	if len(sequence) > len(array) {
		println("false")
		return false
	}
	i := 0
	mark := 0
	for index, element := range array {

		if element == sequence[i] {
			mark++
			i = mark
			if mark == len(sequence) {
				println("true")
				return true
			}
			if index < len(array)-1 {
				continue
			}

		}
		if i == len(sequence)-1 {
			i = mark
			if mark == len(sequence) {
				println("true")
				return true
			}
			continue
		}
		if mark == len(sequence) {
			println("true")
			return true
		}
	}
	println("false")
	return false
}

func IsValidSubsequenceChatGPT(array []int, sequence []int) bool {
	i, j := 0, 0

	for i < len(array) && j < len(sequence) {
		if array[i] == sequence[j] {
			j++
		}
		i++
	}

	if j == len(sequence) {
		println("true")
		return true
	}

	println("false")
	return false
}
