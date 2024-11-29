// Code generated by MockGen. DO NOT EDIT.
// Source: services/message.go
//
// Generated by this command:
//
//	mockgen -source=services/message.go -destination=services/mocks/mock_message_service.go -package=mocks MessageService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	dtos "sirichai-bank/dtos"
	reflect "reflect"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	gomock "go.uber.org/mock/gomock"
)

// MockMessageService is a mock of MessageService interface.
type MockMessageService struct {
	ctrl     *gomock.Controller
	recorder *MockMessageServiceMockRecorder
}

// MockMessageServiceMockRecorder is the mock recorder for MockMessageService.
type MockMessageServiceMockRecorder struct {
	mock *MockMessageService
}

// NewMockMessageService creates a new mock instance.
func NewMockMessageService(ctrl *gomock.Controller) *MockMessageService {
	mock := &MockMessageService{ctrl: ctrl}
	mock.recorder = &MockMessageServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageService) EXPECT() *MockMessageServiceMockRecorder {
	return m.recorder
}

// DeleteOne mocks base method.
func (m *MockMessageService) DeleteOne(ctx context.Context, filter primitive.D) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", ctx, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOne indicates an expected call of DeleteOne.
func (mr *MockMessageServiceMockRecorder) DeleteOne(ctx, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockMessageService)(nil).DeleteOne), ctx, filter)
}

// FindAll mocks base method.
func (m *MockMessageService) FindAll(ctx context.Context, filter ...primitive.E) ([]dtos.Message, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx}
	for _, a := range filter {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindAll", varargs...)
	ret0, _ := ret[0].([]dtos.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockMessageServiceMockRecorder) FindAll(ctx any, filter ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx}, filter...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMessageService)(nil).FindAll), varargs...)
}

// FindByKeyValue mocks base method.
func (m *MockMessageService) FindByKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) ([]dtos.Message, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, value}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindByKeyValue", varargs...)
	ret0, _ := ret[0].([]dtos.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByKeyValue indicates an expected call of FindByKeyValue.
func (mr *MockMessageServiceMockRecorder) FindByKeyValue(ctx, key, value any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, value}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByKeyValue", reflect.TypeOf((*MockMessageService)(nil).FindByKeyValue), varargs...)
}

// FindOneAndUpdate mocks base method.
func (m *MockMessageService) FindOneAndUpdate(ctx context.Context, filter, update primitive.D) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneAndUpdate", ctx, filter, update)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindOneAndUpdate indicates an expected call of FindOneAndUpdate.
func (mr *MockMessageServiceMockRecorder) FindOneAndUpdate(ctx, filter, update any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneAndUpdate", reflect.TypeOf((*MockMessageService)(nil).FindOneAndUpdate), ctx, filter, update)
}

// FindOneKeyValue mocks base method.
func (m *MockMessageService) FindOneKeyValue(ctx context.Context, key string, value any, opts ...primitive.E) (*dtos.Message, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, key, value}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOneKeyValue", varargs...)
	ret0, _ := ret[0].(*dtos.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneKeyValue indicates an expected call of FindOneKeyValue.
func (mr *MockMessageServiceMockRecorder) FindOneKeyValue(ctx, key, value any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, key, value}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneKeyValue", reflect.TypeOf((*MockMessageService)(nil).FindOneKeyValue), varargs...)
}

// InsertOne mocks base method.
func (m *MockMessageService) InsertOne(ctx context.Context, t dtos.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockMessageServiceMockRecorder) InsertOne(ctx, t any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockMessageService)(nil).InsertOne), ctx, t)
}