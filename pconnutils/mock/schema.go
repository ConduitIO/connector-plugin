// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-connector-protocol/pconnutils (interfaces: SchemaService)
//
// Generated by this command:
//
//	mockgen -typed -destination=mock/schema.go -package=mock -mock_names=SchemaService=SchemaService . SchemaService
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	pconnutils "github.com/conduitio/conduit-connector-protocol/pconnutils"
	gomock "go.uber.org/mock/gomock"
)

// SchemaService is a mock of SchemaService interface.
type SchemaService struct {
	ctrl     *gomock.Controller
	recorder *SchemaServiceMockRecorder
	isgomock struct{}
}

// SchemaServiceMockRecorder is the mock recorder for SchemaService.
type SchemaServiceMockRecorder struct {
	mock *SchemaService
}

// NewSchemaService creates a new mock instance.
func NewSchemaService(ctrl *gomock.Controller) *SchemaService {
	mock := &SchemaService{ctrl: ctrl}
	mock.recorder = &SchemaServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *SchemaService) EXPECT() *SchemaServiceMockRecorder {
	return m.recorder
}

// CreateSchema mocks base method.
func (m *SchemaService) CreateSchema(arg0 context.Context, arg1 pconnutils.CreateSchemaRequest) (pconnutils.CreateSchemaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchema", arg0, arg1)
	ret0, _ := ret[0].(pconnutils.CreateSchemaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSchema indicates an expected call of CreateSchema.
func (mr *SchemaServiceMockRecorder) CreateSchema(arg0, arg1 any) *SchemaServiceCreateSchemaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchema", reflect.TypeOf((*SchemaService)(nil).CreateSchema), arg0, arg1)
	return &SchemaServiceCreateSchemaCall{Call: call}
}

// SchemaServiceCreateSchemaCall wrap *gomock.Call
type SchemaServiceCreateSchemaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SchemaServiceCreateSchemaCall) Return(arg0 pconnutils.CreateSchemaResponse, arg1 error) *SchemaServiceCreateSchemaCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SchemaServiceCreateSchemaCall) Do(f func(context.Context, pconnutils.CreateSchemaRequest) (pconnutils.CreateSchemaResponse, error)) *SchemaServiceCreateSchemaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SchemaServiceCreateSchemaCall) DoAndReturn(f func(context.Context, pconnutils.CreateSchemaRequest) (pconnutils.CreateSchemaResponse, error)) *SchemaServiceCreateSchemaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSchema mocks base method.
func (m *SchemaService) GetSchema(arg0 context.Context, arg1 pconnutils.GetSchemaRequest) (pconnutils.GetSchemaResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSchema", arg0, arg1)
	ret0, _ := ret[0].(pconnutils.GetSchemaResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema.
func (mr *SchemaServiceMockRecorder) GetSchema(arg0, arg1 any) *SchemaServiceGetSchemaCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchema", reflect.TypeOf((*SchemaService)(nil).GetSchema), arg0, arg1)
	return &SchemaServiceGetSchemaCall{Call: call}
}

// SchemaServiceGetSchemaCall wrap *gomock.Call
type SchemaServiceGetSchemaCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *SchemaServiceGetSchemaCall) Return(arg0 pconnutils.GetSchemaResponse, arg1 error) *SchemaServiceGetSchemaCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *SchemaServiceGetSchemaCall) Do(f func(context.Context, pconnutils.GetSchemaRequest) (pconnutils.GetSchemaResponse, error)) *SchemaServiceGetSchemaCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *SchemaServiceGetSchemaCall) DoAndReturn(f func(context.Context, pconnutils.GetSchemaRequest) (pconnutils.GetSchemaResponse, error)) *SchemaServiceGetSchemaCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
