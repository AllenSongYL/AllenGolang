package main

func minArray(numbers []int) int {
	var tmp int
	if len(numbers) == 1 {
		return numbers[0]
	}
	for i, j := 0, 1; j < len(numbers); i++ {
		if numbers[i] < numbers[j] {
			tmp = numbers[i]
		} else {
			tmp = numbers[j]
		}
	}
	return tmp
}

func main() {

}
