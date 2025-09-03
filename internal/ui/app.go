package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Run initializes and starts the GoBox application
func Run() {
	// Create the Fyne app
	gobox := app.New()
	window := gobox.NewWindow("GoBox")

	// Welcome label
	title := widget.NewLabelWithStyle(
		"Welcome to GoBox",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	// Placeholder buttons for future tools
	btnAppend := widget.NewButton("Append Text", func() {
		// TODO: Add append text functionality
	})
	btnPrepend := widget.NewButton("Prepend Text", func() {
		// TODO: Add prepend text functionality
	})
	btnCSV := widget.NewButton("CSV Tools", func() {
		// TODO: Add CSV functionality
	})

	// Layout: vertical box with spacing
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		btnAppend,
		btnPrepend,
		btnCSV,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 300))
	window.ShowAndRun()
}
