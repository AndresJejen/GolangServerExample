package main

import (
	"errors"
	"fmt"
	"strconv"

	"log"
	"net"
	"net/http"
	"os"

	"github.com/AndresJejen/AprendiendoGo/calculadora"
)

func ToFloat(value string) (float32, error) {
	if s, err := strconv.ParseFloat(value, 32); err == nil {
		return float32(s), nil
	}
	return 0, errors.New("El valor ingresado no puede convertirse a flotante")
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sumar" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	sumando, _ := r.URL.Query()["sumando"]
	elotro, _ := r.URL.Query()["otro"]

	number_one, err := ToFloat(sumando[0])

	if err != nil {
		http.Error(w, "sumando is not a number", http.StatusBadRequest)
		return
	}
	number_two, err1 := ToFloat(elotro[0])
	if err1 != nil {
		http.Error(w, "otro is not a number", http.StatusBadRequest)
		return
	}

	resultado := calculadora.Sumar(number_one, number_two)

	fmt.Fprintf(w, "Los numeros que me enviaste son %v y %v - y el resultado es %v", number_one, number_two, resultado)
}

func SubstractHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/restar" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	sumando, _ := r.URL.Query()["sumando"]
	elotro, _ := r.URL.Query()["otro"]

	number_one, err := ToFloat(sumando[0])

	if err != nil {
		http.Error(w, "sumando is not a number", http.StatusBadRequest)
		return
	}
	number_two, err1 := ToFloat(elotro[0])
	if err1 != nil {
		http.Error(w, "otro is not a number", http.StatusBadRequest)
		return
	}

	resultado := calculadora.Restar(number_one, number_two)

	fmt.Fprintf(w, "Los numeros que me enviaste son %v y %v - y el resultado es %v", number_one, number_two, resultado)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/multiplicar" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	sumando, _ := r.URL.Query()["sumando"]
	elotro, _ := r.URL.Query()["otro"]

	number_one, err := ToFloat(sumando[0])

	if err != nil {
		http.Error(w, "sumando is not a number", http.StatusBadRequest)
		return
	}
	number_two, err1 := ToFloat(elotro[0])
	if err1 != nil {
		http.Error(w, "otro is not a number", http.StatusBadRequest)
		return
	}

	resultado := calculadora.Multiplicar(number_one, number_two)

	fmt.Fprintf(w, "Los numeros que me enviaste son %v y %v - y el resultado es %v", number_one, number_two, resultado)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dividir" {
		http.Error(w, "404 not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}

	sumando, _ := r.URL.Query()["Divisor"]
	elotro, _ := r.URL.Query()["Dividendo"]

	number_one, err := ToFloat(sumando[0])

	if err != nil {
		http.Error(w, "Divisor is not a number", http.StatusBadRequest)
		return
	}
	number_two, err1 := ToFloat(elotro[0])
	if err1 != nil {
		http.Error(w, "Dividendo is not a number", http.StatusBadRequest)
		return
	}

	resultado, err_div := calculadora.Dividir(number_one, number_two)

	if err_div != nil {
		http.Error(w, err_div.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Los numeros que me enviaste son %v y %v - y el resultado es %v", number_one, number_two, resultado)
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Serving request %s", r.URL.Path)
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hola Mundo!\n")
	fmt.Fprintf(w, "Version: 1.0.0\n")
	fmt.Fprintf(w, "Hostname %s\n", host)

	addrs, _ := net.InterfaceAddrs()

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Fprintf(w, "IP Addres: %s\n", ipnet.IP.String())
			}
		}
	}
}

func main() {
	http.HandleFunc("/sumar", AddHandler)
	http.HandleFunc("/restar", SubstractHandler)
	http.HandleFunc("/dividir", DivideHandler)
	http.HandleFunc("/server", ServerHandler)
	http.HandleFunc("/multiplicar", MultiplyHandler)

	fmt.Printf("Starting server at port 8081\n")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
