package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var db *sql.DB

// Opens a connection to the db
func initDB() {
	var (
		host     = PG_URL
		port     = PG_PORT
		user     = PG_USER
		password = PG_PASS
		dbname   = PG_NAME
	)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	log.Println("connecting to", psqlInfo)
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

// ******* LOGIN/AUTH HANDLERS ****************//
func handleLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	var loginRequest UserLoginRequest

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := UserLoginResponse{Status: "failure", Error: "Invalid request format"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Fetch salt and hashed password for the user
	var salt, storedHashedPassword string
	err := db.QueryRow("SELECT salt, hashed_password FROM users WHERE username=$1", loginRequest.Username).Scan(&salt, &storedHashedPassword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := UserLoginResponse{Status: "failure", Error: "Database error, username does not exist"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Hash the received password with the fetched salt
	hashedPassword := hashPassword(loginRequest.Password, salt)

	// Check the hashed password
	if hashedPassword != storedHashedPassword {
		w.WriteHeader(http.StatusUnauthorized)
		response := UserLoginResponse{Status: "failure", Error: "Incorrect password"}
		json.NewEncoder(w).Encode(response)
		return
	}

	// If you reach here, the login is successful
	w.WriteHeader(http.StatusOK)
	response := UserLoginResponse{Status: "success"}
	json.NewEncoder(w).Encode(response)
}

func handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest UserLoginRequest

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Failed to parse request", http.StatusBadRequest)
		return
	}

	// Check if the username exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username=$1)", userRequest.Username).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if exists {
		response := UserLoginResponse{Status: "failure", Error: "Username already exists"}
		w.WriteHeader(http.StatusConflict) // 409 Conflict
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate a UUID for the user
	userID := uuid.New().String()

	// Generate a random salt
	salt := generateRandomString(10)

	// Hash the password with the salt
	hashedPassword := hashPassword(userRequest.Password, salt)

	// Insert into the database
	_, err = db.Exec("INSERT INTO users (user_id, username, salt, hashed_password) VALUES ($1, $2, $3, $4)",
		userID, userRequest.Username, salt, hashedPassword)

	if err != nil {
		response := UserLoginResponse{Status: "failure", Error: "Database insertion error"}
		w.WriteHeader(http.StatusInternalServerError) // 409 Conflict
		json.NewEncoder(w).Encode(response)
		return
	}

	response := UserLoginResponse{Status: "success", Error: "", UserID: userID}
	json.NewEncoder(w).Encode(response)
}

// ******* SITE ACTION HANDLERS ****************//

// TODO: Handles a purchase. Inserts into TRANSACTIONS table
func handleBuy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bought product!")
}

// Handles a prompt reaction. Returns appropriate prompt from PROMPTS table
func selectPrompt(w http.ResponseWriter, r *http.Request) {
	var requestPayload ReactionRequest
	if err := json.NewDecoder(r.Body).Decode(&requestPayload); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Fetch the prompt for the reaction, .Scan() takes the first row
	var promptMessage string
	err := db.QueryRow("SELECT prompt_message FROM prompts WHERE prompt = $1", requestPayload.Reaction).Scan(&promptMessage)
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

// Receives a text review from a user, and writes it to the REVIEWS table in the db
func writeReview(w http.ResponseWriter, r *http.Request) {
	var review Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO reviews (user_id, product_id, review_description, review_date) VALUES ($1, $2, $3, $4)",
		review.UserID, review.ProductID, review.ReviewDescription, review.ReviewDate)

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
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is Healthy"))
}

func main() {
	// Connect to Postgres
	initDB()
	defer db.Close()

	// Set up routing
	r := mux.NewRouter()
	r.HandleFunc("/health", healthCheck).Methods("GET")

	r.HandleFunc("/login", handleLogin).Methods("POST")
	r.HandleFunc("/buy", handleBuy).Methods("POST")
	r.HandleFunc("/createUser", handleCreateUser).Methods("POST")

	r.HandleFunc("/writereview", writeReview).Methods("POST")
	r.HandleFunc("/selectprompt", selectPrompt).Methods("POST")

	// Set CORS settings

	corsOrigins := handlers.AllowedOrigins([]string{"*"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	corsHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})

	// Launch Server
	log.Println("Server coming up on port", PORT)
	err := http.ListenAndServe(fmt.Sprintf("localhost:%s", PORT), handlers.CORS(corsOrigins, corsMethods, corsHeaders)(r))
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
