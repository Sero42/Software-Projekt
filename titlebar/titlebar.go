package titlebar

import (
	"../animations"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"image/color"
	"sync"
	"time"
)

const (
	startCountdown = 0
	stopCountdown  = 1
)

type titlebarStruct struct {
	batch      *pixel.Batch
	canvas     *pixelgl.Canvas
	sprite     *pixel.Sprite
	background *imdraw.IMDraw
	mutex      sync.Mutex
	timeLeft   uint16 // Muss wegen des Timers durch ein Mutex geschützt werden!!!
	level      uint8
	points     uint32 // max. 4 Spieler
	life       [4]uint8
	command    chan uint8 // Sendet Befehle an den Watchdog-Prozess
	width      float64
	players    uint8 // Anzahl der Spieler
	num        [10]pixel.Rect
	heads      [4]pixel.Rect
	sadHeads   [4]pixel.Rect
	clock      pixel.Rect
	score      pixel.Rect
	blackBox   pixel.Rect
	whiteBox   pixel.Rect
}

// New() definiert einen neuen Titlebar der Breite width
// width muss ein Vielfaches von 8 sein
func New(width uint16) Titlebar {
	t := new(titlebarStruct)
	t.width = float64(width & 0xFFF8)
	t.canvas = pixelgl.NewCanvas(pixel.R(0, 0, t.width, 32))
	t.batch = pixel.NewBatch(&pixel.TrianglesData{}, animations.ItemImage)
	t.sprite = pixel.NewSprite(animations.ItemImage, animations.ItemImage.Bounds())
	var pic pixel.Picture
	t.background = imdraw.New(pic)
	t.background.Color = color.RGBA{240, 128, 0, 255}
	t.background.Push(t.canvas.Bounds().Min)
	t.background.Push(t.canvas.Bounds().Max)
	t.background.Rectangle(0)
	t.level = 1
	t.players = 1
	t.life = [4]uint8{3, 3, 3, 3}

	// Gepufferter Channel für die nebenläufige Steuerung des Titlebars.
	// Kommandos sind oben über Konstanten definiert.
	t.command = make(chan uint8, 10)

	// Rechtecke der Ziffern im Spriteimage definieren
	for i := 0; i < 10; i++ {
		t.num[i] = pixel.R(float64(i+9)*16, 15*16, float64(i+9)*16+8, 15*16+14)
	}
	// Rechteck der Uhr
	t.clock = pixel.R(9*16-8, 18*16, 9*16+8, 18*16+24)
	// Rechteck der Punkteskala
	t.score = pixel.R(3*16, 18*16+8, 7*16, 19*16+8)
	// Rechtecke der Spielerköpfe
	t.heads[0] = pixel.R(12*16, 17*16, 13*16, 18*16)
	t.heads[1] = pixel.R(12*16, 16*16, 13*16, 17*16)
	t.heads[2] = pixel.R(15*16, 17*16, 16*16, 18*16)
	t.heads[3] = pixel.R(15*16, 16*16, 16*16, 17*16)
	// Rechtecke der traurigen Spielerköpfe
	t.sadHeads[0] = pixel.R(13*16, 17*16, 14*16, 18*16)
	t.sadHeads[1] = pixel.R(13*16, 16*16, 14*16, 17*16)
	t.sadHeads[2] = pixel.R(16*16, 17*16, 17*16, 18*16)
	t.sadHeads[3] = pixel.R(16*16, 16*16, 17*16, 17*16)
	// Rechtecke der Restzeitanzeige
	t.blackBox = pixel.R(8*16, 18*16, 8*16+8, 18*16+8)
	t.whiteBox = pixel.R(9*16+8, 18*16, 10*16, 18*16+8)
	return t
}

func (t *titlebarStruct) DecLife(i uint8) {
	if i <= 4 {
		if t.life[i-1] > 0 {
			t.life[i-1]--
		}
	}
}

func (t *titlebarStruct) Draw(target pixel.Target) {
	t.canvas.Draw(target, pixel.IM)
}

func (t *titlebarStruct) GetSeconds() uint16 {
	return t.timeLeft
}

func (t *titlebarStruct) IncLevel() {
	t.level++
}

func (t *titlebarStruct) IncLife(i uint8) {
	if i <= 4 && t.life[i-1] < 9 {
		t.life[i-1]++
	}
}

func (t *titlebarStruct) SetLevel(level uint8) {
	t.level = level
}

func (t *titlebarStruct) SetLife(lifes ...uint8) {
	for i, val := range lifes {
		if i > 3 {
			break
		}
		if val < 10 {
			t.life[i] = val
		} else {
			t.life[i] = 9
		}
	}
}

func (t *titlebarStruct) SetPlayers(numberOfPlayers uint8) {
	if numberOfPlayers <= 4 {
		t.players = numberOfPlayers
	}
}

func (t *titlebarStruct) SetPoints(points uint32) {
	t.points = points
}

func (t *titlebarStruct) SetSeconds(seconds uint16) {
	t.mutex.Lock()
	t.timeLeft = seconds
	t.mutex.Unlock()
}

func (t *titlebarStruct) StartCountdown() {
	t.command <- startCountdown
}

func (t *titlebarStruct) StopCountdown() {
	t.command <- stopCountdown
}

func (t *titlebarStruct) Manager() {
	var countdownStarted bool = false
	for {
		command := <-t.command
		switch command {
		case startCountdown:
			if !countdownStarted {
				go t.countdown()
			}
			countdownStarted = true
		case stopCountdown:
			t.mutex.Lock()
			t.timeLeft = 0
			t.mutex.Unlock()
			countdownStarted = false
		}
	}
}

// countdown() ist eine nebenläufiger Prozess, der durch Watchdog() gestartet wird.
func (t *titlebarStruct) countdown() {
	last := time.Now()
	for t.timeLeft > 0 {
		t.mutex.Lock()
		if t.timeLeft > 0 {
			t.timeLeft--
		}
		t.mutex.Unlock()
		t.Update()
		time.Sleep(2*time.Second - time.Since(last)) // Etwa alle 2 Sekunden wird der Countdown heruntergezählt
		last = time.Now()
	}
}

func (t *titlebarStruct) Update() {
	t.batch.Clear()
	t.background.Draw(t.batch)

	// 8 Pixel Abstand zum linken Rand des Titlebars
	x := float64(8)

	// Zeichne Rahmen des Punktestands
	t.sprite.Set(animations.ItemImage, t.score)
	t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x, 8)).Moved(t.score.Size().Scaled(0.5)))
	x = x + 8

	// Punktestand anzeigen
	punkte := digits(t.points)
	for i, val := range punkte {
		if i > 5 {
			break
		} // Punktezahl zu hoch zum Anzeigen
		t.sprite.Set(animations.ItemImage, t.num[val])
		t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x+(5-float64(i))*8, 8)).Moved(pixel.V(4, 8)))
	}
	x = x + 4*16

	// Uhr zeichnen
	t.sprite.Set(animations.ItemImage, t.clock)
	t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x, 0)).Moved(t.clock.Size().Scaled(0.5)))
	x = x + 16

	// Kurzzeitanzeige links von der Uhr
	r := t.timeLeft % 10
	for i := uint16(1); i < 10; i++ {
		if i > r {
			t.sprite.Set(animations.ItemImage, t.blackBox)
			t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(float64(10-i)*8, 0)).Moved(pixel.V(4, 4)))
		} else {
			t.sprite.Set(animations.ItemImage, t.whiteBox)
			t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(float64(10-i)*8, 0)).Moved(pixel.V(4, 4)))
		}
	}

	// Langzeitanzeige rechts von der Uhr
	r = t.timeLeft / 10
	for i := uint16(0); i < uint16(t.width-x-8)/8; i++ {
		if i >= r {
			t.sprite.Set(animations.ItemImage, t.blackBox)
			t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x+float64(i)*8, 0)).Moved(pixel.V(4, 4)))
		} else {
			t.sprite.Set(animations.ItemImage, t.whiteBox)
			t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x+float64(i)*8, 0)).Moved(pixel.V(4, 4)))
		}
	}

	x = x + 8
	// Anzeige der Leben der einzelnen Spieler
	for i := 0; i < int(t.players); i++ {
		if t.life[i] == 0 {
			t.sprite.Set(animations.ItemImage, t.sadHeads[i])
		} else {
			t.sprite.Set(animations.ItemImage, t.heads[i])
		}
		t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x, 8)).Moved(t.heads[i].Size().Scaled(0.5)))
		x = x + 16
		t.sprite.Set(animations.ItemImage, t.num[t.life[i]])
		t.sprite.Draw(t.batch, pixel.IM.Moved(pixel.V(x, 8)).Moved(pixel.V(4, 8)))
		x = x + 16
	}

	t.batch.Draw(t.canvas)
}

// digits wandelt eine nicht negative Ganzzahl vom Typ uintXX in einen Slice aus uint8 um,
// welcher die Ziffern im Zehnersystem enthält.
func digits(number interface{}) []uint8 {
	var n uint
	switch val := number.(type) {
	case uint8:
		n = uint(val)
	case uint16:
		n = uint(val)
	case uint32:
		n = uint(val)
	case uint64:
		n = uint(val)
	case uint:
		n = val
	default:
		n = 0
	}

	var digits []uint8

	if n == 0 {
		digits = append(digits, 0)
		return digits
	}

	for i := 0; n > 0; i++ {
		digits = append(digits, uint8(n%10))
		n = n / 10
	}

	return digits
}
