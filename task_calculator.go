package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Мапы для преобразования римских чисел в арабские и обратно
var romanToArabic = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicToRoman = map[int]string{
	100: "C", 90: "XC", 50: "L", 40: "XL",
	30: "XXX", 20: "XX", 10: "X", 9: "IX",
	5: "V", 4: "IV", 1: "I",
}

// Преобразование арабских чисел в римские
func intToRoman(num int) string {
	// arabicToRoman может преобразовывать числа от 1 до 100
	var roman string
	// Создаем срез ключей, для их сортировки
	keys := make([]int, 0, len(arabicToRoman))
	for k := range arabicToRoman {
		keys = append(keys, k)
	}
	// Сортируем ключи в обратном порядке для правильного построения римского числа
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, value := range keys {
		for num >= value {
			roman += arabicToRoman[value]
			num -= value
		}
	}
	return roman
}

// Выполнение арифметической операции
func calculate(a int, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

// Проверка, является ли строка римским числом
func isRoman(nmb string) bool {
	_, exists := romanToArabic[nmb]
	return exists
}

func main() {
	for {
		// Получаем пользовательское выражение
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение ('stop' для выхода):")

		// Считываем строку до символа новой строки '\n'
		line, _ := reader.ReadString('\n')

		// Удаляем символ новой строки и пробелы по краям строки
		expr := strings.TrimSpace(line)

		if expr == "stop" {
			break
		}

		// Найти индекс оператора
		index := strings.IndexAny(expr, "+/*-")
		if index == -1 {
			// Если оператор не найден, вернём ошибку
			panic("Выдача паники, так как строка не является математической операцией.")
		}

		// Разделяем строку на три части: левое число, оператор, правое число
		aStr := strings.TrimSpace(expr[:index])
		operator := string(expr[index])
		bStr := strings.TrimSpace(expr[index+1:])

		var a, b int
		var err error
		isRomanInput := false

		// Определяем тип ввода (римские или арабские числа)
		if isRoman(aStr) && isRoman(bStr) {
			isRomanInput = true
			a = romanToArabic[aStr]
			b = romanToArabic[bStr]
		} else if !isRoman(aStr) && !isRoman(bStr) {
			a, err = strconv.Atoi(aStr)
			if err != nil {
				panic("Неверный формат числа: " + aStr)
			}
			b, err = strconv.Atoi(bStr)
			if err != nil {
				panic("Неверный формат числа: " + bStr)
			}
		} else {
			panic("Нельзя смешивать римские и арабские цифры")
		}

		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Числа должны быть в пределах от 1 до 10 включительно")
		}

		result := calculate(a, b, operator)

		if isRomanInput {
			if result < 1 {
				panic("Результат меньше единицы не может быть представлен римскими цифрами")
			}
			romanResult := intToRoman(result)

			fmt.Println("Результат:", romanResult)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}
