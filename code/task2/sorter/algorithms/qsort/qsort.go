package qsort

//func main(){
//	a :=[]int{12,23,4,23,232,333,2}
//	QuickSort(a)
//	fmt.Println(a)
//}
func QuickSort(values []int) {
	//fmt.Println(values,"BUG1")
	quickSort(values, 0, len(values)-1)
	//fmt.Println(values,"BUG2")
}

// 快速排序，每次先排好基准数，然后分而治之
//可以参考填坑模型http://blog.csdn.net/morewindows/article/details/6684558
func quickSort(values []int, left, right int) {
	temp := values[left]

	p := left
	i, j := left, right

	for i <= j {
		//从右边开始找，找到第一个比temp小的数
		for j >= p && values[j] >= temp {
			j--
		}
		//把[j]填到[p]的位置.此算法默认p永远就是那个坑
		if j >= p {
			values[p] = values[j]
			p = j //p坑的位置发生变化，需要从左向右扫描填坑
		}
		//从左向右找到比temp大的第一个数
		if values[i] <= temp && i <= p {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp //左右反复扫描，temp终于找到自己的萝卜坑了
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}
