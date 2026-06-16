package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type MyCamera struct {
	raylib.Camera3D
}

var (
	cameras      []MyCamera
	activeCamera int = 0
)

func CreateCamera(position raylib.Vector3, target raylib.Vector3, up raylib.Vector3, fov float32, projection raylib.CameraProjection) {
	result := MyCamera{
		Camera3D: raylib.Camera3D{
			Position:   position,
			Target:     target,
			Up:         up,
			Fovy:       fov,
			Projection: projection,
		},
	}
	cameras = append(cameras, result)
}

func (myCamera *MyCamera) GetScreenPositionToXZ(screenPosition raylib.Vector2) raylib.Vector3 {
	ray := raylib.GetScreenToWorldRay(screenPosition, myCamera.Camera3D)
	length := -ray.Position.Y / ray.Direction.Y
	return raylib.Vector3{
		X: ray.Position.X + (ray.Direction.X * length),
		Y: 0,
		Z: ray.Position.Z + (ray.Direction.Z * length),
	}
}

func (camera *MyCamera) Move(delta raylib.Vector3) {
	camera.Camera3D.Position = camera.Camera3D.Position.Add(delta)
	camera.Camera3D.Target = camera.Camera3D.Target.Add(delta)
}

func init() {
	CreateCamera(
		raylib.NewVector3(0.0, 5.0, 4.0),
		raylib.NewVector3(0.0, 0.0, 0.0),
		raylib.NewVector3(0.0, 1.0, 0.0),
		60.0,
		raylib.CameraPerspective,
	)
}

func NextCamera() {
	activeCamera++
	if activeCamera >= len(cameras) {
		activeCamera = 0
	}
}

func GetActiveCamera() *MyCamera {
	return &cameras[activeCamera]
}
