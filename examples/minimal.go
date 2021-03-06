package main

import (
	"fmt"
	"github.com/visualfc/go-ui/ui"
)

func main() {	
	fmt.Println(ui.Version())
	ui.Main(func() {
		w := ui.NewWidget()
		w.SetWindowTitle(ui.Version())
		w.SetSizev(300, 200)
		defer w.Close()
		w.Show()
		ui.Run()
	})
}
