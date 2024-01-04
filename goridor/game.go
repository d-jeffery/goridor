package goridor

import (
	"github.com/beefsack/go-astar"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

const (
	ScreenWidth  = 420
	ScreenHeight = 420
	boardSize    = 9
)

type Game struct {
	board      *Board
	boardImage *ebiten.Image
	turn       int
}

func (g *Game) Update() error {
	if g.doPlayerMove(g.turn) {
		g.turn++
		g.turn %= 2 // TODO: make dynamic
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		g.boardImage = ebiten.NewImage(g.board.Size())
	}
	screen.Fill(backgroundColor)
	g.board.Draw(g.boardImage, g.turn)

	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
	x := (sw-bw)/2 + tileSize/2
	y := (sh-bh)/2 + tileSize/2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func NewGame() (*Game, error) {
	g := &Game{}
	var err error
	g.board, err = NewBoard(boardSize)
	g.turn = 0
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *Game) doPlayerMove(turn int) bool {
	g.board.pawns[turn].tile.occupied = false

	if g.board.pawns[turn].human {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			cx, cy := ebiten.CursorPosition()

			boardX, boardY := g.board.Size()

			mx := ((ScreenWidth - boardX) / 2) + tileSize/2
			my := ((ScreenHeight - boardY) / 2) + tileSize/2

			selected := g.board.Tile((cx-mx)/tileSize, (cy-my)/tileSize)
			if selected != nil {
				return g.board.MovePlayer(turn, selected)
			}
		}
	} else {
		path, _, found := astar.Path(g.board.Tile(0, 0), g.board.pawns[turn].tile)
		if !found {
			log.Println("Could not find path")
			return false
		} else {
			return g.board.MovePlayer(turn, path[1].(*Tile))
		}
	}
	g.board.pawns[turn].tile.occupied = true
	return false
}
