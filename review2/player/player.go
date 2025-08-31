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

type PlayerManager struct {
	Players []*Player
}

type PlayerImpl interface {
	NewPlayer(name string) *Player
	DoMove(otherPlayers []Player) error //dm - player, coord shoot, moves 1-2-1-2 (dlia dvoih)
	TakeMove(int, int)                  //tm - hod
	GiveUp()                            //gu - sdatsia
	StatusPlayer() bool                 //sp - igraet/net
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		Players: make([]*Player, 0),
	}
}

// AddPlayer добавляет игрока в массив
func (pm *PlayerManager) AddPlayer(name string) *Player {
	player := NewPlayer(name)
	pm.Players = append(pm.Players, player)
	return player
}

// TakeMove
func (pm *PlayerManager) TakeMove() *Player {
	for _, player := range pm.Players {
		if player.Active {
			return player
		}
	}
	return nil
}

func (pm *PlayerManager) SwitchMove(current *Player) {
	current.Active = false
	for _, p := range pm.Players {
		if p != current {
			p.Active = true
			break
		}
	}
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
	message, hit, _ := enemy.Board.HandleShoot(x, y)
	// Используем hit для логики смены хода, но возвращаем текстовый результат
	if !hit {
		// ход переходит другому
		p.Active = false
		enemy.Active = true
	}
	// Всегда возвращаем текстовый результат и флаг попадания
	return message, hit
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

//// GetPlayers возвращает массив всех игроков
//func (pm *PlayerManager) GetPlayers() []*Player {
//	return pm.Players
//}
