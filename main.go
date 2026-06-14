package main

import (
	"fmt"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	raylib.SetConfigFlags(raylib.FlagVsyncHint)
	raylib.SetConfigFlags(raylib.FlagMsaa4xHint)
	raylib.InitWindow(1280, 720, "Honeycomb")
	raylib.SetTargetFPS(120)
	LoadAssets()
	GenerateHexModel()

	ambientColorLoc := raylib.GetShaderLocation(defaultShader, "ambientColor")
	if ambientColorLoc == -1 {
		panic("Shader location not found")
	}
	raylib.SetShaderValue(defaultShader, ambientColorLoc, ambientColor, raylib.ShaderUniformVec3)

	viewPosLoc := raylib.GetShaderLocation(defaultShader, "viewPos")

	CreateLight(defaultShader, raylib.NewVector3(1.0, 2.0, 3.0), raylib.Red, 10)
	CreateLight(defaultShader, raylib.NewVector3(-3.0, 6.0, -2.0), raylib.Green, 32)

	defaultModel.GetMaterials()[1].Shader = defaultShader
	hexModel.GetMaterials()[0].Shader = defaultShader

	world := World{}
	world.Generate(3)
	for !raylib.WindowShouldClose() {
		HandleMouse()
		HandleKeyboard()

		raylib.SetShaderValue(
		  defaultShader,
			viewPosLoc,
			[]float32{defaultCamera.Position.X, defaultCamera.Position.Y, defaultCamera.Position.Z},
			raylib.ShaderUniformVec3,
		)

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.SkyBlue)
		raylib.BeginMode3D(defaultCamera)
		world.Draw()

		for i := range MAX_LIGHTS {
			if lights[i].active {
				raylib.DrawSphereEx(lights[i].position, 0.05, 8, 8, lights[i].color)
			}
		}

		raylib.DrawModelEx(
			defaultModel,
			raylib.NewVector3(0.0, 0.0, 0.0),
			raylib.NewVector3(0.0, 1.0, 0.0),
			float32(raylib.GetTime() * 100),
			raylib.NewVector3(1.0, 1.0, 1.0),
			raylib.White,
		)

		raylib.EndMode3D()
		DebugText(fmt.Sprintf("fps: %d", raylib.GetFPS()), 16, 16)
		DebugText(fmt.Sprintf("chunks on screen: %d", world.CnunksRendered), 16, 48)
		raylib.EndDrawing()
	}
	raylib.UnloadMesh(&hexMesh)
	UnloadAssets()
	raylib.CloseWindow()
}
