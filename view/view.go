package view

import (
	"embed"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zserge/lorca"
)

//go:embed www
var fs embed.FS
var UI *View

type View struct {
	lorca.UI
	ln net.Listener
}

func New(width, height int) {
	ui, err := lorca.New("", "", width, height)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
	UI = &View{ui, ln}
}

// func (ui *WebUI) InnerText(id string, val interface{}) {
// 	js := fmt.Sprintf("document.getElementById('%s').innerText='%v'", id, val)
// 	UI.Eval(js)
// }

// func (ui *WebUI) SetValue(id string, val interface{}) {
// 	js := fmt.Sprintf("document.getElementById('%s').value='%v'", id, val)
// 	UI.Eval(js)
// }

// func (ui *WebUI) GetValue(id string) lorca.Value {
// 	js := fmt.Sprintf("document.getElementById('%s').value", id)
// 	return UI.Eval(js)
// }

func (ui *View) Wait() {
	<-UI.Done()
	ui.closeAll()
}

func (ui *View) closeAll() {
	ui.ln.Close()
	UI.Close()
}
