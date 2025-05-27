package main

import "fmt"

// We can modify the values of a map just by passing it to a function
func changeMapValues(m map[string]int) {
	for key, value := range m {
		fmt.Println("Changing the value of key ", key)
		m[key] = value * 100
	}
}

func main() {
	m := map[string]int{
		"Subhayan": 40,
		"Shaayan":  37,
		"Dimpu":    36,
	}
	changeMapValues(m)
	for _, value := range m {
		fmt.Println(value)
	}
}
