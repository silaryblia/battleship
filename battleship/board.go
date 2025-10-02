package main

import (
	"fmt"
	"math/rand"
	"time"
)

type orientation int

const (
	gorizontal = 0
	vertical   = 1
)

const boardSize = 10

// Board представляет игровое поле
type Board struct {
	Cells [boardSize][boardSize]Cell
	Ships []Ship
}

// // ShipInfo информация о корабле
// type ShipInfo struct {
// 	Size  int
// 	Cells []*Cell
// }

type BoardImpl interface {
	NewBoard()
	AddShip(x, y, size int, orientation orientation)
	PlaceShipsRandomly(size int) bool
}

// NewBoard создает новое поле
func (b *Board) NewBoard() {

	// Инициализируем клетки
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			b.Cells[i][j] = Cell{X: i, Y: j, HasShip: false}
		}
	}

	fmt.Println("  A B C D E F G H I J")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i+1)
		for j := 0; j < 10; j++ {
			cell := b.Cells[i][j]
			if cell.HasShip {
				fmt.Print("■ ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// placeShip размещает один корабль заданного размера
func (p *Board) placeShip(size int, o orientation) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		//	orientation := r.Intn(2) == 0

		var x, y int
		if o == 0 {
			x = r.Intn(10)
			y = r.Intn(11 - size)
		} else {
			x = r.Intn(11 - size)
			y = r.Intn(10)
		}

		if p.CanPlaceShip(x, y, size, o) {
			p.AddShip(x, y, size, o)
			break
		}
	}
}

// CanPlaceShip проверяет, можно ли разместить корабль
func (b *Board) CanPlaceShip(x, y, size int, orientation orientation) bool {
	for i := -1; i <= size; i++ {
		for j := -1; j <= 1; j++ {
			var checkX, checkY int

			if orientation == 0 { // 0 - horizontal
				checkX = x + j
				checkY = y + i
			} else { // 1 -  vertical
				checkX = x + i
				checkY = y + j
			}

			if checkX < 0 || checkX >= 10 || checkY < 0 || checkY >= 10 {
				continue
			}

			if b.Cells[checkX][checkY].HasShip {
				return false
			}
		}
	}
	return true
}

// PlaceShipsRandomly размещает все корабли случайным образом
func (b *Board) PlaceShipsRandomly(size int, o orientation) bool {
	for attempts := 0; attempts < 1000; attempts++ {
		//	orientation := rand.Intn(2) // 0 - horizont, 1 - vertical

		var x, y int
		if o == 0 {
			x = rand.Intn(10)
			y = rand.Intn(11 - size)
		} else {
			x = rand.Intn(11 - size)
			y = rand.Intn(10)
		}

		if b.CanPlaceShip(x, y, size, o) {
			ship := Ship{Size: size}
			for i := 0; i < size; i++ {
				var cell *Cell
				if o == 0 {
					cell = &b.Cells[x][y+i]
				} else {
					cell = &b.Cells[x+i][y]
				}
				cell.HasShip = true
				ship.Cells = append(ship.Cells, cell)
			}
			b.Ships = append(b.Ships, ship)
			return true
		}
	}
	return false // не смогли поставить после 1000 попыток
}

// // canPlaceShip — приватная проверка (оставляем вспомогательную)
// func (b *Board) canPlaceShip(x, y, size int, orientation int) bool {
//     for i := -1; i <= size; i++ {
//         for j := -1; j <= 1; j++ {
//             var checkX, checkY int
//             if orientation == 0 {
//                 checkX = x + j
//                 checkY = y + i
//             } else {
//                 checkX = x + i
//                 checkY = y + j
//             }

//             if checkX < 0 || checkX >= 10 || checkY < 0 || checkY >= 10 {
//                 continue
//             }
//             if b.Cells[checkX][checkY].HasShip {
//                 return false
//             }
//         }
//     }
//     return true
// }

func (p *Player) PlaceShipsRandomly(o orientation) {
	ships := []int{4, 3, 3, 2, 2, 2, 1, 1, 1, 1}

	for _, size := range ships {
		if !p.Board.PlaceShipsRandomly(size, o) {
			fmt.Println("❌ Ошибка: не удалось разместить корабль!")
		}
	}
}

// AddShip добавляет корабль на поле
func (b *Board) AddShip(x, y, size int, orientation orientation) {
	ship := Ship{Size: size}
	for i := 0; i < size; i++ {
		var cell *Cell
		if orientation == 0 { // horizontal
			cell = &b.Cells[x][y+i]
		} else { // vertical
			cell = &b.Cells[x+i][y]
		}
		cell.HasShip = true
		ship.Cells = append(ship.Cells, cell)
	}
	b.Ships = append(b.Ships, ship)
}

// IsValidPlacement проверяет корректность расположения всех кораблей
func (b *Board) IsValidPlacement() bool {
	return b.GetShipsCount() == 10
}

// GetShipsCount возвращает количество кораблей
func (b *Board) GetShipsCount() int {
	return len(b.Ships)
}

// GetShipCells возвращает количество клеток кораблей
func (b *Board) GetShipCells() int {
	count := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if b.Cells[i][j].HasShip {
				count++
			}
		}
	}
	return count
}
