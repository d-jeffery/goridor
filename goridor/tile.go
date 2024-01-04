package goridor

import (
	"github.com/beefsack/go-astar"
)

type Path struct {
	from *Tile
	to   *Tile
	Cost float64
}

type Tile struct {
	x, y     int
	neighbor [4]*Tile
	occupied bool
}

func NewTile(x, y int) *Tile {
	return &Tile{
		x:        x,
		y:        y,
		occupied: false,
	}
}

func (t *Tile) Up() *Tile {
	if t.neighbor[UP] == nil {
		return nil
	}
	return t.neighbor[UP]
}

func (t *Tile) Right() *Tile {
	if t.neighbor[RIGHT] == nil {
		return nil
	}
	return t.neighbor[RIGHT]
}

func (t *Tile) Down() *Tile {
	if t.neighbor[DOWN] == nil {
		return nil
	}
	return t.neighbor[DOWN]
}

func (t *Tile) Left() *Tile {
	if t.neighbor[LEFT] == nil {
		return nil
	}
	return t.neighbor[LEFT]
}

func (t *Tile) PathNeighbors() []astar.Pather {
	var moves []astar.Pather
	for _, move := range t.neighbor {
		if move != nil {
			moves = append(moves, move)
		}
	}
	return moves
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	return 1
}
