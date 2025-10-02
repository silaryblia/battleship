package main

import (
	"fmt"
)

// Player представляет игрока
type Player struct {
	Name     string
	Board    *Board
	Opponent *Player
}

type PlayerImpl interface {
	GiveUp()
	DoMove(x int, y int) (hit bool, sunk bool, message string)
	TakeMove(x int, y int) (hit bool, sunk bool, message string)
	NewGame()
	NewPlayer(name string) *Player
}

func (p *Player) GiveUp() {
	fmt.Printf("Игрок %s сдался!\n", p.Name)
}

func (p *Player) NewGame() *Game {
	return &Game{
		Player:  p,
		Running: false,
	}
}

func (p *Player) DoMove(x int, y int) (hit bool, sunk bool, message string) {
	for _, ship := range p.Opponent.Board.Ships {
		if ship.HandleShoot(x, y) {
			hit = true
			if ship.Sunk {
				sunk = true
				return true, true, "Потопил!"
			}
			return true, false, "Попал!"
		}
	}
	return false, false, "Мимо!"
}

func (p *Player) TakeMove(x int, y int) (hit bool, sunk bool, message string) {
	// Просто проверяем попадание в наши корабли
	for _, ship := range p.Board.Ships {
		if ship.HandleShoot(x, y) {
			hit = true
			if ship.Sunk {
				sunk = true
				return true, true, "Потопили ваш корабль!"
			}
			return true, false, "Попали в ваш корабль!"
		}
	}
	return false, false, "Промахнулись!"
}
