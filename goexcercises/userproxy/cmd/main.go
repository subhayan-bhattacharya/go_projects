package main

import (
	"fmt"
	"math/rand"
	"userproxy"
)

func main() {
	database := userproxy.UserList{}
	for i := 0; i < 1000000; i++ {
		n := rand.Int31()
		database = append(database, userproxy.User{Id: n})
	}
	proxy := userproxy.UserListProxy{
		Database:  database,
		Capacity:  2,
		UsedCache: true,
	}
	fmt.Printf("%v+\n", proxy)
}
