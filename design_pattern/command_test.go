package design_pattern

import (
	"fmt"
	"testing"
)

// 命令模式本质是把某个对象的方法调用封装到对象中，方便传递、存储、调用。
//
// 示例中把主板单中的启动(start)方法和重启(reboot)方法封装为命令对象，再传递到主机(box)对象中。于两个按钮进行绑定：
//
// - 第一个机箱(box1)设置按钮1(buttion1) 为开机按钮2(buttion2)为重启。
// - 第二个机箱(box1)设置按钮2(buttion2) 为开机按钮1(buttion1)为重启。

// 从而得到配置灵活性。
//
// 除了配置灵活外，使用命令模式还可以用作：
//
// - 批处理
// - 任务队列
// - undo, redo
// 等把具体命令封装到对象中使用的场合

type MotherBoard struct {

}

func (*MotherBoard) Start() {
	fmt.Println("system starting....")
}

func (*MotherBoard) Reboot() {
	fmt.Println("system rebooting....")
}

func (*MotherBoard) ShutDown() {
	fmt.Println("system shutting down....")
}

type Command interface {
	Execute()
}

type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{
		mb : mb,
	}
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{
		mb:mb,
	}
}

func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

type ShutDownCommand struct {
	mb *MotherBoard
}

func NewShutdownCommand(mb *MotherBoard) *ShutDownCommand {
	return &ShutDownCommand{
		mb:mb,
	}
}

func (c *ShutDownCommand) Execute() {
	c.mb.ShutDown()
}

type Box struct {
	btn1 Command
	btn2 Command
}

func NewBox(btn1, btn2 Command) *Box {
	return &Box{
		btn1: btn1,
		btn2: btn2,
	}
}

func (b *Box) PressBtn1() {
	b.btn1.Execute()
}

func (b *Box) PressBtn2() {
	b.btn2.Execute()
}

func TestCommand(t *testing.T) {
	mb := &MotherBoard{}
	startCommand    := NewStartCommand(mb)
	rebootCommand   := NewRebootCommand(mb)
	shutDownCommand := NewShutdownCommand(mb)

	box1 := NewBox(startCommand, shutDownCommand)
	box2 := NewBox(rebootCommand, startCommand)

	box1.PressBtn1()
	box1.PressBtn2()
	box2.PressBtn1()
	box2.PressBtn2()
}