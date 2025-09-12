package tests

import (
	"awesomeProject2/mocks"
	"awesomeProject2/player"
	"awesomeProject2/ship"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestDoMoveHit(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBoard := mocks.NewMockBoardInterface(ctrl)

	mockBoard.EXPECT().
		HandleShoot(1, 1).
		Return("попадание", true, ship.Hurt)

	enemy := &player.Player{Name: "Enemy", Board: mockBoard}
	current := &player.Player{Name: "Current", Active: true}

	msg, hit := current.DoMove(enemy, 1, 1)

	if !hit || msg != "попадание" {
		t.Errorf("ожидали попадание, получили msg=%s hit=%v", msg, hit)
	}
}

func TestDoMoveNotActive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBoard := mocks.NewMockBoardInterface(ctrl)
	enemy := &player.Player{Name: "Enemy", Board: mockBoard}
	current := &player.Player{Name: "Current", Active: false}

	msg, hit := current.DoMove(enemy, 5, 5)

	if hit {
		t.Error("неактивный игрок не должен стрелять")
	}
	if msg != "не ваш ход" {
		t.Errorf("ожидали 'не ваш ход', получили %s", msg)
	}
}

func TestGiveUp(t *testing.T) {
	p := &player.Player{
		Name:   "Test",
		Alive:  true,
		Active: true,
	}
	p.GiveUp()

	if p.Alive {
		t.Error("игрок должен быть мертв после сдачи")
	}
	if p.Active {
		t.Error("игрок не должен быть активным после сдачи")
	}
}

func TestStatusPlayers(t *testing.T) {
	p := &player.Player{Name: "Test"}

	p.Alive = false
	if p.StatusPlayer() != "игрок сдался" {
		t.Error("ожидали статус 'игрок сдался'")
	}
	p.Alive, p.Active = true, true
	if p.StatusPlayer() != "играйте (ваш ход)" {
		t.Error("ожидали статус 'играйте (ваш ход)'")
	}
	p.Active = false
	if p.StatusPlayer() != "подождите" {
		t.Error("ожидали статус 'подождите'")
	}
}
