package findNumberIn2DArray

/*
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
*/

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for {
		for {
			if i >= len(matrix) || j < 0 {
				return false
			}
			if matrix[i][j] == target {
				return true
			}
			if matrix[i][j] > target {
				j--
			} else if matrix[i][j] < target {
				i++
			}
		}
	}
}
