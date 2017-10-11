package main

import (
	"github.com/therecipe/qt/widgets"
)

type Title struct {
	*widgets.QWidget
}

func NewTitle() *Title {
	title := &Title{}
	title.QWidget = widgets.NewQWidget(nil, 0)

	title.SetObjectName("Title")

	return title
}
