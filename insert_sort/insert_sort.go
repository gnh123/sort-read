package main

import "fmt"

// 插入排序
func insertSort(a []int) {

	// 第1个元素, 不需要比较. 可以认识是排序过的.
	for i := 1; i < len(a); i++ {
		// 取出当前的值当key, 它需要和前面的列表排序
		key := a[i]

		// 和前面的值比较
		j := i - 1

		for j >= 0 && key < a[j] {

			// 把前面的列表往后挪一挪
			a[j+1] = a[j]

			j--

		}
		// 这个位置就是需要插入的位置
		a[j+1] = key
	}
}

func main() {
	a := []int{1, 10, 8, 9, 3}
	fmt.Printf("after:%v\n", a)
	insertSort(a)
	fmt.Printf("after:%v\n", a)

	a = []int{1, 3, 3, -1, -1}
	fmt.Printf("after:%v\n", a)
	insertSort(a)
	fmt.Printf("after:%v\n", a)
}
