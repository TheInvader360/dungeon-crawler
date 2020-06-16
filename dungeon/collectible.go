package dungeon

import (
	resentity "github.com/TheInvader360/dungeon-crawler/res/entity"
	"github.com/hajimehoshi/ebiten"
)

type collectible struct {
	farImg *ebiten.Image
	midImg *ebiten.Image
}

var key collectible
var gold collectible
var potion collectible

func init() {
	kf := EssentialNewImageFromEncoded(resentity.KeyFar_png)
	km := EssentialNewImageFromEncoded(resentity.KeyMid_png)
	key = collectible{farImg: kf, midImg: km}

	gf := EssentialNewImageFromEncoded(resentity.GoldFar_png)
	gm := EssentialNewImageFromEncoded(resentity.GoldMid_png)
	gold = collectible{farImg: gf, midImg: gm}

	pf := EssentialNewImageFromEncoded(resentity.PotionFar_png)
	pm := EssentialNewImageFromEncoded(resentity.PotionMid_png)
	potion = collectible{farImg: pf, midImg: pm}
}
