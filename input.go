package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const mouseZoomSpeed = .25

var previousMousePosition raylib.Vector2

func HandleMouse() {
	activeCamera := GetActiveCamera()
	mousePosition := raylib.GetMousePosition()
	mouseXZ := activeCamera.GetScreenPositionToXZ(mousePosition)
	/* Camera pan */
	if raylib.IsMouseButtonDown(raylib.MouseButtonMiddle) {
		oldMouseXZ := activeCamera.GetScreenPositionToXZ(previousMousePosition)
		deltaMouse := raylib.Vector3Subtract(oldMouseXZ, mouseXZ)
		activeCamera.Move(deltaMouse)
	}
	/* Camera height */
	deltaHeight := raylib.GetMouseWheelMove() * mouseZoomSpeed
	activeCamera.Position.Y = max(activeCamera.Position.Y-deltaHeight, 1.0)

	previousMousePosition = mousePosition
}

func HandleKeyboard() {
	if raylib.IsKeyDown(raylib.KeyW) {
		lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(0.0, 0.0, -0.1))
		lights[0].Update()
	}
	if raylib.IsKeyDown(raylib.KeyA) {
		lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(-0.1, 0.0, 0.0))
		lights[0].Update()
	}
	if raylib.IsKeyDown(raylib.KeyS) {
		lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(0.0, 0.0, 0.1))
		lights[0].Update()
	}
	if raylib.IsKeyDown(raylib.KeyD) {
		lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(0.1, 0.0, 0.0))
		lights[0].Update()
	}
	if raylib.IsKeyPressed(raylib.KeyC) {
		NextCamera()
	}
}
