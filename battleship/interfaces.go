package interfaces

// Renderable интерфейс для отрисовки
type Renderable interface {
	Render()
}

// ShipPlacer интерфейс для размещения кораблей
type ShipPlacer interface {
	PlaceShipsRandomly()
}

// BoardOperations интерфейс для операций с полем
type BoardOperations interface {
	Renderable
	AddShip(shipSize int, cells []struct{ X, Y int })
	CanPlaceShip(x, y, size int, isHorizontal bool) bool
	IsValidPlacement() bool
	GetShipsCount() int
	GetShipCells() int
}

// PlayerOperations интерфейс для операций с игроком
type PlayerOperations interface {
	Renderable
	ShipPlacer
	GetBoard() BoardOperations
}

// Game интерфейс для игры
type Game interface {
	Start()
	Stop()
}
