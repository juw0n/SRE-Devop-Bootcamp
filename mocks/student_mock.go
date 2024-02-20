// Code generated by MockGen. DO NOT EDIT.
// Source: /home/juwon/Desktop/cloudComputingLessons/SRE-Devop-Bootcamp/database/sqlc/querier.go
//
// Generated by this command:
//
//	mockgen -source=/home/juwon/Desktop/cloudComputingLessons/SRE-Devop-Bootcamp/database/sqlc/querier.go -destination=mocks/student_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	sqlc "github.com/juw0n/SRE-Devop-Bootcamp/database/sqlc"
	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateCourse mocks base method.
func (m *MockQuerier) CreateCourse(ctx context.Context, arg sqlc.CreateCourseParams) (sqlc.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCourse", ctx, arg)
	ret0, _ := ret[0].(sqlc.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCourse indicates an expected call of CreateCourse.
func (mr *MockQuerierMockRecorder) CreateCourse(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCourse", reflect.TypeOf((*MockQuerier)(nil).CreateCourse), ctx, arg)
}

// CreateEnrollment mocks base method.
func (m *MockQuerier) CreateEnrollment(ctx context.Context, arg sqlc.CreateEnrollmentParams) (sqlc.Enrollment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnrollment", ctx, arg)
	ret0, _ := ret[0].(sqlc.Enrollment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnrollment indicates an expected call of CreateEnrollment.
func (mr *MockQuerierMockRecorder) CreateEnrollment(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnrollment", reflect.TypeOf((*MockQuerier)(nil).CreateEnrollment), ctx, arg)
}

// CreateStudent mocks base method.
func (m *MockQuerier) CreateStudent(ctx context.Context, arg sqlc.CreateStudentParams) (sqlc.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStudent", ctx, arg)
	ret0, _ := ret[0].(sqlc.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStudent indicates an expected call of CreateStudent.
func (mr *MockQuerierMockRecorder) CreateStudent(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStudent", reflect.TypeOf((*MockQuerier)(nil).CreateStudent), ctx, arg)
}

// DeleteCourse mocks base method.
func (m *MockQuerier) DeleteCourse(ctx context.Context, courseID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCourse", ctx, courseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCourse indicates an expected call of DeleteCourse.
func (mr *MockQuerierMockRecorder) DeleteCourse(ctx, courseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCourse", reflect.TypeOf((*MockQuerier)(nil).DeleteCourse), ctx, courseID)
}

// DeleteEnrollment mocks base method.
func (m *MockQuerier) DeleteEnrollment(ctx context.Context, enrollmentID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEnrollment", ctx, enrollmentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnrollment indicates an expected call of DeleteEnrollment.
func (mr *MockQuerierMockRecorder) DeleteEnrollment(ctx, enrollmentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnrollment", reflect.TypeOf((*MockQuerier)(nil).DeleteEnrollment), ctx, enrollmentID)
}

// DeleteStudent mocks base method.
func (m *MockQuerier) DeleteStudent(ctx context.Context, studentID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStudent", ctx, studentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStudent indicates an expected call of DeleteStudent.
func (mr *MockQuerierMockRecorder) DeleteStudent(ctx, studentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStudent", reflect.TypeOf((*MockQuerier)(nil).DeleteStudent), ctx, studentID)
}

// GetCourse mocks base method.
func (m *MockQuerier) GetCourse(ctx context.Context, courseID int64) (sqlc.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCourse", ctx, courseID)
	ret0, _ := ret[0].(sqlc.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCourse indicates an expected call of GetCourse.
func (mr *MockQuerierMockRecorder) GetCourse(ctx, courseID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCourse", reflect.TypeOf((*MockQuerier)(nil).GetCourse), ctx, courseID)
}

// GetEnrollment mocks base method.
func (m *MockQuerier) GetEnrollment(ctx context.Context, enrollmentID int64) (sqlc.Enrollment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnrollment", ctx, enrollmentID)
	ret0, _ := ret[0].(sqlc.Enrollment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnrollment indicates an expected call of GetEnrollment.
func (mr *MockQuerierMockRecorder) GetEnrollment(ctx, enrollmentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnrollment", reflect.TypeOf((*MockQuerier)(nil).GetEnrollment), ctx, enrollmentID)
}

// GetStudent mocks base method.
func (m *MockQuerier) GetStudent(ctx context.Context, studentID int64) (sqlc.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStudent", ctx, studentID)
	ret0, _ := ret[0].(sqlc.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStudent indicates an expected call of GetStudent.
func (mr *MockQuerierMockRecorder) GetStudent(ctx, studentID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStudent", reflect.TypeOf((*MockQuerier)(nil).GetStudent), ctx, studentID)
}

// ListCourses mocks base method.
func (m *MockQuerier) ListCourses(ctx context.Context, arg sqlc.ListCoursesParams) ([]sqlc.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCourses", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCourses indicates an expected call of ListCourses.
func (mr *MockQuerierMockRecorder) ListCourses(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCourses", reflect.TypeOf((*MockQuerier)(nil).ListCourses), ctx, arg)
}

// ListEnrollment mocks base method.
func (m *MockQuerier) ListEnrollment(ctx context.Context, arg sqlc.ListEnrollmentParams) ([]sqlc.Enrollment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEnrollment", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Enrollment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnrollment indicates an expected call of ListEnrollment.
func (mr *MockQuerierMockRecorder) ListEnrollment(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnrollment", reflect.TypeOf((*MockQuerier)(nil).ListEnrollment), ctx, arg)
}

// ListStudents mocks base method.
func (m *MockQuerier) ListStudents(ctx context.Context, arg sqlc.ListStudentsParams) ([]sqlc.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStudents", ctx, arg)
	ret0, _ := ret[0].([]sqlc.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStudents indicates an expected call of ListStudents.
func (mr *MockQuerierMockRecorder) ListStudents(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStudents", reflect.TypeOf((*MockQuerier)(nil).ListStudents), ctx, arg)
}

// UpdateCourse mocks base method.
func (m *MockQuerier) UpdateCourse(ctx context.Context, arg sqlc.UpdateCourseParams) (sqlc.Course, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCourse", ctx, arg)
	ret0, _ := ret[0].(sqlc.Course)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCourse indicates an expected call of UpdateCourse.
func (mr *MockQuerierMockRecorder) UpdateCourse(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCourse", reflect.TypeOf((*MockQuerier)(nil).UpdateCourse), ctx, arg)
}

// UpdateStudent mocks base method.
func (m *MockQuerier) UpdateStudent(ctx context.Context, arg sqlc.UpdateStudentParams) (sqlc.Student, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStudent", ctx, arg)
	ret0, _ := ret[0].(sqlc.Student)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStudent indicates an expected call of UpdateStudent.
func (mr *MockQuerierMockRecorder) UpdateStudent(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStudent", reflect.TypeOf((*MockQuerier)(nil).UpdateStudent), ctx, arg)
}