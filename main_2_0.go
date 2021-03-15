package main

import (
	"./arena"
	"./characters"
	. "./constants"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"image"
	//"image/png"
	"os"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func fun() {
	var winSizeX float64 = 816
	var winSizeY float64 = 720
	var tileSize float64 = 8

	wincfg := pixelgl.WindowConfig{
		Title:  "Bomberman 2021",
		Bounds: pixel.R(0, 0, winSizeX, winSizeY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		panic(err)
	}
	/*	charactersPic, err := loadPicture("graphics/characters.png")
		if err != nil {
			panic(err)
		}
	*/whiteBomberman := characters.NewPlayer(WhiteBomberman)
	whiteBomberman.Ani().Show()
	mat := pixel.IM
	mat = mat.Moved(pixel.V(winSizeX/2, winSizeY/2))
	mat = mat.ScaledXY(win.Bounds().Center(), pixel.V(3.3, 3.3))

	win.Clear(colornames.Whitesmoke)
	arena.Draw(win)
	whiteBomberman.Ani().Update()

	for !win.Closed() && !win.Pressed(pixelgl.KeyEscape) {
		if win.Pressed(pixelgl.KeyRight) {
			mat = mat.Moved(pixel.V(tileSize, 0))
			whiteBomberman.Ani().SetView(Right)
		}
		if win.Pressed(pixelgl.KeyLeft) {
			mat = mat.Moved(pixel.V(-tileSize, 0))
			whiteBomberman.Ani().SetView(Left)
		}
		if win.Pressed(pixelgl.KeyUp) {
			mat = mat.Moved(pixel.V(0, tileSize))
			whiteBomberman.Ani().SetView(Up)
		}
		if win.Pressed(pixelgl.KeyDown) {
			mat = mat.Moved(pixel.V(0, -tileSize))
			whiteBomberman.Ani().SetView(Down)
		}
		win.Clear(colornames.Whitesmoke)
		arena.Draw(win)
		whiteBomberman.Ani().Update()
		whiteBomberman.Ani().GetSprite().Draw(win, mat)
		win.Update()
	}
}

func main() {
	pixelgl.Run(fun)
}