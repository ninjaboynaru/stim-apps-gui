package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	ACTION_INSTALL     string = "Install"
	ACTION_UNINSTALL   string = "Uninstall"
	APP_PLAY_STORE     string = "Play Store"
	APP_CHROME_BROWSER string = "Chrome Browser"
	APP_YOUTUBE        string = "Youtube"
	APP_YT_MUSIC       string = "YT Music"
)

func main() {
	app := app.New()
	window := app.NewWindow("Stim Apps")
	window.SetFixedSize(true)

	var mainGrid *fyne.Container = container.New(layout.NewVBoxLayout())
	var actionRadioGroup *widget.RadioGroup = widget.NewRadioGroup([]string{ACTION_INSTALL, ACTION_UNINSTALL}, func(value string) {
		fmt.Println("Action: " + value)
	})
	actionRadioGroup.Horizontal = true
	actionRadioGroup.SetSelected(ACTION_UNINSTALL)

	var appCheckGroup *widget.CheckGroup = widget.NewCheckGroup(
		[]string{
			APP_PLAY_STORE,
			APP_CHROME_BROWSER,
			APP_YOUTUBE,
			APP_YT_MUSIC,
		},
		func(value []string) {
			fmt.Println("Apps: ", value)
		},
	)

	appCheckGroup.SetSelected([]string{
		APP_PLAY_STORE,
		APP_CHROME_BROWSER,
		APP_YOUTUBE,
		APP_YT_MUSIC,
	})

	var submitButton *widget.Button = widget.NewButtonWithIcon("Submit", theme.MediaPlayIcon(), func() {
		fmt.Println("Submitted")
	})

	mainGrid.Add(actionRadioGroup)
	mainGrid.Add(appCheckGroup)
	mainGrid.Add(submitButton)

	window.SetContent(mainGrid)
	// window.Resize(fyne.NewSize(400, 400))
	window.Show()
	app.Run()

}
