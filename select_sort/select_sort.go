package selectsort

func selectsort(a []int) {
	for i := 0; i < len(a); i++ {
		// 选择最小值的索引
		min := i
		// 从当前值后面的值往下比较
		for j := i + 1; j < len(a); j++ {
			if a[min] > a[j] {
				min = j
			}
		}

		// 如果i 和min 不相等, 说明找到一个比当前值还小的, 需要交换
		if i != min {
			a[i], a[min] = a[min], a[i]
		}
	}
}
