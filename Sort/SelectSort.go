package Sort

//worst O(n^2)
//best O(n^2)
func SelectSort(data []int)  {
	for i := 0; i < len(data) -1;  i++{
		min := data[i]
		for j := i+1 ; j < len(data); j++ {
			if data[j] < min {
				data[j], min = min, data[j]
			}
		}
		data[i] = min
	}
}
