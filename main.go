package main

import (
	"fmt"
    "encoding/json"
    "log"
    "net/http"
	"strconv"
	"go-learning/customModule"
)

// Define a struct for your response
type Message struct {
    Text string `json:"text"`
	Number int  `json:"num"`
}

type AddRequest struct {
    A int `json:"a"`
    B int `json:"b"`
}

func setHeaderToResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}


func helloHandler(w http.ResponseWriter, r *http.Request) {
    // Prepare response
    msg := Message{Text: "Hello, REST API in Go!", Number: 20}
    
    // Set header to JSON
	setHeaderToResponse(w)

    // Encode response as JSON
    json.NewEncoder(w).Encode(msg)
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{Text: "My name is puneet", Number: 10}
	setHeaderToResponse(w)
	json.NewEncoder(w).Encode(msg)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Get query params as strings
	aStr := r.URL.Query().Get("a")
    bStr := r.URL.Query().Get("b")

	// Convert to int
    a, err1 := strconv.Atoi(aStr)
    b, err2 := strconv.Atoi(bStr)

	setHeaderToResponse(w)

	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid or missing query parameters 'a' and 'b'", http.StatusBadRequest)
		return
	}

	sum := customModule.Add(a, b)

	msg := Message{
		Text:   fmt.Sprintf("Sum of %d and %d", a, b),
		Number: sum,
	}

	json.NewEncoder(w).Encode(msg)
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
        return
    }

	var req AddRequest

	// Decode JSON body
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

	// Call the custom add function
    sum := customModule.Add(req.A, req.B)

	// Prepare response
    response := map[string]interface{}{
        "text":   "Sum is",
        "number": sum,
    }

	setHeaderToResponse(w)
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/hi", hiHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/addPost", addPostHandler)
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
