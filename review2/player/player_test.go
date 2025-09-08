package player

import (
	"awesomeProject2/board"
	"testing"
)

func TestDoMove(t *testing.T) {
	// add 2 players
	attacker := &Player{Name: "Attacker", Active: true, Alive: true}
	defender := &Player{Name: "Defender", Active: false, Alive: true}

	// add boards
	attacker.Board = board.NewBoard()
	defender.Board = board.NewBoard()

	// add ship - defender
	defender.Board.Cells[3][3] = "■"

	// test shot
	result, hit := attacker.DoMove(defender, 3, 3) // попадание
	if !hit || result != "попадание" {
		t.Errorf("Ожидалось попадание, получено: %s, hit=%t", result, hit)
	}

	result, hit = attacker.DoMove(defender, 0, 0) // промах
	if hit || result != "промах" {
		t.Errorf("Ожидался промах, получено: %s, hit=%t", result, hit)
	}
}
