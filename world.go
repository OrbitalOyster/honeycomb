package main

import (
	"log"
)

type World struct {
	Chunks         []*Chunk
	CnunksRendered int
}

func GenerateChunkNeighbours(world *World, chunk *Chunk) {
	for i := range 6 {
		q, r, s, qc, rc, sc := chunk.GetNeighbourCoords(i)
		l := CubicToLinear(qc, rc, sc)
		if l < len(world.Chunks) && world.Chunks[l] == nil {
			newChunk := Chunk{
				Qc: qc,
				Rc: rc,
				Sc: sc,
				Q:  q,
				R:  r,
				S:  s,
			}
			newChunk.Generate()
			world.Chunks[l] = &newChunk
		}
	}
}

func (world *World) Generate(worldRadiusInChunks int) {
	numberOfChunks := GetNumberOfHexes(worldRadiusInChunks)
	numberOfHexes := numberOfChunks * GetNumberOfHexes(ChunkRadius)
	world.Chunks = make([]*Chunk, numberOfChunks)
	log.Printf("Generating %d chunks, %d hexes", numberOfChunks, numberOfHexes)
	firstChunk := Chunk{
		Qc: 0,
		Rc: 0,
		Sc: 0,
		Q:  0,
		R:  0,
		S:  0,
	}
	firstChunk.Generate()
	world.Chunks[0] = &firstChunk
	chunksGenerated := 1
	for radius := 0; radius < worldRadiusInChunks-1; radius++ {
		for i := chunksGenerated - 1; i < chunksGenerated+radius*6; i++ {
			GenerateChunkNeighbours(world, world.Chunks[i])
		}
		chunksGenerated += radius * 6
	}
}

func (world *World) Draw() {
	world.CnunksRendered = 0
	for _, chunk := range world.Chunks {
		if chunk.OnScreen() {
			chunk.Draw()
			world.CnunksRendered++
		}
	}
}
