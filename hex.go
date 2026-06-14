package main

import (
	"math"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Hex struct {
	Q            int
	R            int
	S            int
	Color        raylib.Color
	Parent       *Chunk
	Position raylib.Vector3
	seed         int
}

var hexMesh raylib.Mesh
var hexModel raylib.Model

func GenerateHexMesh() raylib.Mesh {
	mesh := raylib.Mesh{
		VertexCount:   7,
		TriangleCount: 6,
	}

	vertices := []float32{
		0, 0, 0,
		0, 0, -float32(HexRadius),
		float32(math.Sqrt(3.0) * HexRadius / 2.0), 0, -float32(HexRadius / 2.0),
		float32(math.Sqrt(3.0) * HexRadius / 2.0), 0, float32(HexRadius / 2.0),
		0, 0, float32(HexRadius),
		-float32(math.Sqrt(3.0) * HexRadius / 2.0), 0, float32(HexRadius / 2.0),
		-float32(math.Sqrt(3.0) * HexRadius / 2.0), 0, -float32(HexRadius / 2.0),
	}

	indices := []uint16{
		0, 2, 1,
		0, 3, 2,
		0, 4, 3,
		0, 5, 4,
		0, 6, 5,
		0, 1, 6,
	}

	texcoords := []float32{
		.5, .5,
		.5, 0,
		.5 + float32(math.Sqrt(3.0)*.5/2.0), 1.0 / 4.0,
		.5 + float32(math.Sqrt(3.0)*.5/2.0), 1.0 - 1.0/4.0,
		.5, 1.0,
		.5 - float32(math.Sqrt(3.0)*.5/2.0), 1.0 - 1.0/4.0,
		.5 - float32(math.Sqrt(3.0)*.5/2.0), 1.0 / 4.0,
	}

	normals := []float32{
		0, 1.0, 0,
		0, 1.0, 0,
		0, 1.0, 0,
		0, 1.0, 0,
		0, 1.0, 0,
		0, 1.0, 0,
		0, 1.0, 0,
	}

	mesh.Vertices = &vertices[0]
	mesh.Indices = &indices[0]
	mesh.Normals = &normals[0]
	mesh.Texcoords = &texcoords[0]

	raylib.UploadMesh(&mesh, false)
	return mesh
}

func GenerateHexModel() {
	hexMesh = GenerateHexMesh()
	hexModel = raylib.LoadModelFromMesh(hexMesh)
	hexModel.Materials.Maps.Texture = defaultTexture
}

func NewHex(q int, r int, s int, color raylib.Color, parent *Chunk) Hex {
	if q+r+s != 0 {
		panic("Wrong coords")
	}
	result := Hex{
		Q:      q,
		R:      r,
		S:      s,
		Color:  color,
		Parent: parent,
	}
	x, z := (float64(q)+float64(r)/2.0)*HexRadius*math.Sqrt(3.0), float64(r)*HexRadius*3.0/2.0
	x += parent.X
	z += parent.Z
	result.Position = raylib.Vector3{
		X: float32(x),
		Y: 0.0,
		Z: float32(z),
	}
	return result
}

func (hex Hex) Draw() {
	raylib.DrawModel(hexModel, hex.Position, 1, hex.Color)
	// if hex.seed%5 == 0 {
	// 	raylib.DrawModelEx(piney, hex.Position, raylib.Vector3{X: 0.0, Y: 1.0, Z: 0.0}, float32(hex.seed), raylib.Vector3{X: .4, Y: .4, Z: .4}, raylib.White)
	// }
}
