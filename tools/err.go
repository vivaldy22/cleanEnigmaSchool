package tools

import "log"

// FatalErr func
func FatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// PrintlnErr func
func PrintlnErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

// PanicErr func
func PanicErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
