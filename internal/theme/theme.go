package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	PrimaryGreen    = color.NRGBA{R: 0x42, G: 0xb8, B: 0x83, A: 255}
	SecondaryTeal   = color.NRGBA{R: 0x34, G: 0x74, B: 0x74, A: 255}
	BackgroundDark  = color.NRGBA{R: 0x35, G: 0x49, B: 0x5e, A: 255}
	BackgroundLight = color.NRGBA{R: 0x3d, G: 0x54, B: 0x6b, A: 255}
	DialogBg        = color.NRGBA{R: 0x2d, G: 0x3e, B: 0x50, A: 255}
	HighlightCoral  = color.NRGBA{R: 0xff, G: 0x7e, B: 0x67, A: 255}
	HoverGreen      = color.NRGBA{R: 0x4a, G: 0xc9, B: 0x92, A: 255}
	TextWhite       = color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 255}
	PlaceholderText = color.NRGBA{R: 0xaa, G: 0xbb, B: 0xcc, A: 255}
	DisabledButton  = color.NRGBA{R: 0x50, G: 0x65, B: 0x78, A: 255}
	ShadowColor     = color.NRGBA{R: 0x20, G: 0x30, B: 0x40, A: 100}
)

type CustomTheme struct {
	fyne.Theme
}

func New() fyne.Theme {
	return &CustomTheme{Theme: theme.DefaultTheme()}
}

func (c *CustomTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNamePrimary, theme.ColorNameButton, theme.ColorNameFocus, theme.ColorNameSuccess, theme.ColorNameHyperlink:
		return PrimaryGreen

	case theme.ColorNameHover:
		return HoverGreen

	case theme.ColorNamePressed, theme.ColorNameSelection, theme.ColorNameScrollBar, theme.ColorNameInputBorder, theme.ColorNameSeparator, theme.ColorNameHeaderBackground:
		return SecondaryTeal

	case theme.ColorNameBackground:
		return BackgroundDark

	case theme.ColorNameInputBackground:
		return BackgroundLight

	case theme.ColorNameMenuBackground, theme.ColorNameOverlayBackground:
		return DialogBg

	case theme.ColorNameForeground:
		return TextWhite

	case theme.ColorNameDisabled, theme.ColorNamePlaceHolder:
		return PlaceholderText

	case theme.ColorNameDisabledButton:
		return DisabledButton

	case theme.ColorNameError, theme.ColorNameWarning:
		return HighlightCoral

	case theme.ColorNameShadow:
		return ShadowColor
	}

	return c.Theme.Color(name, theme.VariantDark)
}
