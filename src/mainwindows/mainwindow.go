package main

import (
	. "widgets"

	"github.com/therecipe/qt/widgets"
)

type MainWindow struct {
	*widgets.QWidget
	title *Title
}

func NewMainWindow() *MainWindow {
	mainWindow := &MainWindow{}
	mainWindow.QWidget = widgets.NewQWidget(nil, 0)
	mainWindow.title = NewTitle()

	mainWindow.SetWindowTitle("起重机监控中心")
	mainWindow.SetObjectName("MainWindow")
	return mainWindow
}
