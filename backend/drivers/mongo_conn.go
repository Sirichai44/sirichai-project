package drivers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBClient struct {
	*mongo.Client
	*mongo.Database
}

// func MongoDBConn(opt config.DBConfig) (*MongoDBClient, error) {
// 	opts := options.Client()
// 	opts.ApplyURI(opt.URL)
// 	opts.SetAuth(options.Credential{
// 		Username: opt.DBUser,
// 		Password: opt.DBPassword,
// 	})

	

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	client, err := mongo.Connect(ctx, opts)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := client.Ping(ctx, nil); err != nil {
// 		return nil, err
// 	}

// 	dbNames, _ := client.ListDatabaseNames(ctx, nil)

// 	dbExists := false
// 	for _, dbName := range dbNames {
// 		if dbName == opt.DBName {
// 			dbExists = true
// 			break
// 		}
// 	}

// 	dbReturn := &MongoDBClient{Client: client, Database: client.Database(opt.DBName)}
// 	if !dbExists {
// 		log.Println("Database not found, creating ...")
// 		dbRegister(ctx, dbReturn, opt)
// 	}

// 	return dbReturn, nil
// }

// func dbRegister(ctx context.Context, client *MongoDBClient, opt config.DBConfig) {
// 	clt := dtos.DBCollection

// 	for _, c := range clt {
// 		if err := client.Database.CreateCollection(ctx, c); err != nil {
// 			log.Printf("Failed to create collection %s: %v", c, err)
// 		}

// 		log.Printf("Collection %s created", c)

// 	}

// 	err := client.Database.Collection(dtos.DBCollectionUsers).FindOne(ctx, primitive.D{primitive.E{Key: "email", Value: opt.AdminEmail}})
// 	if err.Err() == nil {
// 		log.Printf("Admin has been created")
// 		return
// 	}

// 	// create admin
// 	createAdmin(ctx, client, opt.AdminEmail, opt.AdminPassword)

// 	// create categories
// 	initCategories(ctx, client)
// }

// func createAdmin(ctx context.Context, client *MongoDBClient, email, pass string) {
// 	// create admin
// 	adminID, _ := primitive.ObjectIDFromHex("000000000000000000000000")
// 	passHash, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)

// 	admin := dtos.Register{
// 		ID:        adminID,
// 		Username:  "admin",
// 		Email:     email,
// 		Password:  string(passHash),
// 		Role:      dtos.RoleAdmin,
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	_, err := client.Database.Collection(dtos.DBCollectionUsers).InsertOne(ctx, admin)
// 	if err != nil {
// 		log.Printf("Failed to create admin: %v", err)
// 		return
// 	}

// 	log.Printf("Initial admin created")
// }

// func initCategories(ctx context.Context, client *MongoDBClient) {
// 	categories := []dtos.Category{
// 		{
// 			ID:   primitive.NewObjectID(),
// 			Name: "Sedan",
// 		},
// 		{
// 			ID:   primitive.NewObjectID(),
// 			Name: "SUV",
// 		},
// 	}

// 	for _, c := range categories {
// 		_, err := client.Database.Collection(dtos.DBCollectionCategories).InsertOne(ctx, c)
// 		if err != nil {
// 			log.Printf("Failed to create category: %v", err)
// 			return
// 		}
// 	}

// 	log.Printf("Initial categories created")
// }
