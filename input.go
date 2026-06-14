package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const mouseZoomSpeed = .25

var previousMousePosition raylib.Vector2

func HandleMouse() {
	mousePosition := raylib.GetMousePosition()
	mouseXZ := ScreenPositionToXZ(mousePosition, defaultCamera)
	/* Camera pan */
	if raylib.IsMouseButtonDown(raylib.MouseButtonMiddle) {
		oldMouseXZ := ScreenPositionToXZ(previousMousePosition, defaultCamera)
		deltaMouse := raylib.Vector3Subtract(oldMouseXZ, mouseXZ)
		defaultCamera.Position = raylib.Vector3Add(defaultCamera.Position, deltaMouse)
		defaultCamera.Target = raylib.Vector3Add(defaultCamera.Target, deltaMouse)
	}
	/* Camera height */
	defaultCamera.Position.Y -= raylib.GetMouseWheelMove() * mouseZoomSpeed
	previousMousePosition = mousePosition
}

func HandleKeyboard() {
  if raylib.IsKeyDown(raylib.KeyD) {
    lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(0.1, 0.0, 0.0))
    lights[0].Update()
  }
  if raylib.IsKeyDown(raylib.KeyA) {
    lights[0].position = raylib.Vector3Add(lights[0].position, raylib.NewVector3(-0.1, 0.0, 0.0))
    lights[0].Update()
  }
}
