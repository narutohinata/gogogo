package Sort

func InsertSort(data []int)  {
	for i := 1; i < len(data); i++ {
		key := data[i]
		
		for j := i-1; j >= 0 && data[j] > key{
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}
