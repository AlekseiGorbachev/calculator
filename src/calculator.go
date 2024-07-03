package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("--------------------------------")
	fmt.Println("|    Консольный калькулятор    |")
	fmt.Println("|   арабских и римских чисел   |")
	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")
	fmt.Println("Помощь, введите справка или help")
	fmt.Println("Выход, введите выход или exit   ")
	fmt.Println("--------------------------------")
	fmt.Println()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Введите выражение:")

		expression, _ := reader.ReadString('\n')
		expression = strings.TrimSuffix(expression, "\n")
		expression = strings.TrimSuffix(expression, "\r")

		if expression == "справка" || expression == "help" {
			fmt.Println("1. Введите 1-е число от 1 до 10 или от I до X  и нажмите пробел")
			fmt.Println("2. Введите арифметическое действие: +, -, *, / и нажмите пробел")
			fmt.Println("3. Введите 3-е число от 1 до 10 или от I до X и нажмите ввод(enter)")
			fmt.Println("4. Вводить нужно либо все арабские, либо все римские")
			fmt.Println()
			continue
		}

		if expression == "выход" || expression == "exit" {
			fmt.Println("Вы вышли из калькулятора")
			break
		}

		result := calc(expression)

		fmt.Println("Результат:")
		fmt.Println(result)
	}
}

func calc(expression string) string {

	isArabic := strings.ContainsAny(expression, "123456789")
	isRoman := strings.ContainsAny(expression, "IVX")

	if isArabic && isRoman {
		panic("Используются одновременно разные системы счисления.")
	}

	if isArabic {
		return calcArabic(expression)
	}

	if isRoman {
		return calcRoman(expression)
	}
	panic("Операнды не относятся к арабаским и римским цифрам от 1 до 10 и от I до X")
}

func calcArabic(expression string) string {

	divExpression := strings.Fields(expression)

	if len(divExpression) == 1 {
		panic("Cтрока не является математическим выражением")
	}

	if len(divExpression) != 3 {
		panic("Формат выражени не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	num1, _ := strconv.Atoi(divExpression[0])
	num2, _ := strconv.Atoi(divExpression[2])
	operator := divExpression[1]

	if num1 >= 1 && num1 <= 10 && num2 >= 1 && num2 <= 10 {
		switch operator {
		case "+":
			return strconv.Itoa(num1 + num2)
		case "-":
			return strconv.Itoa(num1 - num2)
		case "*":
			return strconv.Itoa(num1 * num2)
		case "/":
			return strconv.Itoa(num1 / num2)
		default:
			panic("Не правильно введенный оператор, можно использовать только  (+, -, /, *)")
		}
	} else {
		panic("Числа выходят за диапазон от 1 до 10")
	}
}

func calcRoman(expression string) string {

	divExpression := strings.Fields(expression)

	if len(divExpression) == 1 {
		panic("Cтрока не является математическим выражением")
	}

	if len(divExpression) != 3 {
		panic("Формат выражени не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
	}

	num1 := romanToArabic(divExpression[0])
	num2 := romanToArabic(divExpression[2])
	operator := divExpression[1]

	if num1 >= 1 && num1 <= 10 && num2 >= 1 && num2 <= 10 {
		switch operator {
		case "+":
			return arabicToRoman(num1 + num2)
		case "-":
			if num1-num2 < 1 {
				panic("В римской системе нет отрицательных чисел")
			} else {
				return arabicToRoman(num1 - num2)
			}
		case "*":
			return arabicToRoman(num1 * num2)
		case "/":
			if num1/num2 < 1 {
				panic("В римской системе нет отрицательных чисел")
			} else {
				return arabicToRoman(num1 / num2)
			}
		default:
			panic("Не правильно введенный оператор, можно использовать только  (+, -, /, *)")
		}

	} else {
		panic("Числа выходят за диапазон от I до X")
	}

}

func romanToArabic(romanNumeral string) int {

	romanNumbers := []string{
		"I", "II", "III", "IV", "V",
		"VI", "VII", "VIII", "IX", "X",
	}

	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}

	result := 0
	prevVal := 0

	if slices.Contains(romanNumbers, romanNumeral) {
		for _, char := range romanNumeral {

			val := romanNumerals[string(char)]

			if val > prevVal {
				result += val - 2*prevVal
			} else {
				result += val
			}

			prevVal = val
		}

	} else {
		panic("Числа выходят за диапазон от I до X")
	}
	return result
}

func arabicToRoman(arabicNumeral int) string {

	arabicNumerals := map[int]string{
		1:   "I",
		2:   "II",
		3:   "III",
		4:   "IV",
		5:   "V",
		6:   "VI",
		7:   "VII",
		8:   "VIII",
		9:   "IX",
		10:  "X",
		11:  "XI",
		12:  "XII",
		13:  "XIII",
		14:  "XIV",
		15:  "XV",
		16:  "XVI",
		17:  "XVII",
		18:  "XVIII",
		19:  "XIX",
		20:  "XX",
		21:  "XXI",
		22:  "XXII",
		23:  "XXIII",
		24:  "XXIV",
		25:  "XXV",
		26:  "XXVI",
		27:  "XXVII",
		28:  "XXVIII",
		29:  "XXIX",
		30:  "XXX",
		31:  "XXXI",
		32:  "XXXII",
		33:  "XXXIII",
		34:  "XXXIV",
		35:  "XXXV",
		36:  "XXXVI",
		37:  "XXXVII",
		38:  "XXXVIII",
		39:  "XXXIX",
		40:  "XL",
		41:  "XLI",
		42:  "XLII",
		43:  "XLIII",
		44:  "XLIV",
		45:  "XLV",
		46:  "XLVI",
		47:  "XLVII",
		48:  "XLVIII",
		49:  "XLIX",
		50:  "L",
		51:  "LI",
		52:  "LII",
		53:  "LIII",
		54:  "LIV",
		55:  "LV",
		56:  "LVI",
		57:  "LVII",
		58:  "LVIII",
		59:  "LIX",
		60:  "LX",
		61:  "LXI",
		62:  "LXII",
		63:  "LXIII",
		64:  "LXIV",
		65:  "LXV",
		66:  "LXVI",
		67:  "LXVII",
		68:  "LXVIII",
		69:  "LXIX",
		70:  "LXX",
		71:  "LXXI",
		72:  "LXXII",
		73:  "LXXIII",
		74:  "LXXIV",
		75:  "LXXV",
		76:  "LXXVI",
		77:  "LXXVII",
		78:  "LXXVIII",
		79:  "LXXIX",
		80:  "LXXX",
		81:  "LXXXI",
		82:  "LXXXII",
		83:  "LXXXIII",
		84:  "LXXXIV",
		85:  "LXXXV",
		86:  "LXXXVI",
		87:  "LXXXVII",
		88:  "LXXXVIII",
		89:  "LXXXIX",
		90:  "XC",
		91:  "XCI",
		92:  "XCII",
		93:  "XCIII",
		94:  "XCIV",
		95:  "XCV",
		96:  "XCVI",
		97:  "XCVII",
		98:  "XCVIII",
		99:  "XCIX",
		100: "C",
	}

	return arabicNumerals[arabicNumeral]

}
