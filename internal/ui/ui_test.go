package ui

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestNew(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if ui == nil {
		t.Fatal("New() returned nil")
	}

	if ui.window != w {
		t.Error("window not set correctly")
	}
}

func TestGetContent(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	content := ui.GetContent()
	if content == nil {
		t.Fatal("GetContent() returned nil")
	}
}

func TestUIComponentsInitialized(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if ui.selectedFolderLabel == nil {
		t.Error("selectedFolderLabel not initialized")
	}

	if ui.logOutput == nil {
		t.Error("logOutput not initialized")
	}

	if ui.progress == nil {
		t.Error("progress not initialized")
	}

	if ui.statusLabel == nil {
		t.Error("statusLabel not initialized")
	}

	if ui.selectFolderBtn == nil {
		t.Error("selectFolderBtn not initialized")
	}

	if ui.organizeBtn == nil {
		t.Error("organizeBtn not initialized")
	}
}

func TestOrganizeButtonDisabledByDefault(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if !ui.organizeBtn.Disabled() {
		t.Error("organizeBtn should be disabled by default")
	}
}

func TestSelectFolderButtonEnabled(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if ui.selectFolderBtn.Disabled() {
		t.Error("selectFolderBtn should be enabled by default")
	}
}

func TestInitialFolderLabel(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if ui.selectedFolderLabel.Text != "No folder selected" {
		t.Errorf("expected 'No folder selected', got '%s'", ui.selectedFolderLabel.Text)
	}
}

func TestProgressHiddenByDefault(t *testing.T) {
	app := test.NewApp()
	defer app.Quit()

	w := app.NewWindow("Test")
	ui := New(w)

	if !ui.progress.Hidden {
		t.Error("progress bar should be hidden by default")
	}
}
