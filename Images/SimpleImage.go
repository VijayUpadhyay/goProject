package main

import (
	"image"
	"image/color"
	"os"
	"image/png"
)

func main()  {
	img:= image.NewRGBA(image.Rect(0,0,100,100))
	img.Set(2,3,color.RGBA{255,0,0,255})
	f,_:=os.OpenFile("C:/Users/vupadhya/Documents/dot.png",os.O_WRONLY|os.O_CREATE,0600)
	defer f.Close()
	png.Encode(f,img)
}
