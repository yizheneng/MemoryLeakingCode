package widgets

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Title struct {
	*widgets.QWidget
	mainWindows *widgets.QWidget

	mouseFlag bool
	windowPos *core.QPoint
	mousePos  *core.QPoint
}

func NewTitle(mainWindows *widgets.QWidget) *Title {
	title := &Title{}
	title.QWidget = widgets.NewQWidget(nil, 0)
	title.mainWindows = mainWindows
	title.mouseFlag = false
	title.windowPos = nil
	title.mousePos = nil

	title.SetObjectName("Title")

	title.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		title.mouseFlag = true

		if title.windowPos != nil {
			title.windowPos.DestroyQPoint()
		}

		if title.mousePos != nil {
			title.mousePos.DestroyQPoint()
		}

		title.windowPos = mainWindows.Pos()
		title.mousePos = event.GlobalPos()
	})

	title.ConnectMouseReleaseEvent(func(event *gui.QMouseEvent) {
		title.mouseFlag = false
	})

	title.ConnectMouseMoveEvent(func(event *gui.QMouseEvent) {
		if title.mouseFlag {
			mainWindows.Move2(title.windowPos.X()+event.GlobalPos().X()-title.mousePos.X(), title.windowPos.Y()+event.GlobalPos().Y()-title.mousePos.Y())
		}
	})

	return title
}
