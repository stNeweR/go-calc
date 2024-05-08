package Math

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type IntMinusData struct {
	Minus int
	Time  time.Duration
}

type FloatMinusData struct {
	Minus float64
	Time  time.Duration
}

func MinusHandler(count, maxNumber int, value string, w http.ResponseWriter, r *http.Request) {
	//minus, elapsed := Minus(count, maxNumber)

	if value == "int" {
		minus, elapsed := IntMinus(count, maxNumber)
		data := ResultData{
			Result: strconv.Itoa(minus),
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
		minus, elapsed := FLoatMinus(count, maxNumber)
		data := ResultData{
			Result: strconv.FormatFloat(minus, 'f', -1, 64),
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
	//tmpl, err := template.ParseFiles("./templates/minus.html")
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//err = tmpl.Execute(w, data)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
}

func IntMinus(count, maxNumber int) (int, time.Duration) {
	start := time.Now()
	length := count
	array := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	// Заполнение массива числами от 1 до 5
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(maxNumber) + 1
	}
	minus := 0
	for _, value := range array {
		minus -= value
	}
	elapsed := time.Since(start)

	return minus, elapsed
}

func FLoatMinus(count, maxNumber int) (float64, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)
	rand.Seed(time.Now().UnixNano())
	floatMax := float64(maxNumber)
	for i := 0; i < length; i++ {
		array[i] = rand.Float64() * floatMax
	}
	minus := 0.0
	for _, value := range array {
		minus -= value
	}
	elapsed := time.Since(start)
	return minus, elapsed
}
