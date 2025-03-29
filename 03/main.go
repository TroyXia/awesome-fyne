package main

import (
	"fyne.io/fyne/v2"
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
	w1 := a.NewWindow("Hello World")
	// 窗口内容
	w1.SetContent(widget.NewLabel("Hello World"))

	// 设置窗口为主窗口
	w1.SetMaster()
	w1.Show()

	// 创建第二个窗口
	w2 := a.NewWindow("Larger")
	// 窗口内容
	w2.SetContent(widget.NewButton("Open new", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third"))
		w3.Resize(fyne.NewSize(200, 200))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(400, 400))
	w2.Show()

	// 运行应用
	a.Run()
}
