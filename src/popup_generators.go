package main

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func genKeyBoardShortcutPopup() {
	helpPop := a.NewWindow("Help")
	helpPop.SetContent(container.NewVScroll(container.NewVBox(
		widget.NewLabel("Keyboard Shortcuts"),
		widget.NewSeparator(),
		widget.NewLabel("Super+F -> Focus on Inquiry Input Field"),
		widget.NewLabel("Super+B -> Clear the Inquiry Input Field"),
		widget.NewLabel("Super+G -> Open Add View"),
		widget.NewLabel("Super+E -> Open Edit View"),
		widget.NewLabel("Super+I -> Open Inquiry View"),
		widget.NewLabel("Super+R -> Open Switch List View"),
		widget.NewLabel("Super+Up -> Switch List (alphabetically up)"),
		widget.NewLabel("Super+Down -> Switch List (alphabetically down)"),
	)))
	helpPop.Resize(fyne.NewSize(210, 275))

	if deskCanvas, ok := helpPop.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				helpPop.Close()
			}
		})
	}
	helpPop.Show()
}

// used for export to csv and to json
func NewExportPop(filetype string) {
	exp := a.NewWindow("Export " + filetype)
	fullExport := false
	cancel := widget.NewButton("Cancel", func() {
		exp.Close()
	})
	fileName := widget.NewEntry()
	fileName.SetPlaceHolder("File name")
	submit := widget.NewButton("Submit", func() {
		switch filetype {
		case "CSV":
			write_csv(fullExport, fileName.Text)
		case "JSON":
			write_json(fullExport, fileName.Text)
		}
	})
	submit.Disable()
	fileName.Validator = validation.NewRegexp(`^.+$`, "file name cannot be empty")
	fileName.SetOnValidationChanged(func(err error) {
		if err != nil {
			submit.Disable()
		} else {
			submit.Enable()
		}
	})
	check := widget.NewCheck("Full Data", func(value bool) {
		fullExport = value
	})

	exp.SetContent(container.NewVScroll(container.NewVBox(
		check,
		fileName,
		container.NewHBox(submit, cancel),
	)))
	exp.Resize(fyne.NewSize(400, 150))

	if deskCanvas, ok := exp.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				exp.Close()
			}
		})
	}
	exp.Show()
}

// used for exporting of word cloud images
func NewImgExportPop(img image.Image) {
	exp := a.NewWindow("Export Image")
	cancel := widget.NewButton("Cancel", func() {
		exp.Close()
	})
	fileName := widget.NewEntry()
	fileName.SetPlaceHolder("File name")
	submit := widget.NewButton("Submit", func() {
		write_png(img, fileName.Text)
		exp.Close()
	})
	submit.Disable()
	fileName.Validator = validation.NewRegexp(`^.+\.png$`, ".png file name cannot be empty")
	fileName.SetOnValidationChanged(func(err error) {
		if err != nil {
			submit.Disable()
		} else {
			submit.Enable()
		}
	})

	exp.SetContent(container.NewVScroll(container.NewVBox(
		fileName,
		container.NewHBox(submit, cancel),
	)))
	exp.Resize(fyne.NewSize(400, 90))

	if deskCanvas, ok := exp.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {
			if key.Name == fyne.KeyEscape {
				exp.Close()
			}
		})
	}
	exp.Show()
}
