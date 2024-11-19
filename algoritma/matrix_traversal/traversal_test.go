package matrixtraversal

import (
	"fmt"
	"testing"
)

func TestMatrix(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	sr, sc := 1, 1
	color := 2

	fmt.Printf("Input image : %v\n", image)

	res := floodFill(image, sr, sc, color)
	fmt.Printf("Output image: %v\n", res)
	fmt.Printf("Answer      : %v\n", [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	})
}
