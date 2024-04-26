package enemy

import (
	"fmt"
	"math"
	"tester/player"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ENEMY_RADIUS = 10.0
)

type enemy struct {
	Pos      rl.Vector2
	Speed    float32
	Radius   float32
	Health   int
	Rotation float64
	Active   bool
	Damage   float32
	Target   *player.Player
	Color    rl.Color
}

func NewEnemy(pos rl.Vector2, speed float32, damage float32, color rl.Color, target *player.Player) *enemy {
	return &enemy{
		Pos:    pos,
		Speed:  speed,
		Damage: damage,
		Radius: ENEMY_RADIUS,
		Health: 5,
		Target: target,
		Color:  color,
		Active: true,
	}
}

func (e *enemy) Update(frameTime, time float32) {
	if e.Health <= 0 {
		e.Active = false
	}

	hit := rl.CheckCollisionCircles(rl.Vector2{X: e.Target.PosX, Y: e.Target.PosY}, e.Target.Radius, e.Pos, e.Radius)
	if hit {
		e.TakeDamage(100000000000000)
	}

	nums := []int{1, 2, 3, 4, 5}

	for _, val := range nums {
		fmt.Println(val)
	}

	e.Rotation = e.UpdateRotation(e.Target)
	e.Pos.X = e.Pos.X + float32(math.Cos(e.Rotation))*e.Speed*frameTime
	e.Pos.Y = e.Pos.Y + float32(math.Sin(e.Rotation))*e.Speed*frameTime

}

func (e *enemy) UpdateRotation(p *player.Player) float64 {
	return math.Atan2(float64(p.PosY)-float64(e.Pos.Y), float64(p.PosX)-float64(e.Pos.X))
}

func (e *enemy) TakeDamage(damage int) {
	e.Health -= damage
	if e.Health <= 0 {
		e.Active = false
	}
}
