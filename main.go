package main

import (
	controller "WebsiteDataOn/controllers"
	"fmt"
	"net/http"
)

func main() {
	// Auth Login dan Logout
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)

	// catalog
	http.HandleFunc("/catalog", controller.IndexCatalog)

	// orderstatus
	http.HandleFunc("/orderstatus", controller.IndexDistributor)
	http.HandleFunc("/orderstatus/add", controller.AddDistributor)
	http.HandleFunc("/orderstatus/edit", controller.EditDistributor)

	fmt.Println("Server started at port http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
	http.ListenAndServe(":8080", nil)
}
