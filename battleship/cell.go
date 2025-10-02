package main

// Cell представляет клетку поля
type Cell struct {
	X, Y    int
	HasShip bool
}

type CellInfo interface {
	GetStatus(hasShip bool) *Cell
}

func (p *Cell) GetStatus(hasShip bool) *Cell {
	return &Cell{
		HasShip: hasShip,
	}
}
