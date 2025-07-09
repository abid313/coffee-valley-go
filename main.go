package main

import (
	controller "WebsiteDataOn/controllers"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Ambil port dari env (Railway support ini)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback untuk lokal
	}

	// Routing
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/logout", controller.Logout)

	http.HandleFunc("/catalog", controller.IndexCatalog)

	http.HandleFunc("/distributors", controller.IndexDistributor)
	http.HandleFunc("/distributors/add", controller.AddDistributor)
	http.HandleFunc("/distributors/edit", controller.EditDistributor)
	http.HandleFunc("/distributors/delete", controller.DeleteDistributor)

	http.HandleFunc("/orderstatus", controller.IndexOrderStatus)
	http.HandleFunc("/orderstatus/add", controller.AddOrderStatus)
	http.HandleFunc("/orderstatus/edit", controller.EditOrderStatus)
	http.HandleFunc("/orderstatus/delete", controller.DeleteOrderStatus)
	http.HandleFunc("/orderstatus/markdone", controller.MarkDoneOrderStatus)

	// Jalankan server
	fmt.Println("Server started at port :" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
