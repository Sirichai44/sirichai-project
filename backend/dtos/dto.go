package dtos

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	RoleAdmin = "admin"
	RoleUser  = "user"

	DBCollection = []string{
		DBCollectionUsers,
		DBCollectionListings,
		DBCollectionCategories,
		DBCollectionMessages,
	}

	DBCollectionUsers      = "users"
	DBCollectionListings   = "listings"
	DBCollectionCategories = "categories"
	DBCollectionMessages   = "messages"

	KEY_AUTH_STATUS = "auth_status"
	KEY_SECRET      = "secret_key"
	KEY_USER        = "user"
)

type DtoRegister struct {
	Username        string `json:"username" validate:"required" example:"username"`
	Email           string `json:"email" validate:"required,email" example:"test@mail.com"`
	Password        string `json:"password" validate:"required" example:"P@ssw0rd"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password" example:"P@ssw0rd"`
}

type DtoLogin struct {
	Email    string `json:"email" example:"admin@mail.com" validate:"required,email"`
	Password string `json:"password" example:"P@ssw0rd" validate:"required"`
}

type DtoRole struct {
	Role string `json:"role" validate:"required,oneof=user moderator" example:"moderator"`
	ID   string `json:"id" validate:"required" example:"5f9b3b3b7f7b7b7b7b7b7b7b"`
}

type DtoList struct {
	Title    string   `json:"title" example:"This is a title" validate:"required"`
	Content  string   `json:"content" example:"This is a content" validate:"required,max=5000"`
	Pictures [][]byte `json:"pictures" validate:"omitempty,max=10"`
	Category string   `json:"category" example:"5f9b3b3b7f7b7b7b7b7b7b7b" validate:"required"`
	IsPublic bool     `json:"is_public" example:"true" validate:"required"`
	Hide     bool     `json:"hide" example:"false"`
}

type DtoHide struct {
	Hide bool `json:"hide" example:"true"`
}

type DtoMessage struct {
	SenderID   string `json:"sender_id" validate:"required" example:"5f9b3b3b7f7b7b7b7b7b7b7b"`
	ReceiverID string `json:"receiver_id" validate:"required" example:"5f9b3b3b8f8b8b8b8b8b8b8b"`
	Type       string `json:"type" validate:"required,oneof=list message" example:"list"`
	Content    string `json:"content" validate:"omitempty,max=255" example:"This is a message"`
}

type Register struct {
	ID        primitive.ObjectID `bson:"_id"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Listing struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `bson:"title"`
	Content    string             `bson:"content"`
	Pictures   [][]byte           `bson:"pictures"`
	CategoryID primitive.ObjectID `bson:"category_id"`
	IsPublic   bool               `bson:"is_public"`
	UserID     primitive.ObjectID `bson:"user_id"`
	Hide       bool               `bson:"hide"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type Category struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}

type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type Message struct {
	ID         primitive.ObjectID `bson:"_id"`
	SenderID   primitive.ObjectID `bson:"sender_id"`
	ReceiverID primitive.ObjectID `bson:"receiver_id"`
	Content    string             `bson:"content"`
	Type       string             `bson:"type"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

type User struct {
	Email    string
	Username string
}
