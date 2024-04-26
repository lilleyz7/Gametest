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

	p := player.NewPlayer()
	es := enemy.NewEnemyStore()
	state := gamestate.NewGameState(p, es)

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		if rl.IsKeyPressed(rl.KeyEscape) {
			state.Paused = true
		}
		if !state.Paused {
			p.Update()
			player.UpdateProjectiles()
			state.Update()
			es.UpdateEnemies()

		}

		rl.BeginDrawing()

		p.Draw()
		player.DrawProjectiles()
		es.DrawEnemies()

		rl.ClearBackground(rl.DarkGray)

		rl.EndDrawing()
	}
}
