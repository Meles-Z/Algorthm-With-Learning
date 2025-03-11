package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func squareOfNumber(nums []int) []int {
	var result []int
	// traverse on each number and check every node
	for i := range nums {
		result = append(result, nums[i]*nums[i])
	}
	return result
}

func squareRoot(n float64) float64 {
	if n < 0 {
		return -1
	}
	return math.Sqrt(float64(n))
}
func HLine(x1, y, x2 int) {
    for ; x1 <= x2; x1++ {
        img.Set(x1, y, col)
    }
}

// VLine draws a veritcal line
func VLine(x, y1, y2 int) {
    for ; y1 <= y2; y1++ {
        img.Set(x, y1, col)
    }
}
var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
	var col color.Color
// Rect draws a rectangle utilizing HLine() and VLine()
func Rect(x1, y1, x2, y2 int) {
    HLine(x1, y1, x2)
    HLine(x1, y2, x2)
    VLine(x1, y1, y2)
    VLine(x2, y1, y2)
}
func main() {
	fmt.Println("Hello, world!")
	fmt.Println(squareOfNumber([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(squareRoot(8))
	
	
		col = color.RGBA{255, 0, 0, 255} // Red
		HLine(10, 20, 80)
		col = color.RGBA{0, 255, 0, 255} // Green
		Rect(10, 10, 80, 50)
	
		f, err := os.Create("draw.png")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		png.Encode(f, img)
}
