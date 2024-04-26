package player

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SPEED  = 2
	RADIUS = 10.0
	HEALTH = 5
)

type Player struct {
	PosX   float32
	PosY   float32
	Speed  float32
	Radius float32
	Health int32
	Color  rl.Color
	Store  []*projectile
}

func NewPlayer() *Player {
	return &Player{
		PosX:   float32(rl.GetScreenWidth()) / 2,
		PosY:   float32(rl.GetScreenHeight()) / 2,
		Speed:  SPEED,
		Radius: RADIUS,
		Health: HEALTH,
		Color:  rl.Blue,
		Store:  []*projectile{},
	}

}

func (p *Player) Draw() {
	rl.DrawCircle(int32(p.PosX), int32(p.PosY), p.Radius, p.Color)
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.PosY -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.PosY += p.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.PosX -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.PosX += p.Speed
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		p.FireProjectile()
	}
}
