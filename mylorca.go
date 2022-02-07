package main

import (
	"mylorca/ctl"
	"mylorca/view"
)

func main() {
	view.New(800, 600)
	ctl.NewIndexCtl()

	view.UI.Wait()
}
