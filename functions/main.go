// type assertion in Go
package main

import "fmt"

func main() {
	names := []string{"Subhayan", "Shaayan", "Dimpu", "Papa", "Mama"}
	for _, name := range names {
		switch name {
		case "Subhayan", "Dimpu":
			fmt.Println("Couples..")
		case "Shaayan":
			fmt.Println("Brother")
		case "Papa", "Mama":
			fmt.Println("Parents")
		}
	}
}
