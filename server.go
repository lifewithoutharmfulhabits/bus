package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type Response struct {
	Data    interface{}
	Message string
}

func StartServer() {

	// static conent: css, html, js, etc.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/payment/add", handlePageAddPayment)
	http.HandleFunc("/driver/add", handlePageAddDriver)
	http.HandleFunc("/driver/list", handlePageListDrivers)

	http.ListenAndServe(":8087", nil)
}

// Function that handles add payment page
func handlePageAddPayment(w http.ResponseWriter, r *http.Request) {
	renderInLayout(w, "add_payment.tmpl", new(Response))
}

func handlePageAddDriver(w http.ResponseWriter, r *http.Request) {

	response := new(Response)

	if r.Method == "POST" && r.FormValue("firstName") != "" && r.FormValue("lastName") != "" {
		driver := Driver{
			FirstName: r.FormValue("firstName"),
			LastName:  r.FormValue("lastName"),
		}
		DAO.addDriver(&driver)
		response.Message = "Водитель добавлен"
	}

	renderInLayout(w, "add_driver.tmpl", response)
}

func handlePageListDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := DAO.getAllDrivers()
	if err != nil {
		// TODO : diplay error to user in a friendly form
	}
	renderInLayout(w, "list_drivers.tmpl", &Response{Data: drivers})
}

func renderInLayout(w http.ResponseWriter, templateName string, response *Response) {
	// Path to main template file.
	layoutFile := path.Join("templates", "layout.tmpl")
	fmt.Println(layoutFile)

	// Path to payment form file. In is nested to main template.
	addPaymentFile := path.Join("templates", templateName)

	// Parsing templates.
	// We need to specify both template files that are used to render add patment page
	templ, tempErr := template.New("layout").ParseFiles(layoutFile, addPaymentFile)
	if tempErr != nil {
		fmt.Println(tempErr)
	}

	// Process template and write it to browser
	renderErr := templ.ExecuteTemplate(w, "layout", response)
	if renderErr != nil {
		fmt.Println(renderErr)
	}
}
