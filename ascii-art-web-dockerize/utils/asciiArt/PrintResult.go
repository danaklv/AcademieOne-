package utils

import (
	"net/http"
	"os"

	
)

func PrintResult(arr_char [][]string, text string, file *os.File) {
	for i := 0; i < 8; i++ {
		for _, ch := range text {
			if ch >= 32 && ch <= 127 {
				index := int(ch) - 32 //индекс символа в standard.txt
				if index >= 0 && index < len(arr_char) {
					for j := 0; j < len(arr_char[index][i]); j++ {
						file.WriteString((string(arr_char[index][i][j]))) //записываем в файл cтроку символа
					}
				}
			}
		}
		file.WriteString("\r") //новая строка после вывода целого символа
	}
}


func AsciiArt(w http.ResponseWriter, text string, banner string) bool {

	testFile, err := os.Create("sample.txt")
	if err != nil {
		Status500(w)
		return false
	}
		
	if text == "" || banner == "" {
		testFile.Truncate(0) 
		return true
	}

	StyleFile, err := os.ReadFile(banner + ".txt")
	if err != nil {
		Status500(w)
		return false
	}

	test := CheckFile(string(StyleFile))
	if !test {
		Status500(w)
		return false
	}
	standard := string(StyleFile)[1:]
	if banner == "thinkertoy" {
		standard = string(StyleFile)[2:]
	}

	arr_lines := ModNewLine(text)
	arr_chars := ModFile(standard)
	count := 0
	for _, line := range arr_lines {
		if line != "" {
			PrintResult(arr_chars, line, testFile)
		} else if count == len(arr_lines)-1 {
			break
		} else {
			testFile.WriteString("\n")
			count++
		}
	}
	return true
}
