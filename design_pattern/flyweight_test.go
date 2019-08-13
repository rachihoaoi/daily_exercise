package design_pattern

import "fmt"

// 享元模式(Flyweight Pattern)：
// 运用共享技术有效地支持大量细粒度对象的复用。系统只使用少量的对象，而这些对象都很相似，状态变化很小，可以实现对象的多次复用。
// 由于享元模式要求能够共享的对象必须是细粒度对象，因此它又称为轻量级模式，它是一种对象结构型模式。

type ImageFlyWeight struct {
	data string
}

func (i *ImageFlyWeight) Data() string {
	return i.data
}


type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyWeight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyWeightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyWeight),
		}
	}
	return imageFactory
}

func NewImageFlyWeight(filename string) *ImageFlyWeight {
	// Load file
	data := fmt.Sprint("image data %s", filename)
	return &ImageFlyWeight{
		data: data,
	}
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyWeight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyWeight(filename)
		f.maps[filename] = image
	}
	return image
}

type ImageViewer struct {
	*ImageFlyWeight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyWeightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyWeight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display :%s\n", i.Data())
}
