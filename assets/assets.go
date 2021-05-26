package assets

import (
	"embed"
	_ "image/jpeg"
	_ "image/png"
)

//go:embed winddorf
//go:embed sprites
//go:embed fonts
var Data embed.FS
