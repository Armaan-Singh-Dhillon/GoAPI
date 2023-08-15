package models

import "time"

type Owner struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Offer struct {
	ID        string `json:"_id"`
	Discount  string `json:"discount"`
	OfferName string `json:"offerName"`
}

type Review struct {
	ID       string    `json:"_id"`
	Comment  string    `json:"comment"`
	Date     time.Time `json:"date"`
	Rating   int       `json:"rating"`
	Username string    `json:"username"`
}

type Specifications struct {
	Color           string `json:"Color"`
	Dimensions      string `json:"Dimensions"`
	NumberOfShelves int    `json:"Number of Shelves"`
	Weight          string `json:"Weight"`
}

type Product struct {
	ID               string         `json:"_id"`
	Owner            Owner          `json:"owner"`
	V                int            `json:"__v"`
	AverageRating    float64        `json:"averageRating"`
	Category         string         `json:"category"`
	Company          string         `json:"company"`
	CreatedAt        time.Time      `json:"createdAt"`
	Description      string         `json:"description"`
	Image            string         `json:"image"`
	Name             string         `json:"name"`
	Offers           []Offer        `json:"offers"`
	Price            int            `json:"price"`
	Refundable       string         `json:"refundable"`
	Reviews          []Review       `json:"reviews"`
	Shipping         string         `json:"shipping"`
	Specifications   Specifications `json:"specifications"`
	TermiteResistant string         `json:"termiteResistant"`
	UpdatedAt        time.Time      `json:"updatedAt"`
	Used             string         `json:"used"`
	WoodType         []string       `json:"woodType"`
}

type Order struct {
// to be implemented
}

type User struct {
	ID        string    `json:"_id"`
	Address1  string    `json:"address1"`
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Offers    []Offer   `json:"offers"`
	Orders    []Order   `json:"orders"`
	Phone     string    `json:"phone"`
	Products  []string  `json:"products"`
	UpdatedAt time.Time `json:"updatedAt"`
}