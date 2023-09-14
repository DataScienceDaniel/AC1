package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)

}

func insereOrd(v *[M]int, n *int, novoValor int) {
	if *n == M {
		fmt.Println("Overflow")

	} else {
		if buscaOrd(*v, *n, novoValor) == -1 {
			v[*n] = novoValor
			*n++
		} else {
			fmt.Println("Elemento ja existe na tabela")
		}
	}
}

func buscaOrd(v [M]int , n, x int ) int {
	i := 0
	v[n] = x
	for v[i] == x {
		i++
	}
		if i == n + 1 || v[i] != x {
			return -1
		}
		return i
}
