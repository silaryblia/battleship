package game

import "testing"

func TestNewGame(t *testing.T) {
	t.Run("Создание игры с 2 игроками", func(t *testing.T) {
		g := NewGame("Player1", "Player2")

		if len(g.Players) != 2 {
			t.Errorf("Ожидалось 2 игрока, получено %d", len(g.Players))
		}

		// Проверяем что первый игрок активен
		if !g.Players[0].Active {
			t.Errorf("Первый игрок должен быть активен")
		}

		if g.Players[1].Active {
			t.Errorf("Второй игрок должен быть активен")
		}
	})
}

func TestSwitchMove(t *testing.T) {
	g := NewGame("Player1", "Player2")

	// Первывй ход - игрок 1
	current := g.TakeMove()
	if current.Name != "Player1" {
		t.Errorf("Ожидался Player1, получен %s", current.Name)
	}

	// Меняем ход
	g.SwitchMove()
	current = g.TakeMove()
	if current.Name != "Player2" {
		t.Errorf("Ожидался Player2, полу %s", current.Name)
	}
}
