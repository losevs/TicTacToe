package main

import (
	"fmt"
	"tictac/functions"

	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

var sign string

func main() {
endless:
	for {
		functions.Matrix = functions.Matr{{"7", "8", "9"}, {"4", "5", "6"}, {"1", "2", "3"}}
		count := 1
		place := 0
		for {
			screen.Clear()
			screen.MoveTopLeft()
			functions.Matrix.Prnt()
			if !functions.Matrix.IsOver() {
				functions.PrintScore()
				break
			}
			if !functions.Matrix.IsFull() {
				fmt.Println(color.GreenString("Draw!"))
				functions.PrintScore()
				break
			}
			if count%2 != 0 {
				sign = "X"
			} else {
				sign = "O"
			}
			if sign == "X" {
				fmt.Printf("\nWrite 1-9: %s's move\n", color.RedString(sign))
			} else {
				fmt.Printf("\nWrite 1-9: %s's move\n", color.BlueString(sign))
			}
			check := false
			x, y := -1, -1
			for !check {
				_, err := fmt.Scan(&place)
				switch place {
				case 1:
					x, y = 2, 0
				case 2:
					x, y = 2, 1
				case 3:
					x, y = 2, 2
				case 4:
					x, y = 1, 0
				case 5:
					x, y = 1, 1
				case 6:
					x, y = 1, 2
				case 7:
					x, y = 0, 0
				case 8:
					x, y = 0, 1
				case 9:
					x, y = 0, 2
				default:
					fmt.Println("\nTry again:")
					continue
				}
				if place <= 9 && place > 0 && err == nil && functions.Matrix.IsEmpty(x, y) {
					check = true
				} else {
					fmt.Println("\nTry again:")
				}
			}
			functions.Matrix[x][y] = sign
			//functions.Matrix.Prnt()
			count++
		}
		fmt.Println("New Game?\nYes/No")
		answer := ""
		fmt.Scan(&answer)
		switch answer {
		case "Yes":
		case ".":
		default:
			break endless
		}
	}
}
