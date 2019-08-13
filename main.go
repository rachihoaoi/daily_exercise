package main

import "fmt"

type order struct {
	OrderName string
	OrderTime string
	OrderDetail string
	Method OrderMethod
}

type order1 struct {
	OrderName string
	OrderTime string
	OrderDetail string
	OrderAssignee string
}

type OrderMethod interface {
	GetOrderName() string
	GetOrderTime() string
	GetOrderDetail() string
}


func (o *order1) GetOrderName () string {
	return o.OrderName
}

func (o *order1) GetOrderTime () string {
	return o.OrderTime
}


func (o *order1) GetOrderDetail () string {
	return o.OrderDetail
}

func main() {
	a := order{}
	b := order1{"a","b","c",""}
	a.Method = &b
	aaa(a)
	fmt.Println(a.Method.GetOrderDetail())
}

func aaa(a order) string {
	return ""
}