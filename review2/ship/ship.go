package ship

type Status int

const (
	Alive Status = iota // 0 - цел
	Hurt                // 1 - ранен
	Dead                // 2 - потоплен
)

type Ship struct {
	Size     int
	Cells    []*Cell
	Hits     []bool
	HitCount int
	Status   Status
}

type Cell struct {
	X, Y int
}

// Status представляет статус корабля через iota

type ShipImpl interface {
	GetStatus() string
	HandleShoot(x int, y int) (bool, bool)
}

// gs - status ship
// hs - ship orientation, dead/no

func (s *Ship) GetStatus() string {
	switch s.Status {
	case Alive:
		return "цел"
	case Hurt:
		return "ранен"
	case Dead:
		return "убит"
	default:
		return "неизвестно"
	}
}

// обрабатывает выстрел по кораблю
func (s *Ship) HandleShoot(x, y int) (bool, Status) {
	for i, cell := range s.Cells {
		if cell.X == x && cell.Y == y {
			if s.Hits[i] {
				return false, s.Status // уже стреляли
			}
			s.Hits[i] = true
			s.HitCount++

			if s.HitCount == s.Size {
				s.Status = Dead
				return true, Dead // потопление
			}

			s.Status = Hurt
			return true, Hurt // ранен
		}
	}
	return false, s.Status // промах
}
