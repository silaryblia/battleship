package game

import (
	"awesomeProject2/board"
	"awesomeProject2/player"
	"awesomeProject2/ship"
	"fmt"
)

type Game struct {
	Players            []*player.Player
	Running            bool
	currentPlayerIndex int
}

type GameImpl interface {
	Round(x, y int) (string, bool)
	IsEnded() bool
	GetWinner() *player.Player
}

//r - shagi
//ie - end/net

func NewGame(playerNames ...string) *Game {
	if len(playerNames) < 2 {
		panic("для игры нужно минимум 2 игрока")
	}

	var players []*player.Player

	// Создаем каждого игрока напрямую
	for i, name := range playerNames {
		playerBoard := board.NewBoard() // Новая доска для каждого игрока
		playerBoard.PlaceShipRandom()   // Размещаем корабли на доске

		player := &player.Player{
			Name:   name,
			Board:  playerBoard,
			Active: i == 0, // первый игрок активен
			Alive:  true,
		}
		players = append(players, player)
	}

	return &Game{
		Players:            players,
		Running:            true,
		currentPlayerIndex: 0,
	}
}

func (g *Game) AddPlayer(name string) *player.Player {
	playerBoard := board.NewBoard()
	playerBoard.PlaceShipRandom()

	player := &player.Player{
		Name:   name,
		Board:  playerBoard,
		Active: false,
		Alive:  true,
	}
	g.Players = append(g.Players, player)
	return player
}

// TakeMove возвращает текущего активного игрока
func (g *Game) TakeMove() *player.Player {
	return g.Players[g.currentPlayerIndex]
}

// SwitchMove переключает ход на следующего игрока
func (g *Game) SwitchMove() {
	if g.currentPlayerIndex == 0 {
		g.currentPlayerIndex = 1
	} else {
		g.currentPlayerIndex = 0
	}
	for i, player := range g.Players {
		player.Active = (i == g.currentPlayerIndex)
	}
}

// GetEnemy возвращает противника для текущего игрока
func (g *Game) GetEnemy() *player.Player {
	currentIndex := g.currentPlayerIndex
	// Для 2 игроков просто возвращаем другого игрока
	enemyIndex := (currentIndex + 1) % len(g.Players)
	return g.Players[enemyIndex]
}

func (g *Game) Round(x, y int) (string, bool) {
	if !g.Running || g.IsEnded() {
		return "Игра завершена", false
	}
	current := g.TakeMove()
	enemy := g.GetEnemy()
	if current == nil || enemy == nil {
		return "Нет активных игроков", false
	}

	result, hit := current.DoMove(enemy, x, y)

	if !hit {
		g.SwitchMove()
	} else {
		// При попадании проверяем, не потоплен ли корабль
		// Если корабль потоплен, игрок может стрелять снова
		if result == "корабль потоплен" {
			// Можно оставить ход у текущего игрока
			fmt.Printf("⭐ %s получает дополнительный ход!\n", current.Name)
		}
	}
	return result, hit
}

// IsEnded проверяет окончание игры
func (g *Game) IsEnded() bool {
	// Считаем живых игроков
	//aliveCount := 0
	activePlayers := 0
	for _, player := range g.Players {
		// Если игрок сдался
		if player.Alive && !g.allShipsSunk(player) {
			activePlayers++
			//	aliveCount++
		}
	}
	return activePlayers <= 1
	// return aliveCount <= 1
}

// allShipsSunk проверяет потоплены ли все корабли
func (g *Game) allShipsSunk(player *player.Player) bool {
	for _, s := range player.Board.Ships {
		if s.Status != ship.Dead {
			return false
		}
	}
	return true
}

func (g *Game) GetWinner() *player.Player {
	if !g.IsEnded() {
		return nil
	}
	// Ищем игрока с непотопленными кораблями
	for _, player := range g.Players {
		if player.Alive && !g.allShipsSunk(player) {
			return player
		}
	}
	return nil
}
