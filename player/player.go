package player

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	Rect    rl.Rectangle
	Color   rl.Color
	AngleAt float64
}

func NewPlayer(rec rl.Rectangle, color rl.Color) *Player {
	return &Player{Rect: rec, Color: color, AngleAt: 0.0}
}
