package main

import "fmt"

var cont int = 0

func main() {
    n := 3
    hanoi(n, "A", "B", "C")
    fmt.Printf("Fim de Jogo: %d passos no total \n", cont)
}

func hanoi(n int, discoA string, discoB string, discoC string) {
    if n == 1 {
        cont += 1
        fmt.Printf("Passo %d: Disco %d de %s se mover para %s \n", cont, n, discoA, discoC)
    } else {
        hanoi(n - 1, discoA, discoC, discoB)
        cont += 1
        fmt.Printf("Passo %d: Disco %d de %s se mover para %s \n", cont, n, discoA, discoC)

        hanoi(n - 1, discoB, discoA, discoC)
    }

}



