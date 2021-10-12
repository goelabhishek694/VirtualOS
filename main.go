package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/theme"
	"log"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")
	searchBox:=widget.NewEntry()
	searchBox.SetPlaceHolder("Search tex...")
	searchBtn:=widget.NewButtonWithIcon("Search", theme.SearchIcon(),func() {
		log.Println("tapped search button")
		log.Println(searchBox.Text)
	})
	installBtn := widget.NewButtonWithIcon("Install", theme.ContentAddIcon(), func() {
		log.Println("tapped install button")
	})
	grid := container.New(layout.NewGridLayout(2),searchBox,searchBtn)
	content := container.NewVBox(
		widget.NewLabel("Welcome to pepOS"),
		container.NewHBox(
			grid,
			// widget.NewLabel("install app"),
			installBtn,
		),
	)

	// content.Add(widget.NewButton("Add more items", func() {
	// 	content.Add(widget.NewLabel("Added"))
	// }))
	

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
