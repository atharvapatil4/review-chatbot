package main

type Product struct {
	ProductID          int    `json:"product_id"`
	ProductName        string `json:"product_name"`
	ProductImageURL    string `json:"product_image_url"`
	ProductDescription string `json:"product_description"`
}

type ReactionRequest struct {
	UserID    int    `json:"user_id"`
	ProductID int    `json:"product_id"`
	Reaction  string `json:"reaction"` // Should be 'thumbs_up' or 'thumbs_down', for now
}

type ReactionResponse struct {
	Prompt string `json:"prompt"`
}
