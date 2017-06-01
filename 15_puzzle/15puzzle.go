package main

import "fmt"
import "math/rand"
import "time"

const (
	NROWS = 4
	NCOLS = 4
	NSHUFFLES = 100
)

var (
	loRow, loCol int
	cells [NROWS][NCOLS]int
)

type Move int
const (
	MOVE_UP Move = iota
	MOVE_DOWN
	MOVE_LEFT
	MOVE_RIGHT
)

func MoveController() Move {
	var input string
	for {
		fmt.Printf("%s", "Use w/s/a/d : ")
		fmt.Scan(&input)
		switch input {
			case "s":
				return MOVE_UP
			case "w":
				return MOVE_DOWN
			case "a":
				return MOVE_LEFT
			case "d":
				return MOVE_RIGHT
		}
	}
}

func UpdateBoard(move Move) int {
	xAxis := [4]int{ 1, -1, 0, 0 }
	yAxis := [4]int{ 0, 0, -1, 1 }
	i := loRow + xAxis[move]
	j := loCol + yAxis[move]

	if i >= 0 && i < NROWS && j >= 0 && j < NCOLS {
		cells[loRow][loCol] = cells[i][j]
		cells[i][j] = 0
		loRow = i
		loCol = j
		return 1
	}
	return 0
}

func SetupBoard() {
	for i := 0; i < NROWS; i++ {
		for j := 0; j < NCOLS; j++ {
			cells[i][j] =  i * NCOLS + j + 1
		}
	}
	cells[NROWS-1][NCOLS-1] = 0
	loRow = NROWS - 1
	loCol = NCOLS - 1
	k := 0
	for k < NSHUFFLES {
		k += UpdateBoard(Move(rand.Intn(13) % 4))
	}
}

func DrawBoard() {
	var i, j int
	for i = 0; i < NROWS; i++ {
		for j = 0; j < NCOLS; j ++ {
			if cells[i][j] != 0 {
				if j != NCOLS - 1 {
					fmt.Printf(" %2d ", cells[i][j])
				} else {
					fmt.Printf(" %2d \n", cells[i][j])
				}
			} else {
				if j != NCOLS - 1 {
					fmt.Printf(" %2s ", "")
				} else {
					fmt.Printf(" %2s \n", "")
				}
			}
		}
		fmt.Printf("\n")
	}
}

func GameOver() bool {
	var k int = 1
	for i := 0; i < NROWS; i++ {
		for j := 0; j < NCOLS; j++ {
			if k < NROWS * NCOLS && cells[i][j] != k {
				k += 1
				return false
			}
		}
	}
	fmt.Printf("outside")
	return true
}

func main() {
	rand.Seed(time.Now().Unix())
	SetupBoard()
	DrawBoard()
	for !GameOver() {
		UpdateBoard(MoveController())
		DrawBoard()
	}
}