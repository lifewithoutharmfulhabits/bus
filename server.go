package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

func StartServer() {

	// static conent: css, html, js, etc.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/payment/add", handlePageAddPayment)

	http.ListenAndServe(":8087", nil)
}

// Function that handles add payment page
func handlePageAddPayment(w http.ResponseWriter, r *http.Request) {

	// Path to main template file.
	layoutFile := path.Join("static", "html", "layout.tmpl")
	fmt.Println(layoutFile)

	// Path to payment form file. In is nested to main template.
	addPaymentFile := path.Join("static", "html", "add_payment.tmpl")

	// Parsing templates.
	// We need to specify both template files that are used to render add patment page
	templ, tempErr := template.New("layout").ParseFiles(layoutFile, addPaymentFile)
	if tempErr != nil {
		fmt.Println(tempErr)
	}

	// Process template and write it to browser
	renderErr := templ.ExecuteTemplate(w, "layout", nil)
	if renderErr != nil {
		fmt.Println(renderErr)
	}
}
