package main

import "fmt"

func find_max(slice []interface{}, less func(i, j int) bool) interface{}{
	if len(slice) == 0{
		return nil
	}
	max := 0
	for ind := range slice{
		if less(max, ind){
			max = ind
		}
	}
	return slice[max]
}

func print(max interface{})  {
	if max != nil{
		fmt.Println("Max:", max)
	}else {
		fmt.Println("Slice is Empty!")
	}

}

func main() {
	//slice of int
	sl := []interface{}{1, 43, 3, 4}
	max := find_max(sl, func(i, j int) bool {
		return sl[i].(int) < sl[j].(int)
	})
	print(max)

	//slice of float
	sl2 := []interface{}{1.1321, 4.3, 3.4, 4.5}
	max2 := find_max(sl2, func(i, j int) bool {
		return sl2[i].(float64) < sl2[j].(float64)
	})
	print(max2)

	//slice of string
	sl3 := []interface{}{"hallo", "my", "friend", "ivan"}
	max3 := find_max(sl3, func(i, j int) bool {
		return sl3[i].(string) < sl3[j].(string)
	})
	print(max3)

	//empty slice
	print(find_max([]interface{}{}, func(i, j int) bool {
		return sl3[i].(string) < sl3[j].(string)
	}))

}
