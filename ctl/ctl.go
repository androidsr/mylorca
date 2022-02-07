package ctl

import (
	"fmt"
	"mylorca/view"
	"reflect"
)

type Ctl struct {
}

func (my *Ctl) bind(c interface{}) {
	tType := reflect.TypeOf(c)
	vType := reflect.ValueOf(c)
	for i := 0; i < tType.NumMethod(); i++ {
		name := tType.Method(i).Name
		m := vType.MethodByName(name)
		fmt.Println("bind method: ", name)
		view.UI.Bind(name, func() {
			m.Call(nil)
		})
	}
}
