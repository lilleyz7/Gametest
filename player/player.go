package player

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	SPEED  = 2
	RADIUS = 10.0
	HEALTH = 5
)

type Player struct {
	Position rl.Vector2
	Speed    float32
	Radius   float32
	Health   int32
	Color    rl.Color
}

func NewPlayer() *Player {
	return &Player{
		Position: rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2},
		Speed:    SPEED,
		Radius:   RADIUS,
		Health:   HEALTH,
		Color:    rl.Blue,
	}

}

func (p *Player) Draw() {
	rl.DrawCircle(int32(p.Position.X), int32(p.Position.Y), p.Radius, p.Color)
}

func (p *Player) Update() {
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		p.Position.Y -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		p.Position.Y += p.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		p.Position.X -= p.Speed
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		p.Position.X += p.Speed
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		p.FireProjectile()
	}
}
