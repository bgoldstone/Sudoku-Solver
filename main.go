package main

import (
	"fmt"
	"strings"
)

//col(y) start,end, row(x) start,end
var boxes = [9][4]int{{0, 2, 0, 2}, {0, 2, 3, 5}, {0, 2, 6, 8}, {3, 5, 0, 2}, {3, 5, 3, 5}, {3, 5, 6, 8}, {6, 8, 0, 2}, {6, 8, 3, 5}, {6, 8, 6, 8}}

func main() {
	//initialize board
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
	//gets solution
	sudoku(board)
}
func sudoku(board [9][9]int) {
	//Prints initial board
	printBoard(board, "Initial Board")
	// creates an instance of a sudoku struct
	sudoku := Sudoku{board, boxes}
	//prints sudoku board after solving
	solvedBoard := solve(sudoku).board
	printBoard(solvedBoard, "Solved Board")
}

//solves sudoku board using a Sudoku struct
func solve(sudoku Sudoku) Sudoku {
	//finds coordinates of empty box
	xCoord, yCoord := findEmptyBox(sudoku)
	//if not empty coordinates, it solved sudoku.
	if xCoord == -1 && yCoord == -1 {
		return sudoku
	}
	//for 1-9, check if that is a valid move
	for k := 1; k <= 9; k++ {
		if validMove(sudoku, xCoord, yCoord, k) {
			sudoku.board[xCoord][yCoord] = k
			solved := solve(sudoku)
			if isSolved(solved) {
				return solved
			} else {
				continue
			}
		}
	}
	//if not valid move, reset coordinate and backtrack
	sudoku.board[xCoord][yCoord] = 0
	return sudoku
}

//checks if value is a valid move
func validMove(sudoku Sudoku, x int, y int, val int) bool {
	//if row is valid
	if val == 0 {
		return true
	}
	for index := range sudoku.board[x] {
		if sudoku.board[x][index] == val && index != y {
			return false
		}
	}
	//if column is valid
	for index := range sudoku.board[:][y] {
		if sudoku.board[index][y] == val && index != x {
			return false
		}
	}
	//gets box index
	boxIndex := findBox(sudoku, x, y)
	//if box index is not valid, return false
	if boxIndex == -1 {
		return false
	}
	//for each value in the box
	for i := sudoku.boxes[boxIndex][2]; i <= sudoku.boxes[boxIndex][3]; i++ {
		for j := sudoku.boxes[boxIndex][0]; j <= sudoku.boxes[boxIndex][1]; j++ {
			if sudoku.board[i][j] == val && i != x && y != j {
				return false
			}
		}

	}
	return true
}

//finds box index
func findBox(sudoku Sudoku, x int, y int) int {
	boxes := sudoku.boxes
	//for each box, check if in range
	for i := 0; i < len(boxes); i++ {
		if boxes[i][0] <= y && y <= boxes[i][1] && boxes[i][2] <= x && x <= boxes[i][3] {
			return i
		}
	}
	//if no box exists, return -1
	return -1
}

//finds next empty box
func findEmptyBox(sudoku Sudoku) (int, int) {
	//for each row
	for x, row := range sudoku.board {
		//for each column
		for y, box := range row {
			if box == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func isSolved(sudoku Sudoku) bool {
	for i := 0; i < len(sudoku.board); i++ {
		for j := 0; j < len(sudoku.board[i]); j++ {
			if sudoku.board[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
func printBoard(board [9][9]int, printString string) {
	fmt.Printf("\n	 %s\n\n", printString)
	for row, rowList := range board {
		if row%3 == 0 && row != 0 {
			fmt.Print(strings.Repeat("- ", 15))
			fmt.Print("\n")

		}
		for col, value := range rowList {
			if col%3 == 0 && col != 0 {
				fmt.Print("|")
			}
			if value == 0 {
				fmt.Printf(" _ ")
			} else {
				fmt.Printf(" %d ", value)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

//Sudoku struct containing board and box locations
type Sudoku struct {
	board [9][9]int
	boxes [9][4]int
}
