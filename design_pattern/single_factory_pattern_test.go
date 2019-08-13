package design_pattern

import (
	"fmt"
	"testing"
)

type ApiFactory interface {
	Method1(name string) string
}

func NewFactory(t int) ApiFactory {
	if t == 1 {
		return &HiAPI{}
	} else if t ==2 {
		return &HelloApi{}
	}
	return nil
}

type HiAPI struct {

}

func (h *HiAPI) Method1 (name string) string {
	return "ApiFactory1 : " + name
}

type HelloApi struct {

}

func (h *HelloApi) Method1 (name string) string {
	return "ApiFactory2 : " + name
}

func TestAAA(t *testing.T) {
	api := NewFactory(2)
	fmt.Println(api.Method1("a"))
}
