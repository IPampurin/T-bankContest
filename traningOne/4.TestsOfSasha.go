/*
Во время разработки некоторой задачи Саша решил сгенерировать несколько новых тестов. Каждый тест Саши должен представлять собой натуральное число,
не меньшее l и не большее r. Кроме того, натуральное число в тесте обязательно должно состоять из одинаковых цифр.
Например, число  подходит под это требование, а число  — нет.
Какое максимальное число различных тестов сможет создать Саша?

Формат входных данных
В единственной строке вводятся два натуральных числа l,r(1<=l,r<=10^18) — ограничения на тесты Саши.

Обратите внимания, что эти числа не вместятся в 32-битный тип данных, используйте 64-битный при необходимости

Формат выходных данных
Выведите одно число — количество тестов, которое может сделать Саша.

Замечание
В первом тесте Саше подходят номера [4,5,6,7].
Во втором тесте подходят все числа, кратные 11, от 11 до 99.

Примеры данных:
Пример 1
Ввод
4 7
Вывод
4

Пример 2
Ввод
10 100
Вывод
9

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	A = 1  // минимально возможное число разрядов
	B = 18 // максимально возможное число разрядов
)

// inputing считывает данные и возвращает введённые значения
func inputing(sc *bufio.Scanner) (uint64, uint64) {

	// считываем ввод
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	// парсим входные данные
	// левый предел
	leftBorder, err := strconv.Atoi(input[0])
	// правый предел
	rightBorder, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}

	return uint64(leftBorder), uint64(rightBorder)
}

// generateOnceDigitNumbers генерирует все числа с одинаковыми цифрами
func generateOnceDigitNumbers() []uint64 {

	numbers := make([]uint64, 0) // слайс из чисел с одинаковыми цифрами

	// перебираем длину числа от 1 до 18
	for length := A; length <= B; length++ {
		// создаем число из единиц нужной длины
		ones := uint64(0)
		for i := 0; i < length; i++ {
			ones = ones*10 + 1
		}
		// умножаем на цифры от 1 до 9
		for digit := 1; digit <= 9; digit++ {
			numbers = append(numbers, uint64(digit)*ones)
		}
	}

	return numbers
}

// calculation вычисления количества нужных чисел
func calculation(left, right uint64) int {

	allNumbers := generateOnceDigitNumbers()

	result := 0 // счётчик количества нужных чисел

	// фильтруем числа по заданному диапазону
	for _, num := range allNumbers {
		if num >= left && num <= right {
			result++
		}
	}

	return result
}

// outputing выводит результат
func outputing(out *bufio.Writer, x int) {

	fmt.Fprintf(out, "%v", x)
}

func main() {

	// определяем ввод
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// определяем вывод
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// считываем введенные данные
	leftBorder, rightBorder := inputing(scanner)

	// вычисляем количество
	x := calculation(leftBorder, rightBorder)

	// выводим результат
	outputing(out, x)
}
