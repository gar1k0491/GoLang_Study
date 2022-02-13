package main

import "fmt"

func main() {
	var lapNumber int = 4
	var playerName string = "Шумахер"
	var equipmentEngine int = 254
	var equipmentWheel int = 93
	var equipmentSteeringWheel int = 49
	var weatherWind int = 21
	var weatherRain int = 17
	var playerSpeed = equipmentEngine + equipmentWheel + equipmentSteeringWheel - weatherWind - weatherRain
	var playerSpeedResult int = playerSpeed
	fmt.Println("===================")
	fmt.Println("Супер гонки. Круг", lapNumber)
	fmt.Println("===================")
	fmt.Print(playerName, " (", playerSpeedResult, ")\n")
	fmt.Println("===================")
	fmt.Println("Водитель:", playerName)
	fmt.Println("Скорость:", playerSpeedResult)
	fmt.Println("-------------------")
	fmt.Println("Оснащение")
	fmt.Print("Двигатель: +", equipmentEngine, " \n")
	fmt.Print("Колеса: +", equipmentWheel, " \n")
	fmt.Print("Руль: +", equipmentSteeringWheel, " \n")
	fmt.Println("-------------------")
	fmt.Println("Действия плохой погоды")
	fmt.Print("Ветер: -", weatherWind, " \n")
	fmt.Print("Дождь: -", weatherRain, " \n")

}
