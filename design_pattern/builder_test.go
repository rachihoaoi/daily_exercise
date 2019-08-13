package design_pattern

import (
	"fmt"
	"testing"
)

type Order interface {
	AddMainSet()
	AddDrink()
	AddSnack()
	SumValue() int
}

type Director struct {
	order Order
}

func NewDirector(order Order) *Director {
	return &Director{
		order: order,
	}
}

func (d *Director) Construct() {
	d.order.AddMainSet()
	d.order.AddDrink()
	d.order.AddSnack()
	d.order.SumValue()
}

type food struct {
	Name string
	Value int
}

type Order1 struct {
	MainSet food
	Drink food
	Snack food
}

func (o *Order1) AddMainSet() {
	o.MainSet.Name = "板烧鸡腿堡"
	o.MainSet.Value = 20
}

func (o *Order1) AddDrink() {
	o.Drink.Name = "可乐"
	o.Drink.Value = 10
}

func (o *Order1) AddSnack() {
	o.Snack.Name = "原味鸡"
	o.Snack.Value = 11
}

func (o *Order1) SumValue() int {
	return o.Snack.Value + o.Drink.Value + o.MainSet.Value
}

type Order2 struct {
	MainSet food
	Drink food
	Snack food
}

func (o *Order2) AddMainSet() {
	o.MainSet.Name = "板烧鸡腿堡"
	o.MainSet.Value = 20
}

func (o *Order2) AddDrink() {
	o.Drink.Name = "可乐"
	o.Drink.Value = 10
}

func (o *Order2) AddSnack() {
	o.Snack.Name = "原味鸡"
	o.Snack.Value = 11
}

func (o *Order2) SumValue() int {
	return o.Snack.Value + o.MainSet.Value
}


func TestBuilder (t *testing.T) {
	builder := &Order2{}
	director := NewDirector(builder)
	director.Construct()
	res := builder.SumValue()
	fmt.Println(res)
}