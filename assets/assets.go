package assets

import(
	"embed"
	_ "image/png"
	_ "image/jpeg"
)

//go:embed winddorf
//go:embed sprites
//go:embed fonts
var Data embed.FS
