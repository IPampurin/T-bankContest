/*
У Кости есть бумажка, на которой написано n чисел. Также у него есть возможность не больше, чем k раз, взять любое число с бумажки,
после чего закрасить одну из старых цифр, а на ее месте написать новую произвольную цифру.
На какое максимальное значение Костя сможет увеличить сумму всех чисел на листочке?

Формат входных данных
В первой строке входного файла даны два целых числа n,k — количество чисел на бумажке и ограничение на число операций (1 <= n <= 1000, 1 <= k <= 10^4).
Во второй строке записано n чисел ai — числа на бумажке (1 <= ai <= 10^9).

Формат выходных данных
В выходной файл выведите одно число — максимальную разность между конечной и начальной суммой.

Замечание
В первом примере Костя может изменить две единицы на две девятки, в результате чего сумма чисел увеличится на 16.
Во втором примере Костя меняет число 85 на 95.
В третьем примере можно ничего не менять.
Обратите внимание, что ответ может превышать вместимость 32-битного типа данных.

Примеры данных

Пример 1
ввод:
5  2
1  2  1  3  5
вывод:
16

Пример 2
ввод:
3  1
99  5  85
вывод:
10

Пример 3
ввод:
1  10
9999
вывод:
0
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// inputing считывает данные и возвращает введённые значения
func inputing(sc *bufio.Scanner) (int, []int) {

	// считываем ввод с первой строки
	sc.Scan()
	inputOne := strings.Split(sc.Text(), " ")
	// парсим входные данные
	// количество чисел на бумажке
	numbersCount, err := strconv.Atoi(inputOne[0])
	// ограничение на число операций
	limitOperations, err := strconv.Atoi(inputOne[1])
	if err != nil {
		log.Fatal(err)
	}

	// numbers - слайс с числами на бумажке
	numbers := make([]int, 0, numbersCount)

	// считываем ввод со второй строки
	sc.Scan()
	inputTwo := strings.Split(sc.Text(), " ")
	// парсим входные данные и пишем в numbers
	for i := 0; i < len(inputTwo); i++ {
		num, err := strconv.Atoi(inputTwo[i])
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, num)
	}

	return limitOperations, numbers
}

// sumElements вычисляет сумму чисел в слайсе
func sumElements(arr []int) int {

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}

// destroyNumber разбивает число по разрядам
func destroyNumber(num int) []int {

	categories := make([]int, 0) // слайс для хранения цифр по разрядам
	var digit int                // цифра разряда
	for num != 0 {
		digit = num % 10
		categories = append(categories, digit)
		num = num / 10
	}

	// переворачиваем слайс для наглядности
	for i, j := 0, len(categories)-1; i < j; i, j = i+1, j-1 {
		categories[i], categories[j] = categories[j], categories[i]
	}

	return categories
}

// собирает число из цифр по разрядам
func collectNumber(nums []int) int {

	num := 0 // интересующее нас число
	pow := 1 // степень десятки
	// собираем число начиная с хвоста слайса
	for i := len(nums) - 1; i >= 0; i-- {
		num += nums[i] * pow
		pow *= 10
	}

	return num
}

// calculation вычисляет максимальную разность между начальной и конечной суммами
func calculation(limitOperations int, numbers []int) int {

	// вычислим сумму чисел до изменений
	beginSum := sumElements(numbers)
	// проинициализируем сумму после внесения изменений начальным значением суммы элементов
	endSum := beginSum
	// максимальная разница между суммами
	deltaMax := endSum - beginSum
	// в different будем хранить набор всех deltaMax (разницу сумм после изменения и до изменения)
	different := make([]int, 0)
	different = append(different, deltaMax)

	for i := 0; i < len(numbers); i++ {

		// убираем из конечной суммы очередное число перед изменениями
		endSum -= numbers[i]
		// получаем разбитое на цифры очередное число
		digits := destroyNumber(numbers[i])
		// currentDigits переменная для сохранения текущей цифры
		var currentDigit, newNumber int
		// меняем каждую цифру числа и подсчитываем изменение суммы
		for j := 0; j < len(digits); j++ {
			// если цифра разряда максимальна, то менять нечего
			if digits[j] != 9 {
				currentDigit = digits[j]                // сохраняем цифру для возврата после проверки замены
				digits[j] = 9                           // заменяем цифру на максимально возможную
				newNumber = collectNumber(digits)       // собираем число с заменой
				endSum += newNumber                     // вычисляем новое значение суммы
				deltaMax = endSum - beginSum            // вычисляем разность сумм до и после изменнения
				different = append(different, deltaMax) // добавляем результат в хранилище разностей
				digits[j] = currentDigit                // возвращаем цифру после проверки замены
			}
		}
	}

	return slices.Max(different)
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
	limitOperations, numbers := inputing(scanner)

	// вычисляем максимальную разницу
	x := calculation(limitOperations, numbers)

	// выводим результат
	outputing(out, x)
}
