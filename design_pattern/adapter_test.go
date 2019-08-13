package design_pattern

type target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

type AdapteeImpl struct {}


func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee methods"
}

func NewApapter(adaptee Adaptee) target {
	return &adapter {
		Adaptee: adaptee,
	}
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}