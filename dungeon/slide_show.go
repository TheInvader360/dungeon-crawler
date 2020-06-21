package dungeon

import (
	"image/color"

	resfont "github.com/TheInvader360/dungeon-crawler/res/font"
	resslide "github.com/TheInvader360/dungeon-crawler/res/slide"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
)

type slide struct {
	image *ebiten.Image
	msg   string
}

type show struct {
	slides []slide
	next   gameState
}

var (
	smallFont    font.Face
	shows        map[string]show
	currentShow  show
	currentSlide int
	currentPosX  int
	minPosX      int
)

func init() {
	tt, err := truetype.Parse(resfont.M5x7_ttf)
	if err != nil {
		panic(err)
	}
	smallFont = truetype.NewFace(tt, &truetype.Options{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	introSlides := make([]slide, 5)
	introSlides[0] = slide{image: EssentialNewImageFromEncoded(resslide.Intro0_png), msg: "Wizard Zaks has stolen our kingdom's protective Power Orb!"}
	introSlides[1] = slide{image: EssentialNewImageFromEncoded(resslide.Intro1_png), msg: "Zaks is in his dungeon lair, defended by beasts, monsters, and enchantments..."}
	introSlides[2] = slide{image: EssentialNewImageFromEncoded(resslide.Intro2_png), msg: "This amulet grants one-way passage through the dungeon's enchanted archways..."}
	introSlides[3] = slide{image: EssentialNewImageFromEncoded(resslide.Intro3_png), msg: "The enchanted archways disappear once crossed. There can be no escape until Zaks has been defeated!"}
	introSlides[4] = slide{image: EssentialNewImageFromEncoded(resslide.Intro4_png), msg: "Please, take the amulet, enter the dungeon, defeat Zaks, retrieve the Power Orb, and save our kingdom!"}
	introShow := show{slides: introSlides, next: exploration}

	gameOverSlides := make([]slide, 1)
	gameOverSlides[0] = slide{image: EssentialNewImageFromEncoded(resslide.GameOver_png), msg: ""}
	gameOverShow := show{slides: gameOverSlides, next: initialize}

	shows = map[string]show{"intro": introShow, "gameOver": gameOverShow}
}

func setupSlideShow(show string) {
	currentShow = shows[show]
	currentSlide = -1
	nextSlide()
}

func nextSlide() bool {
	if currentSlide < len(currentShow.slides)-1 {
		currentSlide++
		currentPosX = ScreenWidth
		msgBounds, _ := font.BoundString(smallFont, currentShow.slides[currentSlide].msg)
		minPosX = -int((msgBounds.Max.X - msgBounds.Min.X) / 64)
		return true
	}
	return false
}

func renderSlide(v *ebiten.Image) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}
	v.DrawImage(currentShow.slides[currentSlide].image, op)
	text.Draw(v, currentShow.slides[currentSlide].msg, smallFont, currentPosX, ScreenHeight-3, color.White)
	return v
}

func updateSlideShow(g *Game) {
	ticksPerShift := 2
	pixelsPerShift := -1
	if IsPressed(r) {
		ticksPerShift = 1
		pixelsPerShift = -2
	}
	if g.counter%ticksPerShift == 0 {
		currentPosX += pixelsPerShift
		if currentPosX < minPosX {
			currentPosX = minPosX
			nextSlide()
		}
	}
	if IsJustPressed(u) {
		if !nextSlide() {
			g.gameState = currentShow.next
		}
	}
}
