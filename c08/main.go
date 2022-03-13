package main

import (
	"fmt"
	"hfgo/c08/magazine"
)

func main() {
	subscriber := magazine.Subscriber{Name: "Adam East"}
	subscriber.Street = "123 Wayne Manor Lane"
	subscriber.City = "Gotham"
	subscriber.State = "NY"
	subscriber.PostalCode = "01234"
	fmt.Println("Street:", subscriber.Street)
	fmt.Println("City:", subscriber.City)
	fmt.Println("State:", subscriber.State)
	fmt.Println("Postal Code:", subscriber.PostalCode)

	employee := magazine.Employee{Name: "Joy Toodawurld"}
	employee.Street = "246 Even Way"
	employee.City = "Brighton"
	employee.State = "CA"
	employee.PostalCode = "90210"
	fmt.Println("Street:", employee.Street)
	fmt.Println("City:", employee.City)
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)

	// NOTE below you can use the longer variable dot Address dot street form but variable dot Street is much less typing :)
	employee01 := magazine.Employee{Name: "Juan Pickselov", Salary: 87999}
	employee01.Address.Street = "123 Lotus Lane"
	employee01.City = "Battlefront"
	employee01.State = "CO"
	fmt.Println("Street:", employee01.Street)
	fmt.Println("City:", employee01.City)
	fmt.Println("State:", employee01.State)
	fmt.Println("Postal Code:", employee01.PostalCode)
}
