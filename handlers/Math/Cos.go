package Math

import (
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type CosData struct {
	Sum  float64
	Time time.Duration
}

func CosHandler(count, maxNumber int, w http.ResponseWriter, r *http.Request) {
	sum, elapsed := Cos(count, maxNumber)

	data := CosData{
		Sum:  sum,
		Time: elapsed,
	}

	tmpl, err := template.ParseFiles("./templates/cos.html")

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

func Cos(count, maxNumber int) (float64, time.Duration) {
	start := time.Now()
	length := count
	array := make([]float64, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		array[i] = rand.Float64()
	}
	sum := 0.0
	for _, num := range array {
		sum += math.Cos(num)
	}
	elapsed := time.Since(start)
	return sum, elapsed
}
