package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}
func handlerAbout(w http.ResponseWriter, r *http.Request) {
	var message = "About"
	w.Write([]byte(message))
}

func main() {
	//fmt.Println("Hello World")
	//fmt.Println("hello", "world!", "how", "are", "you")

	// var firstName string = "Aris"

	// var lastName string = "Muflihul"

	// nickName := "Aris"

	//fmt.Print("Halo", nickName, firstName, lastName+"!")

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)
	http.HandleFunc("/about", handlerAbout)

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

}
