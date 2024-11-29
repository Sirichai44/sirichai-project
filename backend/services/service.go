package services

import (
	"cloud.google.com/go/firestore"
)

type Service[T any] interface {
	// FindAll(ctx context.Context, filter ...primitive.E) ([]T, error)
	// FindOneKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) (*T, error)
	// FindByKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) ([]T, error)
	// FindOneAndUpdate(ctx context.Context, filter primitive.D, update primitive.D) error
	// DeleteOne(ctx context.Context, filter primitive.D) error
	// InsertOne(ctx context.Context, t T) error
}

// type mgc[T any] struct {
// 	mgc *mongo.Collection
// }
type fbc[T any] struct {
	fbc *firestore.CollectionRef
}

// func NewService[T any](mongoConn *mongo.Collection) Service[T] {
// 	return &mgc[T]{
// 		mgc: mongoConn,
// 	}
// }
func NewService[T any](firebaseConn *firestore.CollectionRef) Service[T] {
	return &fbc[T]{
		fbc: firebaseConn,
	}
}

// func (m *mgc[T]) FindAll(ctx context.Context, filter ...primitive.E) ([]T, error) {
// 	var results []T
// 	ft :=  primitive.E{}
// 	if len(filter) > 0 {
// 		ft = filter[0]
// 	}

// 	cursor, err := m.mgc.Find(ctx, ft)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func (m *mgc[T]) FindOneKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) (*T, error) {
// 	filter := primitive.D{primitive.E{Key: key, Value: value}}
// 	filter = append(filter, opts...)

// 	var t T
// 	if err := m.mgc.FindOne(ctx, filter, options.FindOne().SetSort(bson.M{"created_at": -1})).Decode(&t); err != nil {
// 		return nil, err
// 	}

// 	return &t, nil
// }

// func (m *mgc[T]) FindByKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) ([]T, error) {
// 	filter := primitive.D{primitive.E{Key: key, Value: value}}
// 	filter = append(filter, opts...)

// 	var results []T
// 	cursor, err := m.mgc.Find(ctx, filter)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if err := cursor.All(ctx, &results); err != nil {
// 		return nil, err
// 	}

// 	return results, nil
// }

// func (m *mgc[T]) FindOneAndUpdate(ctx context.Context, filter primitive.D, update primitive.D) error {
// 	var t T
// 	if err := m.mgc.FindOneAndUpdate(ctx, filter, update).Decode(&t); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *mgc[T]) DeleteOne(ctx context.Context, filter primitive.D) error {
// 	if _, err := m.mgc.DeleteOne(ctx, filter); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (m *mgc[T]) InsertOne(ctx context.Context, t T) error {
// 	if _, err := m.mgc.InsertOne(ctx, t); err != nil {
// 		return err
// 	}

// 	return nil
// }
