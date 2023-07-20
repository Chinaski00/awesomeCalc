package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Функция для проверки, является ли строка римским числом.
func isRomanNumber(s string) bool {
	return regexp.MustCompile(`^[IVXLCDM]+$`).MatchString(s)
}

// Функция для перевода римского числа в арабское.
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	for _, r := range roman {
		value, ok := romanNumerals[r]
		if !ok {
			return 0, fmt.Errorf("invalid roman numeral")
		}

		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}

		prevValue = value
	}

	return result, nil
}

// Функция для перевода арабского числа в римское.
func arabicToRoman(arabic int) (string, error) {
	if arabic <= 0 || arabic > 10 {
		return "", fmt.Errorf("invalid arabic number: must be in range [1, 10]")
	}

	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	return romanNumerals[arabic], nil
}

// Функция для проверки, является ли число арабским числом от 1 до 10 включительно.
func isValidArabicNumber(number int) bool {
	return number >= 1 && number <= 10
}

// Функция для выполнения операций калькулятора.
func calculate(a, b int, operator string) (int, error) {
	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		result = a / b
	default:
		return 0, fmt.Errorf("недопустимая арифметическая операция")
	}
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разделение строки на операнды и оператор.
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		fmt.Println("Ошибка: неверный формат математической операции.")
		return
	}

	aStr, operator, bStr := tokens[0], tokens[1], tokens[2]

	// Проверка операндов на валидность.
	a, err := strconv.Atoi(aStr)
	if err != nil {
		if isRomanNumber(aStr) {
			// Если первый операнд - римское число, конвертируем его в арабское.
			a, err = romanToArabic(aStr)
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
		} else {
			fmt.Println("Ошибка: недопустимый операнд.")
			return
		}
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		if isRomanNumber(bStr) {
			// Если второй операнд - римское число, конвертируем его в арабское.
			b, err = romanToArabic(bStr)
			if err != nil {
				fmt.Println("Ошибка:", err)
				return
			}
		} else {
			fmt.Println("Ошибка: недопустимый операнд.")
			return
		}
	}

	// Проверка операндов на допустимый диапазон.
	if !isValidArabicNumber(a) || !isValidArabicNumber(b) {
		fmt.Println("Ошибка: числа должны быть от 1 до 10 включительно.")
		return
	}

	// Выполнение операции.
	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод результата.
	if isRomanNumber(aStr) && isRomanNumber(bStr) {
		// Если оба операнда римские числа, конвертируем результат в римское число.
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		fmt.Println(romanResult)
	} else {
		fmt.Println(result)
	}
}
