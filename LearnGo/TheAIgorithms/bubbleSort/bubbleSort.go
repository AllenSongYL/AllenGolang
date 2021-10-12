package bubbleSort

// 冒泡排序
// 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 针对所有的元素重复以上的步骤，除了最后一个。
// 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

// BubbleSort bubbleSort函数，传入slice，返回从小到大排序的slice
func BubbleSort(arr []int) []int {
	swapped := true
	// 默认发生交换，开始循环。如果没有交换，则说明排序结束，结束循环
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
