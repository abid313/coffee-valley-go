package controllers

import (
	"WebsiteDataOn/config"
	"WebsiteDataOn/entities"
	"WebsiteDataOn/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// UserInput adalah struktur data untuk menerima input username dan password dari form login
type UserInput struct {
	username string // Menyimpan username dari input form
	password string // Menyimpan password dari input form
}

// UserModel adalah instance dari model User yang digunakan untuk mengakses data pengguna
var UserModel = models.NewUserModel()

// Index adalah handler untuk route "/" yang memeriksa status login pengguna
func Index(w http.ResponseWriter, r *http.Request) {
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
			var catalog entities.Catalog
			catalogs, err := UserModel.IndexHome(&catalog)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
			} else {
				fmt.Println("Catalogs data:", catalogs) // Menampilkan data di console
				temp, _ := template.ParseFiles("views/index.html")
				temp.Execute(w, catalogs)
			}
		}
	}
}

// Login adalah handler untuk menangani login pengguna
func Login(w http.ResponseWriter, r *http.Request) {
	// Jika request method adalah GET, tampilkan halaman login
	if r.Method == "GET" {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == "POST" {
		// Proses login: ambil data dari form
		r.ParseForm()
		UserInput := &UserInput{
			username: r.Form.Get("username"),
			password: r.Form.Get("password"),
		}

		// Cari pengguna di database berdasarkan username
		var user entities.User
		UserModel.Where(&user, "username", UserInput.username)

		var message error
		if user.Username == "" {
			// Jika pengguna tidak ditemukan, beri pesan error
			message = errors.New("Username atau Password salah!")
		} else {
			// Verifikasi password menggunakan bcrypt
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.password))
			if errPassword != nil {
				// Jika password tidak cocok, beri pesan error
				message = errors.New("Username atau Password salah!")
			}
		}

		if message != nil {
			// Jika ada error, kirim pesan error ke halaman login
			data := map[string]interface{}{
				"error": message,
			}
			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			// Jika login berhasil, set session pengguna
			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Username
			session.Values["nama_lengkap"] = user.NamaLengkap

			// Simpan session
			session.Save(r, w)

			// Redirect pengguna ke halaman utama
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

// Logout adalah handler untuk menangani logout pengguna
func Logout(w http.ResponseWriter, r *http.Request) {
	// Mengambil sesi pengguna dari penyimpanan sesi
	session, _ := config.Store.Get(r, config.SESSION_ID)

	// Hapus session (logout)
	session.Options.MaxAge = -1

	// Simpan perubahan session
	session.Save(r, w)

	// Redirect pengguna ke halaman login setelah logout
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
