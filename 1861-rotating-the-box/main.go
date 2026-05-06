package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

const (
	cellSize = 80
	border   = 2
)

func main() {
	boxGrid := [][]byte{
		{'.', '#', '.', '*', '#'},
		{'#', '*', '#', '.', '*'},
		{'*', '#', '.', '#', '.'},
	}

	outputGrid := rotateTheBox(boxGrid)

	createGridImage(boxGrid, "input_grid.png")
	createGridImage(outputGrid, "output_grid.png")
}

func createGridImage(grid [][]byte, fileName string) {
	rows := len(grid)
	cols := len(grid[0])

	width := cols * cellSize
	height := rows * cellSize

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Background
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x := c * cellSize
			y := r * cellSize

			drawCell(img, x, y, grid[r][c])
			drawBorder(img, x, y)
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	println("Created:", fileName)
}

func drawCell(img *image.RGBA, x, y int, value byte) {
	cellRect := image.Rect(x, y, x+cellSize, y+cellSize)

	switch value {
	case '#':
		// White background
		draw.Draw(img, cellRect, &image.Uniform{color.White}, image.Point{}, draw.Src)

		// Draw blue diamond/item
		drawDiamond(img, x, y)

	case '*':
		// Gray obstacle
		draw.Draw(img, cellRect, &image.Uniform{color.RGBA{120, 135, 145, 255}}, image.Point{}, draw.Src)

	case '.':
		// Empty cell
		draw.Draw(img, cellRect, &image.Uniform{color.White}, image.Point{}, draw.Src)
	}
}

func drawDiamond(img *image.RGBA, x, y int) {
	centerX := x + cellSize/2
	centerY := y + cellSize/2

	size := 25

	for py := y; py < y+cellSize; py++ {
		for px := x; px < x+cellSize; px++ {
			dx := abs(px - centerX)
			dy := abs(py - centerY)

			if dx+dy <= size {
				img.Set(px, py, color.RGBA{20, 120, 255, 255})
			}

			// small darker border
			if dx+dy >= size-2 && dx+dy <= size {
				img.Set(px, py, color.RGBA{20, 40, 160, 255})
			}
		}
	}
}

func drawBorder(img *image.RGBA, x, y int) {
	borderColor := color.RGBA{90, 90, 90, 255}

	for i := 0; i < cellSize; i++ {
		for b := 0; b < border; b++ {
			img.Set(x+i, y+b, borderColor)
			img.Set(x+i, y+cellSize-1-b, borderColor)
			img.Set(x+b, y+i, borderColor)
			img.Set(x+cellSize-1-b, y+i, borderColor)
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
