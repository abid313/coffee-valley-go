package controllers

import (
	"WebsiteDataOn/config"
	"WebsiteDataOn/entities"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type OrderstatusInput struct {
	Id       int64
	Bean     string
	Price    float64
	Quantity int64
	Total    float64
	Status   string
}

func IndexOrderStatus(w http.ResponseWriter, r *http.Request) {
	// Retrieve user session from session storage
	session, _ := config.Store.Get(r, config.SESSION_ID)

	// If session is empty, user is not logged in, redirect to login page
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// If user is not logged in, redirect to login page
	if session.Values["loggedIn"] != true {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Fetch all order statuses
	var orderstatus entities.OrderStatus
	orderstatuses, err := UserModel.GetAllOrderStatus(&orderstatus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Log the retrieved data
	fmt.Println("Order statuses data:", orderstatuses)
	if len(orderstatuses) == 0 {
		fmt.Println("No order statuses found!")
	}

	// Parse and execute the template with the retrieved data
	temp, err := template.ParseFiles("views/orderstatus.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	err = temp.Execute(w, orderstatuses)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

func AddOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/addorderstatus.html")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		// Process form data
		r.ParseForm()
		// Map of bean prices
		beanPrices := map[string]float64{
			"Cubita":            12.00,
			"Colombian Supermo": 13.50,
			"Pure Kona Fancy":   15.90,
			"Kenyan":            24.00,
			"Costa Rican":       12.30,
		}

		// Get the selected bean and its price
		selectedBean := r.Form.Get("bean")
		price := beanPrices[selectedBean]

		OrderstatusInput := &OrderstatusInput{
			Bean:     selectedBean,
			Price:    price,
			Quantity: parseInt(r.Form.Get("quantity")),
			Total:    price * float64(parseInt(r.Form.Get("quantity"))),
			Status:   r.Form.Get("status"),
		}

		// Add order status to the database
		err := UserModel.AddOrderStatus(OrderstatusInput.Bean, OrderstatusInput.Price, int(OrderstatusInput.Quantity), OrderstatusInput.Total, OrderstatusInput.Status)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/orderstatus", http.StatusSeeOther)
	}
}

func parseFloat(value string) float64 {
	parsedValue, _ := strconv.ParseFloat(value, 64)
	return parsedValue
}

func parseInt(value string) int64 {
	parsedValue, _ := strconv.ParseInt(value, 10, 64)
	return parsedValue
}

func EditOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Parse the template
		temp, err := template.ParseFiles("views/editorderstatus.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		// Get the "Id" query parameter
		idString := r.URL.Query().Get("Id")
		fmt.Println("Berikut Id-nya:", idString)
		if idString == "" {
			http.Error(w, "Missing Id", http.StatusBadRequest)
			return
		}

		// Fetch order status details using the id
		orderstatus, err := UserModel.GetOrderStatusById(idString) // Pass the integer ID
		if err != nil {
			http.Error(w, "Order status not found", http.StatusNotFound)
			return
		}

		// Execute the template with order status data
		err = temp.Execute(w, orderstatus)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}

	} else if r.Method == "POST" {
		// Handle POST request here (e.g., saving the form data)
		r.ParseForm()

		// Get the "Id" query parameter
		idString := r.Form.Get("id")
		if idString == "" {
			http.Error(w, "Missing Id", http.StatusBadRequest)
			return
		}

		// Map of bean prices
		beanPrices := map[string]float64{
			"Cubita":            12.00,
			"Colombian Supermo": 13.50,
			"Pure Kona Fancy":   15.90,
			"Kenyan":            24.00,
			"Costa Rican":       12.30,
		}

		// Get the selected bean and its price
		selectedBean := r.Form.Get("bean")
		price := beanPrices[selectedBean]

		// Update order status details using the id
		err := UserModel.UpdateOrderStatus(
			idString,
			selectedBean,
			price,
			int(parseInt(r.Form.Get("quantity"))),
			price*float64(parseInt(r.Form.Get("quantity"))),
			r.Form.Get("status"),
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the order status list page after updating
		http.Redirect(w, r, "/orderstatus", http.StatusSeeOther)
	}
}

func DeleteOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get the "Id" query parameter
		idString := r.URL.Query().Get("Id")
		fmt.Println("Berikut Id Stringnya:", idString)
		if idString == "" {
			http.Error(w, "Missing Id", http.StatusBadRequest)
			return
		}

		// Delete order status using the id
		err := UserModel.DeleteOrderStatus(idString) // Pass the integer ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the order status list page after deleting
		http.Redirect(w, r, "/orderstatus", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func MarkDoneOrderStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get the "Id" query parameter
		idString := r.URL.Query().Get("Id")
		fmt.Println("Berikut Id Stringnya:", idString)
		if idString == "" {
			http.Error(w, "Missing Id", http.StatusBadRequest)
			return
		}

		// Update order status to "Done" using the id
		err := UserModel.MarkDoneOrderStatus(idString) // Pass the integer ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the order status list page after marking as done
		http.Redirect(w, r, "/orderstatus", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
