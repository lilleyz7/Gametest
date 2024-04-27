package player

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	DAMAGE            = 10
	LIFETIME          = 2.0
	FIRERATE          = 0.4
	PROJECTILE_RADIUS = 5
)

var ProjStore ProjectileStore

type ProjectileStore struct {
	Shots        []*projectile
	LastShotTime float64
}

type projectile struct {
	PosX         float32
	PosY         float32
	Damage       int
	Rotation     float64
	Speed        float32
	Color        rl.Color
	CreationTime float64
	Active       bool
	Radius       int
}

func (p *Player) FireProjectile() {

	if ProjStore.LastShotTime+FIRERATE < rl.GetTime() {
		proj := &projectile{
			PosX:         p.Position.X,
			PosY:         p.Position.Y,
			Damage:       DAMAGE,
			Rotation:     p.getRotation(),
			Speed:        200,
			CreationTime: rl.GetTime(),
			Active:       true,
			Radius:       PROJECTILE_RADIUS,
		}

		ProjStore.Shots = append(ProjStore.Shots, proj)
		ProjStore.LastShotTime = rl.GetTime()
	} else {
		fmt.Println("Not enough time")
	}
}

func (p *Player) getRotation() float64 {
	return math.Atan2(float64(rl.GetMouseY())-float64(p.Position.Y), float64(rl.GetMouseX())-float64(p.Position.X))
}

func DrawProjectiles() {
	for _, proj := range ProjStore.Shots {
		if proj.Active {
			rl.DrawCircle(int32(proj.PosX), int32(proj.PosY), float32(proj.Radius), rl.Brown)
		}
	}
}

func UpdateProjectiles() {
	frametime := rl.GetFrameTime()
	time := rl.GetTime()

	for _, proj := range ProjStore.Shots {
		proj.ProjUpdate(frametime, float32(time))
	}
}

func (proj *projectile) ProjUpdate(frameTime, time float32) {
	if time > float32(proj.CreationTime)+LIFETIME {
		proj.Terminate()
		return
	}
	proj.PosX = proj.PosX + float32(math.Cos(proj.Rotation))*proj.Speed*frameTime
	proj.PosY = proj.PosY + float32(math.Sin(proj.Rotation))*proj.Speed*frameTime

}

func (proj *projectile) Terminate() {
	proj.Active = false
}
