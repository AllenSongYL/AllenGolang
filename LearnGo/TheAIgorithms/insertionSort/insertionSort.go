package insertionSort

// 插入排序

// [99,3,5,9,4]
func InsertionSort(arr []int) []int {

	for currentIndex := 1; currentIndex < len(arr); currentIndex++ {
		temporary := arr[currentIndex] //3
		iterator := currentIndex       //1
		// 不断和前面索引值比较，如果前面索引值大于等于取出的值，将当前索引的值替换成前一个索引的值
		for ; iterator > 0 && arr[iterator-1] >= temporary; iterator-- {
			arr[iterator] = arr[iterator-1]
		}
		arr[iterator] = temporary
	}
	return arr
}
