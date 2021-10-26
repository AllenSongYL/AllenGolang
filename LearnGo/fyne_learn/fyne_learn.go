package main

import (
	"fyne.io/fyne"        // 供所有fyne应用程序代码共用的基础定义，包括数据类型和接口
	"fyne.io/fyne/app"    // 提供创建应用程序的 API
	"fyne.io/fyne/widget" // 窗体控件和交互元素
)

func main() {
	myApp := app.New()

	myWin := myApp.NewWindow("hello world!")
	myWin.SetContent(widget.NewLabel("Hello Fyne!"))
	myWin.Resize(fyne.NewSize(500, 300))
	myWin.ShowAndRun()
}
