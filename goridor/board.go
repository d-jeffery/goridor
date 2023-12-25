package goridor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	tileSize   = 42
	tileMargin = 4
)

type Board struct {
	size  int
	tiles map[int]map[int]*Tile
}

func (b *Board) Init() {
	// Create initial tiles
	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			if b.tiles[i] == nil {
				b.tiles[i] = map[int]*Tile{}
			}
			b.tiles[i][j] = NewTile(i, j)
		}
	}
	// Set neighbors
	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			tile := b.Tile(i, j)
			if i > 0 {
				tile.neighbor[UP] = b.Tile(i-1, j)
			}
			if j < b.size {
				tile.neighbor[RIGHT] = b.Tile(i, j+1)
			}
			if i < b.size {
				tile.neighbor[DOWN] = b.Tile(i+1, j)
			}
			if j > 0 {
				tile.neighbor[LEFT] = b.Tile(i, j-1)
			}
		}
	}
}

func (b *Board) Draw(boardImage *ebiten.Image) {

	vector.DrawFilledRect(boardImage, float32(0), float32(0), float32(tileSize*b.size), float32(tileSize*b.size), frameColor, false)

	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			vector.StrokeRect(boardImage, float32(i*tileSize), float32(j*tileSize), float32(tileSize), float32(tileSize), tileMargin, backgroundColor, false)
		}
	}
}

func (b *Board) Size() (int, int) {
	x := b.size*tileSize + (b.size+1)*tileMargin
	y := x
	return x, y
}

func (b *Board) Tile(x, y int) *Tile {
	if b.tiles[x] == nil {
		return nil
	}
	return b.tiles[x][y]
}

func NewBoard(size int) (*Board, error) {
	b := &Board{
		size:  size,
		tiles: map[int]map[int]*Tile{},
	}
	b.Init()
	var err error
	if err != nil {
		return nil, err
	}
	return b, nil
}
