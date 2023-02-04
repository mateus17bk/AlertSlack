package main

import (
	"alertslack/slack"
)

func main() {
	slack.SendSlack("Alerta de servidor down: Google\nErro: Erro ao conectar no servidor\nHorário: 06/07/2022 15:46")
}
