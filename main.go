package main

import (
	"./animations"
	"./arena"
	"./characters"
	. "./constants"
	"./items"
	"./sounds"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	//"golang.org/x/image/colornames"
	"time"
	"math"
)

var bombs []items.Bombe
var turfNtreesArena arena.Arena
var tempAniSlice [][]interface{} // [Animation][Matrix]

// Vor: ...
// Eff: Ist der Counddown der Bombe abgelaufen passiert folgendes:
//     		- Eine neue Explosionsanimation ist erstellt und an die Position der ehemaligen bombe gesetzt.
//      	- Es ertönt der Explosionssound.
//      Ist der Countdown nicht abgelaufen, passiert nichts.

func checkForExplosions () []int {		
	
	//var bomben []items.Bombe
	var indexe []int
	
	for index,item := range(bombs) {
		if ((item).GetTimeStamp()).Before(time.Now()) {
			//bomben = append (bomben,item)
			indexe = append (indexe,index)
				
			x, y := turfNtreesArena.GetFieldCoord(item.GetPos())
			power := int(item.GetPower())
			l, r, u, d := power, power, power, power
			if l > x {
				l = x
			}
			if 12-r < x {
				r = 12 - x
			}
			if d > y {
				d = y
			}
			if 10-u < y {
				u = 10 - y
			}

			if x > 0 {
				for i := 1; i <= int(l); i++ {
					if turfNtreesArena.IsTile(x-i,y) {
						if turfNtreesArena.RemoveTiles(x-i, y) {
							l = i
						} else {
							l = i-1
						}
						break
					}
				}
			}
			if x < 12 {
				for i := 1; i <= int(r); i++ {
					if turfNtreesArena.IsTile(x+i,y) {
						if turfNtreesArena.RemoveTiles(x+i, y) {
							r = i
						} else {
							r = i-1
						}
						break
					}
				}
			}
			if y < 10 {
				for i := 1; i <= int(u); i++ {
					if turfNtreesArena.IsTile(x,y+i) {
						if turfNtreesArena.RemoveTiles(x, y+i) {
							u = i
						} else {
							u = i-1
						}
						break
					}
				}
			}
			if y > 0 {
				for i := 1; i <= int(d); i++ {
					if turfNtreesArena.IsTile(x,y-i) {
						if turfNtreesArena.RemoveTiles(x, y-i) {
							d = i
						} else {
							d = i-1
						}
						break
					}
				}
			}
			
			
			for i:=1; i<=l ;i++ {
				b,bom := isThereABomb(item.GetPos().Add(pixel.V(float64(-i)*TileSize,0))) 
				if b {
					bom.SetTimeStamp(time.Now())
				}
			}
			for i:=1; i<=r ;i++ {
				b,bom := isThereABomb(item.GetPos().Add(pixel.V(float64(i)*TileSize,0))) 
				if b {
					bom.SetTimeStamp(time.Now())
				}
			}
			for i:=1; i<=u ;i++ {
				b,bom := isThereABomb(item.GetPos().Add(pixel.V(0,float64(i)*TileSize)))
				if b {
					bom.SetTimeStamp(time.Now())
				}
			}
			for i:=1; i<=d ;i++ {
				b,bom := isThereABomb(item.GetPos().Add(pixel.V(0,float64(-i)*TileSize)))
				if b {
					bom.SetTimeStamp(time.Now())
				}
			}
			
			fmt.Println("")

			ani := animations.NewExplosion(uint8(l), uint8(r), uint8(u), uint8(d))
			ani.Show()
			tempAni := make([]interface{}, 2)
			tempAni[0] = ani
			tempAni[1] = (item.GetMatrix()).Moved(ani.ToCenter())
			tempAniSlice = append(tempAniSlice, tempAni)
			s2 := sounds.NewSound(Deathflash)
			go s2.PlaySound()	
		}
	}
	
	return indexe
}


// Vor.:...
// Eff.: Explodiere Bomben sind aus dem slice bombs gelöscht

func removeExplodedBombs (indexe []int) {
	for i:=len(indexe)-1;i>=0;i-- {
		index:=indexe[i]
		if len(bombs) == 1 {
			bombs = bombs[:0]
		} else {
			bombs = append(bombs[0:index], bombs[index+1:]...)
		}
	}
}

func showExpolosions (win *pixelgl.Window) []int {
	var indexe []int
	for index,a := range tempAniSlice {
		ani := (a[0]).(animations.Animation)
		ani.Update()
		mtx := (a[1]).(pixel.Matrix)
		(ani.GetSprite()).Draw(win, mtx)
		if !ani.IsVisible() {
			indexe = append(indexe,index)
		}
	}
	return indexe
}

func clearExplosions (indexe []int) {
	for _,index := range(indexe) {
		if len(tempAniSlice)!=0 {
			if len(tempAniSlice) == 1 {
				tempAniSlice = tempAniSlice[:0]
			} else {
				tempAniSlice = tempAniSlice[index:]
			}
		}
	}	
}

func isThereABomb (v pixel.Vec) (bool,items.Bombe) {
	for _,item := range (bombs) {
		if item.GetPos() == v { 
			return true,item 
		}
	}
	return false,nil
}

func sun() {
	const zoomFactor = 3
	const pitchWidth = 13
	const pitchHeight = 11
	var winSizeX float64 = zoomFactor * ((3 + pitchWidth)*TileSize)		// TileSize = 16
	var winSizeY float64 = zoomFactor * ((1 + pitchHeight)*TileSize)
	var stepSize float64 = 1

	wincfg := pixelgl.WindowConfig{
		Title:  "Bomberman 2021",
		Bounds: pixel.R(0, 0, winSizeX, winSizeY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(wincfg)
	if err != nil {
		panic(err)
	}

	turfNtreesArena = arena.NewArena(pitchWidth, pitchHeight)

	whiteBomberman := characters.NewPlayer(WhiteBomberman)
	whiteBomberman.Ani().Show()
	whiteBomberman.IncPower()
	whiteBomberman.IncPower()

	// Put character at free space with at least two free neighbours in a row
A:
	for i := 2 * turfNtreesArena.GetWidth(); i < len(turfNtreesArena.GetPassability())-2*turfNtreesArena.GetWidth(); i++ { // Einschränkung des Wertebereichs von i um index out of range Probleme zu vermeiden
		if turfNtreesArena.GetPassability()[i] && turfNtreesArena.GetPassability()[i-1] && turfNtreesArena.GetPassability()[i-2] || // checke links, rechts, oben, unten
			turfNtreesArena.GetPassability()[i] && turfNtreesArena.GetPassability()[i+1] && turfNtreesArena.GetPassability()[i+2] ||
			turfNtreesArena.GetPassability()[i] && turfNtreesArena.GetPassability()[i+turfNtreesArena.GetWidth()] &&
				turfNtreesArena.GetPassability()[i+2*turfNtreesArena.GetWidth()] ||
			turfNtreesArena.GetPassability()[i] && turfNtreesArena.GetPassability()[i-turfNtreesArena.GetWidth()] &&
				turfNtreesArena.GetPassability()[i-2*turfNtreesArena.GetWidth()] {
			whiteBomberman.MoveTo(turfNtreesArena.GetLowerLeft().Add(pixel.V(float64(i%turfNtreesArena.GetWidth())*
				TileSize, float64(i/turfNtreesArena.GetWidth())*TileSize)))
			break A
		}
	}

	win.SetMatrix(pixel.IM.Scaled(pixel.V(0, 0), zoomFactor))
	win.Update()

	for !win.Closed() && !win.Pressed(pixelgl.KeyEscape) {
		grDir := turfNtreesArena.GrantedDirections(whiteBomberman.GetPosBox()) // [4]bool left-right-up-down granted?

		if win.Pressed(pixelgl.KeyLeft) && grDir[0] {
			whiteBomberman.Move(pixel.V(-stepSize, 0))
			whiteBomberman.Ani().SetView(Left)
		}
		if win.Pressed(pixelgl.KeyRight) && grDir[1] {
			whiteBomberman.Move(pixel.V(stepSize, 0))
			whiteBomberman.Ani().SetView(Right)
		}
		if win.Pressed(pixelgl.KeyUp) && grDir[2] {
			whiteBomberman.Move(pixel.V(0, stepSize))
			whiteBomberman.Ani().SetView(Up)
		}
		if win.Pressed(pixelgl.KeyDown) && grDir[3] {
			whiteBomberman.Move(pixel.V(0, -stepSize))
			whiteBomberman.Ani().SetView(Down)
		}
		if win.JustPressed(pixelgl.KeyB) {
			pb := characters.Player(whiteBomberman).GetPosBox()
			b,_ := isThereABomb (pixel.Vec{math.Round(pb.Center().X/TileSize) * TileSize, math.Round(pb.Center().Y/TileSize) * TileSize})
			if !b {
				var item items.Bombe
				item = items.NewBomb(characters.Player(whiteBomberman))
				bombs = append(bombs, item)
				fmt.Println(item.GetPos())
			}
		}
		
		turfNtreesArena.GetCanvas().Draw(win, *(turfNtreesArena.GetMatrix()))
		
		removeExplodedBombs(checkForExplosions())
		
		for _, item := range bombs {
			item.Draw(win)
		}
		
		clearExplosions(showExpolosions(win))
		
		whiteBomberman.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(sun)
}
