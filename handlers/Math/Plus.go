package Math

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ResultData struct {
	Result string
	Time   time.Duration
}

func PlusHandler(count, maxNumber int, value string, w http.ResponseWriter, r *http.Request) {

	if value == "int" {
		sum, elapsed := AddInt(count, maxNumber)
		data := ResultData{
			Result: strconv.Itoa(sum),
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
		sum, elapsed := AddFloat(count, maxNumber)

		data := ResultData{
			Result: strconv.FormatFloat(sum, 'f', -1, 64),
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

func AddInt(count, maxNumber int) (int, time.Duration) {
	start := time.Now()
	length := count
	array := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(maxNumber) + 1
	}
	sum := 0
	for _, value := range array {
		sum += value
	}
	elapsed := time.Since(start)
	return sum, elapsed
}

func AddFloat(count, maxNumber int) (float64, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)
	rand.Seed(time.Now().UnixNano())
	floatMax := float64(maxNumber)
	for i := 0; i < length; i++ {
		array[i] = rand.Float64() * floatMax
	}
	sum := 0.0
	for _, value := range array {
		sum += value
	}
	elapsed := time.Since(start)
	return sum, elapsed
}
