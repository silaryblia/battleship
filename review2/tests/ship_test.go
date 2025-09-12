package tests

import (
	"testing"

	"awesomeProject2/ship"
)

func TestGetStatus(t *testing.T) {
	s := &ship.Ship{Status: ship.Alive}
	if s.GetStatus() != "цел" {
		t.Errorf("Ожидалось 'цел', получили %s", s.GetStatus())
	}

	s.Status = ship.Hurt
	if s.GetStatus() != "ранен" {
		t.Errorf("Ожидалось 'ранен', получили %s", s.GetStatus())
	}

	s.Status = ship.Dead
	if s.GetStatus() != "убит" {
		t.Errorf("Ожидалось 'убит', получили %s", s.GetStatus())
	}
}

func TestHandleShootHitAndSunk(t *testing.T) {
	s := &ship.Ship{
		Size:   2,
		Cells:  []*ship.Cell{{X: 0, Y: 0}, {X: 0, Y: 1}},
		Hits:   make([]bool, 2),
		Status: ship.Alive,
	}

	//первый выстрел — попали, корабль
	hit, status := s.HandleShoot(0, 0)
	if !hit || status != ship.Hurt {
		t.Errorf("Ожидали попадание и Hurt, получили hit=%v status=%v", hit, status)
	}

	// второй выстрел — добили, корабль потоплен
	hit, status = s.HandleShoot(0, 1)
	if !hit || status != ship.Dead {
		t.Errorf("Ожидали потопление, получили hit=%v status=%v", hit, status)
	}
}

func TestHandleShootMiss(t *testing.T) {
	s := &ship.Ship{
		Size:   1,
		Cells:  []*ship.Cell{{X: 5, Y: 5}},
		Hits:   make([]bool, 1),
		Status: ship.Alive,
	}

	hit, status := s.HandleShoot(0, 0)
	if hit {
		t.Error("ожидали промах")
	}
	if status != ship.Alive {
		t.Errorf("Ожидали Alive, получили %v", status)
	}
}

func TestHandleShootTwice(t *testing.T) {
	s := &ship.Ship{
		Size:   1,
		Cells:  []*ship.Cell{{X: 5, Y: 5}},
		Hits:   make([]bool, 1),
		Status: ship.Alive,
	}

	// первый раз — попали
	hit, _ := s.HandleShoot(5, 5)
	if !hit {
		t.Error("ожидали попадание")
	}
	// второй раз — в ту же клетку, не должно засчитаться
	hit, status := s.HandleShoot(5, 5)
	if hit {
		t.Errorf("повторный выстрел в ту же клетку не должен считаться")
	}
	if status != ship.Hurt && status != ship.Dead {
		t.Errorf("Ожидали Hurt или Dead, получили %v", status)
	}
}
