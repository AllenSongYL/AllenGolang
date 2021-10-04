package selectionSort

// 选择排序
// O(n²) 的时间复杂度
// 所以用到它的时候，数据规模越小越好。唯一的好处可能就是不占用额外的内存空间了吧。

// 首先在未排序序列中找到最小元素，存放到排序序列的起始位置。
// 再从剩余未排序元素中继续寻找最小元素，然后放到已排序序列的末尾。
// 重复第二步，直到所有元素均排序完毕。

func SelectionSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			// 当j索引的数更小时，交换
			if arr[j] < arr[min] {
				min = j
			}
		}

		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
