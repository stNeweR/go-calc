package Math

import (
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"time"
)

func TanHandler(count, maxNumber int, value string, w http.ResponseWriter, r *http.Request) {

	if value == "int" {
		sum, elapsed := IntTan(count, maxNumber)

		data := CosData{
			Sum:  sum,
			Time: elapsed,
		}

		tmpl, err := template.ParseFiles("./templates/tan.html")

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
		sum, elapsed := FloatTan(count, maxNumber)

		data := CosData{
			Sum:  sum,
			Time: elapsed,
		}

		tmpl, err := template.ParseFiles("./templates/tan.html")

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

func IntTan(count, maxNumber int) (float64, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		array[i] = float64(rand.Intn(maxNumber) + 1)
	}

	sum := 0.0
	for _, num := range array {
		sum += math.Tan(num)
	}

	elapsed := time.Since(start)
	return sum, elapsed
}

func FloatTan(count, maxNumber int) (float64, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)

	floatMax := float64(maxNumber)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		array[i] = rand.Float64() * floatMax
	}

	sum := 0.0
	for _, num := range array {
		sum += math.Tan(num)
	}

	elapsed := time.Since(start)
	return sum, elapsed
}
