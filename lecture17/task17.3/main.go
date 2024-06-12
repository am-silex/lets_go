package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Прочитал, что это очень плохая идея, размещать логгер на уровне пакета
// - это создает жесткую зависимость от типа логгера
// Рекомендуют создавать свой тип реализующий интерфей HandlerFunc но с добавкой параметра
// - логгера, как в статье https://www.angus-morrison.com/blog/go-logger-http-handler-injection
var l = myLogger()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	logHandler := logMiddleware()
	httpServer := &http.Server{Addr: ":8080", Handler: authHandler(logHandler(mux))}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalln(fmt.Errorf("Не удалось запустить сервер: %w", err))
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	msg := "Hello, Go!"
	l.Println("resp:", msg)
	fmt.Fprint(res, msg)
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		xId := r.Header.Get("x-my-app-id")
		if xId != "my_secret" {
			http.Error(w, "пользователь не авторизован", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func logMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			l.Println("url:", r.URL)
			h.ServeHTTP(w, r)
		})
	}
}

func myLogger() *log.Logger {
	logFile, err := os.OpenFile("http_server.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	return log.New(logFile, "", log.LstdFlags)
}
