package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func exit() {
	fmt.Println("bye")
}

func main() {
	// 创建应用
	a := app.New()

	// 应用窗口标题
	window := a.NewWindow("Hello World")

	// 窗口内容
	window.SetContent(widget.NewLabel("Hello fyne!"))

	// 运行应用
	window.ShowAndRun()

	exit()
}
