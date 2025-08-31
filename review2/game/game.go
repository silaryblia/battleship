package game

import (
	"awesomeProject2/player"
)

type Game struct {
	PlayerManager *player.PlayerManager
	Running       bool
}

type GameImpl interface {
	Round(x, y int) (string, bool)
	IsEnded() bool
	GetWinner() *player.Player
}

//r - shagi
//ie - end/net

func NewGame(player1Name string, player2Name string) *Game {
	pm := player.NewPlayerManager()

	// Добавляем игроков в массив
	player1 := pm.AddPlayer(player1Name)
	player2 := pm.AddPlayer(player2Name)

	// первый игрок начинает
	player1.Active = true
	player2.Active = false

	return &Game{
		PlayerManager: pm,
		Running:       true,
	}
}

func (g *Game) Round(x, y int) (string, bool) {
	if !g.Running || g.IsEnded() {
		return "Игра завершена", false
	}
	// определение текущего игрока

	current := g.PlayerManager.TakeMove()
	if current == nil {
		return "Нет активных игроков", false
	}

	// Находим противника
	var enemy *player.Player
	for _, p := range g.PlayerManager.Players {
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
		g.PlayerManager.SwitchMove(current)
	}
	return result, hit
}

// IsEnded проверяет окончание игры
func (g *Game) IsEnded() bool {
	// Игра заканчивается, если:
	// 1. Один из игроков сдался
	// 2. У одного из игроков все корабли потоплены
	for _, player := range g.PlayerManager.Players {
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

	for _, player := range g.PlayerManager.Players {
		if !player.Alive || g.allShipsSunk(player) {
			// Возвращаем другого игрока
			for _, p := range g.PlayerManager.Players {
				if p != player {
					return p
				}
			}
		}
	}
	return nil
}

//func (g *Game) GetEnemyPlayer(current *player.Player) *player.Player {
//	for _, p := range g.Players {
//		if p != current {
//			return p
//		}
//	}
//	return nil
//}
