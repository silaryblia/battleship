package ship

type Ship struct {
	Size     int
	Cells    []*Cell
	Hits     []bool
	HitCount int
	Sunk     bool
}

type Cell struct {
	X, Y int
}

type ShipImpl interface {
	GetStatus() string
	HandleShoot(x int, y int) (bool, bool)
}

// gs - status ship
// hs - ship orientation, dead/no

func (s *Ship) GetStatus() string {
	if s.Sunk {
		return "потоплен"
	}
	if s.HitCount > 0 {
		return "ранен"
	}
	return "цел"
}

//
//	if b.Cells[x][y] == "*" {
//		return "промах"
//	}
//	if b.Cells[x][y] == "X" {
//		// Проверяем, остались ли рядом живые части "■"
//		directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
//		for _, dir := range directions {
//			nx, ny := x+dir[0], y+dir[1]
//			for nx >= 0 && nx < 10 && ny >= 0 && ny < 10 {
//				if b.Cells[nx][ny] == "■" {
//					return "попадание"
//				}
//				if b.Cells[nx][ny] == "." || b.Cells[nx][ny] == "*" {
//					break
//				}
//				nx += dir[0]
//				ny += dir[1]
//			}
//		}
//		return "мертв"
//	}
//	return "unknown"
//}

// обрабатывает выстрел по кораблю
func (s *Ship) HandleShoot(x, y int) (bool, bool) {
	for i, cell := range s.Cells {
		if cell.X == x && cell.Y == y {
			if s.Hits[i] {
				return false, false // уже стреляли
			}
			s.Hits[i] = true
			s.HitCount++

			if s.HitCount == s.Size {
				s.Sunk = true
				return true, true // попадание и потопление
			}
			return true, false // попадание
		}
	}
	return false, false // промах
}
