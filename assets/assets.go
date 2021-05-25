package assets

import(
	"embed"
	_ "image/png"
	_ "image/jpeg"
)

//go:embed winddorf
//go:embed sprites
var Data embed.FS
