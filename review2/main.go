package main

import (
	"awesomeProject2/board"
	"awesomeProject2/game"
	"awesomeProject2/player"

	"fmt"
	"math/rand"
	"time"
)

func printBoard(b *board.Board) {
	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 10; j++ {
			fmt.Printf("%s ", b.Cells[i][j])
		}
		fmt.Println()
	}
}

func printEnemyBoard(b *board.Board) {
	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 10; j++ {
			// Показываем только выстрелы, но не корабли противника
			if b.Cells[i][j] == "X" || b.Cells[i][j] == "*" {
				fmt.Printf("%s ", b.Cells[i][j])
			} else {
				fmt.Printf(". ") // скрываем непотопленные корабли
			}
		}
		fmt.Println()
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.NewGame("Игрок 1", "Игрок 2")
	fmt.Println("🚀 Морской бой начался! ")
	fmt.Println("======================")

	// Показываем начальные доски
	fmt.Println("Доска Игрока 1:")
	printBoard(g.Players[0].Board)

	fmt.Println("Доска Игрока 2:")
	printBoard(g.Players[1].Board)

	// Бесконечный цикл до конца игры
	moveCount := 0
	for !g.IsEnded() {
		moveCount++

		// Получаем текущего игрока
		currentPlayer := g.Players[0] // временно, пока не реализован TakeMove
		if !currentPlayer.Active {
			currentPlayer = g.Players[1]
		}

		// Находим противника
		var enemy *player.Player
		for _, p := range g.Players {
			if p != currentPlayer {
				enemy = p
				break
			}
		}

		fmt.Printf("\n🎯 Ход %d:\n", moveCount)
		fmt.Printf("Ходит: %s\n", currentPlayer.Name)
		fmt.Printf("Статус: %s\n", currentPlayer.StatusPlayer())

		// Генерируем случайные координаты
		x := rand.Intn(10)
		y := rand.Intn(10)

		// Выполняем ход
		result, hit := g.Round(x, y)
		fmt.Printf("Выстрел в [%d,%d]: %s\n", x, y, result)
		fmt.Printf("Попадание: %t\n", hit)

		// Показываем поле противника
		if enemy != nil {
			fmt.Println("Поле противника:")
			printEnemyBoard(enemy.Board)
		}

		// показываем свою доску
		fmt.Println("Ваша доска:")
		printBoard(currentPlayer.Board)

		//	time.Sleep(100 * time.Millisecond) // небольшая пауза
	}

	// Завершение игры
	winner := g.GetWinner()
	if winner != nil {
		fmt.Printf("\n🎉 Победитель: %s!\n", winner.Name)
	} else {
		fmt.Println("\n🤝 Ничья!")
	}

	fmt.Println("\nСпасибо за игру!")

	// Показываем финальные доски
	fmt.Println("\nФинальная доска Игрока 1:")
	printBoard(g.Players[0].Board)

	fmt.Println("\nФинальная доска Игрока 2:")
	printBoard(g.Players[1].Board)

	//	time.Sleep(5 * time.Second)
}
