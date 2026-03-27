package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strings"
)

// ============================================================================
// Example 1: ComplexImage.go
// ============================================================================

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Brightness(x, y float64) uint8 {
	var dx, dy float64 = c.X - x, c.Y - y
	d := math.Sqrt(dx*dx+dy*dy) / c.R
	if d > 1 {
		return 0
	} else {
		// to blur:=   return 255 to uint8((1 - math.Pow(d, 5)) * 255)
		return 255
		//return uint8((1 - math.Pow(d, 5)) * 255)
	}
}

func example1_compleximage() {
	var w, h int = 280, 240
	var hw, hh float64 = float64(w / 2), float64(h / 2)
	r := 40.0
	θ := 2 * math.Pi / 3
	cr := &Circle{hw - r*math.Sin(0), hh - r*math.Cos(0), 60}
	cg := &Circle{hw - r*math.Sin(θ), hh - r*math.Cos(θ), 60}
	cb := &Circle{hw - r*math.Sin(-θ), hh - r*math.Cos(-θ), 60}

	m := image.NewRGBA(image.Rect(0, 0, w, h))
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			c := color.RGBA{
				cr.Brightness(float64(x), float64(y)),
				cg.Brightness(float64(x), float64(y)),
				cb.Brightness(float64(x), float64(y)),
				255,
			}
			m.Set(x, y, c)
		}
	}

	f, err := os.OpenFile("C:/Users/vupadhya/Documents/rgb.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	png.Encode(f, m)
}

// ============================================================================
// Example 2: ImagePkgUse.go
// ============================================================================

func example2_imagepkguse() {
	fmt.Println("")
	i := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(i)
	fmt.Println(i.ColorModel())
	fmt.Println(i.Bounds())
}

// ============================================================================
// Example 3: SimpleImage.go
// ============================================================================

func example3_simpleimage() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	img.Set(2, 3, color.RGBA{255, 0, 0, 255})
	f, _ := os.OpenFile("C:/Users/vupadhya/Documents/dot.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
	fmt.Println("Completed!!")
}

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 3 examples from Images")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_compleximage ---")
	example1_compleximage()

	fmt.Println("\n--- example2_imagepkguse ---")
	example2_imagepkguse()

	fmt.Println("\n--- example3_simpleimage ---")
	example3_simpleimage()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")
}
