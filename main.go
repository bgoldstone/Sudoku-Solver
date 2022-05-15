package main

import "fmt"

//col(y) start,end, row(x) start,end
var boxes = [9][4]int{{0, 2, 0, 2}, {0, 2, 3, 5}, {0, 2, 6, 8}, {3, 5, 0, 2}, {3, 5, 3, 5}, {3, 5, 6, 8}, {6, 8, 0, 2}, {6, 8, 3, 5}, {6, 8, 6, 8}}
var visited [][]int

func main() {
	var board [9][9]int
	board = [9][9]int{
		{3, 5, 0, 6, 0, 2, 0, 0, 4},
		{0, 0, 7, 0, 4, 0, 0, 1, 3},
		{0, 6, 9, 8, 3, 1, 0, 0, 7},
		{5, 0, 3, 0, 0, 0, 0, 9, 6},
		{0, 0, 0, 3, 0, 0, 7, 4, 5},
		{9, 4, 6, 0, 0, 0, 8, 0, 0},
		{6, 9, 2, 4, 0, 0, 0, 0, 8},
		{8, 0, 0, 7, 0, 3, 0, 0, 0},
		{0, 0, 4, 0, 2, 0, 0, 0, 1},
	}
	sudoku := Sudoku{board, boxes}
	for _, v := range solve(sudoku).board {
		fmt.Println(v)
	}
}

func solve(sudoku Sudoku) Sudoku {
	coords := findEmptyBox(sudoku)
	xCoord := coords[0]
	yCoord := coords[1]
	if xCoord == -1 && yCoord == -1 {
		return sudoku
	}
	if sudoku.board[xCoord][yCoord] == 0 {
		for k := 1; k <= 9; k++ {
			if checkConditions(sudoku, xCoord, yCoord, k) {
				sudoku.board[xCoord][yCoord] = k
				visited = append(visited, []int{xCoord, yCoord})
				return solve(sudoku)
			}
		}
		if sudoku.board[xCoord][yCoord] == 0 {
			if len(visited) > 2 {
				coord := visited[len(visited)-2]
				sudoku.board[coord[0]][coord[1]] = 0
				visited = visited[:len(visited)-2]
			}
			return sudoku
		}
	}
	findEmpty := findEmptyBox(sudoku)
	if findEmpty[0] != -1 && findEmpty[1] != -1 {
		return solve(sudoku)
	} else {
		return sudoku
	}

}

func checkConditions(sudoku Sudoku, x int, y int, val int) bool {
	for index, _ := range sudoku.board[x] {
		if sudoku.board[x][index] == val && index != y {
			return false
		}
	}
	for index, _ := range sudoku.board[:][y] {
		if sudoku.board[index][y] == val && index != x {
			return false
		}
	}
	boxIndex := findBox(sudoku, x, y)
	if boxIndex == -1 {
		return false
	}
	for i := sudoku.boxes[boxIndex][2]; i <= sudoku.boxes[boxIndex][3]; i++ {
		for j := sudoku.boxes[boxIndex][0]; j <= sudoku.boxes[boxIndex][1]; j++ {
			if sudoku.board[i][j] == val {
				return false
			}
		}

	}
	return true
}

func findBox(sudoku Sudoku, x int, y int) int {
	boxes := sudoku.boxes
	for i := 0; i < len(boxes); i++ {
		if boxes[i][0] <= y && y <= boxes[i][1] && boxes[i][2] <= x && x <= boxes[i][3] {
			return i
		}
	}
	return -1
}
func findEmptyBox(sudoku Sudoku) []int {
	for x, row := range sudoku.board {
		for y, v := range row {
			if v == 0 {
				return []int{x, y}
			}
		}
	}
	return []int{-1, -1}
}

type Sudoku struct {
	board [9][9]int
	boxes [9][4]int
}
