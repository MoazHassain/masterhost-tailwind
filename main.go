package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", dashboardHome)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("./dist/assets"))))

	fmt.Println("starting server on: 8008")
	http.ListenAndServe(":8008", nil)

}

func dashboardHome(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("dist/template/index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp.Execute(w, nil)

}

// func signUp(w http.ResponseWriter, r *http.Request) {

// 	ptmp, err := template.ParseFiles("wpage/login.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp, err = ptmp.ParseFiles("wpage/signup.gohtml")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	ptmp.Execute(w, nil)

// }
