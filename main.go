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
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/logout", controller.Logout)

	// catalog
	http.HandleFunc("/catalog", controller.IndexCatalog)

	// Distributors
	http.HandleFunc("/distributors", controller.IndexDistributor)
	http.HandleFunc("/distributors/add", controller.AddDistributor)
	http.HandleFunc("/distributors/edit", controller.EditDistributor)
	http.HandleFunc("/distributors/delete", controller.DeleteDistributor)

	// Order Status
	http.HandleFunc("/orderstatus", controller.IndexOrderStatus)
	http.HandleFunc("/orderstatus/add", controller.AddOrderStatus)
	http.HandleFunc("/orderstatus/edit", controller.EditOrderStatus)
	http.HandleFunc("/orderstatus/delete", controller.DeleteOrderStatus)
	http.HandleFunc("/orderstatus/markdone", controller.MarkDoneOrderStatus)

	fmt.Println("Server started at port http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
	http.ListenAndServe(":8080", nil)
}
