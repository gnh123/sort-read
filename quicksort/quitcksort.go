package quicksort

func p(a []int, low int, high int) int {

	// 先取基准值
	pivot := a[low]

	// 双指针确认
	for low < high {
		// 从右往左走, 找一个小于基准值的值
		for low < high && a[high] >= pivot {
			high--
		}

		// >> 这里一开始写错了

		//把小于基准值的值的给a[low], 就是小的往左放
		a[low] = a[high]

		// 从左往右走, 找一个大于基准值的值
		for low < high && a[low] <= pivot {
			low++
		}

		//把大于基准值的给到a[high], 就是大的往右放
		a[high] = a[low]
	}
	a[low] = pivot
	return low
}

func quicksortCore(a []int, low int, high int) {
	if low < high {
		pivot := p(a, low, high)
		quicksortCore(a, low, pivot-1)
		quicksortCore(a, pivot+1, high)
	}
}

func quicksort(a []int) {
	quicksortCore(a, 0, len(a)-1)
}
