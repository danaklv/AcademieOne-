package functions

import (
	"math"
	"fmt"
)

var BOARD [][]byte

//Минимальный размер поля
func FindBoartMinSize(tetrominoesList []Tetrominoes) int {

	blockCounter := math.Sqrt(float64(len(tetrominoesList) * 4))

	_, frac := math.Modf(blockCounter)
	if frac != 0 {
		blockCounter = math.Floor(blockCounter) + 1
	}
	return int(blockCounter)

}

//Создание пустого поля с минимальным размером
func CreateBoard(size int) {
	BOARD = nil
	for i := 0; i < size; i++ {
		BOARD = append(BOARD, nil)
		for k := 0; k < size; k++ {
			BOARD[i] = append(BOARD[i], 0)
		}
	}
}


//Финальная версия поля
func PrintBoard() {
	for i := 0; i < len(BOARD); i++ {
		for k := 0; k < len(BOARD); k++ {
			if BOARD[i][k] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print(string(BOARD[i][k]))
		}
		fmt.Println()
	}
}
