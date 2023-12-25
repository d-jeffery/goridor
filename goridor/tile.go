package goridor

import "github.com/beefsack/go-astar"

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

type Path struct {
	from *Tile
	to   *Tile
	Cost float64
}

type Tile struct {
	x, y     int
	neighbor [4]*Tile
}

func NewTile(x, y int) *Tile {
	return &Tile{
		x: x,
		y: y,
	}
}

func (t *Tile) Up() *Tile {
	return t.neighbor[UP]
}

func (t *Tile) Right() *Tile {
	return t.neighbor[RIGHT]
}

func (t *Tile) Down() *Tile {
	return t.neighbor[DOWN]
}

func (t *Tile) Left() *Tile {
	return t.neighbor[LEFT]
}

func (t *Tile) PathNeighbors() []astar.Pather {
	return []astar.Pather{
		t.Up(),
		t.Right(),
		t.Down(),
		t.Left(),
	}
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return 1
}
