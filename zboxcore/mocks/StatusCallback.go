// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	"github.com/0chain/gosdk/core/transaction"
	mock "github.com/stretchr/testify/mock"
)

// StatusCallback is an autogenerated mock type for the StatusCallback type
type StatusCallback struct {
	mock.Mock
}

// CommitMetaCompleted provides a mock function with given fields: request, response, err
func (_m *StatusCallback) CommitMetaCompleted(request string, response string, txn *transaction.Transaction, err error) {
	_m.Called(request, response, err)
}

// Completed provides a mock function with given fields: allocationId, filePath, filename, mimetype, size, op
func (_m *StatusCallback) Completed(allocationId string, filePath string, filename string, mimetype string, size int, op int) {
	_m.Called(allocationId, filePath, filename, mimetype, size, op)
}

// Error provides a mock function with given fields: allocationID, filePath, op, err
func (_m *StatusCallback) Error(allocationID string, filePath string, op int, err error) {
	_m.Called(allocationID, filePath, op, err)
}

// InProgress provides a mock function with given fields: allocationId, filePath, op, completedBytes, data
func (_m *StatusCallback) InProgress(allocationId string, filePath string, op int, completedBytes int, data []byte) {
	_m.Called(allocationId, filePath, op, completedBytes, data)
}

// RepairCompleted provides a mock function with given fields: filesRepaired
func (_m *StatusCallback) RepairCompleted(filesRepaired int) {
	_m.Called(filesRepaired)
}

// Started provides a mock function with given fields: allocationId, filePath, op, totalBytes
func (_m *StatusCallback) Started(allocationId string, filePath string, op int, totalBytes int) {
	_m.Called(allocationId, filePath, op, totalBytes)
}
