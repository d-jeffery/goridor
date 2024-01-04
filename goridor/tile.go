package goridor

import "github.com/beefsack/go-astar"

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
