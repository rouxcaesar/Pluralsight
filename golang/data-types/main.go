package main

import (
	"fmt"
	"pluralsight/golang/data-types/organization"
)

func main() {
	p := organization.NewPerson("James", "Wilson")

	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s\n", err.Error())
	}

	fmt.Printf("%T\n", organization.TwitterHandler("test"))

	println(p.ID())
	println(p.FullName())
	println(p.TwitterHandler())
	println(p.TwitterHandler().RedirectUrl())
}
