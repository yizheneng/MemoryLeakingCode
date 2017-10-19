package main

import (
	. "widgets"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QWidget
	title *Title
}

func NewMainWindow() *MainWindow {
	mainWindow := &MainWindow{}
	mainWindow.QWidget = widgets.NewQWidget(nil, 0)
	mainWindow.title = NewTitle(mainWindow.QWidget)
	mainWindow.SetWindowFlags(core.Qt__FramelessWindowHint)
	mainWindow.SetMinimumSize2(800, 600)

	mainWidget := widgets.NewQWidget(nil, 0)
	mainWidget.SetObjectName("mainWidget")

	layout := widgets.NewQVBoxLayout()
	layout.AddWidget(mainWindow.title, 1, 0)
	layout.AddWidget(mainWidget, 5, 0)

	mainWindow.SetWindowTitle("起重机监控中心")
	mainWindow.SetObjectName("MainWindow")
	mainWindow.SetLayout(layout)
	return mainWindow
}
