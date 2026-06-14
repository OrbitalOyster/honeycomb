package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

var dedugTextShadowDelta = raylib.Vector2{X: 1, Y: 1}

func DebugText(text string, x int, y int) {
	textPosition := raylib.Vector2{X: float32(x), Y: float32(y)}
	raylib.DrawTextEx(defaultFont, text, raylib.Vector2Add(textPosition, dedugTextShadowDelta), float32(defaultFontSize), 1, raylib.Black)
	raylib.DrawTextEx(defaultFont, text, textPosition, float32(defaultFontSize), 1, raylib.White)
}
