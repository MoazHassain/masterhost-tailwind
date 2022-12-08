package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.HandleFunc("/component", component)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/dashboard-home", dashboardHome)
	http.HandleFunc("/billing", billing)
	http.HandleFunc("/profile", profile)
	http.HandleFunc("/add-domain", addDomain)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./dist/assets"))))

	fmt.Println("starting server on: 8500")
	http.ListenAndServe(":8500", nil)

}

func component(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/components.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func signup(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/sign-up.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func login(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/sign-in.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func dashboardHome(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)
}

func billing(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/billing.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func profile(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/profile.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addDomain(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/add-domain.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
