package icon

import (
	_ "embed"

	"fyne.io/fyne/v2"
)

//go:embed logo.svg
var logoData []byte

func Resource() fyne.Resource {
	return fyne.NewStaticResource("logo.svg", logoData)
}
