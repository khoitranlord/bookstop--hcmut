// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Book struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Subtitle      *string  `json:"subtitle"`
	Authors       []string `json:"authors"`
	Description   string   `json:"description"`
	PublishedYear int      `json:"publishedYear"`
	ImageURL      *string  `json:"imageUrl"`
	Isbn          *string  `json:"isbn"`
}

type Browse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	ImageURL    *string    `json:"imageUrl"`
	StartedAt   *time.Time `json:"startedAt"`
	EndedAt     *time.Time `json:"endedAt"`
}

type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Href        string    `json:"href"`
	UserID      string    `json:"userId"`
	User        *User     `json:"user"`
	StartedAt   time.Time `json:"startedAt"`
	EndedAt     time.Time `json:"endedAt"`
}

type Exchange struct {
	ID            string    `json:"id"`
	UserBookIDOld string    `json:"userBookIdOld"`
	UserBookOld   *UserBook `json:"userBookOld"`
	UserBookIDNew string    `json:"userBookIdNew"`
	UserBookNew   *UserBook `json:"userBookNew"`
	ExchangedAt   time.Time `json:"exchangedAt"`
}

type Inventory struct {
	ID         string    `json:"id"`
	UserBookID string    `json:"userBookId"`
	UserBook   *UserBook `json:"userBook"`
	LocationID string    `json:"locationId"`
	Location   *Location `json:"location"`
	Removed    bool      `json:"removed"`
}

type InventoryClaim struct {
	ID          string     `json:"id"`
	InventoryID string     `json:"inventoryId"`
	Inventory   *Inventory `json:"inventory"`
	ClaimedAt   time.Time  `json:"claimedAt"`
}

type Location struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	ParentName  *string `json:"parentName"`
	AddressLine string  `json:"addressLine"`
}

type Thought struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    string    `json:"userId"`
	User      *User     `json:"user"`
	BookID    *string   `json:"bookId"`
	Book      *Book     `json:"book"`
}

type User struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description"`
	ProfileImageURL *string   `json:"profileImageUrl"`
	CreatedAt       time.Time `json:"createdAt"`
	Credit          *int      `json:"credit"`
}

type UserBook struct {
	ID                 string  `json:"id"`
	BookID             string  `json:"bookId"`
	Book               *Book   `json:"book"`
	UserID             string  `json:"userId"`
	User               *User   `json:"user"`
	StartedAt          *string `json:"startedAt"`
	EndedAt            *string `json:"endedAt"`
	OriginalUserBookID *string `json:"originalUserBookId"`
}
