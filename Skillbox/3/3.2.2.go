package main

import "fmt"

func main() {
	fmt.Printf("Укажите стоимость товара: ")
	var productPrice int
	fmt.Scan(&productPrice)

	fmt.Printf("Укажите стоимость доставки: ")
	var deliveryPrice int
	fmt.Scan(&deliveryPrice)

	fmt.Printf("Укажите вашу скидку: ")
	var discountValue int
	fmt.Scan(&discountValue)
	fmt.Println("====================================")

	result := productPrice + deliveryPrice - discountValue

	fmt.Println("Стоимость товара:", productPrice)
	fmt.Println("Стоимость доставки:", deliveryPrice)
	fmt.Println("Размер скидки:", discountValue)
	fmt.Println("---------")
	fmt.Println("Итого:", result)

}
