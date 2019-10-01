package Sort

// Base Version
// O(n^2)
func BubbleSort(data []int)  {
	for i := len(data) - 1; i > 0; i-- { // n
		for j := 0; j < i; j++{ // n
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
			}
		}
	}
}

func BubbleSortVersion2(data []int)  {
	for i := len(data) - 1; i > 0; i-- { // n
		sorted := true // global sorted
		for j := 0; j < i; j++{ // n
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
				sorted = false
			}
		}
		// if global sorted is true stop bubble
		if sorted {
			return
		}
	}
}

func BubbleSortVersion3(data []int)  {
	for i := len(data) - 1; i > 0; i-- { // n
		sortedIndex := i
		for j := 0; j < i; j++{ // n
			if data[j] > data[j+1] {
				temp := data[j]
				data[j] = data[j+1]
				data[j+1] = temp
				sortedIndex = j+1
			}
		}
		i = sortedIndex
	}
}
