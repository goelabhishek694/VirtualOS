package main

import (
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	// "fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("pepOS")
	myWindow.Resize(fyne.NewSize(1080,1920))
	showDesktop:=widget.NewButtonWithIcon("Show Desktop",theme.ComputerIcon(),func() {
		log.Println("show desktop clicked")
	});
	searchBox:=widget.NewEntry()
	searchBox.SetPlaceHolder("Search text...")
	searchBtn:=widget.NewButtonWithIcon("Search", theme.SearchIcon(),func() {
		log.Println("tapped search button")
		log.Println(searchBox.Text)
	})
	installBtn := widget.NewButtonWithIcon("Install", theme.ContentAddIcon(), func() {
		fileDialog:=dialog.NewFileOpen(func(r fyne.URIReadCloser,_ error){
			data,_:=ioutil.ReadAll(r)
			//display data in a new multiline netry
			result:=fyne.NewStaticResource("file",data)
			entry:=widget.NewMultiLineEntry() 
			entry.SetText(string(result.StaticContent))
			myWindow:=fyne.CurrentApp().NewWindow(
				string(result.StaticName))
			myWindow.SetContent(container.NewScroll(entry))
			myWindow.Show()
		},myWindow)
		// fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		fileDialog.Show()
		log.Println("tapped install button")
	})
	grid := container.New(layout.NewGridLayout(2),searchBox,searchBtn)
	content := container.NewVBox(
		showDesktop,
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
