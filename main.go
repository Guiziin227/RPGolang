package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Criar um RPG de prompt onde a pessoa vai poder criar seu personagem.
Personagem --- Escolher o sexo, nome, classe;

Comand help para saber a lista de comando
Sistema de atacks
Vai ter monstros randomicos
XP

*/

// FUNCOES
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func criarPersonagem() cPerson {
	var name string

	nivel := 1
	vida := 50
	exp := 0

	fmt.Println("Qual será seu nome?")
	fmt.Scan(&name)

	person := cPerson{
		Name:  name,
		Nivel: nivel,
		Vida:  vida,
		Exp:   exp,
	}

	return person
}

//func criarMontro()

// STRUCTS
type cPerson struct {
	Name   string
	Ataque int
	Nivel  int
	Vida   int
	Exp    int
}

type Enemy struct {
	Name   string
	HP     int
	Damage int
}

func main() {

	fmt.Println("Bem vindo ao RPGolang")
	person := criarPersonagem()
	fmt.Printf("Muito bem %s, seja bem vindo(a)\n\nSeu nível inicial é %d voce tem %d de vida", person.Name, person.Nivel, person.Vida)

	enemies := []Enemy{
		{Name: "Teemo", HP: random(43, 58), Damage: random(3, 5)},
		{Name: "Renekton", HP: random(43, 58), Damage: random(0, 8)},
		{Name: "Evellyn", HP: random(43, 58), Damage: random(2, 6)},
	}

	fmt.Println("\nParece que uma batalha começou")
	time.Sleep(time.Second * 2)

	for {
		rand.Shuffle(len(enemies), func(i, j int) { enemies[i], enemies[j] = enemies[j], enemies[i] })

		// Selecionando um inimigo aleatório
		enemy := enemies[0]
		fmt.Printf("\n\nInimigo encontrado!\nNome: %s\nVida do Inimigo: %d\nDano do Inimigo: %d\n", enemy.Name, enemy.HP, enemy.Damage)

		for person.Vida > 0 && enemy.HP > 0 {
			fmt.Println("\n\nO que você gostaria de fazer?")
			fmt.Println("1. Atacar o Inimigo")
			fmt.Println("2. Tentar fugir")

			var choice int
			fmt.Scan(&choice)

			switch choice {
			case 1:
				damage := random(10, 20)
				enemy.HP -= damage
				fmt.Printf("Você deu %d de dano", damage)

				if enemy.HP < 0 {
					fmt.Println("Você derrotou o inimigo\n")
				}

				person.Vida -= enemy.Damage
				fmt.Printf("O inimigo causou %d de dano em você!\n", enemy.Damage)
				fmt.Printf("Sua vida atual: %d\nVida do Inimigo: %d\n", person.Vida, enemy.HP)

			case 2:
				// Tentar fugir
				fmt.Println("Você tenta fugir...")

				// Chance de 50% de sucesso na fuga
				if random(1, 101) <= 50 {
					fmt.Println("Você conseguiu fugir!")
					time.Sleep(5 * time.Second) // Espera 5 segundos
					break
				} else {
					fmt.Println("Você não conseguiu fugir!")

					person.Vida -= enemy.Damage
					fmt.Printf("O inimigo bloqueou sua fuga e te deu %d de dano!", enemy.Damage)
					fmt.Printf("Sua vida atual: %d\nVida do Inimigo: %d\n", person.Vida, enemy.HP)
				}

				// Verificando se o jogador foi derrotado

			} //final switch
			if enemy.HP <= 0 {
				fmt.Println("\n\nPARABENS VOCE DERROTOU O INIMIGO.")

			}
			if person.Vida <= 0 {
				fmt.Println("\n\nVocê foi derrotado! Fim de jogo.")
				return
			}

		} //final loop

		if person.Vida > 0 {
			fmt.Println("\n\nProcurando por outro inimigo...")
			time.Sleep(5 * time.Second) // Espera 5 segundos
		}

	}
}
