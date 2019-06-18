package main

import "fmt"

func find_max(slice []interface{}, less func(i, j int) bool) interface{}{
	max := 0
	for ind, _ := range slice{
		if less(max, ind){
			max = ind
		}
	}
	return slice[max]
}

func main() {
	sl := []interface{}{1, 43, 3, 4}
	max := find_max(sl, func(i, j int) bool {
		return sl[i].(int) < sl[j].(int)
	})
	fmt.Println("Max:", max)

	sl2 := []interface{}{1.1321, 4.3, 3.4, 4.5}
	max2 := find_max(sl2, func(i, j int) bool {
		return sl2[i].(float64) < sl2[j].(float64)
	})
	fmt.Println("Max:", max2)

	sl3 := []interface{}{"hallow", "my", "friend", "ivan"}
	max3 := find_max(sl3, func(i, j int) bool {
		return sl3[i].(string) < sl3[j].(string)
	})
	fmt.Println("Max:", max3)

}
