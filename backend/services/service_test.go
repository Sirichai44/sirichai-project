package services

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
// )

// type ServiceMock struct {
// 	ID   primitive.ObjectID `bson:"_id" json:"id"`
// 	Name string             `bson:"name" json:"name"`
// }

// func TestFindALL(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		expected := []ServiceMock{
// 			{ID: primitive.NewObjectID(), Name: "service.mocking"},
// 			{ID: primitive.NewObjectID(), Name: "service.mocking"},
// 		}

// 		first := mtest.CreateCursorResponse(1, "mock.new.service", mtest.FirstBatch, bson.D{
// 			{Key: "_id", Value: expected[0].ID}, {Key: "name", Value: expected[0].Name},
// 		})

// 		second := mtest.CreateCursorResponse(0, "mock.new.service", mtest.NextBatch, bson.D{
// 			{Key: "_id", Value: expected[1].ID}, {Key: "name", Value: expected[1].Name},
// 		})

// 		killCursors := mtest.CreateCursorResponse(0, "mock.new.service", mtest.NextBatch)
// 		mt.AddMockResponses(first, second, killCursors)

// 		res, err := srv.FindAll(context.Background())
// 		assert.Nil(t, err)
// 		assert.Equal(t, expected, res)
// 	})

// 	mt.Run("Error", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
// 			Message: "mock error",
// 		}))

// 		res, err := srv.FindAll(context.Background())
// 		assert.NotNil(t, err)
// 		assert.Nil(t, res)
// 	})
// }

// func TestFindOneKeyValue(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		expected := ServiceMock{ID: primitive.NewObjectID(), Name: "mock.new.service"}

// 		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mock.new.service", mtest.FirstBatch, bson.D{
// 			{Key: "_id", Value: expected.ID}, {Key: "name", Value: expected.Name},
// 		}))

// 		res, err := srv.FindOneKeyValue(context.Background(), "name", expected.Name)
// 		assert.Nil(t, err)
// 		assert.Equal(t, &expected, res)
// 	})

// 	mt.Run("Error", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
// 			Message: "mock error",
// 		}))

// 		res, err := srv.FindOneKeyValue(context.Background(), "name", "mock.new.service")
// 		assert.NotNil(t, err)
// 		assert.Nil(t, res)
// 	})
// }

// func TestFindByKeyValue(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		expected := []ServiceMock{
// 			{ID: primitive.NewObjectID(), Name: "service.mocking"},
// 			{ID: primitive.NewObjectID(), Name: "service.mocking"},
// 		}

// 		first := mtest.CreateCursorResponse(1, "mock.new.service", mtest.FirstBatch, bson.D{
// 			{Key: "_id", Value: expected[0].ID}, {Key: "name", Value: expected[0].Name},
// 		})

// 		second := mtest.CreateCursorResponse(0, "mock.new.service", mtest.NextBatch, bson.D{
// 			{Key: "_id", Value: expected[1].ID}, {Key: "name", Value: expected[1].Name},
// 		})

// 		killCursors := mtest.CreateCursorResponse(0, "mock.new.service", mtest.NextBatch)
// 		mt.AddMockResponses(first, second, killCursors)

// 		res, err := srv.FindByKeyValue(context.Background(), "name", expected[0].Name)
// 		assert.Nil(t, err)
// 		assert.Equal(t, expected, res)
// 	})

// 	mt.Run("Error", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
// 			Message: "mock error",
// 		}))

// 		res, err := srv.FindByKeyValue(context.Background(), "name", "mock.new.service")
// 		assert.NotNil(t, err)
// 		assert.Nil(t, res)
// 	})
// }

// func TestFindOneAndUpdate(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {

// 	})

// 	mt.Run("NoDocument", func(mt *mtest.T) {
// 			srv := NewService[ServiceMock](mt.Coll)

// 			// Mock response for no document found
// 			mt.AddMockResponses(mtest.CreateCursorResponse(0, "mock.new.service", mtest.FirstBatch))

// 			filter := bson.D{{Key: "_id", Value: primitive.NewObjectID()}}
// 			update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "mock.new.service.update"}}}}
// 			err := srv.FindOneAndUpdate(context.Background(), filter, update)
// 			assert.NotNil(t, err)
// 			assert.Equal(t, "mongo: no documents in result", err.Error())
// 	})
// }

// func TestDeleteOne(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		expected := ServiceMock{ID: primitive.NewObjectID(), Name: "mock.new.service"}

// 		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mock.new.service", mtest.FirstBatch, bson.D{
// 			{Key: "_id", Value: expected.ID}, {Key: "name", Value: expected.Name},
// 		}))

// 		filter := bson.D{{Key: "_id", Value: expected.ID}}
// 		err := srv.DeleteOne(context.Background(), filter)
// 		assert.Nil(t, err)
// 	})

// 	mt.Run("Error", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
// 			Message: "mock error",
// 		}))

// 		err := srv.DeleteOne(context.Background(), bson.D{})
// 		assert.NotNil(t, err)
// 	})
// }

// func TestInsertOne(t *testing.T) {
// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

// 	mt.Run("Success", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		expected := ServiceMock{ID: primitive.NewObjectID(), Name: "mock.new.service"}

// 		mt.AddMockResponses(mtest.CreateCursorResponse(1, "mock.new.service", mtest.FirstBatch, bson.D{
// 			{Key: "_id", Value: expected.ID}, {Key: "name", Value: expected.Name},
// 		}))

// 		err := srv.InsertOne(context.Background(), expected)
// 		assert.Nil(t, err)
// 	})

// 	mt.Run("Error", func(mt *mtest.T) {
// 		srv := NewService[ServiceMock](mt.Coll)

// 		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
// 			Message: "mock error",
// 		}))

// 		err := srv.InsertOne(context.Background(), ServiceMock{})
// 		assert.NotNil(t, err)
// 	})
// }
