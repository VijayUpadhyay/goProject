package main

import (
	"fmt"
	"image"
)

func main()  {
	fmt.Println("")
	i:=image.NewRGBA(image.Rect(0,0,100,100))
	fmt.Println(i)
	fmt.Println(i.ColorModel())
	fmt.Println(i.Bounds())
}
