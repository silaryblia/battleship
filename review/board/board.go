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

func (b *Board) HandleShoot(x, y int) (string, bool) {
	// Реализация обработки выстрела
	if x < 0 || x >= 10 || y < 0 || y >= 10 {
		return "неверные координаты", false
	}

	// Проверяем, уже стреляли ли сюда
	if b.Cells[x][y] == "*" || b.Cells[x][y] == "X" {
		return "уже стреляли сюда", false
	}

	if b.Cells[x][y] == "." {
		b.Cells[x][y] = "*"
		return "промах", false // ← ВАЖНО: возвращаем "промах" текстом
	}

	if b.Cells[x][y] == "■" {
		b.Cells[x][y] = "X"
		return "попадание", true
	}

	return "неизвестный результат", false
}

func (b *Board) canPlaceShip(x, y, size int, horizontal bool) bool {
	if horizontal {
		if y+size > 10 {
			return false
		}
		for i := 0; i < size; i++ {
			if b.Cells[x][y+i] != "." {
				return false
			}
		}
	} else {
		if x+size > 10 {
			return false
		}
		for i := 0; i < size; i++ {
			if b.Cells[x+i][y] != "." {
				return false
			}
		}
	}
	return true
}

func (b *Board) PrintBoard(x, y, size int, horizontal bool) {
	if horizontal {
		for i := 0; i < size; i++ {
			b.Cells[x][y+i] = "■"
		}
	} else {
		for i := 0; i < size; i++ {
			{
				b.Cells[x+i][y] = "■"
			}
		}
	}
	var board [10][10]string
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j] = "."
		}
	}
	fmt.Println("   A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d ", i+1)
		for j := 0; j < 10; j++ {
			fmt.Printf("%s ", board[i][j])
		}
		fmt.Println()
	}
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

	// positions := make([][2]int, size)
	if horizontal {
		for i := 0; i < size; i++ {
			b.Cells[x+i][y] = "■" // помечаем клетку
			newShip.Cells[i] = &ship.Cell{X: x, Y: y + i}
			//		positions[i] = [2]int{x + i, y + i}
		}
	} else {
		for i := 0; i < size; i++ {
			b.Cells[x+i][y] = "■" // помечаем клетку
			newShip.Cells[i] = &ship.Cell{X: x + i, Y: y}
			//		positions[i] = [2]int{x + i, y}
		}
	}
	return newShip
}
