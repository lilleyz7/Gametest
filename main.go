package main

import (
	"tester/enemy"
	gamestate "tester/gameState"
	"tester/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Shapes Runner Stuff")
	rl.SetTargetFPS(60)

	cameraOffset := rl.Vector2{X: float32(rl.GetScreenWidth() / 2), Y: float32(rl.GetScreenHeight() / 2)}
	camRotation := float32(0.0)
	camZoom := float32(1.0)

	p := player.NewPlayer()
	camera := rl.NewCamera2D(cameraOffset, p.Position, camRotation, camZoom)

	es := enemy.NewEnemyStore()
	state := gamestate.NewGameState(p, es)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyP) {
			state.Paused = !state.Paused
		}
		if !state.Paused {
			p.Update()
			camera.Target = p.Position
			player.UpdateProjectiles()
			state.Update()
			es.UpdateEnemies()

		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		rl.BeginMode2D(camera)

		p.Draw()
		player.DrawProjectiles()
		es.DrawEnemies()

		rl.EndMode2D()

		rl.EndDrawing()
	}
}
