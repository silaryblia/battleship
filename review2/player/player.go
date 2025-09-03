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
	//Hit    bool
}

type PlayerImpl interface {
	NewPlayer(name string) *Player
	DoMove(otherPlayers []Player) error //dm - player, coord shoot, moves 1-2-1-2 (dlia dvoih)
	TakeMove(int, int)                  //tm - hod
	GiveUp()                            //gu - sdatsia
	StatusPlayer() bool                 //sp - igraet/net
}

func (p *Player) DoMove(enemy *Player, x, y int) (string, bool) {
	if !p.Active {
		return "не ваш ход", false
	}
	// Стреляем по доске ПРОТИВНИКА
	message, hit, _ := enemy.Board.HandleShoot(x, y)
	return message, hit
}

// GiveUp - игрок сдался
func (p *Player) GiveUp() {
	p.Alive = false
	p.Active = false
	fmt.Printf("%s сдался!\n", p.Name)
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
