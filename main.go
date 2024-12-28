package main

import (
	"fmt"
	"math"
	"os"
	"raylibtest/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	x = float32((800 / 2))
	y = float32((600 / 2))

	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	STEPCT        = 10
)

func main() {
	rl.InitWindow(800, 600, "Shoota")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	player := player.NewPlayer(rl.NewRectangle(x-5, y-(100/3.5), 25, 100), rl.Black)
	origin := rl.NewVector2(25/2, 80)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// handle key
		if rl.IsKeyDown(rl.KeyA) {
			player.AngleAt -= float64(75 * rl.GetFrameTime())
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.AngleAt += float64(75 * rl.GetFrameTime())
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			bulletOrigin := rl.NewVector2(15/2, 40/2)
			bullet := rl.NewRectangle(player.Rect.X, player.Rect.Y, 15, 40)
			rl.DrawText("You shot", 0, 0, 20, rl.Black)
			positions, err := GetStepAdditions(bullet, player.AngleAt, STEPCT)
			if err != nil {
				os.Exit(1)
			}

			// Initial
			bullet.X = positions[0].X
			bullet.Y = positions[0].Y

			fmt.Println("Bullet initial position: ", bullet.X, bullet.Y)
			var frameCt int

			for i := 1; i < STEPCT; i++ {
				if frameCt%10 == 0 {
					rl.DrawRectanglePro(bullet, bulletOrigin, float32(player.AngleAt), rl.RayWhite)
					bullet.X = positions[i].X
					bullet.Y = positions[i].Y
					fmt.Println("Bullet position at step ", i, ": ", positions[i].X, positions[i].Y)
					rl.DrawRectanglePro(bullet, bulletOrigin, float32(player.AngleAt), rl.Gold)
				}
				frameCt++
			}

			rl.DrawRectanglePro(bullet, bulletOrigin, float32(player.AngleAt), rl.RayWhite)

		}

		// Clamp to 360 degrees
		if player.AngleAt > 360 {
			player.AngleAt = 0
		}

		rl.DrawRectanglePro(player.Rect, origin, float32(player.AngleAt), rl.Black)

		rl.EndDrawing()
	}
}

func GetStepAdditions(rect rl.Rectangle, angle float64, stepCt int) ([]rl.Vector2, error) {

	initX := rect.X
	initY := rect.Y

	if stepCt <= 0 {
		return nil, fmt.Errorf("nonzero step count given")
	}

	angleRad := (angle * rl.Deg2rad) * 0.1

	dx := float32(math.Cos(angleRad))
	dy := float32(math.Sin(angleRad))

	distToEdge := GetDistanceToScreen(rect, angle)

	stepSize := distToEdge / float64(stepCt)
	positions := []rl.Vector2{rl.NewVector2(rect.X, rect.Y)}

	for i := 0; i < stepCt; i++ {
		initX += float32(stepSize) * dx
		initY += float32(stepSize) * dy
		positions = append(positions, rl.NewVector2(initX, initY))
	}

	return positions, nil
}

func GetDistanceToScreen(rect rl.Rectangle, angle float64) float64 {

	angleRad := angle * rl.Deg2rad

	dx := float32(math.Cos(angleRad))
	dy := float32(math.Sin(angleRad))

	left := float32(math.Inf(1))
	if dx != 0 {
		left = rect.X / dx
	}

	right := float32(math.Inf(1))
	if dx != 0 {
		right = (SCREEN_WIDTH - rect.X) / dx
	}

	top := float32(math.Inf(1))
	if dy != 0 {
		top = rect.Y / dy
	}

	bottom := float32(math.Inf(1))
	if dy != 0 {
		bottom = (SCREEN_HEIGHT - rect.Y) / dy
	}

	minOne := math.Min(float64(left), float64(right))
	minTwo := math.Min(float64(top), float64(bottom))
	return math.Min(minOne, minTwo)

}
