package main

import (
	"fmt"
	"unsafe"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	raylib.SetConfigFlags(raylib.FlagVsyncHint)
	raylib.SetConfigFlags(raylib.FlagMsaa4xHint)
	// raylib.InitWindow(720, 720, "Honeycomb")
	raylib.InitWindow(1920, 1080, "Honeycomb")
	raylib.ToggleFullscreen()
	raylib.SetTargetFPS(120)
	LoadAssets()
	GenerateHexModel()

	ambientColorLoc := raylib.GetShaderLocation(defaultShader, "ambientColor")
	if ambientColorLoc == -1 {
		// panic("Shader location not found")
	}
	raylib.SetShaderValue(defaultShader, ambientColorLoc, ambientColor, raylib.ShaderUniformVec3)

	viewPosLoc := raylib.GetShaderLocation(defaultShader, "viewPos")

	CreateLight(defaultShader, raylib.NewVector3(1.0, 2.0, 3.0), raylib.Red, 10)
	CreateLight(defaultShader, raylib.NewVector3(-1.0, 2.0, -3.0), raylib.Green, 32)

	cube := raylib.GenMeshCube(1.0, 1.0, 1.0)
	skybox := raylib.LoadModelFromMesh(cube)
	skyboxImg := raylib.LoadImage("assets/textures/skybox.png")
	skyboxTexture := raylib.LoadTextureCubemap(skyboxImg, raylib.CubemapLayoutAutoDetect)
	raylib.SetMaterialTexture(skybox.Materials, raylib.MapCubemap, skyboxTexture)
	skybox.GetMaterials()[0].Shader = defaultShader

	setShaderIntValue(defaultShader, "emap", raylib.MapCubemap)

	raylib.UnloadImage(skyboxImg)

	defaultModel.GetMaterials()[1].Shader = defaultShader
	hexModel.GetMaterials()[0].Shader = defaultShader

	world := World{}
	world.Generate(5)

	CreateCamera(
		raylib.NewVector3(0.0, 8.0, 1.0),
		raylib.NewVector3(0.0, 0.0, 0.0),
		raylib.NewVector3(0.0, 1.0, 0.0),
		60.0,
		raylib.CameraPerspective,
	)

	for !raylib.WindowShouldClose() {
		HandleMouse()
		HandleKeyboard()

		activeCamera := GetActiveCamera()
		// raylib.UpdateCamera(&activeCamera.Camera3D, raylib.CameraFirstPerson)

		raylib.SetShaderValue(
			defaultShader,
			viewPosLoc,
			[]float32{activeCamera.Position.X, activeCamera.Position.Y, activeCamera.Position.Z},
			raylib.ShaderUniformVec3,
		)

		raylib.BeginDrawing()
		raylib.ClearBackground(raylib.SkyBlue)
		raylib.BeginMode3D(activeCamera.Camera3D)

		raylib.DisableBackfaceCulling()
		raylib.DisableDepthMask()
		raylib.DrawModel(skybox, raylib.NewVector3(0.0, 0.0, 0.0), 1.0, raylib.White)
		raylib.EnableBackfaceCulling()
		raylib.EnableDepthMask()

		world.Draw()

		for i := range MAX_LIGHTS {
			if lights[i].enabled {
				raylib.DrawSphereEx(lights[i].position, 0.05, 8, 8, lights[i].color)
			}
		}

		raylib.DrawModelEx(
			defaultModel,
			raylib.NewVector3(0.0, 0.0, 0.0),
			raylib.NewVector3(0.0, 1.0, 0.0),
			float32(raylib.GetTime()*100),
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

func setShaderIntValue(shader raylib.Shader, name string, value int32) {
	raylib.SetShaderValue(
		shader,
		raylib.GetShaderLocation(shader, name),
		unsafe.Slice((*float32)(unsafe.Pointer(&value)), 4),
		raylib.ShaderUniformInt,
	)
}
