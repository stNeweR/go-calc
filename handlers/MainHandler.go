package handlers

import (
	"calculator/handlers/Math"
	"net/http"
	"strconv"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Ошибка разбора формы", http.StatusInternalServerError)
		return
	}
	count := r.FormValue("count")
	maxNumber := r.FormValue("max")
	operation := r.FormValue("operation")
	value := r.FormValue("type")

	num, err := strconv.Atoi(count)
	maxNum, err := strconv.Atoi(maxNumber)

	if operation == "plus" {
		// +
		Math.PlusHandler(num, maxNum, value, w, r)
	} else if operation == "minus" {
		// +
		Math.MinusHandler(num, maxNum, value, w, r)
	} else if operation == "multiplication" {
		// +
		Math.MultiplicationHandler(num, maxNum, value, w, r)
	} else if operation == "divide" {
		// +
		Math.DivideHandler(num, maxNum, value, w, r)
	} else if operation == "cos" {
		Math.CosHandler(num, maxNum, w, r)
	} else if operation == "tan" {
		// +
		Math.TanHandler(num, maxNum, value, w, r)
	}

}
