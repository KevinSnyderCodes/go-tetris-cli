package main

import (
	"log"
	"time"

	termbox "github.com/nsf/termbox-go"
)

// Position represents a location in a 2D plane.
type Position struct {
	x int
	y int
}

// Game represents a Tetris game.
type Game struct {
	fg termbox.Attribute
	bg termbox.Attribute

	active []*Position
	blocks [][]bool
}

// Tick advances the Tetris game by one tick.
func (o *Game) Tick() {
	for _, position := range o.active {
		if position.y == 19 {
			o.place()
			return
		}
		if o.blocks[position.y+1][position.x] {
			o.place()
			return
		}
	}
	for _, position := range o.active {
		position.y++
	}
}

// Rotate rotates the active piece in the Tetris game.
func (o *Game) Rotate() {}

// MoveDown moves the active piece in the Tetris game down.
func (o *Game) MoveDown() {
	for _, position := range o.active {
		if position.y == 19 {
			o.place()
			return
		}
		if o.blocks[position.y+1][position.x] {
			o.place()
			return
		}
	}
	for _, position := range o.active {
		position.y++
	}
}

// MoveLeft moves the active piece in the Tetris game left.
func (o *Game) MoveLeft() {
	for _, position := range o.active {
		if position.x == 0 {
			return
		}
		if o.blocks[position.y][position.x-1] {
			return
		}
	}
	for _, position := range o.active {
		position.x--
	}
}

// MoveRight moves the active piece in the Tetris game right.
func (o *Game) MoveRight() {
	for _, position := range o.active {
		if position.x == 9 {
			return
		}
		if o.blocks[position.y][position.x+1] {
			return
		}
	}
	for _, position := range o.active {
		position.x++
	}
}

// Drop drops the active piece in the Tetris game.
func (o *Game) Drop() {}

// place places the active piece in its final location.
func (o *Game) place() {
	// Move active piece to blocks
	for _, position := range o.active {
		o.blocks[position.y][position.x] = true
	}
	o.active = []*Position{
		&Position{x: 3, y: 3},
		&Position{x: 3, y: 4},
		&Position{x: 3, y: 5},
		&Position{x: 4, y: 5},
	}

	// Remove lines
	for y := range o.blocks {
		completeRow := true
		for x := range o.blocks[y] {
			completeRow = o.blocks[y][x]
			if !completeRow {
				break
			}
		}
		if completeRow {
			for y > 0 {
				o.blocks[y] = o.blocks[y-1]
				y--
			}
		}
	}
}

// Draw draws the Tetris game.
func (o *Game) Draw() {
	termbox.Clear(o.fg, o.bg)

	// Draw box
	termbox.SetCell(0, 0, '╔', o.fg, o.bg)
	termbox.SetCell(11, 0, '╗', o.fg, o.bg)
	termbox.SetCell(0, 21, '╚', o.fg, o.bg)
	termbox.SetCell(11, 21, '╝', o.fg, o.bg)
	for i := 0; i < 10; i++ {
		termbox.SetCell(1+i, 0, '═', o.fg, o.bg)
		termbox.SetCell(1+i, 21, '═', o.fg, o.bg)
	}
	for i := 0; i < 20; i++ {
		termbox.SetCell(0, 1+i, '║', o.fg, o.bg)
		termbox.SetCell(11, 1+i, '║', o.fg, o.bg)
	}

	// Draw active piece
	for _, position := range o.active {
		termbox.SetCell(1+position.x, 1+position.y, '█', o.fg, o.bg)
	}
	// Draw blocks
	for y := range o.blocks {
		for x := range o.blocks[y] {
			if o.blocks[y][x] {
				termbox.SetCell(1+x, 1+y, '█', o.fg, o.bg)
			}
		}
	}

	termbox.HideCursor()
	termbox.Flush()
}

func getKey(key chan termbox.Key) {
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			key <- ev.Key
		}
	}
}

func run() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	defer termbox.Close()

	ticker := time.NewTicker(500 * time.Millisecond)

	key := make(chan termbox.Key)
	go getKey(key)

	game := &Game{
		fg: termbox.ColorWhite,
		bg: termbox.ColorBlack,
		active: []*Position{
			&Position{x: 3, y: 3},
			&Position{x: 3, y: 4},
			&Position{x: 3, y: 5},
			&Position{x: 4, y: 5},
		},
		blocks: [][]bool{
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
			[]bool{false, false, false, false, false, false, false, false, false, false},
		},
	}

	for {
		select {
		case <-ticker.C:
			game.Tick()
			game.Draw()
		case k := <-key:
			switch k {
			case termbox.KeyEsc:
				return nil
			case termbox.KeyArrowUp:
				game.Rotate()
			case termbox.KeyArrowDown:
				game.MoveDown()
			case termbox.KeyArrowLeft:
				game.MoveLeft()
			case termbox.KeyArrowRight:
				game.MoveRight()
			case termbox.KeySpace:
				game.Drop()
			}
			game.Draw()
		}
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
