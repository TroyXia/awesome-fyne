package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("2006-01-02 15:04:05")
	clock.SetText(formatted)
}

func main() {
	// 创建应用
	a := app.New()

	// 应用窗口标题
	window := a.NewWindow("Hello World")

	// 窗口内容
	clock := widget.NewLabel("")

	// 初始化窗口内容
	updateTime(clock)
	window.SetContent(clock)

	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	// 运行应用
	window.ShowAndRun()
}
