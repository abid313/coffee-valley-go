package controllers

import (
	"WebsiteDataOn/config"
	"WebsiteDataOn/entities"
	"fmt"
	"html/template"
	"net/http"
)

func IndexCatalog(w http.ResponseWriter, r *http.Request) {
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
			// Cari pengguna di database berdasarkan username
			var catalog entities.Catalog
			catalogs, err := UserModel.GetAllCatalog(&catalog)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				fmt.Println("Catalogs data:", catalogs) // Menampilkan data di console
				if len(catalogs) == 0 {
					fmt.Println("No catalogs found!")
				}
				temp, _ := template.ParseFiles("views/catalog.html")
				temp.Execute(w, catalogs)
			}

		}
	}
}
