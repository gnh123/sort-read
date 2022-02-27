package bubblesort

func bubbleSort(a []int) {

	// 决定外层的循环次数
	for i := 0; i < len(a); i++ {
		// 决定内层的循环次数, 右侧是逐渐稳定的排序, 所以要慢慢减去
		for j := 0; j < len(a)-i-1; j++ {

			if a[j] > a[j+1] {
				// 两两个交换
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
