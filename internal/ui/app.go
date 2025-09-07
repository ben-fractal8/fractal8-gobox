package ui

import (
	"fractal8-gobox/internal/tools"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

	//List tools tab
	btnAppend := widget.NewButton("Append/Prepend List", func() {
		input := widget.NewMultiLineEntry()
		prefix := widget.NewEntry()
		suffix := widget.NewEntry()
		form := widget.NewForm(
			widget.NewFormItem("List (one item per line)", input),
			widget.NewFormItem("Prefix", prefix),
			widget.NewFormItem("Suffix", suffix),
		)

		dialog.ShowCustomConfirm("Append/Prepend", "Run", "Cancel", form, func(ok bool) {
			if ok {
				lines := strings.Split(input.Text, "\n")
				result := tools.AppendOrPrependList(lines, prefix.Text, suffix.Text)
				dialog.ShowInformation("Result", strings.Join(result, "\n"), window)
			}
		}, window)
	})
	descAppend := widget.NewLabel("Add text to the start or end of each item in a list")

	btnListToCSV := widget.NewButton("List<->CSV", func() {
		input := widget.NewMultiLineEntry()
		mode := widget.NewSelect([]string{"List → CSV", "CSV → List"}, func(_ string) {})

		form := widget.NewForm(
			widget.NewFormItem("Input", input),
			widget.NewFormItem("Mode", mode),
		)

		dialog.ShowCustomConfirm("List ↔ CSV", "Convert", "Cancel", form, func(ok bool) {
			if ok {
				if mode.Selected == "List → CSV" {
					lines := strings.Split(input.Text, "\n")
					result := tools.ListToCSV(lines)
					dialog.ShowInformation("CSV Result", result, window)
				} else if mode.Selected == "CSV → List" {
					result, err := tools.CSVToList(input.Text)
					if err != nil {
						dialog.ShowError(err, window)
					} else {
						dialog.ShowInformation("List Result", strings.Join(result, "\n"), window)
					}
				}
			}
		}, window)
	})
	descListToCSV := widget.NewLabel("Convert a list to CSV or vice versa")

	listTab := container.NewVBox(
		btnAppend,
		descAppend,
		widget.NewSeparator(),
		btnListToCSV,
		descListToCSV,
	)

	//Layout Tabs

	tabs := container.NewAppTabs(
		container.NewTabItem("Lists", listTab),
		container.NewTabItem("Strings", widget.NewLabel("String tools here...")),
		container.NewTabItem("Numbers", widget.NewLabel("CSV tools here...")),
	)

	//Main VBox
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		tabs,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(400, 300))
	window.ShowAndRun()
}
