package main

import (
	"fmt"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const MAX_LIGHTS = 8

var (
	ambientColor = []float32{0.1, 0.1, 0.1}
	lights       = make([]Light, MAX_LIGHTS)
	lightCount   = 0
)

type Light struct {
	shader    raylib.Shader
	active    bool
	position  raylib.Vector3
	color     raylib.Color
	radius    float32
	activeLoc int32
	posLoc    int32
	colorLoc  int32
	radiusLoc int32
}

func CreateLight(shader raylib.Shader, position raylib.Vector3, color raylib.Color, radius float32) {
	result := Light{
		active:    true,
		shader:    shader,
		position:  position,
		color:     color,
		radius:    radius,
		activeLoc: raylib.GetShaderLocation(shader, fmt.Sprintf("lights[%d].active", lightCount)),
		posLoc:    raylib.GetShaderLocation(shader, fmt.Sprintf("lights[%d].position", lightCount)),
		colorLoc:  raylib.GetShaderLocation(shader, fmt.Sprintf("lights[%d].color", lightCount)),
		radiusLoc: raylib.GetShaderLocation(shader, fmt.Sprintf("lights[%d].radius", lightCount)),
	}
	lights[lightCount] = result
	lightCount++
	result.Update()
}

func (light *Light) Update() {
	/* Activation */
	var active float32 = 0.0
	if light.active {
		active = 1.0
	}
	raylib.SetShaderValue(light.shader, light.activeLoc, []float32{active}, raylib.ShaderUniformInt)
	/* Position */
	raylib.SetShaderValue(light.shader, light.posLoc, []float32{light.position.X, light.position.Y, light.position.Z}, raylib.ShaderUniformVec3)
	/* Color */
	raylib.SetShaderValue(
		light.shader,
		light.colorLoc,
		[]float32{
			float32(light.color.R) / 255,
			float32(light.color.G) / 255,
			float32(light.color.B) / 255,
		},
		raylib.ShaderUniformVec3,
	)
	/* Radius */
	raylib.SetShaderValue(
		light.shader,
		light.radiusLoc,
		[]float32{light.radius},
		raylib.ShaderUniformFloat,
	)
}
