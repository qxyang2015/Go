package sort

import "fmt"

//冒泡排序
func BubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	return arr
}
