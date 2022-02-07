package ctl

import (
	"mylorca/element"
)

type IndexCtl struct {
	Ctl
	input *element.Input
	sp1   *element.Span
}

func NewIndexCtl() *IndexCtl {
	c := &IndexCtl{}
	c.input = &element.Input{Id: "input1"}
	c.sp1 = &element.Span{Id: "sp1"}
	c.bind(c)
	return c
}

func (c *IndexCtl) Hello() {
	c.input.SetValue(c.input.GetValue() + "s")
	c.sp1.SetValue(c.input.GetValue())
}
