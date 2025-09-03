package board

import (
	"awesomeProject2/ship"
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	Cells [10][10]string
	Ships []*ship.Ship
}

type BoardImpl interface {
	PrintBoard(x, y, size int, horizontal bool) // PlaceFleet
	PlaceShipRandom()
	HandleShoot(x, y int)
	canPlaceShip(x, y, size int, horizontal bool) //PlaceShipCoords
	placeShip(x, y, size int, horizontal bool)
}

func NewBoard() *Board {
	b := &Board{}
	// Инициализируем клетки
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Cells[i][j] = "."
		}
	}
	return b
}

func (b *Board) HandleShoot(x, y int) (string, bool, ship.Status) {
	// Реализация обработки выстрела
	if x < 0 || x >= 10 || y < 0 || y >= 10 {
		return "неверные координаты", false, ship.Alive
	}

	// Проверяем, уже стреляли ли сюда
	if b.Cells[x][y] == "*" || b.Cells[x][y] == "X" {
		return "уже стреляли сюда", false, ship.Alive
	}

	// Проверяем попадание в корабли
	for _, s := range b.Ships {
		hit, status := s.HandleShoot(x, y)
		if hit {
			b.Cells[x][y] = "X"
			// Используем существующий GetStatus для сообщения
			message := "попадание"
			if status == ship.Dead {
				message = "корабль потоплен"
			}
			return message, true, status
		}
	}

	b.Cells[x][y] = "*"
	return "промах", false, ship.Alive
}

func (b *Board) canPlaceShip(x, y, size int, horizontal bool) bool {
	dx, dy := 0, 1
	if !horizontal {
		dx, dy = 1, 0
	}
	// Один цикл для всех проверок
	for i := 0; i < size; i++ {
		cellX, cellY := x+i*dx, y+i*dy
		// Проверка границ
		if cellX < 0 || cellX >= 10 || cellY < 0 || cellY >= 10 {
			return false
		}
		// Проверка клетки и всех её соседей
		for checkX := cellX - 1; checkX <= cellX+1; checkX++ {
			for checkY := cellY - 1; checkY <= cellY+1; checkY++ {
				if checkX >= 0 && checkX < 10 && checkY >= 0 && checkY < 10 {
					if b.Cells[checkX][checkY] != "." {
						return false
					}
				}
			}
		}
	}
	return true
}

// правильно размещает корабли
func (b *Board) PlaceShipRandom() {
	fleet := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}
	rand.Seed(time.Now().UnixNano())

	for _, size := range fleet {
		placed := false
		for !placed {
			x := rand.Intn(10)
			y := rand.Intn(10)
			//
			horizontal := rand.Intn(2) == 0
			if b.canPlaceShip(x, y, size, horizontal) {
				// СОЗДАЕМ И РАЗМЕЩАЕМ КОРАБЛЬ
				ship := b.placeShip(x, y, size, horizontal)
				b.Ships = append(b.Ships, ship)
				placed = true
			}
		}
	}
}

// создает и размещает корабль на доске
func (b *Board) placeShip(x, y, size int, horizontal bool) *ship.Ship {
	newShip := &ship.Ship{
		Size:  size,
		Cells: make([]*ship.Cell, size),
		Hits:  make([]bool, size),
	}
	dx, dy := 0, 1 // По умолчанию горизонтально
	if !horizontal {
		dx, dy = 1, 0 // Вертикально
	}
	// 121322133213
	for i := 0; i < size; i++ {
		cellX, cellY := x+i*dx, y+i*dy
		b.Cells[cellX][cellY] = "■"
		newShip.Cells[i] = &ship.Cell{X: cellX, Y: cellY}
	}
	return newShip
}

func (b *Board) PrintBoard(showShips bool) {
	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 10; j++ {
			if showShips {
				//	Показываем все клетки (для своей доски)
				fmt.Printf("%s ", b.Cells[i][j])
			} else {
				// Показываем только выстрелы (для доски противника
				if b.Cells[i][j] == "X" || b.Cells[i][j] == "*" {
					fmt.Printf("%s ", b.Cells[i][j])
				} else {
					fmt.Printf(". ") // скрываем непотопленные корабли
				}
			}
		}
		fmt.Println()
	}
}
