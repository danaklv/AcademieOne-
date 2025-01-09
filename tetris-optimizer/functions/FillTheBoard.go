package functions

import "os"

// TryPosition пытается разместить тетромино на доске, используя рекурсию 
func TryPosition(square int, tetrominoesList []Tetrominoes, size int) {
	// Перебираем строки доски
	for y := 0; y < len(BOARD); y++ {
		// Перебираем столбцы доски
		for x := 0; x < len(BOARD); x++ {
			// Проверяем, можно ли разместить тетромино в текущей позиции
			if CheckPosition(y, x, square, tetrominoesList) {
				// Если достигнут конец доски или размещены все тетромино, выводим решение и завершаем программу
				if y == len(BOARD)-1 || square == len(tetrominoesList)-1 {
					PrintBoard()
					os.Exit(0)
				} else {
					// Рекурсивно вызываем TryPosition для следующего тетромино
					TryPosition(square+1, tetrominoesList, size)
				}
				// Очищаем позицию на доске после попытки размещения
				ClearPosition(y, x, square, tetrominoesList)
			}
		}
	}

	// Если square == 0, то значит мы прошли все возможные позиции для первого тетромино и не нашли решение.
	// Увеличиваем размер доски и пытаемся разместить тетромино снова.
	if square == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		TryPosition(0, tetrominoesList, increaseSize)
	}
}

// CheckPosition проверяет, можно ли разместить тетромино в указанной позиции на доске
func CheckPosition(y int, x int, square int, tetrominoesList []Tetrominoes) bool {
	// Проверяем, не выходит ли тетромино за границы доски
	for i := 0; i < len(tetrominoesList[square].form); i++ {
		if len(tetrominoesList[square].form)+y > len(BOARD) || len(tetrominoesList[square].form[i])+x > len(BOARD) {
			return false
		}
	}

	// Перебираем блоки тетромино
	for a := y; a < (len(tetrominoesList[square].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[square].form[a-y]) + x); b++ {
			// Пропускаем пустые блоки
			if tetrominoesList[square].form[a-y][b-x] == 46 {
				continue
			}

			// Проверяем, свободен ли блок на доске для размещения тетромино
			if BOARD[a][b] == 0 {
				// Если свободен, размещаем тетромино
				BOARD[a][b] = byte(65 + square)
			} else {
				// Если блок уже занят, отменяем размещение и возвращаем false
				ClearPosition(y, x, square, tetrominoesList)
				return false
			}
		}
	}

	return true
}

// ClearPosition очищает блоки на доске, занятые тетромино после размещения
func ClearPosition(y int, x int, square int, tetrominoesList []Tetrominoes) {
	// Перебираем блоки тетромино
	for a := y; a < (len(tetrominoesList[square].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[square].form[a-y]) + x); b++ {
			// Пропускаем пустые блоки тетромино
			if tetrominoesList[square].form[a-y][b-x] == 46 {
				continue
			}
			// Если блок был занят тетромино, очищаем его
			if BOARD[a][b] == byte(65+square) {
				BOARD[a][b] = 0
			}
		}
	}
}
