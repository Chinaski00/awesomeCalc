package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// isArabicNumber проверяет, является ли строка арабским числом (от 1 до 10)
func isArabicNumber(number string) bool {
	num, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	return num >= 1 && num <= 10
}

// isRomanNumber проверяет, является ли строка римским числом (от I до X)
func isRomanNumber(number string) bool {
	romanNumerals := map[string]bool{
		"I":    true,
		"II":   true,
		"III":  true,
		"IV":   true,
		"V":    true,
		"VI":   true,
		"VII":  true,
		"VIII": true,
		"IX":   true,
		"X":    true,
	}
	return romanNumerals[number]
}

// arabicToRoman преобразует арабское число в римское
func arabicToRoman(num int) string {
	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder
	for _, roman := range romanNumerals {
		for num >= roman.Value {
			result.WriteString(roman.Symbol)
			num -= roman.Value
		}
	}
	return result.String()
}

// romanToArabic преобразует римское число в арабское
func romanToArabic(roman string) (int, error) {
	romanNumerals := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	convertedNumber, exists := romanNumerals[roman]
	if !exists {
		return 0, fmt.Errorf("недопустимое римское число")
	}
	return convertedNumber, nil
}

// calculate выполняет операцию между двумя числами
func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("недопустимая арифметическая операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разделение строки на операнды и оператор
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		fmt.Println("Ошибка: неверный формат математической операции.")
		return
	}

	aStr, operator, bStr := tokens[0], tokens[1], tokens[2]

	// Проверка операндов на валидность
	var a, b int
	var err error
	if isArabicNumber(aStr) && isArabicNumber(bStr) {
		a, err = strconv.Atoi(aStr)
		if err != nil {
			fmt.Println("Ошибка: недопустимый операнд.")
			return
		}
		b, err = strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("Ошибка: недопустимый операнд.")
			return
		}
	} else if isRomanNumber(aStr) && isRomanNumber(bStr) {
		a, err = romanToArabic(aStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		b, err = romanToArabic(bStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
	} else {
		fmt.Println("Ошибка: используются разные системы счисления.")
		return
	}

	// Выполнение операции
	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	// Вывод результата
	if isArabicNumber(aStr) && isArabicNumber(bStr) {
		fmt.Println(result)
	} else {
		fmt.Println(arabicToRoman(result))
	}
}
