package assets

import(
	"embed"
	_ "image/png"
	_ "image/jpeg"
)

//go:embed winddorf
var Data embed.FS

