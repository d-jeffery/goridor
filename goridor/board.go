package goridor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const (
	tileSize   = 42
	tileMargin = 4
)

type Pawn struct {
	size      int
	tile      *Tile
	pawnColor color.RGBA
	human     bool
}

type Board struct {
	size  int
	tiles map[int]map[int]*Tile
	pawns [4]*Pawn
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
			if j > 0 {
				tile.neighbor[UP] = b.Tile(i, j-1)
			}
			if i < b.size {
				tile.neighbor[RIGHT] = b.Tile(i+1, j)
			}
			if j < b.size {
				tile.neighbor[DOWN] = b.Tile(i, j+1)
			}
			if i > 0 {
				tile.neighbor[LEFT] = b.Tile(i-1, j)
			}
		}
	}

	b.pawns[0] = &Pawn{10, b.tiles[4][8], player1, true}
	b.pawns[1] = &Pawn{10, b.tiles[4][0], player2, false}
}

func (b *Board) Draw(boardImage *ebiten.Image, playerTurn int) {

	vector.DrawFilledRect(boardImage, float32(0), float32(0), float32(tileSize*b.size), float32(tileSize*b.size), frameColor, false)

	location := b.pawns[playerTurn].tile
	for _, moves := range location.neighbor {
		if moves == nil {
			continue
		}
		if b.IsOccupied(moves) {
			vector.DrawFilledRect(boardImage, float32(moves.x*tileSize), float32(moves.y*tileSize), float32(tileSize), float32(tileSize), blockedTiles, false)
		} else {
			vector.DrawFilledRect(boardImage, float32(moves.x*tileSize), float32(moves.y*tileSize), float32(tileSize), float32(tileSize), neighbourTiles, false)
		}
	}

	for j := 0; j < b.size; j++ {
		for i := 0; i < b.size; i++ {
			vector.StrokeRect(boardImage, float32(i*tileSize), float32(j*tileSize), float32(tileSize), float32(tileSize), tileMargin, backgroundColor, false)
		}
	}

	for _, p := range b.pawns {
		if p == nil {
			continue
		}
		vector.DrawFilledCircle(boardImage, float32((p.tile.x*tileSize)+(tileSize/2)), float32((p.tile.y*tileSize)+(tileSize/2)), float32(p.size), p.pawnColor, false)
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
		pawns: [4]*Pawn{}, //[4]*Pawn{{4, 8, 10, player1}, {4, 0, 10, player2}, nil, nil},
	}
	b.Init()
	var err error
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Board) IsOccupied(tile *Tile) bool {
	for _, p := range b.pawns {
		if p == nil {
			continue
		}
		if p.tile == tile {
			return true
		}
	}
	return false
}

func (b *Board) MovePlayer(turn int, tile *Tile) bool {
	location := b.pawns[turn].tile
	for _, n := range location.neighbor {
		if n == tile && !b.IsOccupied(tile) {
			b.pawns[turn].tile = tile
			return true
		}
	}
	return false
}
