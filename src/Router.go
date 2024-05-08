package src

import (
	Home "calculator/handlers"
	Main "calculator/handlers"
	"fmt"
	"net/http"
)

type GetRoutes struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

func Router() *http.ServeMux {
	routes := []GetRoutes{
		{Path: "/", Handler: Home.Handler},
		//{Path: "/plus/", Handler: Plus.PlusHandler},
		//{Path: "/multiplication/", Handler: Multiplication.MultiplicationHandler},
		//{Path: "/minus/", Handler: Minus.MinusHandler},
		//{Path: "/divide/", Handler: Divide.DivideHandler},
		//{Path: "/cos/", Handler: Cos.CosHandler},
		//{Path: "/tan/", Handler: Tan.TanHandler},
		{Path: "/main/", Handler: Main.MainHandler},
	}

	router := http.NewServeMux()

	for _, route := range routes {
		router.HandleFunc(route.Path, route.Handler)
	}

	return router
}

// Обработчик страницы "Контакты"
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Страница 'Контакты'")
}
