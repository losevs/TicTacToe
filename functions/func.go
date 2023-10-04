package functions

import (
	"fmt"
	"unicode"

	"github.com/fatih/color"
)

type Score struct {
	Red  int
	Blue int
}

var WinScore Score

type Matr [][]string

var Matrix = Matr{{"7", "8", "9"}, {"4", "5", "6"}, {"1", "2", "3"}}

func (mat Matr) Prnt() {
	fmt.Println(" - - -")
	for i := 0; i < len(Matrix); i++ {
		for j := 0; j < len(Matrix[i]); j++ {
			if Matrix[i][j] == "X" {
				fmt.Printf("|%s", color.RedString(Matrix[i][j]))
			} else if Matrix[i][j] == "O" {
				fmt.Printf("|%s", color.BlueString(Matrix[i][j]))
			} else {
				fmt.Printf("|%s", color.HiBlackString(Matrix[i][j]))
			}
		}
		fmt.Println("|")
		fmt.Println(" - - -")
	}
}

func (mat Matr) IsOver() bool {
	switch {
	//horizontal
	case mat[0][0] == mat[0][1] && mat[0][1] == mat[0][2] && mat[0][2] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][0] == mat[0][1] && mat[0][1] == mat[0][2] && mat[0][2] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	case mat[1][0] == mat[1][1] && mat[1][1] == mat[1][2] && mat[1][2] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[1][0] == mat[1][1] && mat[1][1] == mat[1][2] && mat[1][2] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	case mat[2][0] == mat[2][1] && mat[2][1] == mat[2][2] && mat[2][2] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[2][0] == mat[2][1] && mat[2][1] == mat[2][2] && mat[2][2] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
		//vertical
	case mat[0][0] == mat[1][0] && mat[1][0] == mat[2][0] && mat[2][0] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][0] == mat[1][0] && mat[1][0] == mat[2][0] && mat[2][0] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	case mat[0][1] == mat[1][1] && mat[1][1] == mat[2][1] && mat[2][1] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][1] == mat[1][1] && mat[1][1] == mat[2][1] && mat[2][1] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	case mat[0][2] == mat[1][2] && mat[1][2] == mat[2][2] && mat[2][2] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][2] == mat[1][2] && mat[1][2] == mat[2][2] && mat[2][2] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
		//diagonal
	case mat[0][0] == mat[1][1] && mat[1][1] == mat[2][2] && mat[2][2] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][0] == mat[1][1] && mat[1][1] == mat[2][2] && mat[2][2] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	case mat[0][2] == mat[1][1] && mat[1][1] == mat[2][0] && mat[2][0] == "X":
		fmt.Println(color.RedString("X won!"))
		WinScore.Red++
		return false
	case mat[0][2] == mat[1][1] && mat[1][1] == mat[2][0] && mat[2][0] == "O":
		fmt.Println(color.BlueString("O won!"))
		WinScore.Blue++
		return false
	}
	return true
}

func (mat Matr) IsEmpty(i, j int) bool {
	return unicode.IsDigit(rune(mat[i][j][0]))
}

func PrintScore() {
	fmt.Println(color.RedString(fmt.Sprintf("%d", WinScore.Red)), "-", color.BlueString(fmt.Sprintf("%d", WinScore.Blue)))
}

func (mat Matr) IsFull() bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if unicode.IsDigit(rune(mat[i][j][0])) {
				return true
			}
		}
	}
	return false
}
