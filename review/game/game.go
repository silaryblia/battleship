package game

import (
	"awesomeProject2/player"
)

type Game struct {
	Players []*player.Player
	Running bool
}

type GameImpl interface {
	Round(x, y int) (string, bool)
	IsEnded() bool
	GetWinner() *player.Player
}

//r - shagi
//ie - end/net

func NewGame(player1Name string, player2Name string) *Game {
	player1 := player.NewPlayer(player1Name)
	player2 := player.NewPlayer(player2Name)

	// первый игрок начинает
	player1.Active = true
	player2.Active = false

	return &Game{
		Players: []*player.Player{player1, player2},
		Running: true,
	}
}

func (g *Game) Round(x, y int) (string, bool) {
	if !g.Running || g.IsEnded() {
		return "Игра завершена", false
	}
	// определение текущего игрока

	current := player.TakeMove(g.Players...)
	if current == nil {
		return "Нет активных игроков", false
	}

	// Находим противника
	var enemy *player.Player
	for _, p := range g.Players {
		if p != current {
			enemy = p
			break
		}
	}
	if enemy == nil {
		return "Не найден противник", false

	}
	result, hit := current.DoMove(enemy, x, y)

	// Если промах - переключаем ход
	if !hit {
		for _, p := range g.Players {
			p.Active = (p != current)
		}
	}
	return result, hit
}

//	// Проверяем окончание игры
//	if g.IsEnded() {
//		g.Running = false
//		winner := g.GetWinner()
//		if winner != nil {
//			return fmt.Sprintf("%s! %s побеждает", result, winner.Name)
//		}
//		return fmt.Sprintf("%s! Игра завершена", result)
//	}
//	return result
//}

// IsEnded проверяет окончание игры
func (g *Game) IsEnded() bool {
	// Игра заканчивается, если:
	// 1. Один из игроков сдался
	// 2. У одного из игроков все корабли потоплены
	for _, player := range g.Players {
		// Если игрок сдался
		if !player.Alive {
			return true
		}
		// Если у игрока все корабли потоплены
		if g.allShipsSunk(player) {
			return true
		}
	}
	return false
}

// allShipsSunk проверяет потоплены ли все корабли
func (g *Game) allShipsSunk(player *player.Player) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if player.Board.Cells[i][j] == "■" {
				return false
			}
		}
	}
	return true
}

func (g *Game) GetWinner() *player.Player {
	if !g.IsEnded() {
		return nil
	}

	for _, player := range g.Players {
		if !player.Alive || g.allShipsSunk(player) {
			// Возвращаем другого игрока
			for _, p := range g.Players {
				if p != player {
					return p
				}
			}
		}
	}
	return nil
}

func (g *Game) GetEnemyPlayer(current *player.Player) *player.Player {
	for _, p := range g.Players {
		if p != current {
			return p
		}
	}
	return nil
}
