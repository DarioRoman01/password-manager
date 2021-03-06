package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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

func SeePasswordsTab(pwd string) *container.TabItem {
	tab := container.NewTabItem("See Passwords", widget.NewLabel("Content of tab 1"))
	label := widget.NewLabel("Passwords")
	tab.Content = container.NewVBox(
		label,
		widget.NewButton("see passwords", func() {
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

func SetPasswordTab(key string) *container.TabItem {
	tab := container.NewTabItem("Set Password", widget.NewLabel("Content of tab 1"))
	label := widget.NewLabel("Set Password")

	Pwdentry := widget.Entry{PlaceHolder: "new password"}

	tab.Content = container.NewVBox(
		label,
		&Pwdentry,
		widget.NewButton("add password", func() {
			pwd := fmt.Sprintf("\n%s", Pwdentry.Text)
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

func makeAppTabsTab(key string) fyne.CanvasObject {
	tabs := container.NewAppTabs(
		makeFirstTab(),
		SeePasswordsTab(key),
		SetPasswordTab(key),
	)

	return container.NewBorder(nil, nil, nil, nil, tabs)
}

func confirmCallback(w fyne.Window, label *widget.Label, text string) func(bool) {
	return func(response bool) {
		if !response {
			return
		}

		if err := encrypt.NewFile([]byte(text)); err != nil {
			label.SetText(err.Error())
			return
		}

		w.SetContent(makeAppTabsTab(text))
	}
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

			w.SetContent(makeAppTabsTab(entry.Text))
		}),
		widget.NewButton("New", func() {
			cnf := dialog.NewConfirm("Confimation", "Are you sure to create a new file?", confirmCallback(w, label, entry.Text), w)
			cnf.SetConfirmText("Sure")
			cnf.SetDismissText("Nah")
			cnf.Show()

		}),
	))

	w.ShowAndRun()
}
