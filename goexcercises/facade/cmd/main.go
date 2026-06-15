package main

import (
	"facade"
	"fmt"
	"net/http"
)

func main() {
	facade := facade.WeatherFacade{
		Client: &http.Client{},
	}
	response, err := facade.GetByCityAndCountryCode("Kolkata", "IN")
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%+v\n", response)
}
