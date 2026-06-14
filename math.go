package main

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func ScreenPositionToXZ(screenPosition raylib.Vector2, camera raylib.Camera) raylib.Vector3 {
	ray := raylib.GetScreenToWorldRay(screenPosition, camera)
	length := -ray.Position.Y / ray.Direction.Y
	return raylib.Vector3{
		X: ray.Position.X + (ray.Direction.X * length),
		Y: 0,
		Z: ray.Position.Z + (ray.Direction.Z * length),
	}
}

func GetNumberOfHexes(radius int) int {
	if radius == 0 {
		return 0
	}
	return 3*radius*radius - 3*radius + 1
}

func GetRadius(q int, r int, s int) int {
	if q+r+s != 0 {
		panic("Wrong coords")
	}
	q_a := math.Abs(float64(q))
	r_a := math.Abs(float64(r))
	s_a := math.Abs(float64(s))
	return int((q_a + r_a + s_a) / 2)
}

func LinearToRadius(linear int) int {
	if linear == 0 {
		return 0
	}
	return int(math.Floor((3 + math.Sqrt(12*(float64(linear)-1)+9)) / 6))
}

func LinearToCubic(linear int) (int, int, int) {
	if linear == 0 {
		return 0, 0, 0
	}
	radius := LinearToRadius(linear)
	d := linear - GetNumberOfHexes(radius)
	side := d / radius
	var q, r, s int
	switch side {
	case 0:
		q = radius
		s = -d % radius
		r = -q - s
	case 1:
		s = -radius
		r = d % radius
		q = -r - s
	case 2:
		r = radius
		q = -d % radius
		s = -q - r
	case 3:
		q = -radius
		s = d % radius
		r = -q - s
	case 4:
		s = radius
		r = -d % radius
		q = -r - s
	case 5:
		r = -radius
		q = d % radius
		s = -q - r
	default:
		panic("Oh no!")
	}
	return q, r, s
}

func CubicToLinear(q int, r int, s int) int {
	if q+r+s != 0 {
		panic("Wrong coords")
	}
	/* Edge case */
	if q == 0 && r == 0 && s == 0 {
		return 0
	}
	radius := GetRadius(q, r, s)
	hexes := GetNumberOfHexes(radius)
	switch q {
	case radius: /* Top-right */
		return hexes - s
	case -radius: /* Bottom-left */
		return hexes + radius*3 + s
	default:
		switch r {
		case radius: /* Bottom */
			return hexes + radius*2 - q
		case -radius: /* Top */
			return hexes + radius*5 + q
		default:
			switch s {
			case radius: /* Top-left */
				return hexes + radius*4 - r
			case -radius: /* Bottom-right */
				return hexes + radius + r
			}
		}
	}
	/* Should not get here */
	panic("On no!")
}

func CartesianToCubic(x float64, y float64) (int, int, int) {
	x = x / HexRadius
	y = y / HexRadius

	q := math.Sqrt(3.0)/3.0*float64(x) - 1.0/3.0*float64(y)
	r := 2.0 / 3.0 * y
	s := 0.0 - q - r

	q_rounded := math.Round(q)
	r_rounded := math.Round(r)
	s_rounded := math.Round(s)

	dq := math.Abs(q_rounded - q)
	dr := math.Abs(r_rounded - r)
	ds := math.Abs(s_rounded - s)

	if dq > dr && dq > ds {
		q_rounded = 0.0 - r_rounded - s_rounded
	} else if dr > dq && dr > ds {
		r_rounded = 0.0 - q_rounded - s_rounded
	} else {
		s_rounded = 0.0 - q_rounded - r_rounded
	}

	return int(q_rounded), int(r_rounded), int(s_rounded)
}

func CubicToCartesian(q int, r int, s int) (float64, float64) {
	x := (float64(q) + float64(r)/2.0) * HexRadius * math.Sqrt(3.0)
	z := float64(r) * HexRadius * 3.0 / 2.0
	return x, z
}
