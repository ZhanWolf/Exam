package main

import "fmt"

func main() {
	floor1 := []int{2, 2, 3, 4}
	floor1ture := []int{0, 1, 1, 1}
	floor2 := []int{2, 1, 2, 3}
	floor2ture := []int{0, 1, 1, 1}
	i := 1
	temp1 := floor1[i]
	temp2 := floor1[i]

	for {
		if floor1ture[i] == 1 {
			temp2--
			if temp2 == 0 {
				break
			}
			if i == 3 {
				i = 0
			}
			i++
		} else {
			if i == 3 {
				i = 0
			}
			i++
		}
	}
	temp3 := floor2[i]
	temp4 := floor2[i]
	for {
		if floor2ture[i] == 1 {
			temp4--
			if temp4 == 0 {
				break
			}
			if i == 3 {
				i = 0
			}
			i++
		} else {
			if i == 3 {
				i = 0
			}
			i++

		}
	}

	sum := temp1 + temp3
	fmt.Println(sum)
}
