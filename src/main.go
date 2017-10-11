package main

import (
	"log"
	"os"

	_ "yizheneng/MonitoringCenter/others"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

const (
	QSS_FILE_PATH = ":/style/style.css"
)

func main() {
	log.Println("System running!!!")
	app := widgets.NewQApplication(len(os.Args), os.Args)

	/// 主题设置
	log.Println("style path:", QSS_FILE_PATH)
	defaultQssFile := core.NewQFile2(QSS_FILE_PATH)
	if defaultQssFile.Open(core.QIODevice__ReadOnly) {
		app.SetStyleSheet(defaultQssFile.ReadAll().Data())
		defaultQssFile.Close()
	} else {
		log.Println("style file open failed!")
	}

	mainWindow := NewMainWindow()
	mainWindow.Show()

	app.Exec()
}
