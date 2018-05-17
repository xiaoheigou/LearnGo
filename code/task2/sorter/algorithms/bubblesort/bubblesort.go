package bubblesort

func BubbleSort(values []int) {
	// 具体可参考http://blog.csdn.net/yanxiaolx/article/details/51622286
	for i := 0; i < len(values); i++ {
		for j := len(values) - 1; j > i; j-- {
			if values[j] < values[j-1] {
				tmp := values[j-1]
				values[j-1] = values[j]
				values[j] = tmp
			}
		}
	}
}
