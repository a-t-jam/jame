package travel

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/a-t-jam/jame/assets"
	"github.com/a-t-jam/jame/game/combat"
	"github.com/a-t-jam/jame/game/dialog"
	"github.com/a-t-jam/jame/game/scene"
	"github.com/a-t-jam/jame/game/win"
	"github.com/a-t-jam/jame/ui"
)

var (
	playerSprite *ebiten.Image
	playerNode   ui.Node
	playerPos    int
	isWalking    bool
)

func init() {
	playerPos = 0
	playerNode = ui.Node{
		X:       1280.0 / 2.0,
		Y:       720.0 - 200.0,
		Align:   ui.AlignCenter,
		Surface: scene.WalkDuckSurface,
	}
}

func updateAnim() {
	// update anim here
	print(playerPos)
	if playerNode.Surface.CurrentFrameIx == 3 {
		isWalking = false
	}
}

func Update(scene_ *scene.Scene) error {
	if isWalking {
		updateAnim()
		return nil
	}

	if playerPos == 16 {
		scene_.State = scene.WinState
		win.Enter(scene_)
		return nil
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		playerPos += 1

		isWalking = true
		assets.Bubble.Rewind()
		assets.Bubble.Play()

		return dialog.Update(scene_, dialog.Dialogs["on_move_forward"])
	}

	if playerPos == 0 {
		return dialog.Update(scene_, dialog.Dialogs["moving_instruction"])
	}

	if playerPos%5 == 0 {
		// FIXME: hack to move forward AFTER the battle
		playerPos += 1

		// enter combat scene
		scene_.State = scene.CombatState
		// TODO: with corresponding enemy for this positiion
		combat.Enter(scene_, combat.Enemy1)

		return nil
	}

	return dialog.Update(scene_, dialog.Dialogs["on_move_forward"])
}

func Draw(scene *scene.Scene, screen *ebiten.Image) {
	assets.DrawOcean1(screen)

	playerNode.Draw(screen)

	if isWalking {
		updateAnims(scene, screen)
	}
	dialog.Draw(scene, screen)

	debugDraw(scene, screen)
}

func debugDraw(scene *scene.Scene, screen *ebiten.Image) {
	message := fmt.Sprintf("FPS: %v", ebiten.CurrentFPS())
	text.Draw(screen, message, assets.PixelFont, 100.0, 300.0, color.White)
}

func updateAnims(scene_ *scene.Scene, screen *ebiten.Image) {
	// the duck animation
	s := playerNode.Surface

	elapsed := time.Since(scene.StartTime)
	n := elapsed.Milliseconds() / (1000 * 8 / 60)

	n_frames := len(s.Uvs)
	n_pingpong := (n_frames)*2 - 1

	frame := int(n) % n_pingpong
	if frame >= n_frames {
		frame -= n_frames
	}

	s.CurrentFrameIx = frame
}
