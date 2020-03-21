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

////冒泡排序优化第一版
//
//5 public class BubbleSoerOpt1 {
//6     public static void main(String[] args) {
//7         int[] list = {5,4,3,1,2};
//8         int temp = 0; // 开辟一个临时空间, 存放交换的中间值
//9         // 要遍历的次数
//10         for (int i = 0; i < list.length-1; i++) {
//11             int flag = 1; //设置一个标志位
//12             //依次的比较相邻两个数的大小，遍历一次后，把数组中第i小的数放在第i个位置上
//13             for (int j = 0; j < list.length-1-i; j++) {
//14                 // 比较相邻的元素，如果前面的数小于后面的数，交换
//15                 if (list[j] < list[j+1]) {
//16                     temp = list[j+1];
//17                     list[j+1] = list[j];
//18                     list[j] = temp;
//19                     flag = 0;  //发生交换，标志位置0
//20                 }
//21             }
//22             System.out.format("第 %d 遍最终结果：", i+1);
//23             for(int count:list) {
//24                 System.out.print(count);
//25             }
//26             System.out.println("");
//27             if (flag == 1) {//如果没有交换过元素，则已经有序
//28                 return;
//29             }
//30
//31         }
//32     }
//33 }
