package functions

import (
	"fmt"
	"os"

)

func CheckTetromino(tetrominoesList []Tetrominoes)  {
for _, tetromino := range tetrominoesList {
	found := false
	for _, tetro := range TETROMINO {
		if compareTetrominoes(tetromino.form, tetro) {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Wrong tetromino")
		os.Exit(0)
	}
}

}

func compareTetrominoes(t1, t2 [][]byte) bool {
if len(t1) != len(t2) {
	return false
}
for i := range t1 {
	if len(t1[i]) != len(t2[i]) {
		return false
	}
	for j := range t1[i] {
		if t1[i][j] != t2[i][j] {
			return false
		}
	}
}
return true
}


var TETROMINO [][][]byte = [][][] byte { //19 возможных фигур тетрамино 
	{{35,35,35,35}},
	{{35, 35}, {35, 35}},
	{{35}, {35}, {35}, {35}},
	{{46, 35, 35}, {35, 35, 46}},
	{{35, 35, 35}, {35, 46, 46}},
	{{35, 35, 46}, {46, 35, 35}},
	{{35, 35, 35}, {46, 46, 35}},
	{{35, 46, 46}, {35, 35, 35}},
	{{46, 46, 35}, {35, 35, 35}},
	{{46, 35, 46}, {35, 35, 35}},
	{{35, 35, 35}, {46, 35, 46}},
	{{46, 35}, {35, 35}, {35, 46}},
	{{35, 35}, {46, 35}, {46, 35}},
	{{35, 46}, {35, 35}, {46, 35}},
	{{46, 35}, {46, 35}, {35, 35}},
	{{35, 46}, {35, 35}, {35, 46}},
	{{46, 35}, {35, 35}, {46, 35}},
	{{35, 35}, {35, 46}, {35, 46}},
	{{35, 46}, {35, 46}, {35, 35}},
}

type Tetrominoes struct {
	form [][]byte
}

func FindTetro(content [][]byte) []Tetrominoes { //ищет тетрамино 
	var tetrominoList []Tetrominoes //список тетрамино
	var res [][]byte //одно тетрамино 

	for i := 0; i < len(content); i++ {
		if len(content[i]) == 0 { //конец текущего тетрамино
			if len(res) > 0 {
				tetromino := Tetrominoes{form: processSquare(res)} //обработка тетрамино
				tetrominoList = append(tetrominoList, tetromino) //добавляем в список
			}
			res = nil //обновление
			continue
		}

		res = append(res, content[i]) //каждая строка добавляется в тек тетрамино
	}

	// Проверка последнего тетрамино
	if len(res) > 0 {
		tetromino := Tetrominoes{form: processSquare(res)}
		tetrominoList = append(tetrominoList, tetromino)
	}

	return tetrominoList
}

func processSquare(square [][]byte) [][]byte {
	var res [][]byte 

	// Проверка наличия 35 в каждой строке
	for i := 0; i < len(square); i++ {
		has35 := false
		for j := 0; j < len(square[i]); j++ {
			if square[i][j] == 35 {
				has35 = true
				break
			}
		}
		if has35 {
			res = append(res, square[i])
		}
	}
	//Если нет строк с 35, выходим
	if len(res) == 0 {
		fmt.Println("error format")
		os.Exit(0)
		// return nil
	}
	// Проверка наличия 35 в каждом столбце
	// Массив в который будем добавлять индексы столбцов, содержащих 35
	var columnsWith35 []int
	for j := 0; j < len(res[0]); j++ {
		has35 := false
		for i := 0; i < len(res); i++ {
			if res[i][j] == 35 {
				has35 = true
				break
			}
		}
		if has35 {
			columnsWith35 = append(columnsWith35, j)
		}
	}
	// Создание нового списка, содержащего только столбцы с 35
	var finalRes [][]byte
	for _, row := range res {
		var newRow []byte
		for _, columnIndex := range columnsWith35 {
			newRow = append(newRow, row[columnIndex])
		}
		finalRes = append(finalRes, newRow)
	}

	return finalRes
}


func CheckFormat(transformedContent [][]byte) {
	cubiki := 0
	count := 0
	for i := 0; i < len(transformedContent); i++ {
		count++
		if len(transformedContent[i]) == 0 { //конец тетрамино
			if cubiki != 4 && count != 5 { //если нет 4 # и квадрат не 4 на 4 
				fmt.Println("Wrong format")
				os.Exit(0)
			} else {
				count = 0
				cubiki = 0
			}
		} else if len(transformedContent[i]) != 4 {
			fmt.Println("Wrong format")
			os.Exit(0)
		} else {
			for k := 0; k < len(transformedContent[i]); k++ {
				if transformedContent[i][k] != 35 && transformedContent[i][k] != 46 { //если что то кроме точки и хештега 
					fmt.Println("Wrong format")
					os.Exit(0)
				}
				if transformedContent[i][k] == 35 {
					cubiki++
				}
			}
		}
	}

	if count != 4 && count != 0 {
		fmt.Println("Wrong format")
		os.Exit(0)
	}

	if cubiki != 4 && cubiki != 0 {
		fmt.Println("Wrong format")
		os.Exit(0)
	}
}
