package main

import (
	"fmt"
	"log"
	"net/http"
)

// var period = flag.Duration("period", 1*time.Second, "sleep period")

// var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	// flag.Parse()
	// fmt.Println(*temp)

	// fmt.Printf("Sleeping for %v...", *period)
	// time.Sleep(*period)
	// fmt.Println()
	db := database{"shoes": 50, "socks": 5}

	// log.Fatal(http.ListenAndServe("localhost:8000", db))

	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))
	// log.Fatal(http.ListenAndServe("localhost:8000", mux)) // *http.ServerMux implemented http.Handler

	// http package already has default mux, we can leverage it
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// type Celsius float64
// type Fahrenheit float64

// func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
// func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// type celsiusFlag struct{ Celsius }

// func (f *celsiusFlag) String() string {
// 	return f.String()
// }
// func (f *celsiusFlag) Set(s string) error {
// 	var unit string
// 	var value float64
// 	fmt.Sscanf(s, "%f%s", &value, &unit)

// 	switch unit {
// 	case "C", "°C":
// 		f.Celsius = Celsius(value)
// 	case "F", "°F":
// 		f.Celsius = FToC(Fahrenheit(value))
// 		return nil
// 	}

// 	return fmt.Errorf("invalid temperature %q", s)
// }

// func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
// 	f := celsiusFlag{value}
// 	fmt.Println(f)
// 	flag.CommandLine.Var(&f, name, usage)
// 	fmt.Println("After ", f)
// 	return &f.Celsius
// }

type dollars float32

// Implement Stringer interface
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}

}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)

}

// Implement Handler interface
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}

	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such pate: %s\n", req.URL)
	}
}
