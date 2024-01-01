package goridor

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 420
	boardSize    = 9
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		cx, cy := ebiten.CursorPosition()

		boardX, boardY := g.board.Size()

		mx := ((ScreenWidth - boardX) / 2) + tileSize/2
		my := ((ScreenHeight - boardY) / 2) + tileSize/2

		selected := g.board.Tile((cx-mx)/tileSize, (cy-my)/tileSize)
		g.board.MovePlayer(0, selected)

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw-bw)/2 + tileSize/2
	y := (sh-bh)/2 + tileSize/2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
	ebitenutil.DebugPrint(screen, "Hello, World")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}
