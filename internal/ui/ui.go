package ui

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/dale-tomson/declutter/internal/icon"
	"github.com/dale-tomson/declutter/internal/organizer"
)

type App struct {
	window              fyne.Window
	selectedFolder      string
	selectedFolderLabel *widget.Label
	logOutput           *widget.Entry
	progress            *widget.ProgressBar
	statusLabel         *widget.Label
	selectFolderBtn     *widget.Button
	organizeBtn         *widget.Button
}

func New(w fyne.Window) *App {
	app := &App{window: w}
	app.setupUI()
	return app
}

func (a *App) GetContent() fyne.CanvasObject {
	return a.buildLayout()
}

func (a *App) setupUI() {
	a.selectedFolderLabel = widget.NewLabel("No folder selected")
	a.selectedFolderLabel.Wrapping = fyne.TextWrapWord

	a.logOutput = widget.NewMultiLineEntry()
	a.logOutput.Wrapping = fyne.TextWrapWord
	a.logOutput.Disable()
	a.logOutput.SetMinRowsVisible(12)

	a.progress = widget.NewProgressBar()
	a.progress.Hide()

	a.statusLabel = widget.NewLabel("")
	a.statusLabel.Alignment = fyne.TextAlignCenter

	a.organizeBtn = widget.NewButton("Organize Files", nil)
	a.organizeBtn.Importance = widget.HighImportance
	a.organizeBtn.Disable()
	a.organizeBtn.OnTapped = a.onOrganize

	a.selectFolderBtn = widget.NewButton("Select Folder", a.onSelectFolder)
	a.selectFolderBtn.Importance = widget.MediumImportance
}

func (a *App) buildLayout() fyne.CanvasObject {
	logoImg := canvas.NewImageFromResource(icon.Resource())
	logoImg.FillMode = canvas.ImageFillOriginal
	logoImg.ScaleMode = canvas.ImageScaleSmooth
	logoImg.SetMinSize(fyne.NewSize(100, 80))

	logoBorder := canvas.NewRectangle(color.Transparent)
	logoBorder.StrokeColor = color.White
	logoBorder.StrokeWidth = 2
	logoBorder.CornerRadius = 5
	logoWithBorder := container.NewStack(logoBorder, container.NewPadded(logoImg))

	titleLabel := canvas.NewText("Declutter", nil)
	titleLabel.TextSize = 28
	titleLabel.TextStyle = fyne.TextStyle{Bold: true}

	descLabel := widget.NewLabelWithStyle(
		"Organize your files into Year/Month folders\nbased on their timestamps",
		fyne.TextAlignLeading,
		fyne.TextStyle{},
	)

	titleSection := container.NewVBox(titleLabel, descLabel)
	headerContent := container.NewHBox(logoWithBorder, container.NewCenter(titleSection))

	folderSection := container.NewVBox(
		widget.NewLabel("Selected Folder:"),
		container.NewBorder(nil, nil, nil, nil, a.selectedFolderLabel),
	)

	buttons := container.NewHBox(
		a.selectFolderBtn,
		a.organizeBtn,
	)

	logSection := container.NewVBox(
		widget.NewLabel("Activity Log:"),
		container.NewMax(a.logOutput),
	)

	content := container.NewBorder(
		container.NewVBox(
			headerContent,
			widget.NewSeparator(),
			folderSection,
			buttons,
			a.progress,
			a.statusLabel,
		),
		nil,
		nil,
		nil,
		logSection,
	)

	return container.NewPadded(content)
}

func (a *App) log(message string) {
	fyne.Do(func() {
		current := a.logOutput.Text
		if current != "" {
			current += "\n"
		}
		a.logOutput.SetText(current + message)
		a.logOutput.CursorRow = strings.Count(a.logOutput.Text, "\n")
	})
}

func (a *App) onSelectFolder() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, a.window)
			return
		}
		if uri == nil {
			return
		}

		a.selectedFolder = uri.Path()
		a.selectedFolderLabel.SetText("📂 " + a.selectedFolder)
		a.organizeBtn.Enable()
		a.logOutput.SetText("")
		a.statusLabel.SetText("")

		org := organizer.New(a.selectedFolder, nil)
		files, err := org.GetFiles()
		if err != nil {
			a.log(fmt.Sprintf("Error reading folder: %v", err))
			return
		}
		a.log(fmt.Sprintf("Found %d files to organize", len(files)))
	}, a.window)
}

func (a *App) onOrganize() {
	if a.selectedFolder == "" {
		dialog.ShowInformation("Info", "Please select a folder first", a.window)
		return
	}

	dialog.ShowConfirm("Confirm Organization",
		fmt.Sprintf("This will organize all files in:\n%s\n\nFiles will be moved to Year/Month folders based on their modification dates.\n\nContinue?", a.selectedFolder),
		func(confirmed bool) {
			if !confirmed {
				return
			}
			a.performOrganization()
		}, a.window)
}

func (a *App) performOrganization() {
	a.logOutput.SetText("")
	a.progress.Show()
	a.progress.SetValue(0)
	a.selectFolderBtn.Disable()
	a.organizeBtn.Disable()
	a.statusLabel.SetText("Organizing...")

	go func() {
		org := organizer.New(a.selectedFolder, func(msg string) {
			a.log(msg)
		})

		files, err := org.GetFiles()
		if err != nil {
			a.log(fmt.Sprintf("Error: %v", err))
			fyne.Do(func() {
				a.progress.Hide()
				a.selectFolderBtn.Enable()
				a.organizeBtn.Enable()
				a.statusLabel.SetText("Error occurred")
			})
			return
		}

		if len(files) == 0 {
			a.log("No files found to organize")
			fyne.Do(func() {
				a.progress.Hide()
				a.selectFolderBtn.Enable()
				a.organizeBtn.Enable()
				a.statusLabel.SetText("No files to organize")
			})
			return
		}

		fyne.Do(func() {
			a.progress.SetValue(0.2)
		})
		a.log("Starting organization...")

		moved, skipped, err := org.OrganizeFiles(files)

		fyne.Do(func() {
			a.progress.SetValue(1.0)
		})

		if err != nil {
			a.log(fmt.Sprintf("Error during organization: %v", err))
		}

		a.log("─────────────────────────────")
		a.log(fmt.Sprintf("✅ Complete! Moved: %d, Skipped: %d", moved, skipped))

		fyne.Do(func() {
			a.progress.Hide()
			a.selectFolderBtn.Enable()
			a.organizeBtn.Enable()
			a.statusLabel.SetText(fmt.Sprintf("Done! %d files moved, %d skipped", moved, skipped))
		})
	}()
}
