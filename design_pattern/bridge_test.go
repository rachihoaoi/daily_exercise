package design_pattern

import (
	"fmt"
	"testing"
)

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

// sms type

type MessageSMS struct {}

func (m *MessageSMS) Send(text,to string) {
	fmt.Printf("send %v by sms to %v", text, to)
}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

// email type

type MessageEMail struct {}


func (m *MessageEMail) Send(text,to string) {
	fmt.Printf("send %v by email to %v \n", text, to)
}

func ViaEmail() MessageImplementer {
	return &MessageEMail{}
}

type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage)SendMessage(text, to string) {
	m.method.Send(text, to)
}

type AquaMessage struct {
	method MessageImplementer
}

func NewAquaMessage(method MessageImplementer) *AquaMessage {
	return &AquaMessage{
		method: method,
	}
}

func (m *AquaMessage)SendMessage(text,to string) {
	m.method.Send(fmt.Sprintf("[湊あきあ] %s", text), to)
}

func TestAqua(t *testing.T) {
	m := NewAquaMessage(ViaEmail())
	m.method.Send("wo ai ni", "me")

	fmt.Println(4<<(^uint(0)>>63))
	fmt.Println(4<<(^uint8(0)>>7))
	fmt.Println(^uint8(0))




}