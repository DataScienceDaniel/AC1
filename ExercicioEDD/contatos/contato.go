package contatos

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail

}

func (c *Contato) EditarEmail(indice int, email string, contatos *[5]Contato) {
	contatos[indice].AlterarEmail(email)
}