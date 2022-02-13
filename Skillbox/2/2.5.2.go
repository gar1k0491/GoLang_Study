package main

import "fmt"

func main() {
	productPrice := 6400
	deliveryPrice := 350
	discountValue := 700

	result := productPrice + deliveryPrice - discountValue

	fmt.Println("Стоимость товара:", productPrice)
	fmt.Println("Стоимость доставки:", deliveryPrice)
	fmt.Println("Размер скидки:", discountValue)
	fmt.Println("---------")
	fmt.Println("Итого:", result)

}
