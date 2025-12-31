package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const width = 800
const height = 600

var BACKGROUND = rl.Color{R: 16, G: 16, B: 16, A: 255}  //"#101010"
var FOREGROUND = rl.Color{R: 80, G: 250, B: 80, A: 255} //"#101010"
var angle float32 = 0
var dz float32 = 1

type Point2d struct {
	x float32
	y float32
}
type Point3d struct {
	x float32
	y float32
	z float32
}

func point(p Point2d) {
	var s float32 = 20
	x := int32(p.x - s/2)
	y := int32(p.y - s/2)
	rl.DrawRectangle(x, y, int32(s), int32(s), FOREGROUND)
}

func line(p1 Point2d, p2 Point2d) {
	rl.DrawLineEx(rl.Vector2{X: p1.x, Y: p1.y}, rl.Vector2{X: p2.x, Y: p2.y}, 3, FOREGROUND)
}
func screen(p Point2d) Point2d {
	x := (p.x + 1) / 2 * float32(width)
	y := (1 - (p.y+1)/2) * float32(height)

	return Point2d{x: x, y: y}
}

func project(p Point3d) Point2d {
	return Point2d{
		x: p.x / p.z,
		y: p.y / p.z,
	}
}

func translate_z(p Point3d, dz float32) Point3d {
	return Point3d{x: p.x, y: p.y, z: p.z + dz}
}

func rotate_xz(p Point3d, angle float32) Point3d {
	c := math.Cos(float64(angle))
	s := math.Sin(float64(angle))
	return Point3d{
		x: p.x*float32(c) - p.z*float32(s),
		y: p.y,
		z: p.x*float32(s) + p.z*float32(c),
	}
}

func main() {
	vs := []Point3d{
		{x: 0.25, y: 0.25, z: 0.25},
		{x: -0.25, y: 0.25, z: 0.25},
		{x: -0.25, y: -0.25, z: 0.25},
		{x: 0.25, y: -0.25, z: 0.25},

		{x: 0.25, y: 0.25, z: -0.25},
		{x: -0.25, y: 0.25, z: -0.25},
		{x: -0.25, y: -0.25, z: -0.25},
		{x: 0.25, y: -0.25, z: -0.25},
	}
	fs := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{0, 4},
		{1, 5},
		{2, 6},
		{3, 7},
	}

	rl.InitWindow(width, height, "3d demistified attempt")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	dt := float32(1.0 / 60.0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(BACKGROUND)
		angle += float32(math.Pi) * dt
		for _, f := range fs {
			for i := range f {
				a := vs[f[i]]
				b := vs[f[(i+1)%len(f)]]
				line(screen(project(translate_z(rotate_xz(a, angle), dz))),
					screen(project(translate_z(rotate_xz(b, angle), dz))))
			}
		}

		rl.EndDrawing()
	}
}
