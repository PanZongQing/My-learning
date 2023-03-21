package main

import (
	"My-learning/pkg/setting"
	"My-learning/routers"
	"fmt"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriterTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
