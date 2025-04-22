package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Daha sonra düzeltmeyi planlıyorum bu kısıma bakacağım
// Beliki bunlara kat verebilirim mesela km 1000 gibi
var acceptedLength = map[string]float32{
	"mm":  0.001,
	"cm":  0.01,
	"m":   1,
	"km":  0.001,
	"inc": 39.37,
	"ft":  3.28,
	"yd":  1.09,
	"mi":  0.00062,
}
var acceptedWeight = map[string]float32{
	"mg":  100,
	"g":   1,
	"kg":  0.001,
	"ons": 0.035,
	"lb":  0.0022,
}

var acceptedTemperature = map[string]float32{
	"C": 1,
	"F": 33.8,
	"K": 274.15,
}

type Converter struct {
	Val  float32 `json:"val"`
	From string  `json:"from"`
	To   string  `json:"to"`
}

func lengthHandler(w http.ResponseWriter, r *http.Request) {
	var converter Converter
	err := json.NewDecoder(r.Body).Decode(&converter)
	if err != nil {
		http.Error(w, "Invalid JSON value", http.StatusBadRequest)
		return
	}

	err = validVal(converter.From, converter.To, acceptedLength)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := convertMap(converter.From, converter.To, converter.Val, acceptedLength)
	fmt.Fprintf(w, "Result : %.2f\n", res)

	fmt.Fprintf(w, "Length: %f, From: %s, To : %s\n", converter.Val, converter.From, converter.To)
}

func weightHandler(w http.ResponseWriter, r *http.Request) {
	var converter Converter
	err := json.NewDecoder(r.Body).Decode(&converter)
	if err != nil {
		http.Error(w, "Invalid JSON value", http.StatusBadRequest)
		return
	}

	err = validVal(converter.From, converter.To, acceptedWeight)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := convertMap(converter.From, converter.To, converter.Val, acceptedWeight)
	fmt.Fprintf(w, "Result : %.2f\n", res)

	fmt.Fprintf(w, "Weight: %f, From: %s, To : %s\n", converter.Val, converter.From, converter.To)
}

func temperatureHandler(w http.ResponseWriter, r *http.Request) {
	var converter Converter
	err := json.NewDecoder(r.Body).Decode(&converter)
	if err != nil {
		http.Error(w, "Invalid JSON value", http.StatusBadRequest)
		return
	}

	err = validVal(converter.From, converter.To, acceptedTemperature)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := convertMap(converter.From, converter.To, converter.Val, acceptedTemperature)
	fmt.Fprintf(w, "Result : %.2f\n", res)

	fmt.Fprintf(w, "Temperature: %f, From: %s, To : %s\n", converter.Val, converter.From, converter.To)
}

func convertMap(from string, to string, val float32, accepted map[string]float32) float32 {
	temp := val / accepted[from]
	return temp * accepted[to]
}

func validVal(from string, to string, accepted map[string]float32) error {
	var i int = 0
	for val, _ := range accepted {
		if val == from || val == to {
			i++
		}
	}
	if i == 2 {
		return nil
	}
	return fmt.Errorf("invalid unit")
}

func main() {
	http.HandleFunc("/length", lengthHandler)
	http.HandleFunc("/weight", weightHandler)
	http.HandleFunc("/temperature", temperatureHandler)

	fmt.Println("Server is running")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server did not start")
	}
}
