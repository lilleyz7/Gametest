package enemy

import (
	"tester/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	OFFSET = 10
	SPEED  = 100.0
	DAMAGE = 10.0
)

type EnemyStore struct {
	ActiveEnemies []*enemy
}

func NewEnemyStore() *EnemyStore {
	return &EnemyStore{}
}

func (es *EnemyStore) SpawnEnemy(p *player.Player) {
	side := rl.GetRandomValue(1, 4)
	switch side {
	case 1:
		xDis := rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
		enemy := NewEnemy(rl.NewVector2(float32(xDis), OFFSET*-1), SPEED, DAMAGE, rl.Black, p)
		es.ActiveEnemies = append(es.ActiveEnemies, enemy)
	case 2:
		yDis := rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
		enemy := NewEnemy(rl.NewVector2(float32(rl.GetScreenWidth()+OFFSET), float32(yDis)), SPEED, DAMAGE, rl.Black, p)
		es.ActiveEnemies = append(es.ActiveEnemies, enemy)
	case 3:
		xDis := rl.GetRandomValue(0, int32(rl.GetScreenWidth()))
		enemy := NewEnemy(rl.NewVector2(float32(xDis), float32(rl.GetScreenHeight())+OFFSET), SPEED, DAMAGE, rl.Black, p)
		es.ActiveEnemies = append(es.ActiveEnemies, enemy)
	case 4:
		yDis := rl.GetRandomValue(0, int32(rl.GetScreenHeight()))
		enemy := NewEnemy(rl.NewVector2(OFFSET*-1, float32(yDis)), SPEED, DAMAGE, rl.Black, p)
		es.ActiveEnemies = append(es.ActiveEnemies, enemy)
	}
}

func (es *EnemyStore) DrawEnemies() {
	for _, e := range es.ActiveEnemies {
		if e.Active {
			rl.DrawCircle(int32(e.Pos.X), int32(e.Pos.Y), e.Radius, e.Color)
		}
	}
}

func (es *EnemyStore) UpdateEnemies() {
	frametime := rl.GetFrameTime()
	time := rl.GetTime()

	for _, e := range es.ActiveEnemies {
		e.Update(frametime, float32(time))
	}
}
