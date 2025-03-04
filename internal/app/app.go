package app

import (
	"image/color"
	"to-do-app/views"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func Run() {
	myApp := app.NewWithID("com.dohd.todo-app")
	window := myApp.NewWindow("TO-DO LIST")
	window.Resize(fyne.NewSize(400, 400))

	tabs := container.NewAppTabs(
		container.NewTabItem("\tPersonal\t\t", views.PersonalTab()),
		container.NewTabItem("\tProfessional\t", views.ProfessionalTab()),
	)

	// light gray background
	bg := canvas.NewRectangle(color.RGBA{R: 90, G: 90, B: 90, A: 255})
	window.SetContent(container.NewStack(bg, tabs))
	window.ShowAndRun()
}
