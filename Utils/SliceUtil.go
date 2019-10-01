package Utils

import "fmt"

func PrintSlice(data []int)  {
	for i, v := range data {
		fmt.Printf("data[%d] is %d \n", i, v)
	}
}
