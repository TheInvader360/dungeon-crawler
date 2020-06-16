package dungeon

import (
	resentity "github.com/TheInvader360/dungeon-crawler/res/entity"
	"github.com/hajimehoshi/ebiten"
)

type enemy struct {
	farImg  *ebiten.Image
	midImg  *ebiten.Image
	nearImg *ebiten.Image
}

var enemies []enemy

func init() {
	f := EssentialNewImageFromEncoded(resentity.Enemy0Far_png)
	m := EssentialNewImageFromEncoded(resentity.Enemy0Mid_png)
	n := EssentialNewImageFromEncoded(resentity.Enemy0Near_png)
	e0 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy1Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy1Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy1Near_png)
	e1 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy2Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy2Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy2Near_png)
	e2 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy3Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy3Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy3Near_png)
	e3 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy4Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy4Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy4Near_png)
	e4 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy5Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy5Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy5Near_png)
	e5 := enemy{farImg: f, midImg: m, nearImg: n}

	f = EssentialNewImageFromEncoded(resentity.Enemy6Far_png)
	m = EssentialNewImageFromEncoded(resentity.Enemy6Mid_png)
	n = EssentialNewImageFromEncoded(resentity.Enemy6Near_png)
	e6 := enemy{farImg: f, midImg: m, nearImg: n}

	enemies = []enemy{e0, e1, e2, e3, e4, e5, e6}
}
