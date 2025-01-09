package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("There should be 2 arguments")
		os.Exit(1)

	}
	args := os.Args[1:]

	inputFile, err := os.ReadFile(os.Args[1])
	if len(inputFile) == 0 {
		fmt.Println("no data")
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	data := splitWords(string(inputFile))

	moddata := modifydata(data)

	dataWithSpace := ""

	for i, word := range moddata {
		if i != 0 {
			if strings.Contains(string(moddata[i-1]), "\r") {
				dataWithSpace += ""
			} else {
				dataWithSpace += " "
			}
		}
		dataWithSpace += word
	}

	FinalRes := ""

	Punctuation := handleSingleQuotes(dataWithSpace)
	temp := ""

	countBrackets := 0
	for i, char := range Punctuation {
		switch char {
		case '(':
			temp += string(char)
			countBrackets++
		case ')':
			countBrackets--
		default:
			if i == len(Punctuation)-1 {
				FinalRes += string(char)
			}
			if countBrackets == 0 && !(i == len(Punctuation)-1 || (Punctuation[i+1] == '(' && char == ' ')) {
				FinalRes += string(char)
			}
		}
	}

	

	fmt.Println(FinalRes)

	
	err = os.WriteFile(args[1], []byte(FinalRes), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func cap(str string) string {
	str = strings.ToLower(str)
	return up(str[:1]) + str[1:]
}

func splitWords(s string) []string {
	var result []string
	var word string
	var temp []rune

	for _, char := range s {
		switch char {
		case ' ', '\n':
			if len(temp) == 0 {

				if word != "" {
					result = append(result, word)
				}
				word = ""
				continue
			}
		case '(':
			temp = append(temp, char)
		case ')':
			if len(temp) > 0 {
				temp = temp[:len(temp)-1]
			} else {
				word += string(char)
				continue
			}
		}

		word += string(char)
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}

func bin(binary string) string {
	decimal := 0
	multiplier := 1
	for i := len(binary) - 1; i >= 0; i-- {
		bit := string(binary[i])
		if bit == "1" {
			decimal += multiplier
		} else if bit != "0" {
			return ""
		}
		multiplier *= 2
	}
	return strconv.Itoa(decimal)
}

func hex(word string) string {
	result := 0
	for _, char := range word {
		value := 0
		switch {
		case char >= '0' && char <= '9':
			value = int(char - '0')
		case char >= 'a' && char <= 'f':
			value = int(char-'a') + 10
		case char >= 'A' && char <= 'F':
			value = int(char-'A') + 10
		default:
			return ""
		}
		result = result*16 + value
	}
	return strconv.Itoa(result)
}

func up(word string) string {
	res := strings.ToUpper(string(word))
	return res
}

func low(word string) string {
	res := strings.ToLower(string(word))
	return res
}

func modifydata(data []string) []string {
	if !(strings.HasPrefix(data[0], ("(cap")) || strings.HasPrefix(data[0], ("(up")) || strings.HasPrefix(data[0], ("(low")) || strings.HasPrefix(data[0], ("(bin")) || strings.HasPrefix(data[0], ("(hex"))) {
		for i, word := range data {

			if word == "(hex)" {

				data[i-1] = hex(data[i-1])
			}
			if word == "(bin)" {

				data[i-1] = bin(data[i-1])
			}

			if strings.HasPrefix(word, "(up") {

				if word[3] == ')' {
					data[i-1] = up(data[i-1])
				} else if word[3] == ',' {
					count := ""
					for j := 5; word[j] >= '0' && word[j] <= '9'; j++ {
						count += string(word[j])
					}

					number, err := strconv.Atoi(count)

					if err != nil {
						fmt.Println(count, " not a number")
					}

					if (len(data)-1)-number < 0 {
						fmt.Println("NUMERIC VALUE IS GREATER THAN WORDS CAN BE CONBERTED")
						os.Exit(1)
					}
					for k := i - number; k < i; k++ {

						if strings.HasPrefix(data[k], ("(cap")) || strings.HasPrefix(data[k], ("(up")) || strings.HasPrefix(data[k], ("(low")) || strings.HasPrefix(data[k], ("(hex)")) || strings.HasPrefix(data[k], ("(bin)")) {
							k = i - (number + 1)
							i--
						}
						data[k] = up(data[k])
					}

				}
			}
			if strings.HasPrefix(word, "(low") {

				if word[4] == ')' {
					data[i-1] = low(data[i-1])
				} else if word[4] == ',' {
					count := ""
					for j := 6; word[j] >= '0' && word[j] <= '9'; j++ {
						count += string(word[j])
					}
					number, err := strconv.Atoi(count)
					if err != nil {
						fmt.Println(count, " not a number")
					}
					if (len(data)-1)-number < 0 {
						fmt.Println("NUMERIC VALUE IS GREATER THAN WORDS CAN BE CONBERTED")
						os.Exit(1)
					}
					for k := i - number; k < i; k++ {
						if strings.HasPrefix(data[k], ("(cap")) || strings.HasPrefix(data[k], ("up")) || strings.HasPrefix(data[k], ("(low")) || strings.HasPrefix(data[k], ("(hex)")) || strings.HasPrefix(data[k], ("(bin)")) {
							k = i - (number + 1)
							i--
						}

						data[k] = low(data[k])
					}
				}
			}
			if strings.HasPrefix(word, "(cap") {

				if word[4] == ')' {
					data[i-1] = cap(data[i-1])
				} else if word[4] == ',' {
					count := ""
					for j := 6; word[j] >= '0' && word[j] <= '9'; j++ {
						count += string(word[j])
					}
					number, err := strconv.Atoi(count)
					if err != nil {
						fmt.Println("error")

					}

					if (len(data)-1)-number < 0 {
						fmt.Println("NUMERIC VALUE IS GREATER THAN WORDS CAN BE CONBERTED")
						os.Exit(1)
					}
					for k := i - number; k < i; k++ {
						if strings.HasPrefix(data[k], ("(cap")) || strings.HasPrefix(data[k], ("(up")) || strings.HasPrefix(data[k], ("(low")) || strings.HasPrefix(data[k], ("(hex)")) || strings.HasPrefix(data[k], ("(bin)")) {
							k = i - (number + 1)
							i--
						}
						data[k] = cap(data[k])
					}
				}

			}
		}
	}
	return data
}


func handleSingleQuotes(text string) string {
	var k, an bool
	ind := 0
	rns := []rune(text)
	ans := []rune{}
	for i := 0; i < len(rns); i++ {
		if i != len(rns)-1 && rns[i] == ' ' && (rns[i+1] == ' ' || strings.ContainsRune(".,!?:;", rns[i+1]) || (k && ((len(ans) != 0 && string(ans[len(ans)-1]) == "'") || string(rns[i+1]) == "'")) || (i != 0 && string(rns[i-1]) == "'" || len(ans) == 0)) {
			continue
		} else if string(rns[i]) == "'" {
			k = !k
		} else if (rns[i] == 'a' || rns[i] == 'A') && ((i == 0 && len(rns) > 1 && rns[i+1] == ' ') || (i != len(rns)-1 && i > 0 && (rns[i-1] == ' ' || string(rns[i-1]) == "'" || strings.ContainsRune(".,!?:;", rns[i-1])) && rns[i+1] == ' ')) {
			ind = i
			an = true
		}
		if i != ind && an && rns[i-1] == ' ' && strings.ContainsRune("aeiouhAEIOUH", rns[i]) {
			ans[len(ans)-1] = 'n'
			ans = append(ans, ' ')
			an = false
		}
		ans = append(ans, rns[i])
		if (strings.ContainsRune(".,!?:;", rns[i]) && i != len(rns)-1 && rns[i+1] != ' ' && !strings.ContainsRune(".,!?:;", rns[i+1])) || (string(rns[i]) == "'" && i != len(rns)-1 && !strings.ContainsRune(".,!?:;", rns[i+1]) && !k) {
			ans = append(ans, ' ')
		}
	}
	return string(ans)
}
