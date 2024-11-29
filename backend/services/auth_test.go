package services

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// func TestNewAuthService(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		collection := mt.Coll
// 		srvAuth := NewAuthService(collection)

// 		assert.NotNil(t, srvAuth)
// 		assert.IsType(t, &authService{}, srvAuth)
// 	})
// }
