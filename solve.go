package main

import (
	"fmt"
	"strings"

	"github.com/bgoldstone/Sudoku-Solver/sudoku"
)

//col(y) start,end, row(x) start,end
var boxes = [9][4]int{
	{0, 2, 0, 2}, {0, 2, 3, 5}, {0, 2, 6, 8},
	{3, 5, 0, 2}, {3, 5, 3, 5}, {3, 5, 6, 8},
	{6, 8, 0, 2}, {6, 8, 3, 5}, {6, 8, 6, 8}}

func main() {
	//initialize board
	var board = [9][9]int{
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
	solveSudoku(board)
}

//prints boards and gets a solution
func solveSudoku(board [9][9]int) {
	//Prints initial board
	printBoard(&board, "Initial Board")
	// creates an instance of a sudoku struct
	sudokuPuzzle := sudoku.Sudoku{Board: board, Boxes: boxes}
	//prints sudoku board after solving
	solvedBoard := sudoku.Solve(sudokuPuzzle).Board
	printBoard(&solvedBoard, "Solved Board")
}

//Prints board
func printBoard(board *[9][9]int, header string) {
	//Header
	fmt.Printf("\n	 %s\n\n", header)
	//for each row
	for row, rowList := range *board {
		//creates box borders
		if row%3 == 0 && row != 0 {
			fmt.Print(strings.Repeat("- ", 15))
			fmt.Print("\n")

		}
		//for each column
		for col, value := range rowList {
			//creates box borders
			if col%3 == 0 && col != 0 {
				fmt.Print("|")
			}
			//if empty value, put '_', else print value
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
