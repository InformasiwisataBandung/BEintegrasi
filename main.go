package main

import (
	"net/http"

	"github.com/InformasiwisataBandung/BEintegrasi/Controller"
	"github.com/InformasiwisataBandung/BEintegrasi/Login"
	"github.com/InformasiwisataBandung/BEintegrasi/Signup"
)

func main() {

	Controller.Auth()
	// Menghubungkan rute HTTP dari package login
	// Mendaftarkan rute HTTP dari package login
	Login.RegisterLoginRoutes()
	// Mendaftarkan rute HTTP dari package signup
	http.HandleFunc("/Signup", Signup.SignupHandler)
	//Mendaftarkan Fungsi GCF
	// Melayani form login
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/login.html")
		} else {
			http.Error(w, "Metode tidak diizinkancoy", http.StatusMethodNotAllowed)
		}
	})

	// Mulai server HTTP
	http.ListenAndServe(":8989", nil)
}
