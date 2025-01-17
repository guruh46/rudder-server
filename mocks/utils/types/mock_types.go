// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rudderlabs/rudder-server/utils/types (interfaces: UserSuppression,Reporting)
//
// Generated by this command:
//
//	mockgen -destination=../../mocks/utils/types/mock_types.go -package mock_types github.com/rudderlabs/rudder-server/utils/types UserSuppression,Reporting
//

// Package mock_types is a generated GoMock package.
package mock_types

import (
	context "context"
	reflect "reflect"

	model "github.com/rudderlabs/rudder-server/enterprise/suppress-user/model"
	tx "github.com/rudderlabs/rudder-server/utils/tx"
	types "github.com/rudderlabs/rudder-server/utils/types"
	gomock "go.uber.org/mock/gomock"
)

// MockUserSuppression is a mock of UserSuppression interface.
type MockUserSuppression struct {
	ctrl     *gomock.Controller
	recorder *MockUserSuppressionMockRecorder
	isgomock struct{}
}

// MockUserSuppressionMockRecorder is the mock recorder for MockUserSuppression.
type MockUserSuppressionMockRecorder struct {
	mock *MockUserSuppression
}

// NewMockUserSuppression creates a new mock instance.
func NewMockUserSuppression(ctrl *gomock.Controller) *MockUserSuppression {
	mock := &MockUserSuppression{ctrl: ctrl}
	mock.recorder = &MockUserSuppressionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserSuppression) EXPECT() *MockUserSuppressionMockRecorder {
	return m.recorder
}

// GetSuppressedUser mocks base method.
func (m *MockUserSuppression) GetSuppressedUser(workspaceID, userID, sourceID string) *model.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSuppressedUser", workspaceID, userID, sourceID)
	ret0, _ := ret[0].(*model.Metadata)
	return ret0
}

// GetSuppressedUser indicates an expected call of GetSuppressedUser.
func (mr *MockUserSuppressionMockRecorder) GetSuppressedUser(workspaceID, userID, sourceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuppressedUser", reflect.TypeOf((*MockUserSuppression)(nil).GetSuppressedUser), workspaceID, userID, sourceID)
}

// MockReporting is a mock of Reporting interface.
type MockReporting struct {
	ctrl     *gomock.Controller
	recorder *MockReportingMockRecorder
	isgomock struct{}
}

// MockReportingMockRecorder is the mock recorder for MockReporting.
type MockReportingMockRecorder struct {
	mock *MockReporting
}

// NewMockReporting creates a new mock instance.
func NewMockReporting(ctrl *gomock.Controller) *MockReporting {
	mock := &MockReporting{ctrl: ctrl}
	mock.recorder = &MockReportingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReporting) EXPECT() *MockReportingMockRecorder {
	return m.recorder
}

// DatabaseSyncer mocks base method.
func (m *MockReporting) DatabaseSyncer(c types.SyncerConfig) types.ReportingSyncer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseSyncer", c)
	ret0, _ := ret[0].(types.ReportingSyncer)
	return ret0
}

// DatabaseSyncer indicates an expected call of DatabaseSyncer.
func (mr *MockReportingMockRecorder) DatabaseSyncer(c any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseSyncer", reflect.TypeOf((*MockReporting)(nil).DatabaseSyncer), c)
}

// Report mocks base method.
func (m *MockReporting) Report(ctx context.Context, metrics []*types.PUReportedMetric, tx *tx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Report", ctx, metrics, tx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Report indicates an expected call of Report.
func (mr *MockReportingMockRecorder) Report(ctx, metrics, tx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Report", reflect.TypeOf((*MockReporting)(nil).Report), ctx, metrics, tx)
}

// Stop mocks base method.
func (m *MockReporting) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockReportingMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockReporting)(nil).Stop))
}
