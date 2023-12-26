package goridor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type Pawn struct {
	x, y, size int
	pawnColor  color.RGBA
}

func (p *Pawn) Draw(boardImage *ebiten.Image) {
	vector.DrawFilledCircle(boardImage, float32((p.x*tileSize)+(tileSize/2)), float32((p.y*tileSize)+(tileSize/2)), float32(p.size), p.pawnColor, false)
	//vector.DrawFilledCircle(boardImage, float32(tileSize/2), float32(tileSize/2), float32(p.size), player1, false)
}

func (p *Pawn) Move() int {
	for {
		switch {
		case inpututil.IsKeyJustPressed(ebiten.KeyW):
			return UP
		case inpututil.IsKeyJustPressed(ebiten.KeyS):
			return DOWN
		case inpututil.IsKeyJustPressed(ebiten.KeyA):
			return LEFT
		case inpututil.IsKeyJustPressed(ebiten.KeyD):
			return RIGHT
		}
	}
}
