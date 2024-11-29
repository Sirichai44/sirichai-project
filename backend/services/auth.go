package services

import (
	"sirichai-bank/dtos"

	"go.mongodb.org/mongo-driver/mongo"
)

type AuthService interface {
	Service[dtos.Register]
}

// type authService struct {
// 	Service[dtos.Register]
// 	mgc *mongo.Collection
// }
type authService struct {
	Service[dtos.Register]
	mgc *mongo.Collection
}

// func NewAuthService(mongoDB *mongo.Collection) AuthService {
// 	return &authService{
// 		Service: NewService[dtos.Register](mongoDB),
// 		mgc:     mongoDB,
// 	}
// }


