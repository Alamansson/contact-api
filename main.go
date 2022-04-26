package main

// https://tproger.ru/translations/deploy-a-secure-golang-rest-api/
// https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b

import (
	"fmt"
	"github.com/alamansson/contact-api/app"
	"github.com/alamansson/contact-api/controllers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)
	router.HandleFunc("api/user/new", controllers.CreateAccount).Methods("GET")
	router.HandleFunc("api/user/login", controllers.Authenticate).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		fmt.Print(err)
	}
}
