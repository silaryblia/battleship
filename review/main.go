package main

import (
	"awesomeProject2/board"
	"awesomeProject2/game"
	"fmt"
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

func main() {

	g := game.NewGame("Игрок 1", "Игрок 2")
	fmt.Println("🚀 Морской бой начался! ")
	fmt.Println("======================")

	// Пример игрового цикла
	coords := [][2]int{
		{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}, {5, 5}, {6, 6}, {7, 7}, {8, 8}, {9, 9},
		{0, 9}, {1, 8}, {2, 7}, {3, 6}, {4, 5}, {5, 4}, {6, 3}, {7, 2}, {8, 1}, {9, 0},
		{2, 3}, {5, 6}, {1, 8}, {7, 2}, {4, 4}, {9, 1}, {3, 7}, {6, 5}, {0, 9}, {8, 0},
	}

	fmt.Println("Доска Игрока 1:")
	printBoard(g.Players[0].Board)

	fmt.Println("Доска Игрока 2:")
	printBoard(g.Players[1].Board)

	for i, coord := range coords {
		if g.IsEnded() {
			break
		}
		fmt.Printf("\n🎯 Ход %d:\n", i+1)

		// Получаем текущего игрока
		currentPlayer := g.Players[0] // временно, пока не реализован TakeMove
		if !currentPlayer.Active {
			currentPlayer = g.Players[1]
		}

		fmt.Printf("Ходит: %s\n", currentPlayer.Name)
		fmt.Printf("Статус: %s\n", currentPlayer.StatusPlayer())

		// Выполняем ход
		result, hit := g.Round(coord[0], coord[1])
		fmt.Printf("Выстрел в [%d,%d]: %s\n", coord[0], coord[1], result)
		fmt.Printf("Попадание: %t\n", hit)

		// Показываем поле противника (упрощенно)
		fmt.Println("Поле противника:")
		enemy := g.GetEnemyPlayer(currentPlayer) // нужно добавить этот метод
		if enemy != nil {
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					fmt.Printf("%s ", enemy.Board.Cells[i][j]) // ← поле противника
				}
				fmt.Println()
			}
		}

		// Завершение игры
		if g.IsEnded() {
			winner := g.GetWinner()
			if winner != nil {
				fmt.Printf("\n🎉 Победитель: %s!\n", winner.Name)
			} else {
				fmt.Println("\n🤝 Ничья!")
			}
		} else {
			fmt.Println("\n⏸️ Игра прервана")
		}

		fmt.Println("\nСпасибо за игру!")

		//	time.Sleep(5 * time.Second)
	}
}
