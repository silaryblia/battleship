package tests

import (
	"awesomeProject2/game"
	"awesomeProject2/mocks"

	"testing"

	//"github.com/golang/mock/gomock"
	"go.uber.org/mock/gomock"
)

func TestBattle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPlayer1 := mocks.NewMockPlayerInterface(ctrl)
	mockPlayer2 := mocks.NewMockPlayerInterface(ctrl)

	// Настраиваем DoMove только для первого игрока
	mockPlayer1.EXPECT().
		DoMove(gomock.Any(), 2, 3).
		Return("попадание", true)

	// Имена игроков — просто строки
	player1Name := "Игрок 1"
	player2Name := "Игрок 2"

	// Пример использования в тесте
	result, hit := mockPlayer1.DoMove(nil, 2, 3)
	t.Logf("%s сделал ход: результат=%s, попадание=%v",
		player1Name, result, hit)

	t.Logf("%s ждёт хода", player2Name)

	mockPlayer2.EXPECT().
		DoMove(gomock.Any(), 2, 1).
		Return("промах", true)

	// Пример использования в тесте
	result1, hit1 := mockPlayer2.DoMove(nil, 2, 1)
	t.Logf("%s сделал ход: результат=%s, промах=%v",
		player2Name, result1, hit1)

	t.Logf("%s ждёт хода", player1Name)
}

func TestSwitchMove(t *testing.T) {
	g := game.NewGame("A", "B")

	first := g.TakeMove()
	g.SwitchMove()
	second := g.TakeMove()

	if first == second {
		t.Error("после SwitchMove должен быть другой игрок")
	}
}

func TestAddPlayer(t *testing.T) {
	g := game.NewGame("A", "B")
	g.AddPlayer("C")

	if len(g.Players) != 3 {
		t.Errorf("ожидали 3 игроков, получили %d", len(g.Players))
	}
}

func TestRoundAndIsEnded(t *testing.T) {
	g := game.NewGame("A", "B")

	// Стреляем по всем кораблям второго игрока
	for _, s := range g.Players[1].Board.GetShips() {
		for _, c := range s.Cells {
			g.Round(c.X, c.Y)
		}
	}

	if !g.IsEnded() {
		t.Error("игра должна завершиться после уничтожения всех кораблей второго игрока")
	}

	winner := g.GetWinner()
	if winner == nil || winner.Name != "A" {
		t.Errorf("ожидали победителя A, получили %+v", winner)
	}
}

func TestTakeMove(t *testing.T) {
	g := game.NewGame("A", "B")

	current := g.TakeMove()
	if current == nil || current.Name != "A" {
		t.Errorf("ожидали, что первый ход у A, получили %+v", current)
	}
}

func TestGetEnemy(t *testing.T) {
	g := game.NewGame("A", "B")

	current := g.TakeMove()
	enemy := g.GetEnemy()

	if current == enemy {
		t.Error("противник не должен совпадать с текущим игроком")
	}
	if enemy.Name != "B" {
		t.Errorf("ожидали врага B, получили %s", enemy.Name)
	}
}

func TestAllShipsSunk(t *testing.T) {
	g := game.NewGame("A", "B")

	if g.AllShipsSunk(g.Players[0]) {
		t.Error("корабли игрока A не должны быть потоплены")
	}

	for _, s := range g.Players[0].Board.GetShips() {
		s.Status = 2 // dead
	}

	if !g.AllShipsSunk(g.Players[0]) {
		t.Error("ожидали, что все корабли игрока A потоплены")
	}
}

//func TestIsEnded(t *testing.T) {
//	g := game.NewGame("A", "B")
//
//	for _, s := range g.Players[1].Board.GetShips() {
//		s.Status = 2 // dead
//	}
//
//	if !g.IsEnded() {
//		t.Error("игра должна закончиться, если все корабли второго игрока уничтожены")
//	}
//
//	winner := g.GetWinner()
//	if winner == nil || winner.Name != "P1" {
//		t.Error("ожидали победителя P1")
//	}
//}

//AddPlayer(name string) *player.Player
//TakeMove() *player.Player //tm - hod
//SwitchMove()
//GetEnemy() *player.Player
//Round(x, y int) (string, bool)
//IsEnded() bool
//allShipsSunk(player *player.Player) bool
//GetWinner() *player.Player
