package main

import "fmt"

func main() {
	fmt.Println("Добро пожаловать! Для продолжения работы с сайтом Вам необходимо авторизироваться. Пожалуйста следуйте инструкциям ниже.")
	fmt.Printf("Введите ваше имя пользователя: ")
	var userName string
	fmt.Scan(&userName)
	fmt.Printf("Введите ваш пароль: ")
	var userPass string
	fmt.Scan(&userPass)

	fmt.Println("-----")

	fmt.Printf("[%s], вы успешно зашли!", userName)

}
