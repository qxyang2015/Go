package sort

/*
算法步骤
	将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
	从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/

//插入排序
func InsertSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 0; j-- {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

//逆向思想,插入排序
func InsertSort2(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
