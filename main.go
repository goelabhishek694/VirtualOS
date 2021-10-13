package main

import (
	"fmt"
	"io/ioutil"
	"log"
	// "os"
	"time"
	// "fmt"
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
	myWindow.Resize(fyne.NewSize(250,400))
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	// myWindow.SetContent(tabs)
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
		fileDialog:=dialog.NewFolderOpen(func(f fyne.ListableURI,_ error){
			fmt.Printf("typeof f=%T\n",f.Path())
			fmt.Println(f.Path())
			files,_:=ioutil.ReadDir(f.Path())
			isPresent:=false
			for _, file:=range files{
				if file.Name() == "metadata.json" {
					isPresent=true
					dir:=f.Path()+"/metadata.json"
					checkIfCorrectFormat(dir)
				}
				fmt.Println("hello ",file.Name())
			}
			if !isPresent{
				//show a error window
				go showAnother(myApp)
			}
			//display data in a new multiline entry
			// result:=fyne.NewStaticResource("file",files)
			// entry:=widget.NewMultiLineEntry() 
			// entry.SetText(string(result.StaticContent))
			// myWindow:=fyne.CurrentApp().NewWindow(
			// 	string(result.StaticName))
			// myWindow.SetContent(container.NewScroll(entry))
			// myWindow.Show()
		},myWindow)
		// fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		fileDialog.Show()
		log.Println("tapped install button")
	})
	grid := container.New(layout.NewGridLayout(2),searchBox,searchBtn)
	
	content := container.NewVBox(
		container.NewHBox(showDesktop,tabs),
		
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

func checkIfCorrectFormat(dir string){
	content ,_:=ioutil.ReadFile(dir)

	fmt.Println("content ", string(content))
}

func showAnother(a fyne.App) {
	time.Sleep(time.Second * 2)

	win := a.NewWindow("Error Window")
	win.SetContent(widget.NewLabel("App packages not installed properly. Reinstall them and try again later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()

	// time.Sleep(time.Second * 5)
	// win.Close()
}