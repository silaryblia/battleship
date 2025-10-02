package main

//	"battleship/game"

func main() {

	// Создаем игрока через фабрику
	var player PlayerImpl = NewPlayer("Игрок 1")

	// Создаем игру (через интерфейс)
	var game GameImpl = player.NewGame()

	// Запускаем игру
	game.Start()

	// defer чтобы остановить в конце
	defer game.Stop()
}

// player := PlayerImpl.NewPlayer("Игрок 1")
// // Создаем игру
// game := NewGame(player)

// // Запускаем игру
// game.Start()

// // Останавливаем игру
// defer game.Stop()

/*
	// Create player
	p := player.NewPlayer("Player 1")

	// Размещаем корабли случайным образом
	p.PlaceShipsRandomly()

	// Выводим поле с кораблями
	fmt.Printf("Поле игрока: %s\n", p.Name)
	p.Render()
*/
