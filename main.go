// Vivaldy Andhira Suwandhi - Challenge 6.3 Enigma School API #3

package main

import (
	"fmt"

	"github.com/vivaldy22/cleanEnigmaSchool/configs"
)

func main() {
	fmt.Println("Hello World")
	db := configs.InitDB()
	router := configs.CreateRouter()
	configs.InitRouters(db, router)
	configs.RunServer(router)
}
