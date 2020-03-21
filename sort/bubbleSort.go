package sort

//冒泡排序
func BubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
