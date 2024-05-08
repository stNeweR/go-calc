package Math

import (
	"fmt"
	"html/template"
	"math/big"
	"math/rand"
	"net/http"
	"time"
)

type MultiplicationData struct {
	Number string
	Time   time.Duration
}

func MultiplicationHandler(count, maxNumber int, value string, w http.ResponseWriter, r *http.Request) {

	if value == "int" {
		number, elapsed := IntMultiplication(count, maxNumber)
		data := ResultData{
			Result: number,
			Time:   elapsed,
		}
		tmpl, err := template.ParseFiles("./templates/result.html")
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
		fmt.Println(value)
		number, elapsed := FloatMultiplication(count, maxNumber)
		data := ResultData{
			Result: number,
			Time:   elapsed,
		}
		tmpl, err := template.ParseFiles("./templates/result.html")
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

func IntMultiplication(count, maxNumber int) (string, time.Duration) {
	numberOfElements := count
	maxNum := int64(maxNumber)
	array := make([]int64, numberOfElements)
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	for i := 0; i < numberOfElements; i++ {
		array[i] = rand.Int63n(maxNum) + 1
	}
	product := big.NewInt(1)
	temp := new(big.Int)
	for _, num := range array {
		temp.SetInt64(num)
		product.Mul(product, temp)
	}
	result := product.String()
	elapsed := time.Since(start)
	return result, elapsed
}

func FloatMultiplication(count, maxNumber int) (string, time.Duration) {
	numberOfElements := count
	maxNum := float64(maxNumber)
	array := make([]float64, numberOfElements)
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	for i := 0; i < numberOfElements; i++ {
		array[i] = rand.Float64() * maxNum
	}
	product := big.NewFloat(1.0)
	temp := new(big.Float)
	for _, num := range array {
		temp.SetFloat64(num)
		product.Mul(product, temp)
	}
	result := product.Text('g', -1)
	elapsed := time.Since(start)
	return result, elapsed
}
