package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// Car represents a simple demo structure for car details.
type Car struct {
	Make  string
	Model string
}

func main() {
	// Print a greeting to the console when the program starts.
	fmt.Println("Hello World")

	// Define a handler function for the root ("/") route.
	handleFunc1 := func(w http.ResponseWriter, r *http.Request) {
		// Parse the HTML template file.
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Create a map of cars to be displayed in the template.
		cars := map[string][]Car{
			"Cars": {
				{Make: "Chevrolet", Model: "Corvette"},
				{Make: "Tesla", Model: "Model S"},
				{Make: "BMW", Model: "3 Series"},
			},
		}

		// Execute the template, passing the map of cars to the response writer.
		tmpl.Execute(w, cars)
	}

	// Define a handler function for the "/add-car/" route.
	handleFunc2 := func(w http.ResponseWriter, r *http.Request) {
		// Simulate a delay to showcase asynchronous behavior.
		time.Sleep(1 * time.Second)

		// Retrieve form values from the POST request.
		make := r.PostFormValue("make")
		model := r.PostFormValue("model")

		// Print the details of the added car to the console.
		fmt.Println("Make:", make, "Model:", model, "added to the list")

		// Parse the HTML template file.
		tmpl := template.Must(template.ParseFiles("index.html"))

		// Execute a specific template within the file, passing a new Car instance.
		tmpl.ExecuteTemplate(w, "car-list-element", Car{Make: make, Model: model})
	}

	// Define a handler for serving static files (e.g., CSS, JavaScript) from the "/static/" path.
	// Important: This line ensures that files in the "static" directory are served correctly.
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register the handler functions for different routes.
	http.HandleFunc("/", handleFunc1)
	http.HandleFunc("/add-car/", handleFunc2)

	// Start the HTTP server and listen on port 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
