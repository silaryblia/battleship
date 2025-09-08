package board

import (
	"testing"
)

func TestPlaceShipRandom(t *testing.T) {
	t.Run("Размещение всех кораблей", func(t *testing.T) {
		board := NewBoard()
		board.PlaceShipRandom()

		// Проверяем количество кораблей
		expectedShips := 10 // 4+3+3+2+2+2+1+1+1+1
		if len(board.Ships) != expectedShips {
			t.Errorf("Ожидалось %d кораблей, получено %d", expectedShips, len(board.Ships))
		}

		// Проверяем общее количество клеток кораблей
		shipCells := 0
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				if board.Cells[i][j] == "■" {
					shipCells++
				}
			}
		}

		expectedCells := 20 // 4+3+3+2+2+2+1+1+1+1
		if shipCells != expectedCells {
			t.Errorf("Ожидалось %d клеток кораблей, получено %d", expectedCells, shipCells)
		}
	})
}

func TestCanPlaceShip(t *testing.T) {
	board := NewBoard()

	tests := []struct {
		name        string
		x, y, size  int
		horizontal  bool
		expected    bool
		description string
	}{
		{"Валидное размещение", 0, 0, 2, true, true, "Корабль помещается"},
		{"За пределами", 9, 9, 2, true, false, "Выход за границы"},
		{"Пересечение", 5, 5, 1, true, true, "Одиночный корабль"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := board.canPlaceShip(tt.x, tt.y, tt.size, tt.horizontal)
			if result != tt.expected {
				t.Errorf("%s: ожидалось %t, получено %t", tt.description, tt.expected, result)
			}
		})
	}
}

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

//func TestMax(t *testing.T) {
//	// Arrange
//	numbers := &board.Board{}
//	expected := 15
//	// Act
//	result := board.NewBoard()
//	// Assert
//	if result != numbers {
//		t.Errorf("Incorrect result Expect %d, got", expected)
//	}
//}
