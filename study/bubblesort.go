package study

func TestSort() {
	arr := []int{101, 32, 77, 4, -5, 88, 23, 11}
	mySort(arr)
	for _, i := range arr {
		println(i)
	}
	var arr2 [8]*int
	for i := range arr {
		arr2[i] = &arr[i]
	}
	for _, i := range arr2 {
		println(*i)
	}
}

func mySort(arr []int) {
	var flag bool
	for i := 0; i < len(arr)-1; i++ {
		flag = true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {
			break
		}
	}
}
