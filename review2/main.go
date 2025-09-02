package main

import (
	"awesomeProject2/game"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	g := game.NewGame("Игрок 1", "Игрок 2")
	fmt.Println("🚀 Морской бой начался! ")
	fmt.Println("======================")

	//g.AddPlayer("Игрок 4")

	// Показываем доски всех игроков
	for i, player := range g.Players {
		fmt.Printf("Доска %s:\n", player.Name)
		player.Board.PrintBoard()
		if i < len(g.Players)-1 {
			fmt.Println()
		}
	}

	// Карта уже сделанных выстрелов для каждого игрока
	shotHistory := make(map[string]map[[2]int]bool)
	for _, player := range g.Players {
		shotHistory[player.Name] = make(map[[2]int]bool)
	}

	// Бесконечный цикл до конца игры
	moveCount := 0
	for !g.IsEnded() {
		moveCount++
		current := g.TakeMove()
		enemy := g.GetEnemy()
		fmt.Printf("\nХодит: %s\n", current.Name)

		var x, y int
		var shot [2]int
		found := false

		for attempts := 0; attempts < 100; attempts++ {
			x, y = rand.Intn(10), rand.Intn(10)
			shot = [2]int{x, y}
			if !shotHistory[current.Name][shot] {
				found = true
				break
			}
		}

		if !found {
			for x = 0; x < 10; x++ {
				for y = 0; y < 10; y++ {
					shot = [2]int{x, y}
					if !shotHistory[current.Name][shot] {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}

		if !found {
			fmt.Println("Нет свободных клеток для выстрела!")
			break
		}

		//// Генерируем координаты пока не найдем непростреленную клетку
		//for {
		//	x, y = rand.Intn(10), rand.Intn(10)
		//	shot = [2]int{x, y}
		//	if !shotHistory[current.Name][shot] {
		//		break
		//	}
		//	fmt.Printf("Клетка [%d,%d] уже прострелена, ищем другую...\n", x, y)
		//}

		// Помечаем клетку как простреленную
		shotHistory[current.Name][shot] = true
		result, hit := g.Round(x, y)
		fmt.Printf("Выстрел в [%d,%d]: %s\n", x, y, result)
		if hit {
			fmt.Println("✅ Попадание!")
		} else {
			fmt.Println("❌ Промах!")
		}

		// Показываем поле противника после хода
		enemy = g.GetEnemy()
		fmt.Printf("Поле %s:\n", enemy.Name)
		enemy.Board.PrintEnemyBoard()

		// Небольшая пауза для читаемости
		// time.Sleep(100 * time.Millisecond)
	}

	//
	if winner := g.GetWinner(); winner != nil {
		fmt.Printf("\n🎉 Победитель: %s!\n", winner.Name)
		fmt.Printf("Всего ходов: %d\n", moveCount)
	}

	fmt.Println("\nСпасибо за игру!")

	// Показываем финальные доски
	fmt.Println("\nФинальные доски:")
	for _, player := range g.Players {
		fmt.Printf("Доска %s:\n", player.Name)
		player.Board.PrintBoard()
		fmt.Println()
		//	time.Sleep(5 * time.Second)
	}
}
