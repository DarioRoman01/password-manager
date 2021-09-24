package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/DarioRoman01/password_manager/encrypt"
	"github.com/DarioRoman01/password_manager/passwords"
)

func makeFirstTab() *container.TabItem {
	tab := container.NewTabItem("Generate Password", widget.NewLabel("Content of tab 1"))
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

func SeePasswordsTab() *container.TabItem {
	tab := container.NewTabItem("See Passwords", widget.NewLabel("Content of tab 1"))
	label := widget.NewLabel("Passwords")
	entry := widget.Entry{PlaceHolder: "encryption password"}
	tab.Content = container.NewVBox(
		label,
		&entry,
		widget.NewButton("see passwords", func() {
			pwd := entry.Text
			content, err := encrypt.Desencrypt([]byte(pwd))
			if err != nil {
				label.Text = "Wrong password"
				return
			}

			label.SetText(string(content))
		}),
	)

	return tab
}

func SetPasswordTab() *container.TabItem {
	tab := container.NewTabItem("Set Password", widget.NewLabel("Content of tab 1"))
	label := widget.NewLabel("Set Password")

	keyEntry := widget.Entry{PlaceHolder: "Key"}
	Pwdentry := widget.Entry{PlaceHolder: "new password"}

	tab.Content = container.NewVBox(
		label,
		&keyEntry,
		&Pwdentry,
		widget.NewButton("add password", func() {
			key := keyEntry.Text
			pwd := Pwdentry.Text
			err := encrypt.AddPassword([]byte(pwd), []byte(key))
			if err != nil {
				label.SetText("Something wrong happend")
				return
			}

			label.SetText("password added succesfully")
		}),
	)

	return tab
}

func InitSessionView(w fyne.Window) (err error) {
	label := widget.NewLabel("Insert encryption password")
	entry := widget.Entry{PlaceHolder: "password"}
	w.SetContent(container.NewVBox(
		label,
		&entry,
		widget.NewButton("verify", func() {
			_, err = encrypt.Desencrypt([]byte(entry.Text))
			if err != nil {
				label.SetText("wrong password")
				return
			}

			err = nil
		}),
	))

	return nil
}

func makeAppTabsTab() fyne.CanvasObject {
	tabs := container.NewAppTabs(
		makeFirstTab(),
		SeePasswordsTab(),
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

	label := widget.NewLabel("Insert encryption password")
	entry := widget.Entry{PlaceHolder: "password", Password: true}
	w.SetContent(container.NewVBox(
		label,
		&entry,
		widget.NewButton("verify", func() {
			_, err := encrypt.Desencrypt([]byte(entry.Text))
			if err != nil {
				label.SetText("wrong password")
				return
			}

			w.SetContent(makeAppTabsTab())
		}),
	))

	w.ShowAndRun()
}
