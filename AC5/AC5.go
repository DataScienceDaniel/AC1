package main

import "fmt"

// func main() {
	array := []int{2,4,6,8,10,30,40}
	alvo := 70

	resultado := achaArray(array, alvo)
	fmt.Println(resultado)

}

func achaArray(array []int, alvo int) []int{
	y := len(array) - 1
	for i := 0; i < len(array); {
		if (array[i] + array[y]) > alvo{
			y--
		}else if (array[i] + array[y]) < alvo{
			i++
		}else{
			return []int{array[i], array[y]}
		}
	}
	return []int{-1, -1}
}

