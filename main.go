// Vivaldy Andhira Suwandhi - Challenge 6.3 Enigma School API #3

package main

import (
	"github.com/vivaldy22/cleanEnigmaSchool/configs"
)

func main() {
	db := configs.InitDB()
	router := configs.CreateRouter()
	configs.InitRouters(db, router)
	configs.RunServer(router)
}
