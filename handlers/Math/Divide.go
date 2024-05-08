package Math

import (
	"html/template"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

type DivideData struct {
	Result *big.Float
	Time   time.Duration
}

func DivideHandler(count, maxNumber int, value string, w http.ResponseWriter, r *http.Request) {

	if value == "int" {
		divide, elapsed := IntDivide(count, maxNumber)

		data := DivideData{
			Result: divide,
			Time:   elapsed,
		}

		tmpl, err := template.ParseFiles("./templates/divide.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if value == "float" {
		divide, elapsed := FloatDivide(count, maxNumber)

		data := DivideData{
			Result: divide,
			Time:   elapsed,
		}

		tmpl, err := template.ParseFiles("./templates/divide.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func IntDivide(count, maxNumber int) (*big.Float, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)

	for i := 0; i < length; i++ {
		array[i] = float64(rand.Intn(maxNumber) + 1)
	}

	divide := new(big.Float).SetFloat64(float64(0))
	for _, value := range array {
		ratValue := new(big.Float).SetFloat64(value)
		divide.Quo(divide, ratValue)
	}

	result := new(big.Float)
	result.SetPrec(100).Set(divide)
	elapsed := time.Since(start)

	return result, elapsed
}

func FloatDivide(count, maxNumber int) (*big.Float, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)

	floatMax := float64(maxNumber)

	for i := 0; i < length; i++ {
		array[i] = rand.Float64() * floatMax
	}

	divide := new(big.Float).SetFloat64(float64(0))
	for _, value := range array {
		ratValue := new(big.Float).SetFloat64(value)
		divide.Quo(divide, ratValue)
	}

	result := new(big.Float)
	result.SetPrec(100).Set(divide)
	elapsed := time.Since(start)

	return result, elapsed
}
