package main

import (
	"fmt"
	ctt "ExercicioEDD/contatos"
	"ExercicioEDD/arrays"
	"bufio"
	"os"
)

func main() {
	var contatos [5]ctt.Contato
	var nome, email, opcaoContato, opcaoEmail string
	leitor := bufio.NewReader(os.Stdin)
	c1 := ctt.Contato {
		Nome: "Daniel",
		Email: "ddd",
	}

	fmt.Println(contatos)
	arrays.AdicionaContato(c1, &contatos)
	fmt.Println(contatos)

	// Loop para adicionar, remover ou sair da adição de contato
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcaoContato)

		switch opcaoContato {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = arrays.AdicionaContato((contatos *[5]ctt.Contato) {Nome: nome, Email: email}, contatos)
		case "2":
			contatos = arrays.ExcluiContato(contatos)
		default:
			return
		}

		fmt.Println(contatos)

		// Loop para continuar ou editar com seu email
		for  {
			fmt.Print("Digite (1) para continuar com o email, (2) para editar o email: ")
				fmt.Scanln(&opcaoEmail)

			switch opcaoEmail {
			case "1":
			fmt.Println("Voce Continuou com o seu email: ", email)
			case "2":
				fmt.Println("Informe seu novo email: ")
				fmt.Scanln(&email)
			default:
				return
			}

			fmt.Println(contatos)
			}


	}



}

