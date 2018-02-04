package main

import (
	"log"
	"net/http"

	"github.com/art-frela/go-ini"
)

var DNS map[string]string

func main() {
	DNS = make(map[string]string)
	file, err := ini.LoadFile("config.ini")
	if err != nil {
		panic(err)
	}
	DNS = file.Section("db")
	// for el, eval := range DNS {
	// 	fmt.Printf("%s\t%s\n", el, eval)
	// }
	//fmt.Printf("%s", DNS)
	// name, ok := file.Get("db", "server")
	// if !ok {
	// 	panic("'server' variable missing from 'db' section")
	// }
	// fmt.Printf("Name is: %s\n", name)

	// for key, value := range file["mysection"] {
	// 	fmt.Printf("%s => %s\n", key, value)
	// }
	// for name2, section := range file {
	// 	fmt.Printf("Section name: %s:%s\n", section, name2)
	// }

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8086", router))
}
