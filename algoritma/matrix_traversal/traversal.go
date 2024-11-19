package matrixtraversal

/*
Flood Fill ( Leetcode #733 )

You are given an image represented by an m x n grid of integers image, where image[i][j] represents the pixel value of the image.
You are also given three integers sr, sc, and color. Your task is to perform a flood fill on the image starting from the pixel image[sr][sc].

To perform a flood fill:
1. Begin with the starting pixel and change its color to color.
2. Perform the same process for each pixel that is directly adjacent (pixels that share a side with the original pixel, either horizontally or vertically) and shares the same color as the starting pixel.
3. Keep repeating this process by checking neighboring pixels of the updated pixels and modifying their color if it matches the original color of the starting pixel.
4. The process stops when there are no more adjacent pixels of the original color to update.

Return the modified image after performing the flood fill.

Example 1:
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]

Example 2:
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
Output: [[0,0,0],[0,0,0]]
*/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	dfs(image, sr, sc, image[sr][sc], newColor)
	return image
}

func dfs(image [][]int, sr int, sc int, sourceColor int, newColor int) {

	if sr < 0 || sc < 0 || sr >= len(image) || sc >= len(image[0]) || image[sr][sc] != sourceColor {
		return
	}

	image[sr][sc] = newColor
	dfs(image, sr-1, sc, sourceColor, newColor)
	dfs(image, sr+1, sc, sourceColor, newColor)
	dfs(image, sr, sc-1, sourceColor, newColor)
	dfs(image, sr, sc+1, sourceColor, newColor)
}
