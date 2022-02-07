package element

import (
	"fmt"
	"mylorca/view"
)

const (
	setValueStr  = "document.getElementById('%s').value='%v'"
	getValueStr  = "document.getElementById('%s').value"
	setInnerText = "document.getElementById('%s').innerText='%v'"
	getInnerText = "document.getElementById('%s').innerText"
)

type Input struct {
	Id    string
	value string
}

func (m *Input) SetValue(value string) {
	js := fmt.Sprintf(setValueStr, m.Id, value)
	view.UI.Eval(js)
}

func (m *Input) GetValue() string {
	js := fmt.Sprintf(getValueStr, m.Id)
	value := view.UI.Eval(js)
	m.value = value.String()
	return m.value
}

type Span struct {
	Id    string
	value string
}

func (m *Span) SetValue(value string) {
	js := fmt.Sprintf(setInnerText, m.Id, value)
	view.UI.Eval(js)
}

func (m *Span) GetValue() string {
	js := fmt.Sprintf(getInnerText, m.Id)
	value := view.UI.Eval(js)
	m.value = value.String()
	return m.value
}
