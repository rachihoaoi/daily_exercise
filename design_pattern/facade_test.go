package design_pattern

import (
	"fmt"
	"testing"
)

func NewSumApi() SumApi{
	return &SumApiImpl{
		M1: NewMethodA(),
		M2: NewMethodB(),
	}
}

type SumApi interface {
	GetInfo() string
}

type SumApiImpl struct {
	M1 MethodA
	M2 MethodB
}

func (s *SumApiImpl) GetInfo() string {
	return s.M1.DoA() + "," + s.M2.DoB()
}



type MethodA interface {
	DoA() string
}

type MethodAImpl struct {}

func NewMethodA() MethodA {
	return &MethodAImpl{}
}

func (m *MethodAImpl) DoA() string {
	return "Doing Method A"
}



type MethodB interface {
	DoB() string
}

func NewMethodB() MethodB {
	return &MethodBImpl{}
}

type MethodBImpl struct {}

func (m *MethodBImpl) DoB() string {
	return "Doing Method B"
}

func TestFacade(t *testing.T) {
	fmt.Println(1)
	sumApi := NewSumApi()
	fmt.Println(sumApi.GetInfo())
}