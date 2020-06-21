# Dungeon Crawler

A minimal dungeon crawling adventure game (RPG) written in go using the ebiten library.

* 2D model, pseudo 3D view.
* Super low res graphics (60x60).
* Simple controls (cursor or wasd keys / xbox one d-pad / mouse clicks).

## Browser Demo

Play the unfinished game here https://theinvader360.github.io/dungeon-crawler/

## Local Setup

    git clone https://github.com/TheInvader360/dungeon-crawler
    cd dungeon-crawler/
    go test ./...
    go run main.go

## Tasks / Features

### Minimal Features

- [x] Grid based dungeon map
- [x] Player movement - turn left, turn right, move forward
- [x] First person exploration view (e.g. [1](https://en.wikipedia.org/wiki/Maze_War)/[2](https://en.wikipedia.org/wiki/Wizardry:_Proving_Grounds_of_the_Mad_Overlord)/[3](https://en.wikipedia.org/wiki/3D_Monster_Maze)/[4](https://en.wikipedia.org/wiki/Dungeons_of_Daggorath)/[5](https://en.wikipedia.org/wiki/The_Bard%27s_Tale_(1985_video_game))/[6](https://en.wikipedia.org/wiki/Might_and_Magic_Book_One:_The_Secret_of_the_Inner_Sanctum)/[7](https://en.wikipedia.org/wiki/Dungeon_Master_(video_game))/[8](https://en.wikipedia.org/wiki/Eye_of_the_Beholder_(video_game)))
- [x] Top down mini map (locked to player position)
- [x] Permanent blocked cells (walls)
- [x] Removable blocked cells (breakable/unlockable)
- [x] Enemies
- [x] First person billboards (enemies/collectibles/etc)
- [x] Key collectibles
- [x] Loot collectibles
- [x] Health collectibles
- [x] Load next dungeon on exit
- [x] Game over state
- [ ] Game completed state
- [ ] Player and enemy stats
- [ ] Combat simulation (turn based? rhythm based? qte based?)

### Feature Creep Corral

- [x] Plot and MacGuffins (the thinner and schlockier the better!)
- [x] Slideshows (intro, game over, game completed)
- [x] Support mouse input
- [x] Support xbox one gamepad input
- [ ] Support other gamepad input
- [ ] Support touchscreen input
- [ ] Mobile build
- [ ] Only show visited/seen cells on mini map
- [ ] Map collectible (reveal full dungeon on mini map)
- [ ] More enemies and bosses
- [ ] Stat modifier collectibles (e.g. hpmax+/atk+/def+)
- [ ] Traps (e.g. instant death/lose health/modify stat)
- [ ] Bestiary (unlock entry after defeating enemy)
- [ ] Merchants
- [ ] Save and load progress
- [ ] Character select or creation
- [ ] Animation
- [ ] Enemy movement (basic fight or flight AI with pathfinding)
- [ ] Sound effects
- [ ] Music
