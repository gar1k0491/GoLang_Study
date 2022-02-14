package main

import "fmt"

func main() {

	fmt.Println("Программа симуляции работы маршрутного такси.")
	fmt.Println("Прибываем на остановку «Улица Чкалова».")
	var passangersIn int
	var passangersOut int
	var passangersResult int
	var passangersAll int
	passangersAll += passangersIn
	fmt.Println("В салоне пассажиров:", passangersResult)
	passangersResult -= passangersOut
	fmt.Println("Сколько пассажиров вышло на остановке:", passangersOut)
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&passangersIn)
	passangersResult += passangersIn
	passangersAll += passangersIn
	fmt.Println("В салоне пассажиров:", passangersResult)

	var cash int
	fmt.Println("Отправляемся с остановки «Улица Чкалова».")
	fmt.Println("====================================================")

	fmt.Println("Прибываем на остановку «Улица Кроссовочка».")
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&passangersOut)
	passangersResult -= passangersOut
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&passangersIn)
	passangersResult += passangersIn
	passangersAll += passangersIn
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Println("Отправляемся с остановки «Улица Кроссовочка».")
	fmt.Println("====================================================")

	fmt.Println("Прибываем на остановку «Улица Пенделя».")
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Print("Сколько пассажиров вышло на остановке: ")
	fmt.Scan(&passangersOut)
	passangersResult -= passangersOut
	fmt.Print("Сколько пассажиров вошло на остановке: ")
	fmt.Scan(&passangersIn)
	passangersResult += passangersIn
	passangersAll += passangersIn
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Println("Отправляемся с остановки «Улица Пенделя».")
	fmt.Println("====================================================")

	fmt.Println("Прибываем на остановку «Улица Очкариков» (Конечная).")
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Println("Сколько пассажиров вышло на остановке:", passangersResult)
	passangersResult -= passangersResult
	fmt.Println("Сколько пассажиров вошло на остановке: 0")
	fmt.Println("В салоне пассажиров:", passangersResult)
	fmt.Println("====================================================")

	cash = 20 * passangersAll
	fmt.Print("Всего заработали: ", cash, " руб.\n")
	driverCash := cash / 4
	fmt.Print("Зарплата водителя: ", driverCash, " руб.\n")
	toplivo := cash / 5
	fmt.Print("Расходы на топливо: ", toplivo, " руб.\n")
	nalog := cash / 5
	fmt.Print("Налоги: ", nalog, " руб.\n")
	repair := cash / 5
	fmt.Print("Расходы на ремонт машины: ", repair, " руб.\n")
	dohod := cash - (cash / 4) - (cash / 5) - (cash / 5) - (cash / 5)
	fmt.Print("Итого доход: ", dohod, " руб.")

}
