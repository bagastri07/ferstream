// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kumparan/ferstream (interfaces: JetStream)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	nats "github.com/nats-io/nats.go"
)

// MockJetStream is a mock of JetStream interface.
type MockJetStream struct {
	ctrl     *gomock.Controller
	recorder *MockJetStreamMockRecorder
}

// MockJetStreamMockRecorder is the mock recorder for MockJetStream.
type MockJetStreamMockRecorder struct {
	mock *MockJetStream
}

// NewMockJetStream creates a new mock instance.
func NewMockJetStream(ctrl *gomock.Controller) *MockJetStream {
	mock := &MockJetStream{ctrl: ctrl}
	mock.recorder = &MockJetStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJetStream) EXPECT() *MockJetStreamMockRecorder {
	return m.recorder
}

// AddStream mocks base method.
func (m *MockJetStream) AddStream(arg0 *nats.StreamConfig, arg1 ...nats.JSOpt) (*nats.StreamInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddStream", varargs...)
	ret0, _ := ret[0].(*nats.StreamInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStream indicates an expected call of AddStream.
func (mr *MockJetStreamMockRecorder) AddStream(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStream", reflect.TypeOf((*MockJetStream)(nil).AddStream), varargs...)
}

// ConsumerInfo mocks base method.
func (m *MockJetStream) ConsumerInfo(arg0, arg1 string, arg2 ...nats.JSOpt) (*nats.ConsumerInfo, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConsumerInfo", varargs...)
	ret0, _ := ret[0].(*nats.ConsumerInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConsumerInfo indicates an expected call of ConsumerInfo.
func (mr *MockJetStreamMockRecorder) ConsumerInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumerInfo", reflect.TypeOf((*MockJetStream)(nil).ConsumerInfo), varargs...)
}

// GetNATSConnection mocks base method.
func (m *MockJetStream) GetNATSConnection() *nats.Conn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNATSConnection")
	ret0, _ := ret[0].(*nats.Conn)
	return ret0
}

// GetNATSConnection indicates an expected call of GetNATSConnection.
func (mr *MockJetStreamMockRecorder) GetNATSConnection() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNATSConnection", reflect.TypeOf((*MockJetStream)(nil).GetNATSConnection))
}

// Publish mocks base method.
func (m *MockJetStream) Publish(arg0 string, arg1 []byte, arg2 ...nats.PubOpt) (*nats.PubAck, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*nats.PubAck)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockJetStreamMockRecorder) Publish(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockJetStream)(nil).Publish), varargs...)
}

// QueueSubscribe mocks base method.
func (m *MockJetStream) QueueSubscribe(arg0, arg1 string, arg2 nats.MsgHandler, arg3 ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueueSubscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueueSubscribe indicates an expected call of QueueSubscribe.
func (mr *MockJetStreamMockRecorder) QueueSubscribe(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueSubscribe", reflect.TypeOf((*MockJetStream)(nil).QueueSubscribe), varargs...)
}

// Subscribe mocks base method.
func (m *MockJetStream) Subscribe(arg0 string, arg1 nats.MsgHandler, arg2 ...nats.SubOpt) (*nats.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(*nats.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockJetStreamMockRecorder) Subscribe(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockJetStream)(nil).Subscribe), varargs...)
}
