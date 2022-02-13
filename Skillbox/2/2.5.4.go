package main

import "fmt"

func main() {
	fullPrice := 4000000
	pcsHallwayHouse := 10
	pcsApartmentsHallway := 40

	pcsApartmentsHouse := pcsHallwayHouse * pcsApartmentsHallway
	priceEachApartments := fullPrice / (pcsHallwayHouse * pcsApartmentsHallway)

	fmt.Println("Стоимость кап ремонта дома составляет:", fullPrice, "рублей")
	fmt.Println("Количество квартир в доме:", pcsApartmentsHouse)
	fmt.Println("Оплата за кап ремонт с каждой квартиры составит:", priceEachApartments, "рублей")
}
