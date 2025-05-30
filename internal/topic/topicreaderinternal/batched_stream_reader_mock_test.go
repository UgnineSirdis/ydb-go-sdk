// Code generated by MockGen. DO NOT EDIT.
// Source: batched_stream_reader_interface.go
//
// Generated by this command:
//
//	mockgen -source batched_stream_reader_interface.go --typed -destination batched_stream_reader_mock_test.go -package topicreaderinternal -write_package_comment=false
package topicreaderinternal

import (
	context "context"
	reflect "reflect"

	topicreadercommon "github.com/ydb-platform/ydb-go-sdk/v3/internal/topic/topicreadercommon"
	tx "github.com/ydb-platform/ydb-go-sdk/v3/internal/tx"
	gomock "go.uber.org/mock/gomock"
)

// MockbatchedStreamReader is a mock of batchedStreamReader interface.
type MockbatchedStreamReader struct {
	ctrl     *gomock.Controller
	recorder *MockbatchedStreamReaderMockRecorder
}

// MockbatchedStreamReaderMockRecorder is the mock recorder for MockbatchedStreamReader.
type MockbatchedStreamReaderMockRecorder struct {
	mock *MockbatchedStreamReader
}

// NewMockbatchedStreamReader creates a new mock instance.
func NewMockbatchedStreamReader(ctrl *gomock.Controller) *MockbatchedStreamReader {
	mock := &MockbatchedStreamReader{ctrl: ctrl}
	mock.recorder = &MockbatchedStreamReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockbatchedStreamReader) EXPECT() *MockbatchedStreamReaderMockRecorder {
	return m.recorder
}

// CloseWithError mocks base method.
func (m *MockbatchedStreamReader) CloseWithError(ctx context.Context, err error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", ctx, err)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockbatchedStreamReaderMockRecorder) CloseWithError(ctx, err any) *MockbatchedStreamReaderCloseWithErrorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockbatchedStreamReader)(nil).CloseWithError), ctx, err)
	return &MockbatchedStreamReaderCloseWithErrorCall{Call: call}
}

// MockbatchedStreamReaderCloseWithErrorCall wrap *gomock.Call
type MockbatchedStreamReaderCloseWithErrorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderCloseWithErrorCall) Return(arg0 error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderCloseWithErrorCall) Do(f func(context.Context, error) error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderCloseWithErrorCall) DoAndReturn(f func(context.Context, error) error) *MockbatchedStreamReaderCloseWithErrorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Commit mocks base method.
func (m *MockbatchedStreamReader) Commit(ctx context.Context, commitRange topicreadercommon.CommitRange) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", ctx, commitRange)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockbatchedStreamReaderMockRecorder) Commit(ctx, commitRange any) *MockbatchedStreamReaderCommitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockbatchedStreamReader)(nil).Commit), ctx, commitRange)
	return &MockbatchedStreamReaderCommitCall{Call: call}
}

// MockbatchedStreamReaderCommitCall wrap *gomock.Call
type MockbatchedStreamReaderCommitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderCommitCall) Return(arg0 error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderCommitCall) Do(f func(context.Context, topicreadercommon.CommitRange) error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderCommitCall) DoAndReturn(f func(context.Context, topicreadercommon.CommitRange) error) *MockbatchedStreamReaderCommitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PopMessagesBatchTx mocks base method.
func (m *MockbatchedStreamReader) PopMessagesBatchTx(ctx context.Context, tx tx.Transaction, opts ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PopMessagesBatchTx", ctx, tx, opts)
	ret0, _ := ret[0].(*topicreadercommon.PublicBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PopMessagesBatchTx indicates an expected call of PopMessagesBatchTx.
func (mr *MockbatchedStreamReaderMockRecorder) PopMessagesBatchTx(ctx, tx, opts any) *MockbatchedStreamReaderPopMessagesBatchTxCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PopMessagesBatchTx", reflect.TypeOf((*MockbatchedStreamReader)(nil).PopMessagesBatchTx), ctx, tx, opts)
	return &MockbatchedStreamReaderPopMessagesBatchTxCall{Call: call}
}

// MockbatchedStreamReaderPopMessagesBatchTxCall wrap *gomock.Call
type MockbatchedStreamReaderPopMessagesBatchTxCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderPopMessagesBatchTxCall) Return(arg0 *topicreadercommon.PublicBatch, arg1 error) *MockbatchedStreamReaderPopMessagesBatchTxCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderPopMessagesBatchTxCall) Do(f func(context.Context, tx.Transaction, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderPopMessagesBatchTxCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderPopMessagesBatchTxCall) DoAndReturn(f func(context.Context, tx.Transaction, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderPopMessagesBatchTxCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ReadMessageBatch mocks base method.
func (m *MockbatchedStreamReader) ReadMessageBatch(ctx context.Context, opts ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadMessageBatch", ctx, opts)
	ret0, _ := ret[0].(*topicreadercommon.PublicBatch)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadMessageBatch indicates an expected call of ReadMessageBatch.
func (mr *MockbatchedStreamReaderMockRecorder) ReadMessageBatch(ctx, opts any) *MockbatchedStreamReaderReadMessageBatchCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadMessageBatch", reflect.TypeOf((*MockbatchedStreamReader)(nil).ReadMessageBatch), ctx, opts)
	return &MockbatchedStreamReaderReadMessageBatchCall{Call: call}
}

// MockbatchedStreamReaderReadMessageBatchCall wrap *gomock.Call
type MockbatchedStreamReaderReadMessageBatchCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderReadMessageBatchCall) Return(arg0 *topicreadercommon.PublicBatch, arg1 error) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderReadMessageBatchCall) Do(f func(context.Context, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderReadMessageBatchCall) DoAndReturn(f func(context.Context, ReadMessageBatchOptions) (*topicreadercommon.PublicBatch, error)) *MockbatchedStreamReaderReadMessageBatchCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// TopicOnReaderStart mocks base method.
func (m *MockbatchedStreamReader) TopicOnReaderStart(consumer string, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "TopicOnReaderStart", consumer, err)
}

// TopicOnReaderStart indicates an expected call of TopicOnReaderStart.
func (mr *MockbatchedStreamReaderMockRecorder) TopicOnReaderStart(consumer, err any) *MockbatchedStreamReaderTopicOnReaderStartCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopicOnReaderStart", reflect.TypeOf((*MockbatchedStreamReader)(nil).TopicOnReaderStart), consumer, err)
	return &MockbatchedStreamReaderTopicOnReaderStartCall{Call: call}
}

// MockbatchedStreamReaderTopicOnReaderStartCall wrap *gomock.Call
type MockbatchedStreamReaderTopicOnReaderStartCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderTopicOnReaderStartCall) Return() *MockbatchedStreamReaderTopicOnReaderStartCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderTopicOnReaderStartCall) Do(f func(string, error)) *MockbatchedStreamReaderTopicOnReaderStartCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderTopicOnReaderStartCall) DoAndReturn(f func(string, error)) *MockbatchedStreamReaderTopicOnReaderStartCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// WaitInit mocks base method.
func (m *MockbatchedStreamReader) WaitInit(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitInit", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitInit indicates an expected call of WaitInit.
func (mr *MockbatchedStreamReaderMockRecorder) WaitInit(ctx any) *MockbatchedStreamReaderWaitInitCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitInit", reflect.TypeOf((*MockbatchedStreamReader)(nil).WaitInit), ctx)
	return &MockbatchedStreamReaderWaitInitCall{Call: call}
}

// MockbatchedStreamReaderWaitInitCall wrap *gomock.Call
type MockbatchedStreamReaderWaitInitCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockbatchedStreamReaderWaitInitCall) Return(arg0 error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockbatchedStreamReaderWaitInitCall) Do(f func(context.Context) error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockbatchedStreamReaderWaitInitCall) DoAndReturn(f func(context.Context) error) *MockbatchedStreamReaderWaitInitCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
