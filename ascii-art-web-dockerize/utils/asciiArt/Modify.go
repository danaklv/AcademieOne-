package utils

import "strings"

func ModFile(font string) [][]string {
    font = strings.ReplaceAll(font, "\r", "")

    char := strings.Split(font, "\n\n") //разделение файла на символы (каждый символ разделяется двумя "\n")

    arr := make([][]string, len(char)) //разделение символа на строки
    for i, symbols := range char {
        lines := strings.Split(symbols, "\n")
        arr[i] = lines
    }
    return arr
}
func ModNewLine(text string) []string {
	byteString := []byte(text)
	for i := 0; i < len(byteString); i++ {
		if byteString[i] == 10 {
			byteString = append(byteString[:i], append([]byte{'\\', 'n'}, byteString[i+1:]...)...)
			i++
		}
	}
	text = string(byteString)
	arr_text := strings.Split(text, "\\n")
	return arr_text
}
