package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var db *sql.DB

// Opens a connection to the db
func initDB() {
	var (
		host     = PG_URL
		port     = PG_PORT
		user     = "postgres" // hardcoded from seed.py
		password = "password" // hardcoded from seed.py
		dbname   = "postgres" // hardcoded from seed.py
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	println(psqlInfo)
	var err error
	// Open with postgres driver
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Handles Auth request
func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logged in!")
}

// TODO: Handles a purchase. Inserts into TRANSACTIONS table
func handleBuy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bought product!")
}

// Handles a reaction. Returns appropriate prompt from PROMPTS table
func handleReaction(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Insert a new row into user_products with the reaction
	_, err := db.Exec("INSERT INTO products (user_id, product_id, reaction) VALUES ($1, $2, $3)",
		requestPayload.UserID, requestPayload.ProductID, requestPayload.Reaction)
	if err != nil {
		http.Error(w, "Failed to insert reaction", http.StatusInternalServerError)
		return
	}

	// Fetch the prompt for the reaction, .Scan() takes the first row
	var promptMessage string
	err = db.QueryRow("SELECT prompt_message FROM prompts WHERE prompt = $1", requestPayload.Reaction).Scan(&promptMessage)
	if err != nil {
		http.Error(w, "Failed to retrieve prompt", http.StatusInternalServerError)
		return
	}

	responsePayload := ReactionResponse{
		Prompt: promptMessage,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responsePayload)
}

// Receives a text review from a user, and writes it to the PRODUCTS table in the db
func handleResponse(w http.ResponseWriter, r *http.Request) {
	var response struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// For demonstration, let's just insert the review text into a hypothetical `reviews` table.
	_, err := db.Exec("INSERT INTO reviews (review_text) VALUES ($1)", response.Text)
	if err != nil {
		http.Error(w, "Failed to store review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("health check hit")
}

func main() {

	initDB()
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")
	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/buy", handleBuy).Methods("POST")
	r.HandleFunc("/response", handleResponse).Methods("POST")
	r.HandleFunc("/reaction", handleReaction).Methods("POST")

	log.Println("Server coming up on port", PORT)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", PORT), r)
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
