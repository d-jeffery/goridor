package goridor

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

func (t *Tile) Move(dir int) *Tile {
	if t.neighbor[dir] == nil {
		return nil
	}
	return t.neighbor[dir]
}

func (t *Tile) Up() *Tile {
	return t.Move(UP)
}

func (t *Tile) Right() *Tile {
	return t.Move(RIGHT)
}

func (t *Tile) Down() *Tile {
	return t.Move(DOWN)
}

func (t *Tile) Left() *Tile {
	return t.Move(LEFT)
}
