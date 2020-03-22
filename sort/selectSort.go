package sort

/*
算法步骤:
  首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
  再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
  重复第二步，直到所有元素均排序完毕。
*/

func SelectionSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
