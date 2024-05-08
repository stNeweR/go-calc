package main

import (
	Home "calculator/handlers"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Home.Handler(rr, req)
	fmt.Println("Начальная страница отвечает")
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидался статус-код %v, получен %v", http.StatusOK, status)
	}
}

func TestAdd(t *testing.T) {

	_, numbers, _ := Add()

	file, err := os.Create("number.txt")
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	for _, number := range numbers {
		_, err := fmt.Fprint(file, number, " ")
		if err != nil {
			fmt.Println("Ошибка при записи числа в файл:", err)
			return
		}
	}
	t.Log("Результаты записаны в файлы number.txt")
}

func Add() (int, []int, time.Duration) {
	start := time.Now()
	length := 3000
	array := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		array[i] = rand.Intn(7) + 1
	}
	sum := 0
	for _, value := range array {
		sum += value
	}
	elapsed := time.Since(start)
	return sum, array, elapsed
}
