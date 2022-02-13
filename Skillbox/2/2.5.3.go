package main

import "fmt"

func main() {
	restaurantName := "``Крылышки и ножки``"
	workTimeCashier := 480
	clientTime := 2
	cashierTime := 4

	waitTime := clientTime + cashierTime
	howManyClients := workTimeCashier / (clientTime + cashierTime)

	fmt.Println("Вас приветствует ресторан", restaurantName)
	fmt.Println("Приблизительное время обслуживания:", waitTime, "минут.")
	fmt.Println("За смену обслуживается", howManyClients, "клиентов!")
}
