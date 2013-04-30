package main

import (
	"./convert"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func (m Image) ShowImage(s string) {
	toimg, _ := os.Create("./" + s + ".png") //Create the file to which we will save the image
	png.Encode(toimg, m)                     //Encode the image as a png
}

type Image struct{ Width, Height int } //Implement the image.Image interface

func (m Image) Mandelbrot(x, y int) (int, int, int) {
	x0 := float64(x)/float64(m.Width)*(rMax-rMin) + rMin  //Scale the pixel grid to the bounds of the Mandelbrot function
	y0 := float64(y)/float64(m.Height)*(iMax-iMin) + iMin //Likewise for y axis
	a := complex(float64(x0), float64(y0))                //Use complex number x+yi
	z := complex(0, 0)                                    //Initial seed value: 0+0i
	i := 0
	for ; cmplx.Abs(z) < 2 && i < iterations; i++ { //Continue until the absolute value passes 2 (indicating that the number diverges)
		z = z*z + a
	}
	if i == iterations {
		return 0, 0, 0
	}
	return convert.HSVtoRGB(int(math.Abs(float64(i)*360.)/float64(iterations)), 100, 100) //Color the area with hue related to "closeness" to the set, ie numbers the diverge slower have higher hue
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(m.Width), int(m.Height)) //Define the bounds of the image
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel //Use the RGBA color model (package image/color does not have native support for HSV)
}

func (m Image) At(x, y int) color.Color {
	r, b, g := m.Mandelbrot(x, y)
	return color.RGBA{uint8(r), uint8(b), uint8(g), 255} //Convert the int from the set to unit8 (as image.At() wants to return this type)
}

var (
	rMin       = -2.5 //Bounds of the complex plane
	rMax       = 1.
	iMin       = -1.
	iMax       = 1.
	width      int
	height     int
	name       string
	iterations int
)

func init() {

	flag.StringVar(&name, "name", "", "Name of the image")
	flag.IntVar(&width, "size", 1000, "Width of the image")
	flag.IntVar(&iterations, "iterations", 20, "Number of iterations")
	flag.Parse()
}

func main() {
	height = int(float64(width) * (iMax - iMin) / (rMax - rMin))
	m := Image{width, height}
	if name == "" {
		fmt.Println("Please give the image a name!")
		return
	}
	m.ShowImage(name) //Compute and save the file to the first argument
}
