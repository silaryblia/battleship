package tests

import (
	board2 "awesomeProject2/board"
	"awesomeProject2/ship"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestCanPlaceShip(t *testing.T) {
	board := board2.NewBoard()

	// test valid placement
	if !board.CanPlaceShip(0, 0, 3, true) {
		t.Error("Должны можно разместить корабль в начале доски")
	}

	// test invalid placement (за границами)
	if board.CanPlaceShip(8, 8, 3, true) {
		t.Error("Не должны разместить корабль за границами")
	}

	// test overlapping ships
	board.PlaceShip(2, 2, 2, true)

	if board.CanPlaceShip(2, 2, 3, true) {
		t.Error("Не должны можно разместить корабль поверх существующего")
	}
}

func TestGetCells(t *testing.T) {
	ctrl := gomock.NewController(t) //
	defer ctrl.Finish()             //
	board := board2.NewBoard()
	cells := board.GetCells()

	if len(cells) != 10 || len(cells[0]) != 10 {
		t.Errorf("ожидался массив 10x10, а получили %v", cells)
	}
}

func TestHandleShoot(t *testing.T) {
	b := board2.NewBoard()
	_ = b.PlaceShip(0, 0, 2, true)

	// попадание
	msg, hit, status := b.HandleShoot(0, 0)
	if !hit || status != ship.Hurt || msg != "попадание" {
		t.Errorf("ожидали попадание, получили %s %v %v", msg, hit, status)
	}
	// Промах
	msg, hit, status = b.HandleShoot(5, 5)
	if msg != "промах" || hit {
		t.Errorf("ожидали промах, получили %s %v %v", msg, hit, status)
	}
}

func TestPlaceShip(t *testing.T) {
	b := board2.NewBoard()

	err := b.PlaceShip(0, 0, 3, true)
	if err != nil {
		t.Errorf("не ожидали ошибку при пересечении кораблей")
	}
}

func TestPlaceShipRandom(t *testing.T) {
	b := board2.NewBoard()
	b.PlaceShipRandom()

	if len(b.GetShips()) != 10 {
		t.Errorf("ожидали 10 кораблей, получили %d", len(b.GetShips()))
	}
}

func TestPrintBoard(t *testing.T) {
	b := board2.NewBoard()
	_ = b.PlaceShip(0, 0, 2, true)

	// проверим, что метод вызывается без паники
	b.PrintBoard(true)
	b.PrintBoard(false)
}

func TestGetShips(t *testing.T) {
	b := board2.NewBoard()
	if len(b.GetShips()) != 0 {
		t.Errorf("ожидали пустой список кораблей, получили %d", len(b.GetShips()))
	}

	_ = b.PlaceShip(0, 0, 2, true)
	ships := b.GetShips()
	if len(ships) != 1 {
		t.Errorf("ожидали 1 корабль, получили %d", len(ships))
	}
	if ships[0].Size != 2 {
		t.Errorf("ожидали корабль размера 2, получили %d", ships[0].Size)
	}
}

//func TestGetShips(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//	board := board2.NewBoard()
//	ships := board.GetShips()
//
//}

//HandleShoot(x, y int) (message string, hit bool, sunk ship.Status)
//**GetCells() [10][10]string
//GetShips() []*ship.Ship
//PlaceShip(x, y, size int, horizontal bool) error
//PlaceShipRandom()
//**CanPlaceShip(x, y, size int, horizontal bool) bool
//PrintBoard(showShips bool) // PlaceFleet

//mockCells := mocks.NewMockBoardInterface(ctrl)
//
//board := board2.NewBoard()
//
//expected := [10][10]string{}
//expected[0][0] = "."
//mockCells.EXPECT().GetCells().Return(expected)
//cells := board.GetCells()
//if cells[0][0] != "." {
//	t.Errorf("ожидали '.' в [0][0], а получили %v", cells[0][0])
//}

//if len(cells) != 2 || len(cells[0]) != 2 {
//	t.Errorf("Ожидается массив 2х2 : %v", cells)
//}
//
//if len(cells) != 10 || len(cells[0]) != 10 {
//	t.Errorf("ожидался массив 10x10, а получили %v", cells)
//}

//if len(cells) == 10 {
//	t.Logf("доска правильного размера :%v", cells)
//} else {
//	t.Errorf("Ожидается массив 10х10 :%v", cells)
//}

//func (m *MockBoardInterface) GetCells() [10][10]string {
//	m.ctrl.T.Helper()
//	ret := m.ctrl.Call(m, "GetCells")
//	ret0, _ := ret[0].([10][10]string)
//	return ret0

//// GetCells возвращает массив клеток доски
//func (b *Board) GetCells() [10][10]string {
//	return b.Cells
//}

//func TestPlaceShipRandom(t *testing.T) {
//	t.Run("Размещение всех кораблей", func(t *testing.T) {
//		board := NewBoard()
//		board.PlaceShipRandom()
//
//		// Проверяем количество кораблей
//		expectedShips := 10 // 4+3+3+2+2+2+1+1+1+1
//		if len(board.Ships) != expectedShips {
//			t.Errorf("Ожидалось %d кораблей, получено %d", expectedShips, len(board.Ships))
//		}
//
//		// Проверяем общее количество клеток кораблей
//		shipCells := 0
//		for i := 0; i < 10; i++ {
//			for j := 0; j < 10; j++ {
//				if board.Cells[i][j] == "■" {
//					shipCells++
//				}
//			}
//		}
//
//		expectedCells := 20 // 4+3+3+2+2+2+1+1+1+1
//		if shipCells != expectedCells {
//			t.Errorf("Ожидалось %d клеток кораблей, получено %d", expectedCells, shipCells)
//		}
//	})
//}
//
//func TestCanPlaceShip(t *testing.T) {
//	board := NewBoard()
//
//	tests := []struct {
//		name        string
//		x, y, size  int
//		horizontal  bool
//		expected    bool
//		description string
//	}{
//		{"Валидное размещение", 0, 0, 2, true, true, "Корабль помещается"},
//		{"За пределами", 9, 9, 2, true, false, "Выход за границы"},
//		{"Пересечение", 5, 5, 1, true, true, "Одиночный корабль"},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			result := board.CanPlaceShip(tt.x, tt.y, tt.size, tt.horizontal)
//			if result != tt.expected {
//				t.Errorf("%s: ожидалось %t, получено %t", tt.description, tt.expected, result)
//			}
//		})
//	}
//}

//func TestHandleShoot(t *testing.T)
//	board := NewBoard()
//
//	// Создаем тестовый корабль
//	ship := &Ship{
//		Size:  1,
//		Cells: []*Cell{{X: 5, Y: 5}},
//		Hits:  make([]bool, 1),
//	}
//	board.Ships = []*Ship{ship}
//	board.Cells[5][5] = "■"
//
//	tests := []struct {
//		x, y     int
//		expected string
//	}{
//		{5, 5, "попадание"},    // Попадание
//		{0, 0, "промах"},       // Промах
//		{5, 5, "уже стреляли"}, // Повторный выстрел
//	}
//
//	for i, tt := range tests {
//		result, _, _ := board.HandleShoot(tt.x, tt.y)
//		if result != tt.expected {
//			t.Errorf("Тест %d: ожидалось '%s', получено '%s'", i+1, tt.expected, result)
//		}
//	}
//}
