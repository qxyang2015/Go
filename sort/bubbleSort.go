package sort

/*
算法步骤
比较相邻的元素。如果第一个比第二个大，就交换他们两个。
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
针对所有的元素重复以上的步骤，除了最后一个。
持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/

//向下冒泡，加改进
func BubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		flag := 1
		for j := i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 0
			}
		}
		if flag == 1 {
			break
		}
	}
	return arr
}

//向上冒泡
func BubbleSort2(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
