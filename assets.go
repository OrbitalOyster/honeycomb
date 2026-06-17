package main

import (
	"path/filepath"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

var (
	defaultFontFilename    = filepath.Join("assets", "fonts", "GolosText-Bold.ttf")
	defaultFontSize        = 32
	defaultFont            raylib.Font
	defaultTextureFilename = filepath.Join("assets", "textures", "ground.png")
	defaultTexture         raylib.Texture2D
	noiseTextureFilename   = filepath.Join("assets", "textures", "noise.png")
	noiseTexture           raylib.Texture2D
	defaultModelFilename   = filepath.Join("assets", "models", "monkey.gltf")
	defaultModel           raylib.Model
	defaultVShaderFilename = filepath.Join("assets", "shaders", "default.vs")
	defaultFShaderFilename = filepath.Join("assets", "shaders", "default.fs")
	defaultShader          raylib.Shader
)

func LoadAssets() {
	defaultFont = raylib.LoadFontEx(defaultFontFilename, int32(defaultFontSize), nil)

	var tmpImage *raylib.Image

	tmpImage = raylib.LoadImage(defaultTextureFilename)
	defaultTexture = raylib.LoadTextureFromImage(tmpImage)

	tmpImage = raylib.LoadImage(noiseTextureFilename)
	noiseTexture = raylib.LoadTextureFromImage(tmpImage)
	raylib.UnloadImage(tmpImage)

	defaultModel = raylib.LoadModel(defaultModelFilename)
	defaultShader = raylib.LoadShader(defaultVShaderFilename, defaultFShaderFilename)
}

func UnloadAssets() {
	raylib.UnloadFont(defaultFont)
	raylib.UnloadTexture(defaultTexture)
	raylib.UnloadModel(defaultModel)
	raylib.UnloadShader(defaultShader)
}
