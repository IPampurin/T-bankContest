/*
1 задание
Костя подключен к мобильному оператору «Мобайл». Абонентская плата Кости составляет A рублей в месяц. За эту стоимость Костя получает B мегабайт интернет-трафика.
Если Костя выйдет за лимит трафика, то каждый следующий мегабайт будет стоить ему C рублей.
Костя планирует потратить D мегабайт интернет-трафика в следующий месяц. Помогите ему сосчитать, во сколько рублей ему обойдется интернет.
Формат входных данных
Вводится 4 целых положительных числа A,B,C,D(1≤A,B,C,D≤100) — стоимость тарифа Кости, размер тарифа Кости, стоимость каждого лишнего мегабайта,
размер интернет-трафика Кости в следующем месяце. Числа во входном файле разделены пробелами.
Формат выходных данных
Выведите одно натуральное число — суммарные расходы Кости на интернет.

Замечание
В первом примере Костя сначала оплатит пакет интернета, после чего потратит на 5 мегабайт больше, чем разрешено по тарифу.
Следовательно, за 5 мегабайт он дополняет отдельно, получившаяся стоимость 100+12×5=160 рублей.
Во втором примере Костя укладывается в тарифный план, поэтому платит только за него.

Примеры данных
Пример 1
100  10  12  15
160
Пример 2
100  10  12  1
100
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

// inputing считывает данные и возвращает введённые значения
func inputing(sc *bufio.Scanner) (int, int, int, int) {

	// считываем ввод
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	// парсим входные данные
	// стоимость тарифа
	costTariff, err := strconv.Atoi(input[0])

	// размер тарифа
	sizeTariff, err := strconv.Atoi(input[1])

	// стоимость мегабайта сверх лимита
	costMBOverLimit, err := strconv.Atoi(input[2])

	// плановый трафик
	planTraffic, err := strconv.Atoi(input[3])

	if err != nil {
		log.Fatal(err)
	}

	return costTariff, sizeTariff, costMBOverLimit, planTraffic
}

// calculation вычисляет сумму к оплате
func calculation(costTariff, sizeTariff, costMBOverLimit, planTraffic int) int {

	x := costTariff                       // сначала сумма равна тарифу
	overLimit := planTraffic - sizeTariff // ожидаемое превышение трафика
	// если превышения не предполагается, то и к тарифу прибавлять нечего
	if overLimit < 0 {
		overLimit = 0
	}
	// итого сумма
	x += overLimit * costMBOverLimit

	return x
}

// outputing выводит результат
func outputing(out *bufio.Writer, x int) {

	fmt.Fprintf(out, "%v", x)
}

func main() {

	// пределяем ввод
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	// определяем вывод
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	// считываем введенные данные
	costTariff, sizeTariff, costMBOverLimit, planTraffic := inputing(scanner)

	// вычисляем сумму
	x := calculation(costTariff, sizeTariff, costMBOverLimit, planTraffic)

	// выводим результат
	outputing(out, x)
}
