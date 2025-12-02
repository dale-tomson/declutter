package theme

import (
	"image/color"
	"testing"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// TestNew verifies that New creates a CustomTheme
func TestNew(t *testing.T) {
	th := New()
	if th == nil {
		t.Error("New() returned nil")
	}

	customTheme, ok := th.(*CustomTheme)
	if !ok {
		t.Error("New() did not return *CustomTheme")
	}
	if customTheme.Theme == nil {
		t.Error("CustomTheme.Theme is nil")
	}
}

// TestColorPrimary verifies the primary color
func TestColorPrimary(t *testing.T) {
	th := New()

	c := th.Color(theme.ColorNamePrimary, theme.VariantDark)
	nrgba, ok := c.(color.NRGBA)
	if !ok {
		t.Fatal("Color is not NRGBA")
	}

	if nrgba.R != 0x42 || nrgba.G != 0xb8 || nrgba.B != 0x83 {
		t.Errorf("Expected primary green (#42b883), got #%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
	}
}

// TestColorBackground verifies the background color
func TestColorBackground(t *testing.T) {
	th := New()

	c := th.Color(theme.ColorNameBackground, theme.VariantDark)
	nrgba, ok := c.(color.NRGBA)
	if !ok {
		t.Fatal("Color is not NRGBA")
	}

	if nrgba.R != 0x35 || nrgba.G != 0x49 || nrgba.B != 0x5e {
		t.Errorf("Expected background dark (#35495e), got #%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
	}
}

// TestColorError verifies the error/warning color
func TestColorError(t *testing.T) {
	th := New()

	c := th.Color(theme.ColorNameError, theme.VariantDark)
	nrgba, ok := c.(color.NRGBA)
	if !ok {
		t.Fatal("Color is not NRGBA")
	}

	if nrgba.R != 0xff || nrgba.G != 0x7e || nrgba.B != 0x67 {
		t.Errorf("Expected highlight coral (#ff7e67), got #%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
	}
}

// TestColorForeground verifies the foreground/text color
func TestColorForeground(t *testing.T) {
	th := New()

	c := th.Color(theme.ColorNameForeground, theme.VariantDark)
	nrgba, ok := c.(color.NRGBA)
	if !ok {
		t.Fatal("Color is not NRGBA")
	}

	if nrgba.R != 0xff || nrgba.G != 0xff || nrgba.B != 0xff {
		t.Errorf("Expected text white (#ffffff), got #%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
	}
}

// TestColorSecondary verifies colors using secondary teal
func TestColorSecondary(t *testing.T) {
	th := New()

	testCases := []struct {
		name      string
		colorName fyne.ThemeColorName
	}{
		{"Pressed", theme.ColorNamePressed},
		{"ScrollBar", theme.ColorNameScrollBar},
		{"InputBorder", theme.ColorNameInputBorder},
		{"Selection", theme.ColorNameSelection},
		{"Separator", theme.ColorNameSeparator},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := th.Color(tc.colorName, theme.VariantDark)
			nrgba, ok := c.(color.NRGBA)
			if !ok {
				t.Fatal("Color is not NRGBA")
			}

			if nrgba.R != 0x34 || nrgba.G != 0x74 || nrgba.B != 0x74 {
				t.Errorf("Expected secondary teal (#347474), got #%02x%02x%02x", nrgba.R, nrgba.G, nrgba.B)
			}
		})
	}
}

// TestExportedColors verifies the exported color variables
func TestExportedColors(t *testing.T) {
	tests := []struct {
		name     string
		color    color.NRGBA
		expected color.NRGBA
	}{
		{"PrimaryGreen", PrimaryGreen, color.NRGBA{R: 0x42, G: 0xb8, B: 0x83, A: 255}},
		{"SecondaryTeal", SecondaryTeal, color.NRGBA{R: 0x34, G: 0x74, B: 0x74, A: 255}},
		{"BackgroundDark", BackgroundDark, color.NRGBA{R: 0x35, G: 0x49, B: 0x5e, A: 255}},
		{"HighlightCoral", HighlightCoral, color.NRGBA{R: 0xff, G: 0x7e, B: 0x67, A: 255}},
		{"BackgroundLight", BackgroundLight, color.NRGBA{R: 0x3d, G: 0x54, B: 0x6b, A: 255}},
		{"TextWhite", TextWhite, color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 255}},
		{"DialogBg", DialogBg, color.NRGBA{R: 0x2d, G: 0x3e, B: 0x50, A: 255}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.color != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, tt.color)
			}
		})
	}
}
