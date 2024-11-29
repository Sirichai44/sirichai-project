package apis

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"sirichai-bank/apis/middleware"
// 	"sirichai-bank/dtos"
// 	"sirichai-bank/services"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"golang.org/x/crypto/bcrypt"
// )

// type auth struct {
// 	srv     services.AuthService
// 	signKey string
// }

// func NewHandlerAuth(f fiber.Router, srv services.AuthService, signKey string) {
// 	auth := auth{srv: srv, signKey: signKey}

// 	g := f.Group(apiVersion).Group("auth")
// 	// g.Get("/test", func(c *fiber.Ctx) error {
// 	// 	return c.JSON(fiber.Map{
// 	// 		"status":  fiber.StatusOK,
// 	// 		"message": "OK",
// 	// 		"results": "Auth API is working",
// 	// 	})
// 	// })

// 	g.Post("/register", HandleBodyParser(auth.Register))
// 	g.Post("/login", HandleBodyParser(auth.Login))
// 	g.Post("/role", middleware.IsAuthenticated, HandleBodyParser(auth.Role))
// }

// // Register godoc
// // @Summary Register a new user
// // @Description Register a new user for the application
// // @Tags auth
// // @Accept  json
// // @Produce  json
// // @Param register body dtos.DtoRegister true "Register"
// // @Success 201 {object} dtos.Context{status=int,message=string,results=dtos.RegisterResponse}
// // @Failure 400 {object} error "Bad Request"
// // @Failure 406 {object} error "Not Acceptable"
// // @Router /api/v1/auth/register [post]
// func (a *auth) Register(dto dtos.DtoRegister, c *fiber.Ctx) (*dtos.Context, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	findEmail, err := a.srv.FindOneKeyValue(context.Background(), "email", dto.Email)
// 	if err != nil && err != mongo.ErrNoDocuments {
// 		log.Println("Register", "FindOneKeyValue", "failed")
// 		return nil, err
// 	}

// 	if findEmail != nil {
// 		log.Println("Register", "FindOneKeyValue", "email already exists")
// 		return nil, fiber.NewError(fiber.StatusBadRequest, "Email already exists")
// 	}

// 	passHash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		log.Println("Register", "GenerateFromPassword", "failed")
// 		return nil, fiber.ErrNotAcceptable
// 	}

// 	rgt := dtos.Register{
// 		ID:        primitive.NewObjectID(),
// 		Username:  dto.Username,
// 		Email:     dto.Email,
// 		Password:  string(passHash),
// 		Role:      dtos.RoleUser,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	if err := a.srv.InsertOne(ctx, rgt); err != nil {
// 		log.Println("Register", "Register", "failed")
// 		return nil, err
// 	}

// 	token, err := genToken(rgt.Username, rgt.Email, a.signKey)
// 	if err != nil {
// 		log.Println("Register", "genToken", "failed")
// 		return nil, err
// 	}

// 	log.Println("Register", rgt.ID.Hex(), "successfully")

// 	return NewContext(fiber.StatusCreated, "register successfully", fiber.Map{
// 		"token": token,
// 	}), nil
// }

// // Login godoc
// // @Summary Login a user
// // @Description Login a user to the application
// // @Tags auth
// // @Accept  json
// // @Produce  json
// // @Param register body dtos.DtoLogin true "Login"
// // @Success 200 {object} dtos.Context{status=int,message=string,results=dtos.RegisterResponse}
// // @Failure 400 {object} error "Bad Request"
// // @Failure 406 {object} error "Not Acceptable"
// // @Router /api/v1/auth/login [post]
// func (a *auth) Login(dto dtos.DtoLogin, c *fiber.Ctx) (*dtos.Context, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	rgt, err := a.srv.FindOneKeyValue(ctx, "email", dto.Email)
// 	if err != nil && err != mongo.ErrNoDocuments {
// 		log.Println("Login", "FindOneKeyValue", "failed", err)
// 		return nil, err
// 	}

// 	if rgt == nil {
// 		log.Println("Login", "FindOneKeyValue", "email not found")
// 		return nil, fiber.NewError(fiber.StatusBadRequest, "Email not found")
// 	}

// 	if err := bcrypt.CompareHashAndPassword([]byte(rgt.Password), []byte(dto.Password)); err != nil {
// 		log.Println("Login", "CompareHashAndPassword", "failed")
// 		return nil, fiber.NewError(fiber.StatusBadRequest, "Password is incorrect")
// 	}

// 	token, err := genToken(rgt.Username, rgt.Email, a.signKey)
// 	if err != nil {
// 		log.Println("Login", "genToken", "failed")
// 		return nil, err
// 	}

// 	filter := primitive.D{primitive.E{Key: "_id", Value: rgt.ID}}
// 	update := primitive.D{primitive.E{Key: "$set", Value: primitive.D{primitive.E{Key: "updated_at", Value: time.Now()}}}}
// 	if err = a.srv.FindOneAndUpdate(ctx, filter, update); err != nil {
// 		log.Println("Login", "FindOneAndUpdate", "failed")
// 		return nil, err
// 	}

// 	log.Println("Login", rgt.ID.Hex(), "successfully")

// 	return NewContext(fiber.StatusOK, "login successfully", fiber.Map{
// 		"token": token,
// 	}), nil
// }

// // ChangeRole godoc
// // @Summary Change a user role
// // @Description Change a user role to moderator or user
// // @Tags auth
// // @Accept  json
// // @Produce  json
// // @Security ApiKeyAuth
// // @Param Authorization header string true "Bearer token"
// // @Param Role body dtos.DtoRole true "Role"
// // @Success 200 {object} dtos.Context{status=int,message=string,results=nil}
// // @Failure 400 {object} error "Bad Request"
// // @Failure 406 {object} error "Not Acceptable"
// // @Router /api/v1/auth/role [post]
// func (a *auth) Role(dto dtos.DtoRole, c *fiber.Ctx) (*dtos.Context, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	user := c.Locals(dtos.KEY_USER).(*dtos.User)
// 	userInfo, err := a.srv.FindOneKeyValue(ctx, "email", user.Email)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			log.Println("Role", "FindOneKeyValue", "no user found")
// 		} else {
// 			log.Println("Role", "FindOneKeyValue", "failed")
// 		}
// 		return nil, err
// 	}

// 	if userInfo.Role != dtos.RoleAdmin {
// 		log.Println("Role", "not an admin")
// 		return nil, fiber.NewError(fiber.StatusUnauthorized, "You are not an admin")
// 	}

// 	find, err := a.srv.FindOneKeyValue(ctx, "_id", dto.ID)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			log.Println("Role", "FindOneKeyValue", "no user found")
// 			return nil, fiber.NewError(fiber.StatusBadRequest, "User not found")
// 		} else {
// 			log.Println("Role", "FindOneKeyValue", "failed")
// 			return nil, err
// 		}
// 	}

// 	if find.Role == dto.Role {
// 		log.Println("Role", "FindOneKeyValue", "role already exists")
// 		return nil, fiber.NewError(fiber.StatusBadRequest, "Role already exists")
// 	}

// 	if find.Role == dtos.RoleAdmin && find.ID.Hex() == userInfo.ID.Hex() {
// 		log.Println("Role", "FindOneKeyValue", "cannot change own role")
// 		return nil, fiber.NewError(fiber.StatusBadRequest, "Cannot change own role")
// 	}

// 	filter := primitive.D{primitive.E{Key: "_id", Value: find.ID}}
// 	update := primitive.D{primitive.E{Key: "$set", Value: primitive.D{primitive.E{Key: "role", Value: dto.Role}}}}
// 	if err := a.srv.FindOneAndUpdate(ctx, filter, update); err != nil {
// 		log.Println("Role", "FindOneAndUpdate", "failed")
// 		return nil, err
// 	}

// 	log.Println("Role updated", dto.Role, find.ID.Hex(), "successfully")

// 	return NewContext(fiber.StatusOK, "role updated successfully", nil), nil
// }

// func genToken(username, email, signKey string) (string, error) {
// 	claims := dtos.Claims{
// 		Email:    email,
// 		Username: username,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	tokenString, err := token.SignedString([]byte(signKey))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
