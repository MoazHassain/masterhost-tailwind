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
	http.HandleFunc("/dashboard-home/dns", dns)
	http.HandleFunc("/dashboard-home/dns/add", addDomain)
	http.HandleFunc("/dashboard-home/packages", packages)
	http.HandleFunc("/dashboard-home/packages/add", addPackages)
	http.HandleFunc("/dashboard-home/databases", databases)
	http.HandleFunc("/dashboard-home/databases/add", addDatabase)
	http.HandleFunc("/dashboard-home/databases/add-user", addDbUser)
	http.HandleFunc("/billing", billing)
	http.HandleFunc("/profile", profile)

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
func dns(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/products/dns.html")
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
func packages(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/products/packages.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addPackages(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/add-package.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

func databases(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/products/databases.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addDatabase(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/add-database.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}
func addDbUser(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/dashboard-home.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("dist/wpage/add-dbUser.html")
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
