package sudoku

//solves sudoku board using a Sudoku struct
func Solve(sudoku Sudoku) Sudoku {
	//finds coordinates of empty box
	xCoord, yCoord := findEmptyBox(&sudoku)
	//if not empty coordinates, it solved sudoku.
	if xCoord == -1 && yCoord == -1 {
		return sudoku
	}
	//for 1-9, check if that is a valid move
	for k := 1; k <= 9; k++ {
		if validMove(&sudoku, xCoord, yCoord, k) {
			sudoku.Board[xCoord][yCoord] = k
			solved := Solve(sudoku)
			//if board is solved, return solution
			if isSolved(&solved) {
				return solved
			}
		}
	}
	//if not valid move, reset coordinate and backtrack
	sudoku.Board[xCoord][yCoord] = 0
	return sudoku
}

//checks if value is a valid move
func validMove(sudoku *Sudoku, x int, y int, val int) bool {
	//if row is valid
	for index := range sudoku.Board[x] {
		if sudoku.Board[x][index] == val && index != y {
			return false
		}
	}
	//if column is valid
	for index := range sudoku.Board[:][y] {
		if sudoku.Board[index][y] == val && index != x {
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
	for i := sudoku.Boxes[boxIndex][2]; i <= sudoku.Boxes[boxIndex][3]; i++ {
		for j := sudoku.Boxes[boxIndex][0]; j <= sudoku.Boxes[boxIndex][1]; j++ {
			//if another value in box that is the same as the chosen value, return false
			if sudoku.Board[i][j] == val && i != x && y != j {
				return false
			}
		}

	}
	return true
}

//finds box index
func findBox(sudoku *Sudoku, x int, y int) int {
	boxes := sudoku.Boxes
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
func findEmptyBox(sudoku *Sudoku) (int, int) {
	//for each row
	for x, row := range sudoku.Board {
		//for each column
		for y, box := range row {
			//if box unsolved, return that
			if box == 0 {
				return x, y
			}
		}
	}
	//else return -1,-1 if solved
	return -1, -1
}

//checks if sudoku is solved.
func isSolved(sudoku *Sudoku) bool {
	//for each row
	for row := 0; row < len(sudoku.Board); row++ {
		//for each column
		for col := 0; col < len(sudoku.Board[row]); col++ {
			if sudoku.Board[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

//Sudoku struct containing board and box locations
type Sudoku struct {
	Board [9][9]int
	Boxes [9][4]int
}
