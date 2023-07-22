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
			return 0, fmt.Errorf("недопустимое римское число")
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
	if arabic <= 0 {
		return "", fmt.Errorf("недопустимое арабское число: должно быть положительным")
	}

	if arabic > 3999 {
		return "", fmt.Errorf("недопустимое арабское число: превышено максимальное значение (3999)")
	}

	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	romanNumber := ""
	for value := range romanNumerals {
		for arabic >= value {
			romanNumber += romanNumerals[value]
			arabic -= value
		}
	}

	return romanNumber, nil
}

// Функция для выполнения операций калькулятора
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

// Функция для проверки, является ли строка числом от 1 до 10.
func isValidNumber(s string) bool {
	number, err := strconv.Atoi(s)
	return err == nil && number >= 1 && number <= 10
}

// Функция для парсинга операнда.
func parseOperand(operand string) (int, error) {
	if isValidNumber(operand) {
		return strconv.Atoi(operand)
	}

	if isRomanNumber(operand) {
		return romanToArabic(operand)
	}

	return 0, fmt.Errorf("недопустимый операнд")
}

// Функция для выполнения арифметических операций.
func calculateExpression(input string) (string, error) {
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		return "", fmt.Errorf("неверный формат математической операции")
	}

	a, err := parseOperand(tokens[0])
	if err != nil {
		return "", err
	}

	operator := tokens[1]

	b, err := parseOperand(tokens[2])
	if err != nil {
		return "", err
	}

	// Проверка, что оба операнда имеют одинаковый тип (арабские или римские числа).
	if (isValidNumber(tokens[0]) && isRomanNumber(tokens[2])) || (isRomanNumber(tokens[0]) && isValidNumber(tokens[2])) {
		return "", fmt.Errorf("используются одновременно разные системы счисления")
	}

	// Выполнение операции
	result, err := calculate(a, b, operator)
	if err != nil {
		return "", err
	}

	// Вывод результата
	if isRomanNumber(tokens[0]) && isRomanNumber(tokens[2]) {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	}

	return strconv.Itoa(result), nil
}

func main() {
	fmt.Println("Калькулятор. Введите математическое выражение (например, 1 + 2) или 'exit' для выхода.")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		result, err := calculateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println("Результат:", result)
	}
}
