package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"
)

func PrintScene(knots []*Knot) {
	xOffset := 11
	yOffset := 15

	gridSizeX := 26
	gridSizeY := 21

	lines := []string{}
	for i := 0; i < gridSizeY; i++ {
		lines = append(lines, strings.Repeat(".", gridSizeX))
	}

	for i := len(knots) - 1; i >= 0; i-- {
		knot := knots[i]
		knotName := fmt.Sprintf("%d", i)
		if i == 0 {
			knotName = "H"
		} else if i == len(knots)-1 {
			knotName = "T"
		}

		x := knot.X + xOffset
		y := knot.Y + yOffset
		line := lines[y]

		line = line[:x] + knotName + line[x+1:]
		lines[y] = line
	}

	if lines[yOffset][xOffset] == '.' {
		lines[yOffset] = lines[yOffset][:xOffset] + "s" + lines[yOffset][xOffset+1:]
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Printf("\n\n")
}

var frameNumber = 0

func GenerateFrame(knots []*Knot, visited map[string]bool) {
	minX := -258
	minY := -356
	maxX := 188
	maxY := 94

	xOffset := -minX
	yOffset := -minY

	width := maxX - minX
	height := maxY - minY

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if _, ok := visited[fmt.Sprintf("%d,%d", x-xOffset, y-yOffset)]; ok {
				img.Set(x, y, color.RGBA{192, 192, 192, 255}) // visited
			} else {
				img.Set(x, y, color.RGBA{255, 255, 255, 255}) // background
			}
		}
	}

	// set start pos to black
	img.Set(xOffset, yOffset, color.RGBA{0, 0, 0, 255})

	for i := 0; i < len(knots); i++ {
		c := color.RGBA{0, 0, 255, 255}
		if i == 0 {
			c = color.RGBA{255, 0, 0, 255}
		}
		if i == len(knots)-1 {
			c = color.RGBA{0, 255, 0, 255}
		}

		k := knots[i]
		img.Set(k.X+xOffset, k.Y+yOffset, c)
	}

	file, err := os.Create(fmt.Sprintf("images/frame%04d.png", frameNumber))
	if err != nil {
		panic(err)
	}

	frameNumber += 1

	png.Encode(file, img)
	file.Close()
}
