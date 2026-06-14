package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

var defaultCamera = raylib.Camera3D{}

func init() {
	defaultCamera.Position = raylib.NewVector3(0.0, 5.0, 4.0)
	defaultCamera.Target = raylib.NewVector3(0.0, 0.0, 0.0)
	defaultCamera.Up = raylib.NewVector3(0.0, 1.0, 0.0)
	defaultCamera.Fovy = 60.0
	defaultCamera.Projection = raylib.CameraPerspective
}
