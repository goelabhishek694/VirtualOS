// // // // package main

// // // // import (
// // // // 	"fmt"
// // // // 	"fyne.io/fyne/v2/app"
// // // // 	"fyne.io/fyne/v2/widget"
// // // // )

// // // // func main() {
// // // // 	myApp := app.New()
// // // // 	myWindow := myApp.NewWindow("pepOS")
// // // // 	myWindow.SetContent(widget.NewLabel("Hello"))

// // // // 	myWindow.Show()
// // // // 	myApp.Run()
// // // // 	tidyUp()
// // // // }

// // // // func tidyUp() {
// // // // 	fmt.Println("Exited")
// // // // }

// // // package main

// // // import (
// // // 	"time"

// // // 	"fyne.io/fyne/v2"
// // // 	"fyne.io/fyne/v2/app"
// // // 	"fyne.io/fyne/v2/widget"
// // // )

// // // func main() {
// // // 	myApp := app.New()
// // // 	myWindow := myApp.NewWindow("Hello")
// // // 	myWindow.SetContent(widget.NewLabel("Hello"))

// // // 	go showAnother(myApp)
// // // 	myWindow.ShowAndRun()
// // // }

// // // func showAnother(a fyne.App) {
// // // 	time.Sleep(time.Second * 5)

// // // 	win := a.NewWindow("Shown later")
// // // 	win.SetContent(widget.NewLabel("5 seconds later"))
// // // 	win.Resize(fyne.NewSize(200, 200))
// // // 	win.Show()

// // // 	// time.Sleep(time.Second * 2)
// // // 	win.Close()
// // // }

// // package main

// // import (
// // 	"image/color"
// // 	"time"

// // 	"fyne.io/fyne/v2"
// // 	"fyne.io/fyne/v2/app"
// // 	"fyne.io/fyne/v2/canvas"
// // 	"fyne.io/fyne/v2/theme"
// // )

// // func main() {
// // 	myApp := app.New()
// // 	myWindow := myApp.NewWindow("Canvas")
// // 	myCanvas := myWindow.Canvas()

// // 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
// // 	text := canvas.NewText("Text", green)
// // 	text.TextStyle.Bold = true
// // 	myCanvas.SetContent(text)
// // 	go changeContent(myCanvas)

// // 	myWindow.Resize(fyne.NewSize(100, 100))
// // 	myWindow.ShowAndRun()
// // }

// // func changeContent(c fyne.Canvas) {
// // 	time.Sleep(time.Second * 2)

// // 	blue := color.NRGBA{R: 0, G: 0, B: 180, A: 255}
// // 	c.SetContent(canvas.NewRectangle(blue))

// // 	time.Sleep(time.Second * 2)
// // 	c.SetContent(canvas.NewLine(color.Gray{Y: 180}))

// // 	time.Sleep(time.Second * 2)
// // 	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
// // 	circle := canvas.NewCircle(color.White)
// // 	circle.StrokeWidth = 4
// // 	circle.StrokeColor = red
// // 	c.SetContent(circle)

// // 	time.Sleep(time.Second * 2)
// // 	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
// // }

// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	//"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Container")
// 	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
// 	// red := color.NRGBA{R: 255, G: 0, B: 255, A: 255}

// 	text1 := canvas.NewText("Hello", green)
// 	text2 := canvas.NewText("There", green)
// 	text3 := canvas.NewText("ghgh", green)
// 	text4 := canvas.NewText("trrtg", green)
// 	text2.Move(fyne.NewPos(20, 20))
// 	text3.Move(fyne.NewPos(40, 40))
// 	text4.Move(fyne.NewPos(60, 80))
// 	content := container.NewWithoutLayout(text1, text2, text3, text4)
// 	// content := container.New(layout.NewGridLayout(2), text1, text2)

// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Widget")

// 	myWindow.SetContent(widget.NewEntry())
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Box Layout")

// 	text1 := canvas.NewText("Hello", color.White)
// 	text2 := canvas.NewText("There", color.White)
// 	text3 := canvas.NewText("(right)", color.White)
// 	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

// 	text4 := canvas.NewText("centered", color.White)
// 	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
// 	myWindow.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Grid Layout")

// 	text1 := canvas.NewText("1", color.White)
// 	text2 := canvas.NewText("2", color.White)
// 	text3 := canvas.NewText("3", color.White)
// 	grid := container.New(layout.NewGridLayout(3), text1, text2, text3)
// 	myWindow.SetContent(grid)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Border Layout")

// 	top := canvas.NewText("top bar", color.White)
// 	left := canvas.NewText("left", color.White)
// 	middle := canvas.NewText("content", color.White)
// 	content := container.New(layout.NewBorderLayout(top, nil, left, nil),
// 		top, left, middle)
// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"image/color"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/canvas"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/layout"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Form Layout")

// 	label1 := canvas.NewText("Label 1", color.Black)
// 	value1 := canvas.NewText("Value", color.White)
// 	label2 := canvas.NewText("Label 2", color.Black)
// 	value2 := canvas.NewText("Something", color.White)
// 	grid := container.New(layout.NewFormLayout(), label1, value1, label2, value2)
// 	myWindow.SetContent(grid)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	//"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("TabContainer Widget")

// 	tabs := container.NewAppTabs(
// 		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
// 		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
// 	)

// 	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

// 	tabs.SetTabLocation(container.TabLocationLeading)

// 	myWindow.SetContent(tabs)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Entry Widget")

// 	content := container.NewVBox(
// 		widget.NewLabel("The top row of the VBox"),
// 		container.NewHBox(
// 			widget.NewLabel("Label 1"),
// 			widget.NewLabel("Label 2"),
// 		),
// 	)

// 	content.Add(widget.NewButton("Add more items", func() {
// 		content.Add(widget.NewLabel("Added"))
// 	}))

// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"log"
// 	// "fmt"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/widget"
// 	"fyne.io/fyne/v2/theme"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Button Widget")

// 	// content := widget.NewButton("click me", func() {
// 	// 	log.Println("tapped")
// 	// })

// 	content := widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
// 		log.Println("tapped home")
// 	})

// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"log"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Entry Widget")

// 	input := widget.NewEntry()
// 	input.SetPlaceHolder("Enter text...")

// 	content := container.NewVBox(input, widget.NewButton("Save", func() {
// 		log.Println("Content was:", input.Text)
// 	}))

// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"log"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Choice Widgets")

// 	check := widget.NewCheck("Optional", func(value bool) {
// 		log.Println("Check set to", value)
// 	})
// 	radio := widget.NewRadioGroup([]string{"Option 1", "Option 2"}, func(value string) {
// 		log.Println("Radio set to", value)
// 	})
// 	combo := widget.NewSelect([]string{"Option 1", "Option 2"}, func(value string) {
// 		log.Println("Select set to", value)
// 	})

// 	myWindow.SetContent(container.NewVBox(check, radio, combo))
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"time"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("ProgressBar Widget")

// 	progress := widget.NewProgressBar()
// 	infinite := widget.NewProgressBarInfinite()

// 	go func() {
// 		for i := 0.0; i <= 1.0; i += 0.1 {
// 			time.Sleep(time.Millisecond * 250)
// 			progress.SetValue(i)
// 		}
// 	}()

// 	myWindow.SetContent(container.NewVBox(progress, infinite))
// 	myWindow.ShowAndRun()
// }

// package main

// import (
// 	"log"

// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/container"
// 	"fyne.io/fyne/v2/theme"
// 	"fyne.io/fyne/v2/widget"
// )

// func main() {
// 	myApp := app.New()
// 	myWindow := myApp.NewWindow("Toolbar Widget")

// 	toolbar := widget.NewToolbar(
// 		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
// 			log.Println("New document")
// 		}),
// 		widget.NewToolbarSeparator(),
// 		widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
// 		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
// 		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
// 		widget.NewToolbarSpacer(),
// 		widget.NewToolbarAction(theme.HelpIcon(), func() {
// 			log.Println("Display help")
// 		}),
// 	)

// 	content := container.NewBorder(toolbar, nil, nil, nil, widget.NewLabel("Content"))
// 	myWindow.SetContent(content)
// 	myWindow.ShowAndRun()
// }

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("List Widget")

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
