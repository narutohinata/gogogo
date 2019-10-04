package Sort
// 归并排序

func MergeSort(data []int)  {
	_MergeSort(data, 0, len(data)-1)
}

func _MergeSort(data []int, start int, end int)  {
	if start < end {
		mid := (start + end) / 2
		_MergeSort(data, start, mid)
		_MergeSort(data, mid+1, end)
		Merge(data, start, mid, end)
	}
}

func Merge(data []int, start int, mid int, end int)  {
	len1 := mid - start + 1
	len2 := end - mid

	tempArry1 := make([]int, len1)
	tempArry2 := make([]int, len2)

	for i := 0; i < len1;  i++ {
		tempArry1[i] = data[start+i]
	}

	for j := 0; j < len2; j++ {
		tempArry2[j] = data[mid+j+1]
	}

	count1, count2 := 0, 0
	for k := start; k <= end ; k++ {
	//	TODO should implement
		if count1 > (len1 - 1) {
			data[k] = tempArry2[count2]
			count2 += 1
		} else if count2 > (len2 - 1) {
			data[k] = tempArry1[count1]
			count1 += 1
		} else {
			if tempArry1[count1] <= tempArry2[count2]{
				data[k] = tempArry1[count1]
				count1 += 1
			} else {
				data[k] = tempArry2[count2]
				count2 += 1
			}
		}
	}
}