package main

import rl "github.com/gen2brain/raylib-go/raylib"

const height = 800
const width = 600

var BACKGROUND = rl.Color{R: 16, G: 16, B: 16, A: 1}  //"#101010"
var FOREGROUND = rl.Color{R: 80, G: 250, B: 80, A: 1} //"#101010"

type Point2d struct {
	x int32
	y int32
}
type Point3d struct {
	x int32
	y int32
	z int32
}

func point(p Point2d) {
	var s int32 = 20
	x := p.x - s/2
	y := p.y - s/2
	rl.DrawRectangle(x, y, s, s, FOREGROUND)
}

func line(p1 Point2d, p2 Point2d) {
	rl.DrawLineEx(rl.Vector2{X: float32(p1.x), Y: float32(p1.y)}, rl.Vector2{X: float32(p2.x), Y: float32(p2.y)}, 3, FOREGROUND)
}
func screen(p *Point2d) Point2d {
	x := (p.x + 1) / 2 * width
	y := (1 - (p.y + 1) / 2) * height

	return Point2d{x: x, y: y}
}

func project(p Point3d) Point2d{
	return  Point2d{
		x: p.x / p.z,
		y: p.y / p.z,
	}
}


func translate_z(p Point3d, dz int) Point3d{

}

func main() {

	rl.InitWindow(800, 450, "3d demistified attempt")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(BACKGROUND)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
