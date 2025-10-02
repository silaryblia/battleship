package main

// Ship представляет корабль
type Ship struct {
	Size     int
	Cells    []*Cell
	Hits     []bool
	HitCount int
	Sunk     bool
}

type ShipStatus int

const (
	Alive ShipStatus = iota // цел   0
	Hurt                    // ранен   1
	Dead                    // мертв   2
)

type ShipImpl interface {
	GetStatus() ShipStatus
	HandleShoot(x int, y int) bool
}

func (b *Ship) GetStatus() ShipStatus {
	return ShipStatus(Alive) //test
}

func (b *Ship) HandleShoot(x int, y int) bool {
	for i := range b.Cells {
		if b.Cells[i].X == x && b.Cells[i].Y == y {
			// Если уже попадали в эту клетку
			if b.Hits[i] {
				return true
			}

			// Отмечаем попадание
			b.Hits[i] = true
			b.HitCount++

			// Проверяем, потоплен ли корабль
			if b.HitCount == len(b.Cells) {
				b.Sunk = true
			}

			return true
		}
	}
	return false
}
