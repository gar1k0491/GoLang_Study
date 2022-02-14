package main

import "fmt"

func main() {
	fmt.Println("Добро пожаловать! Для продолжения работы с сайтом Вам необходимо пройти регистрацию. Пожалуйста следуйте инструкциям ниже.")
	fmt.Printf("Введите ваше имя пользователя: ")
	var userName string
	fmt.Scan(&userName)
	fmt.Printf("Введите ваш пароль: ")
	var userPass string
	fmt.Scan(&userPass)
	fmt.Printf("Введите ваш возраст: ")
	var userAge int
	fmt.Scan(&userAge)

	fmt.Printf("Поздравляю, %s, теперь вы зарегестрированы!\n", userName)
	fmt.Println("Ваш пароль:", userPass)
	fmt.Println("Ваш возраст:", userAge)
}
