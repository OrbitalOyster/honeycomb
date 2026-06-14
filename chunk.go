package main

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var ChunkRadiusInUnits = HexRadius * (ChunkRadius - .5) * math.Sqrt(3)

type Chunk struct {
	Q int
	R int
	S int

	Qc int
	Rc int
	Sc int

	X float64
	Z float64

	Hexes      []Hex
	Neighbours [6]*Chunk
}

func (chunk *Chunk) OnScreen() bool {
	return true
}

func (chunk Chunk) Draw() {
	for _, h := range chunk.Hexes {
		h.Draw()
	}
}

func GetRandomColor() raylib.Color {
	return raylib.Color{
		R: uint8(raylib.GetRandomValue(0, 255)),
		G: uint8(raylib.GetRandomValue(0, 255)),
		B: uint8(raylib.GetRandomValue(0, 255)),
		A: 255,
	}
}

func (chunk *Chunk) Generate() {
	x, z := CubicToCartesian(chunk.Q, chunk.R, chunk.S)
	chunk.X = x
	chunk.Z = z
	chunkSize := GetNumberOfHexes(ChunkRadius)
	chunk.Hexes = make([]Hex, chunkSize)
	for i := range chunkSize {
		q, r, s := LinearToCubic(i)
		chunk.Hexes[i] = NewHex(q, r, s, raylib.White, chunk)
		chunk.Hexes[i].seed = int(raylib.GetRandomValue(1, 100))
	}
}

func (chunk *Chunk) GetNeighbourCoords(d int) (int, int, int, int, int, int) {
	q, r, s := chunk.Q, chunk.R, chunk.S
	qc, rc, sc := chunk.Qc, chunk.Rc, chunk.Sc
	switch d {
	case 0:
		return q + ChunkRadius*2 - 1, r - ChunkRadius, s - (ChunkRadius - 1), qc + 1, rc - 1, sc
	case 1:
		return q + ChunkRadius, r + ChunkRadius - 1, s - (ChunkRadius*2 - 1), qc + 1, rc, sc - 1
	case 2:
		return q - (ChunkRadius - 1), r + ChunkRadius*2 - 1, s - ChunkRadius, qc, rc + 1, sc - 1
	case 3:
		return q - (ChunkRadius*2 - 1), r + ChunkRadius, s + ChunkRadius - 1, qc - 1, rc + 1, sc
	case 4:
		return q - ChunkRadius, r - (ChunkRadius - 1), s + ChunkRadius*2 - 1, qc - 1, rc, sc + 1
	case 5:
		return q + ChunkRadius - 1, r - (ChunkRadius*2 - 1), s + ChunkRadius, qc, rc - 1, sc + 1
	}
	/* Should not get here */
	panic("Oh no!")
}
