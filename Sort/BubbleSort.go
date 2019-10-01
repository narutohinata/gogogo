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
