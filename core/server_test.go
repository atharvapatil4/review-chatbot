package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleLoginSuccess(t *testing.T) {
	initDB() // dont mock DB for now
	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(`{"username":"admin","password":"pass"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response UserLoginResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "success" {
		t.Errorf("unexpected response status: got %v want %v", response.Status, "success")
	}
}

func TestHandleLoginFailUser(t *testing.T) {
	initDB() // dont mock DB for now
	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(`{"username":"test","password":"pass"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestHandleLoginFailPass(t *testing.T) {
	initDB() // dont mock DB for now
	req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(`{"username":"admin","password":"wrong_pass"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleLogin)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

}
func TestHandleCreateUserSuccess(t *testing.T) {
	initDB() // dont mock DB for now
	req, err := http.NewRequest("POST", "/createuser", bytes.NewBufferString(`{"username":"testuser","password":"testpassword"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	var response UserLoginResponse
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	if response.Status != "success" {
		t.Errorf("unexpected response status: got %v want %v", response.Status, "success")
	}
	// Cleanup db
	_, err = db.Exec("DELETE FROM users WHERE username=$1", "testuser")
	if err != nil {
		t.Error(err)
	}
}

func TestHandleCreateUserFail(t *testing.T) {
	initDB() // dont mock DB for now
	// Create a user that already exists
	req, err := http.NewRequest("POST", "/createuser", bytes.NewBufferString(`{"username":"admin","password":"pass"}`))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handleCreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestWriteReviewSuccess(t *testing.T) {
	initDB() // dont mock DB for now
	user_id, _ := getFirstUserID()
	product_id, _ := getFirstProductID()
	review := Review{
		UserID:            user_id,
		ProductID:         product_id,
		ReviewDescription: "Test review. Delete me!",
		ReviewDate:        time.Now(),
	}
	payload, _ := json.Marshal(review)
	req, err := http.NewRequest("POST", "/writereview", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(writeReview)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	// Cleanup db
	_, err = db.Exec("DELETE FROM users WHERE username=$1", "testuser")
	if err != nil {
		t.Error(err)
	}
}

func TestWriteReviewFailSerialization(t *testing.T) {
	// Setup
	initDB() // dont mock DB for now
	// Send a malformed JSON payload
	req, err := http.NewRequest("POST", "/writereview", bytes.NewBufferString(`{username":"admin","password":"pass"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(writeReview)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestWriteReviewFailInsert(t *testing.T) {
	initDB() // dont mock DB for now
	product_id, _ := getFirstProductID()
	review := Review{
		UserID:            "random_user_id", // use something not in the users table
		ProductID:         product_id,
		ReviewDescription: "Test review. Delete me!",
		ReviewDate:        time.Now(),
	}
	payload, _ := json.Marshal(review)
	req, err := http.NewRequest("POST", "/writereview", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(writeReview)

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestSelectPromptFailSerialization(t *testing.T) {
	// Setup
	initDB() // dont mock DB for now
	// Send a malformed JSON payload
	req, err := http.NewRequest("POST", "/selectprompt", bytes.NewBufferString(`{username":"admin","password":"pass"}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(selectPrompt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestSelectPromptFail(t *testing.T) {
	initDB() // dont mock DB for now
	prompt := ReactionRequest{
		Reaction: "thumbs_sideways",
	}
	payload, _ := json.Marshal(prompt)
	req, err := http.NewRequest("POST", "/selectprompt", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(selectPrompt)

	handler.ServeHTTP(rr, req)

	if status := rr.Result().StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestSelectPromptSuccess(t *testing.T) {
	initDB() // dont mock DB for now
	prompt := ReactionRequest{
		Reaction: "thumbs_up",
	}
	payload, _ := json.Marshal(prompt)
	req, err := http.NewRequest("POST", "/selectprompt", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(selectPrompt)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

// Utils
func getFirstUserID() (string, error) {
	var userID string
	err := db.QueryRow("SELECT user_id FROM users LIMIT 1").Scan(&userID)
	if err != nil {
		return "", err
	}
	return userID, nil
}

func getFirstProductID() (string, error) {
	var productID string
	err := db.QueryRow("SELECT product_id FROM products LIMIT 1").Scan(&productID)
	if err != nil {
		return "", err
	}
	return productID, nil
}
