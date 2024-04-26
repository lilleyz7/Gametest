package gamestate

import (
	"fmt"
	"math"
	"tester/enemy"
	"tester/player"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type gameState struct {
	Paused             bool
	CurrLevel          int
	TimeBtwSpawns      float32
	StartTime          float64
	EnemiesToSpawn     int
	TimeBtwRounds      int
	BaseEnemiesToSpawn int
	Player             *player.Player
	EnemyStore         *enemy.EnemyStore
}

func NewGameState(p *player.Player, es *enemy.EnemyStore) *gameState {
	return &gameState{
		Paused:             false,
		CurrLevel:          1,
		TimeBtwSpawns:      5.0,
		TimeBtwRounds:      2,
		StartTime:          rl.GetTime(),
		BaseEnemiesToSpawn: 10,
		EnemiesToSpawn:     2,
		Player:             p,
		EnemyStore:         es,
	}
}

func (state *gameState) Update() {
	state.CheckCollisions()
	if state.StartTime+float64(state.TimeBtwSpawns) < rl.GetTime() {
		state.EnemyStore.SpawnEnemy(state.Player)
		state.EnemiesToSpawn -= 1
		state.StartTime = rl.GetTime()
	}

	if state.EnemiesToSpawn <= 0 {
		time.Sleep(time.Duration(state.TimeBtwRounds))
		state.CurrLevel += 1
		state.EnemiesToSpawn = int(float64(state.BaseEnemiesToSpawn) * math.Floor(float64(state.CurrLevel)*0.5))
	}
}

func (state *gameState) CheckCollisions() {
	for _, proj := range player.ProjStore.Shots {
		for _, enemy := range state.EnemyStore.ActiveEnemies {
			hit := rl.CheckCollisionCircles(
				rl.Vector2{X: proj.PosX, Y: proj.PosY},
				float32(proj.Radius),
				enemy.Pos,
				enemy.Radius,
			)
			if hit {
				enemy.TakeDamage(proj.Damage)
				proj.Terminate()
			}
			fmt.Println(hit)
		}
	}

}
