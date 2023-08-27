package main

import "time"

type Product struct {
	ProductID          int    `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductImageURL    string `json:"product_image_url"`
	ProductDescription string `json:"product_description"`
}

type ReactionRequest struct {
	Reaction string `json:"reaction"` // Should be 'thumbs_up' or 'thumbs_down', for now
}

type ReactionResponse struct {
	Prompt string `json:"prompt"`
}

type Review struct {
	UserID            string    `json:"user_id"`
	ProductID         string    `json:"product_id"`
	ReviewDescription string    `json:"review_description"`
	ReviewDate        time.Time `json:"review_date"`
}

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Status string `json:"status"`
	UserID string `json:"user_id,omitempty"`
	Error  string `json:"error,omitempty"`
}
