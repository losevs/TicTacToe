package functions

import (
	"fmt"

	"github.com/fatih/color"
)

type Matr [][]string

var Matrix = Matr{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

func (mat Matr) Prnt() {
	fmt.Println(" - - -")
	for i := 0; i < len(Matrix); i++ {
		for j := 0; j < len(Matrix[i]); j++ {
			if Matrix[i][j] == "X" {
				fmt.Printf("|%s", color.RedString(Matrix[i][j]))
			} else {
				fmt.Printf("|%s", color.BlueString(Matrix[i][j]))
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
		fmt.Println("X won!")
		return false
	case mat[0][0] == mat[0][1] && mat[0][1] == mat[0][2] && mat[0][2] == "O":
		fmt.Println("O won!")
		return false
	case mat[1][0] == mat[1][1] && mat[1][1] == mat[1][2] && mat[1][2] == "X":
		fmt.Println("X won!")
		return false
	case mat[1][0] == mat[1][1] && mat[1][1] == mat[1][2] && mat[1][2] == "O":
		fmt.Println("O won!")
		return false
	case mat[2][0] == mat[2][1] && mat[2][1] == mat[2][2] && mat[2][2] == "X":
		fmt.Println("X won!")
		return false
	case mat[2][0] == mat[2][1] && mat[2][1] == mat[2][2] && mat[2][2] == "O":
		fmt.Println("O won!")
		return false
		//vertical
	case mat[0][0] == mat[1][0] && mat[1][0] == mat[2][0] && mat[2][0] == "X":
		fmt.Println("X won!")
		return false
	case mat[0][0] == mat[1][0] && mat[1][0] == mat[2][0] && mat[2][0] == "O":
		fmt.Println("O won!")
		return false
	case mat[0][1] == mat[1][1] && mat[1][1] == mat[2][1] && mat[2][1] == "X":
		fmt.Println("X won!")
		return false
	case mat[0][1] == mat[1][1] && mat[1][1] == mat[2][1] && mat[2][1] == "O":
		fmt.Println("O won!")
		return false
	case mat[0][2] == mat[1][2] && mat[1][2] == mat[2][2] && mat[2][2] == "X":
		fmt.Println("X won!")
		return false
	case mat[0][2] == mat[1][2] && mat[1][2] == mat[2][2] && mat[2][2] == "O":
		fmt.Println("O won!")
		return false
		//diagonal
	case mat[0][0] == mat[1][1] && mat[1][1] == mat[2][2] && mat[2][2] == "X":
		fmt.Println("X won!")
		return false
	case mat[0][0] == mat[1][1] && mat[1][1] == mat[2][2] && mat[2][2] == "O":
		fmt.Println("O won!")
		return false
	case mat[0][2] == mat[1][1] && mat[1][1] == mat[2][0] && mat[2][0] == "X":
		fmt.Println("X won!")
		return false
	case mat[0][2] == mat[1][1] && mat[1][1] == mat[2][0] && mat[2][0] == "O":
		fmt.Println("O won!")
		return false
	}
	return true
}

func (mat Matr) IsEmpty(i, j int) bool {
	return mat[i][j] == " "
}

func (mat Matr) IsFull() bool {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == " " {
				return true
			}
		}
	}
	return false
}
