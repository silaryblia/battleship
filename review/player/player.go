package player

import (
	"awesomeProject2/board"
	"fmt"
)

type Player struct {
	Name   string
	Board  *board.Board
	Active bool // играет или нет
	Alive  bool // сдался или нет
	Hit    bool
}

type PlayerImpl interface {
	NewPlayer(name string) *Player
	DoMove(otherPlayers []Player) error //dm - player, coord shoot, moves 1-2-1-2 (dlia dvoih)
	TakeMove(int, int)                  //tm - hod
	GiveUp()                            //gu - sdatsia
	StatusPlayer() bool                 //sp - igraet/net
}

func NewPlayer(name string) *Player {
	player := &Player{
		Name:   name,
		Board:  board.NewBoard(),
		Active: false,
		Alive:  true,
	}
	player.Board.PlaceShipRandom()
	return player
}

func (p *Player) DoMove(enemy *Player, x, y int) (string, bool) {
	if !p.Active {
		return "не ваш ход", false
	}

	result, hit := enemy.Board.HandleShoot(x, y)

	// Используем hit для логики смены хода, но возвращаем текстовый результат
	if !hit {
		// ход переходит другому
		p.Active = false
		enemy.Active = true
	}

	// Всегда возвращаем текстовый результат и флаг попадания
	return result, hit
}

// TakeMove - чей ход
func TakeMove(players ...*Player) *Player {
	for _, p := range players {
		if p.Active {
			return p
		}
	}
	return nil
}

// GiveUp - игрок сдался
func (p *Player) GiveUp(command string) {
	if command == "сдаюсь" {
		p.Alive = false
		p.Active = false
		fmt.Printf("%s сдался!\n", p.Name)
	}
}

// StatusPlayer - вернёт статус игрока
func (p *Player) StatusPlayer() string {
	if !p.Alive {
		return "игрок сдался"
	}
	if p.Active {
		return "играйте (ваш ход)"
	}
	return "подождите"
}
