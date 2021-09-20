package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/DarioRoman01/password_manager/passwords"
)

func makeFirstTab() *container.TabItem {
	tab := container.NewTabItem("Tab 1", widget.NewLabel("Content of tab 1"))
	label := widget.NewLabel("Generate password")
	tab.Content = container.NewVBox(
		label,
		widget.NewButton("generate password", func() {
			password := passwords.GeneratePassword()
			label.SetText(password)
		}),
	)

	return tab
}

func MakeTabLocationSelect(callback func(container.TabLocation)) *widget.Select {
	locations := widget.NewSelect([]string{"Top", "Bottom", "Leading", "Trailing"}, func(s string) {
		callback(map[string]container.TabLocation{
			"Top":      container.TabLocationTop,
			"Bottom":   container.TabLocationBottom,
			"Leading":  container.TabLocationLeading,
			"Trailing": container.TabLocationTrailing,
		}[s])
	})
	locations.SetSelected("Top")
	return locations
}

func makeAppTabsTab() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		makeFirstTab(),
		container.NewTabItem("Tab 2 bigger", widget.NewLabel("Content of tab 2")),
		container.NewTabItem("Tab 3", widget.NewLabel("Content of tab 3")),
	)

	return container.NewBorder(nil, nil, nil, nil, tabs)
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	w.Resize(fyne.Size{
		Width:  600,
		Height: 600,
	})

	w.SetContent(makeAppTabsTab())
	w.ShowAndRun()
}
