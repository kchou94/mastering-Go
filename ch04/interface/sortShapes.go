package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const min = 1
const max = 5

func rF64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

type Shape3D interface {
	Vol() float64
}

type Cube struct {
	x float64
}

type Cuboid struct {
	x, y, z float64
}

type Sphere struct {
	r float64
}

func (c Cube) Vol() float64 {
	return c.x * c.x * c.x
}

func (c Cuboid) Vol() float64 {
	return c.x * c.y * c.z
}

func (s Sphere) Vol() float64 {
	return 4.0 / 3.0 * s.r * s.r * s.r * math.Pi
}

type shapes []Shape3D

// Implementing sort.Interface
func (s shapes) Len() int {
	return len(s)
}

func (s shapes) Less(i, j int) bool {
	return s[i].Vol() < s[j].Vol()
}

func (s shapes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func PrintShapes(s shapes) {
	for _, v := range s {
		switch v.(type) {
		case Cube:
			fmt.Printf("Cube: volume %.2f\n", v.Vol())
		case Cuboid:
			fmt.Printf("Cuboid: volume %.2f\n", v.Vol())
		case Sphere:
			fmt.Printf("Sphere: volume %.2f\n", v.Vol())
		default:
			fmt.Println("Unknown shape")
		}
	}
	fmt.Println()
}

func main() {
	shape := shapes{}
	rand.Seed(time.Now().Unix())

	for i := 0; i < 3; i++ {
		cube := Cube{rF64(min, max)}
		cuboid := Cuboid{rF64(min, max), rF64(min, max), rF64(min, max)}
		sphere := Sphere{rF64(min, max)}
		shape = append(shape, cube, cuboid, sphere)
	}

	PrintShapes(shape)

	// Sort the shapes
	sort.Sort(shapes(shape))
	PrintShapes(shape)

	// Reverse the shapes
	sort.Sort(sort.Reverse(shapes(shape)))
	PrintShapes(shape)
}
