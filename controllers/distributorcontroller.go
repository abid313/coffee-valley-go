package controllers

import (
	"WebsiteDataOn/config"
	"WebsiteDataOn/entities"
	"fmt"
	"html/template"
	"net/http"
)

type DistributorInput struct {
	Id      int64
	Nama    string
	City    string
	Region  string
	Country string
	Phone   string
	Email   string
}

func IndexDistributor(w http.ResponseWriter, r *http.Request) {
	// Mengambil sesi pengguna dari penyimpanan sesi
	session, _ := config.Store.Get(r, config.SESSION_ID)

	// Jika sesi kosong, pengguna belum login, redirect ke halaman login
	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		// Jika pengguna tidak dalam status login, redirect ke halaman login
		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			var distributor entities.Distributor
			distributors, err := UserModel.GetAllDistributor(&distributor)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				fmt.Println("Distributors data:", distributors) // Menampilkan data di console
				if len(distributors) == 0 {
					fmt.Println("No distributors found!")
				}
				temp, _ := template.ParseFiles("views/distributors.html")
				temp.Execute(w, distributors)
			}

		}
	}
}

func AddDistributor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/adddistributors.html")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		// Proses login: ambil data dari form
		r.ParseForm()
		DistributorInput := &DistributorInput{
			Nama:    r.Form.Get("name"),
			City:    r.Form.Get("city"),
			Region:  r.Form.Get("region"),
			Country: r.Form.Get("country"),
			Phone:   r.Form.Get("phone"),
			Email:   r.Form.Get("email"),
		}

		// Cari pengguna di database berdasarkan username
		err := UserModel.AddDistributor(DistributorInput.Nama, DistributorInput.City, DistributorInput.Region, DistributorInput.Country, DistributorInput.Phone, DistributorInput.Email)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/distributors", http.StatusSeeOther)
	}
}

func EditDistributor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Parse the template
		temp, err := template.ParseFiles("views/editdistributors.html")
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
		fmt.Println(idString)

		// Fetch distributor details using the id
		distributor, err := UserModel.EditDetailDistributor(idString) // Pass the integer ID
		if err != nil {
			http.Error(w, "Distributor not found", http.StatusNotFound)
			return
		}

		// Execute the template with distributor data
		err = temp.Execute(w, distributor)
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
		fmt.Println(idString)
		// Update distributor details using the id
		err := UserModel.UpdateDistributor(idString, r.Form.Get("name"), r.Form.Get("city"), r.Form.Get("region"), r.Form.Get("country"), r.Form.Get("phone"), r.Form.Get("email"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Redirect to the distributor list page after updating
		http.Redirect(w, r, "/distributors", http.StatusSeeOther)
	}
}

func DeleteDistributor(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Get the "Id" query parameter
		idString := r.URL.Query().Get("Id")
		fmt.Println("Berikut Id Stringnya:", idString)
		if idString == "" {
			http.Error(w, "Missing Id", http.StatusBadRequest)
			return
		}

		// Fetch distributor details using the id
		err := UserModel.DeleteDistributor(idString) // Pass the integer ID
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the distributor list page after deleting
		http.Redirect(w, r, "/distributors", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
