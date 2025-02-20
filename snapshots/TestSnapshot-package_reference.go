// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3.Tester -o ./tests/tester_mock_test.go -n TesterMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TesterMock implements minimock.Tester
type TesterMock struct {
	t minimock.Tester

	funcError          func(p1 ...interface{})
	inspectFuncError   func(p1 ...interface{})
	afterErrorCounter  uint64
	beforeErrorCounter uint64
	ErrorMock          mTesterMockError

	funcErrorf          func(format string, args ...interface{})
	inspectFuncErrorf   func(format string, args ...interface{})
	afterErrorfCounter  uint64
	beforeErrorfCounter uint64
	ErrorfMock          mTesterMockErrorf

	funcFailNow          func()
	inspectFuncFailNow   func()
	afterFailNowCounter  uint64
	beforeFailNowCounter uint64
	FailNowMock          mTesterMockFailNow

	funcFatal          func(args ...interface{})
	inspectFuncFatal   func(args ...interface{})
	afterFatalCounter  uint64
	beforeFatalCounter uint64
	FatalMock          mTesterMockFatal

	funcFatalf          func(format string, args ...interface{})
	inspectFuncFatalf   func(format string, args ...interface{})
	afterFatalfCounter  uint64
	beforeFatalfCounter uint64
	FatalfMock          mTesterMockFatalf
}

// NewTesterMock returns a mock for minimock.Tester
func NewTesterMock(t minimock.Tester) *TesterMock {
	m := &TesterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ErrorMock = mTesterMockError{mock: m}
	m.ErrorMock.callArgs = []*TesterMockErrorParams{}

	m.ErrorfMock = mTesterMockErrorf{mock: m}
	m.ErrorfMock.callArgs = []*TesterMockErrorfParams{}

	m.FailNowMock = mTesterMockFailNow{mock: m}

	m.FatalMock = mTesterMockFatal{mock: m}
	m.FatalMock.callArgs = []*TesterMockFatalParams{}

	m.FatalfMock = mTesterMockFatalf{mock: m}
	m.FatalfMock.callArgs = []*TesterMockFatalfParams{}

	return m
}

type mTesterMockError struct {
	mock               *TesterMock
	defaultExpectation *TesterMockErrorExpectation
	expectations       []*TesterMockErrorExpectation

	callArgs []*TesterMockErrorParams
	mutex    sync.RWMutex
}

// TesterMockErrorExpectation specifies expectation struct of the Tester.Error
type TesterMockErrorExpectation struct {
	mock   *TesterMock
	params *TesterMockErrorParams

	Counter uint64
}

// TesterMockErrorParams contains parameters of the Tester.Error
type TesterMockErrorParams struct {
	p1 []interface{}
}

// Expect sets up expected params for Tester.Error
func (mmError *mTesterMockError) Expect(p1 ...interface{}) *mTesterMockError {
	if mmError.mock.funcError != nil {
		mmError.mock.t.Fatalf("TesterMock.Error mock is already set by Set")
	}

	if mmError.defaultExpectation == nil {
		mmError.defaultExpectation = &TesterMockErrorExpectation{}
	}

	mmError.defaultExpectation.params = &TesterMockErrorParams{p1}
	for _, e := range mmError.expectations {
		if minimock.Equal(e.params, mmError.defaultExpectation.params) {
			mmError.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmError.defaultExpectation.params)
		}
	}

	return mmError
}

// Inspect accepts an inspector function that has same arguments as the Tester.Error
func (mmError *mTesterMockError) Inspect(f func(p1 ...interface{})) *mTesterMockError {
	if mmError.mock.inspectFuncError != nil {
		mmError.mock.t.Fatalf("Inspect function is already set for TesterMock.Error")
	}

	mmError.mock.inspectFuncError = f

	return mmError
}

// Return sets up results that will be returned by Tester.Error
func (mmError *mTesterMockError) Return() *TesterMock {
	if mmError.mock.funcError != nil {
		mmError.mock.t.Fatalf("TesterMock.Error mock is already set by Set")
	}

	if mmError.defaultExpectation == nil {
		mmError.defaultExpectation = &TesterMockErrorExpectation{mock: mmError.mock}
	}

	return mmError.mock
}

// Set uses given function f to mock the Tester.Error method
func (mmError *mTesterMockError) Set(f func(p1 ...interface{})) *TesterMock {
	if mmError.defaultExpectation != nil {
		mmError.mock.t.Fatalf("Default expectation is already set for the Tester.Error method")
	}

	if len(mmError.expectations) > 0 {
		mmError.mock.t.Fatalf("Some expectations are already set for the Tester.Error method")
	}

	mmError.mock.funcError = f
	return mmError.mock
}

// Error implements minimock.Tester
func (mmError *TesterMock) Error(p1 ...interface{}) {
	mm_atomic.AddUint64(&mmError.beforeErrorCounter, 1)
	defer mm_atomic.AddUint64(&mmError.afterErrorCounter, 1)

	if mmError.inspectFuncError != nil {
		mmError.inspectFuncError(p1...)
	}

	mm_params := &TesterMockErrorParams{p1}

	// Record call args
	mmError.ErrorMock.mutex.Lock()
	mmError.ErrorMock.callArgs = append(mmError.ErrorMock.callArgs, mm_params)
	mmError.ErrorMock.mutex.Unlock()

	for _, e := range mmError.ErrorMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmError.ErrorMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmError.ErrorMock.defaultExpectation.Counter, 1)
		mm_want := mmError.ErrorMock.defaultExpectation.params
		mm_got := TesterMockErrorParams{p1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmError.t.Errorf("TesterMock.Error got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmError.funcError != nil {
		mmError.funcError(p1...)
		return
	}
	mmError.t.Fatalf("Unexpected call to TesterMock.Error. %v", p1)

}

// ErrorAfterCounter returns a count of finished TesterMock.Error invocations
func (mmError *TesterMock) ErrorAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmError.afterErrorCounter)
}

// ErrorBeforeCounter returns a count of TesterMock.Error invocations
func (mmError *TesterMock) ErrorBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmError.beforeErrorCounter)
}

// Calls returns a list of arguments used in each call to TesterMock.Error.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmError *mTesterMockError) Calls() []*TesterMockErrorParams {
	mmError.mutex.RLock()

	argCopy := make([]*TesterMockErrorParams, len(mmError.callArgs))
	copy(argCopy, mmError.callArgs)

	mmError.mutex.RUnlock()

	return argCopy
}

// MinimockErrorDone returns true if the count of the Error invocations corresponds
// the number of defined expectations
func (m *TesterMock) MinimockErrorDone() bool {
	for _, e := range m.ErrorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrorCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcError != nil && mm_atomic.LoadUint64(&m.afterErrorCounter) < 1 {
		return false
	}
	return true
}

// MinimockErrorInspect logs each unmet expectation
func (m *TesterMock) MinimockErrorInspect() {
	for _, e := range m.ErrorMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TesterMock.Error with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrorMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrorCounter) < 1 {
		if m.ErrorMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TesterMock.Error")
		} else {
			m.t.Errorf("Expected call to TesterMock.Error with params: %#v", *m.ErrorMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcError != nil && mm_atomic.LoadUint64(&m.afterErrorCounter) < 1 {
		m.t.Error("Expected call to TesterMock.Error")
	}
}

type mTesterMockErrorf struct {
	mock               *TesterMock
	defaultExpectation *TesterMockErrorfExpectation
	expectations       []*TesterMockErrorfExpectation

	callArgs []*TesterMockErrorfParams
	mutex    sync.RWMutex
}

// TesterMockErrorfExpectation specifies expectation struct of the Tester.Errorf
type TesterMockErrorfExpectation struct {
	mock   *TesterMock
	params *TesterMockErrorfParams

	Counter uint64
}

// TesterMockErrorfParams contains parameters of the Tester.Errorf
type TesterMockErrorfParams struct {
	format string
	args   []interface{}
}

// Expect sets up expected params for Tester.Errorf
func (mmErrorf *mTesterMockErrorf) Expect(format string, args ...interface{}) *mTesterMockErrorf {
	if mmErrorf.mock.funcErrorf != nil {
		mmErrorf.mock.t.Fatalf("TesterMock.Errorf mock is already set by Set")
	}

	if mmErrorf.defaultExpectation == nil {
		mmErrorf.defaultExpectation = &TesterMockErrorfExpectation{}
	}

	mmErrorf.defaultExpectation.params = &TesterMockErrorfParams{format, args}
	for _, e := range mmErrorf.expectations {
		if minimock.Equal(e.params, mmErrorf.defaultExpectation.params) {
			mmErrorf.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmErrorf.defaultExpectation.params)
		}
	}

	return mmErrorf
}

// Inspect accepts an inspector function that has same arguments as the Tester.Errorf
func (mmErrorf *mTesterMockErrorf) Inspect(f func(format string, args ...interface{})) *mTesterMockErrorf {
	if mmErrorf.mock.inspectFuncErrorf != nil {
		mmErrorf.mock.t.Fatalf("Inspect function is already set for TesterMock.Errorf")
	}

	mmErrorf.mock.inspectFuncErrorf = f

	return mmErrorf
}

// Return sets up results that will be returned by Tester.Errorf
func (mmErrorf *mTesterMockErrorf) Return() *TesterMock {
	if mmErrorf.mock.funcErrorf != nil {
		mmErrorf.mock.t.Fatalf("TesterMock.Errorf mock is already set by Set")
	}

	if mmErrorf.defaultExpectation == nil {
		mmErrorf.defaultExpectation = &TesterMockErrorfExpectation{mock: mmErrorf.mock}
	}

	return mmErrorf.mock
}

// Set uses given function f to mock the Tester.Errorf method
func (mmErrorf *mTesterMockErrorf) Set(f func(format string, args ...interface{})) *TesterMock {
	if mmErrorf.defaultExpectation != nil {
		mmErrorf.mock.t.Fatalf("Default expectation is already set for the Tester.Errorf method")
	}

	if len(mmErrorf.expectations) > 0 {
		mmErrorf.mock.t.Fatalf("Some expectations are already set for the Tester.Errorf method")
	}

	mmErrorf.mock.funcErrorf = f
	return mmErrorf.mock
}

// Errorf implements minimock.Tester
func (mmErrorf *TesterMock) Errorf(format string, args ...interface{}) {
	mm_atomic.AddUint64(&mmErrorf.beforeErrorfCounter, 1)
	defer mm_atomic.AddUint64(&mmErrorf.afterErrorfCounter, 1)

	if mmErrorf.inspectFuncErrorf != nil {
		mmErrorf.inspectFuncErrorf(format, args...)
	}

	mm_params := &TesterMockErrorfParams{format, args}

	// Record call args
	mmErrorf.ErrorfMock.mutex.Lock()
	mmErrorf.ErrorfMock.callArgs = append(mmErrorf.ErrorfMock.callArgs, mm_params)
	mmErrorf.ErrorfMock.mutex.Unlock()

	for _, e := range mmErrorf.ErrorfMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmErrorf.ErrorfMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmErrorf.ErrorfMock.defaultExpectation.Counter, 1)
		mm_want := mmErrorf.ErrorfMock.defaultExpectation.params
		mm_got := TesterMockErrorfParams{format, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmErrorf.t.Errorf("TesterMock.Errorf got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmErrorf.funcErrorf != nil {
		mmErrorf.funcErrorf(format, args...)
		return
	}
	mmErrorf.t.Fatalf("Unexpected call to TesterMock.Errorf. %v %v", format, args)

}

// ErrorfAfterCounter returns a count of finished TesterMock.Errorf invocations
func (mmErrorf *TesterMock) ErrorfAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmErrorf.afterErrorfCounter)
}

// ErrorfBeforeCounter returns a count of TesterMock.Errorf invocations
func (mmErrorf *TesterMock) ErrorfBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmErrorf.beforeErrorfCounter)
}

// Calls returns a list of arguments used in each call to TesterMock.Errorf.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmErrorf *mTesterMockErrorf) Calls() []*TesterMockErrorfParams {
	mmErrorf.mutex.RLock()

	argCopy := make([]*TesterMockErrorfParams, len(mmErrorf.callArgs))
	copy(argCopy, mmErrorf.callArgs)

	mmErrorf.mutex.RUnlock()

	return argCopy
}

// MinimockErrorfDone returns true if the count of the Errorf invocations corresponds
// the number of defined expectations
func (m *TesterMock) MinimockErrorfDone() bool {
	for _, e := range m.ErrorfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrorfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrorfCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErrorf != nil && mm_atomic.LoadUint64(&m.afterErrorfCounter) < 1 {
		return false
	}
	return true
}

// MinimockErrorfInspect logs each unmet expectation
func (m *TesterMock) MinimockErrorfInspect() {
	for _, e := range m.ErrorfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TesterMock.Errorf with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ErrorfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterErrorfCounter) < 1 {
		if m.ErrorfMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TesterMock.Errorf")
		} else {
			m.t.Errorf("Expected call to TesterMock.Errorf with params: %#v", *m.ErrorfMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcErrorf != nil && mm_atomic.LoadUint64(&m.afterErrorfCounter) < 1 {
		m.t.Error("Expected call to TesterMock.Errorf")
	}
}

type mTesterMockFailNow struct {
	mock               *TesterMock
	defaultExpectation *TesterMockFailNowExpectation
	expectations       []*TesterMockFailNowExpectation
}

// TesterMockFailNowExpectation specifies expectation struct of the Tester.FailNow
type TesterMockFailNowExpectation struct {
	mock *TesterMock

	Counter uint64
}

// Expect sets up expected params for Tester.FailNow
func (mmFailNow *mTesterMockFailNow) Expect() *mTesterMockFailNow {
	if mmFailNow.mock.funcFailNow != nil {
		mmFailNow.mock.t.Fatalf("TesterMock.FailNow mock is already set by Set")
	}

	if mmFailNow.defaultExpectation == nil {
		mmFailNow.defaultExpectation = &TesterMockFailNowExpectation{}
	}

	return mmFailNow
}

// Inspect accepts an inspector function that has same arguments as the Tester.FailNow
func (mmFailNow *mTesterMockFailNow) Inspect(f func()) *mTesterMockFailNow {
	if mmFailNow.mock.inspectFuncFailNow != nil {
		mmFailNow.mock.t.Fatalf("Inspect function is already set for TesterMock.FailNow")
	}

	mmFailNow.mock.inspectFuncFailNow = f

	return mmFailNow
}

// Return sets up results that will be returned by Tester.FailNow
func (mmFailNow *mTesterMockFailNow) Return() *TesterMock {
	if mmFailNow.mock.funcFailNow != nil {
		mmFailNow.mock.t.Fatalf("TesterMock.FailNow mock is already set by Set")
	}

	if mmFailNow.defaultExpectation == nil {
		mmFailNow.defaultExpectation = &TesterMockFailNowExpectation{mock: mmFailNow.mock}
	}

	return mmFailNow.mock
}

// Set uses given function f to mock the Tester.FailNow method
func (mmFailNow *mTesterMockFailNow) Set(f func()) *TesterMock {
	if mmFailNow.defaultExpectation != nil {
		mmFailNow.mock.t.Fatalf("Default expectation is already set for the Tester.FailNow method")
	}

	if len(mmFailNow.expectations) > 0 {
		mmFailNow.mock.t.Fatalf("Some expectations are already set for the Tester.FailNow method")
	}

	mmFailNow.mock.funcFailNow = f
	return mmFailNow.mock
}

// FailNow implements minimock.Tester
func (mmFailNow *TesterMock) FailNow() {
	mm_atomic.AddUint64(&mmFailNow.beforeFailNowCounter, 1)
	defer mm_atomic.AddUint64(&mmFailNow.afterFailNowCounter, 1)

	if mmFailNow.inspectFuncFailNow != nil {
		mmFailNow.inspectFuncFailNow()
	}

	if mmFailNow.FailNowMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFailNow.FailNowMock.defaultExpectation.Counter, 1)

		return

	}
	if mmFailNow.funcFailNow != nil {
		mmFailNow.funcFailNow()
		return
	}
	mmFailNow.t.Fatalf("Unexpected call to TesterMock.FailNow.")

}

// FailNowAfterCounter returns a count of finished TesterMock.FailNow invocations
func (mmFailNow *TesterMock) FailNowAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFailNow.afterFailNowCounter)
}

// FailNowBeforeCounter returns a count of TesterMock.FailNow invocations
func (mmFailNow *TesterMock) FailNowBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFailNow.beforeFailNowCounter)
}

// MinimockFailNowDone returns true if the count of the FailNow invocations corresponds
// the number of defined expectations
func (m *TesterMock) MinimockFailNowDone() bool {
	for _, e := range m.FailNowMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FailNowMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFailNowCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFailNow != nil && mm_atomic.LoadUint64(&m.afterFailNowCounter) < 1 {
		return false
	}
	return true
}

// MinimockFailNowInspect logs each unmet expectation
func (m *TesterMock) MinimockFailNowInspect() {
	for _, e := range m.FailNowMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to TesterMock.FailNow")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FailNowMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFailNowCounter) < 1 {
		m.t.Error("Expected call to TesterMock.FailNow")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFailNow != nil && mm_atomic.LoadUint64(&m.afterFailNowCounter) < 1 {
		m.t.Error("Expected call to TesterMock.FailNow")
	}
}

type mTesterMockFatal struct {
	mock               *TesterMock
	defaultExpectation *TesterMockFatalExpectation
	expectations       []*TesterMockFatalExpectation

	callArgs []*TesterMockFatalParams
	mutex    sync.RWMutex
}

// TesterMockFatalExpectation specifies expectation struct of the Tester.Fatal
type TesterMockFatalExpectation struct {
	mock   *TesterMock
	params *TesterMockFatalParams

	Counter uint64
}

// TesterMockFatalParams contains parameters of the Tester.Fatal
type TesterMockFatalParams struct {
	args []interface{}
}

// Expect sets up expected params for Tester.Fatal
func (mmFatal *mTesterMockFatal) Expect(args ...interface{}) *mTesterMockFatal {
	if mmFatal.mock.funcFatal != nil {
		mmFatal.mock.t.Fatalf("TesterMock.Fatal mock is already set by Set")
	}

	if mmFatal.defaultExpectation == nil {
		mmFatal.defaultExpectation = &TesterMockFatalExpectation{}
	}

	mmFatal.defaultExpectation.params = &TesterMockFatalParams{args}
	for _, e := range mmFatal.expectations {
		if minimock.Equal(e.params, mmFatal.defaultExpectation.params) {
			mmFatal.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFatal.defaultExpectation.params)
		}
	}

	return mmFatal
}

// Inspect accepts an inspector function that has same arguments as the Tester.Fatal
func (mmFatal *mTesterMockFatal) Inspect(f func(args ...interface{})) *mTesterMockFatal {
	if mmFatal.mock.inspectFuncFatal != nil {
		mmFatal.mock.t.Fatalf("Inspect function is already set for TesterMock.Fatal")
	}

	mmFatal.mock.inspectFuncFatal = f

	return mmFatal
}

// Return sets up results that will be returned by Tester.Fatal
func (mmFatal *mTesterMockFatal) Return() *TesterMock {
	if mmFatal.mock.funcFatal != nil {
		mmFatal.mock.t.Fatalf("TesterMock.Fatal mock is already set by Set")
	}

	if mmFatal.defaultExpectation == nil {
		mmFatal.defaultExpectation = &TesterMockFatalExpectation{mock: mmFatal.mock}
	}

	return mmFatal.mock
}

// Set uses given function f to mock the Tester.Fatal method
func (mmFatal *mTesterMockFatal) Set(f func(args ...interface{})) *TesterMock {
	if mmFatal.defaultExpectation != nil {
		mmFatal.mock.t.Fatalf("Default expectation is already set for the Tester.Fatal method")
	}

	if len(mmFatal.expectations) > 0 {
		mmFatal.mock.t.Fatalf("Some expectations are already set for the Tester.Fatal method")
	}

	mmFatal.mock.funcFatal = f
	return mmFatal.mock
}

// Fatal implements minimock.Tester
func (mmFatal *TesterMock) Fatal(args ...interface{}) {
	mm_atomic.AddUint64(&mmFatal.beforeFatalCounter, 1)
	defer mm_atomic.AddUint64(&mmFatal.afterFatalCounter, 1)

	if mmFatal.inspectFuncFatal != nil {
		mmFatal.inspectFuncFatal(args...)
	}

	mm_params := &TesterMockFatalParams{args}

	// Record call args
	mmFatal.FatalMock.mutex.Lock()
	mmFatal.FatalMock.callArgs = append(mmFatal.FatalMock.callArgs, mm_params)
	mmFatal.FatalMock.mutex.Unlock()

	for _, e := range mmFatal.FatalMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmFatal.FatalMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFatal.FatalMock.defaultExpectation.Counter, 1)
		mm_want := mmFatal.FatalMock.defaultExpectation.params
		mm_got := TesterMockFatalParams{args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFatal.t.Errorf("TesterMock.Fatal got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmFatal.funcFatal != nil {
		mmFatal.funcFatal(args...)
		return
	}
	mmFatal.t.Fatalf("Unexpected call to TesterMock.Fatal. %v", args)

}

// FatalAfterCounter returns a count of finished TesterMock.Fatal invocations
func (mmFatal *TesterMock) FatalAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFatal.afterFatalCounter)
}

// FatalBeforeCounter returns a count of TesterMock.Fatal invocations
func (mmFatal *TesterMock) FatalBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFatal.beforeFatalCounter)
}

// Calls returns a list of arguments used in each call to TesterMock.Fatal.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFatal *mTesterMockFatal) Calls() []*TesterMockFatalParams {
	mmFatal.mutex.RLock()

	argCopy := make([]*TesterMockFatalParams, len(mmFatal.callArgs))
	copy(argCopy, mmFatal.callArgs)

	mmFatal.mutex.RUnlock()

	return argCopy
}

// MinimockFatalDone returns true if the count of the Fatal invocations corresponds
// the number of defined expectations
func (m *TesterMock) MinimockFatalDone() bool {
	for _, e := range m.FatalMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FatalMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFatalCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFatal != nil && mm_atomic.LoadUint64(&m.afterFatalCounter) < 1 {
		return false
	}
	return true
}

// MinimockFatalInspect logs each unmet expectation
func (m *TesterMock) MinimockFatalInspect() {
	for _, e := range m.FatalMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TesterMock.Fatal with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FatalMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFatalCounter) < 1 {
		if m.FatalMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TesterMock.Fatal")
		} else {
			m.t.Errorf("Expected call to TesterMock.Fatal with params: %#v", *m.FatalMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFatal != nil && mm_atomic.LoadUint64(&m.afterFatalCounter) < 1 {
		m.t.Error("Expected call to TesterMock.Fatal")
	}
}

type mTesterMockFatalf struct {
	mock               *TesterMock
	defaultExpectation *TesterMockFatalfExpectation
	expectations       []*TesterMockFatalfExpectation

	callArgs []*TesterMockFatalfParams
	mutex    sync.RWMutex
}

// TesterMockFatalfExpectation specifies expectation struct of the Tester.Fatalf
type TesterMockFatalfExpectation struct {
	mock   *TesterMock
	params *TesterMockFatalfParams

	Counter uint64
}

// TesterMockFatalfParams contains parameters of the Tester.Fatalf
type TesterMockFatalfParams struct {
	format string
	args   []interface{}
}

// Expect sets up expected params for Tester.Fatalf
func (mmFatalf *mTesterMockFatalf) Expect(format string, args ...interface{}) *mTesterMockFatalf {
	if mmFatalf.mock.funcFatalf != nil {
		mmFatalf.mock.t.Fatalf("TesterMock.Fatalf mock is already set by Set")
	}

	if mmFatalf.defaultExpectation == nil {
		mmFatalf.defaultExpectation = &TesterMockFatalfExpectation{}
	}

	mmFatalf.defaultExpectation.params = &TesterMockFatalfParams{format, args}
	for _, e := range mmFatalf.expectations {
		if minimock.Equal(e.params, mmFatalf.defaultExpectation.params) {
			mmFatalf.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFatalf.defaultExpectation.params)
		}
	}

	return mmFatalf
}

// Inspect accepts an inspector function that has same arguments as the Tester.Fatalf
func (mmFatalf *mTesterMockFatalf) Inspect(f func(format string, args ...interface{})) *mTesterMockFatalf {
	if mmFatalf.mock.inspectFuncFatalf != nil {
		mmFatalf.mock.t.Fatalf("Inspect function is already set for TesterMock.Fatalf")
	}

	mmFatalf.mock.inspectFuncFatalf = f

	return mmFatalf
}

// Return sets up results that will be returned by Tester.Fatalf
func (mmFatalf *mTesterMockFatalf) Return() *TesterMock {
	if mmFatalf.mock.funcFatalf != nil {
		mmFatalf.mock.t.Fatalf("TesterMock.Fatalf mock is already set by Set")
	}

	if mmFatalf.defaultExpectation == nil {
		mmFatalf.defaultExpectation = &TesterMockFatalfExpectation{mock: mmFatalf.mock}
	}

	return mmFatalf.mock
}

// Set uses given function f to mock the Tester.Fatalf method
func (mmFatalf *mTesterMockFatalf) Set(f func(format string, args ...interface{})) *TesterMock {
	if mmFatalf.defaultExpectation != nil {
		mmFatalf.mock.t.Fatalf("Default expectation is already set for the Tester.Fatalf method")
	}

	if len(mmFatalf.expectations) > 0 {
		mmFatalf.mock.t.Fatalf("Some expectations are already set for the Tester.Fatalf method")
	}

	mmFatalf.mock.funcFatalf = f
	return mmFatalf.mock
}

// Fatalf implements minimock.Tester
func (mmFatalf *TesterMock) Fatalf(format string, args ...interface{}) {
	mm_atomic.AddUint64(&mmFatalf.beforeFatalfCounter, 1)
	defer mm_atomic.AddUint64(&mmFatalf.afterFatalfCounter, 1)

	if mmFatalf.inspectFuncFatalf != nil {
		mmFatalf.inspectFuncFatalf(format, args...)
	}

	mm_params := &TesterMockFatalfParams{format, args}

	// Record call args
	mmFatalf.FatalfMock.mutex.Lock()
	mmFatalf.FatalfMock.callArgs = append(mmFatalf.FatalfMock.callArgs, mm_params)
	mmFatalf.FatalfMock.mutex.Unlock()

	for _, e := range mmFatalf.FatalfMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmFatalf.FatalfMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFatalf.FatalfMock.defaultExpectation.Counter, 1)
		mm_want := mmFatalf.FatalfMock.defaultExpectation.params
		mm_got := TesterMockFatalfParams{format, args}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFatalf.t.Errorf("TesterMock.Fatalf got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmFatalf.funcFatalf != nil {
		mmFatalf.funcFatalf(format, args...)
		return
	}
	mmFatalf.t.Fatalf("Unexpected call to TesterMock.Fatalf. %v %v", format, args)

}

// FatalfAfterCounter returns a count of finished TesterMock.Fatalf invocations
func (mmFatalf *TesterMock) FatalfAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFatalf.afterFatalfCounter)
}

// FatalfBeforeCounter returns a count of TesterMock.Fatalf invocations
func (mmFatalf *TesterMock) FatalfBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFatalf.beforeFatalfCounter)
}

// Calls returns a list of arguments used in each call to TesterMock.Fatalf.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFatalf *mTesterMockFatalf) Calls() []*TesterMockFatalfParams {
	mmFatalf.mutex.RLock()

	argCopy := make([]*TesterMockFatalfParams, len(mmFatalf.callArgs))
	copy(argCopy, mmFatalf.callArgs)

	mmFatalf.mutex.RUnlock()

	return argCopy
}

// MinimockFatalfDone returns true if the count of the Fatalf invocations corresponds
// the number of defined expectations
func (m *TesterMock) MinimockFatalfDone() bool {
	for _, e := range m.FatalfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FatalfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFatalfCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFatalf != nil && mm_atomic.LoadUint64(&m.afterFatalfCounter) < 1 {
		return false
	}
	return true
}

// MinimockFatalfInspect logs each unmet expectation
func (m *TesterMock) MinimockFatalfInspect() {
	for _, e := range m.FatalfMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TesterMock.Fatalf with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FatalfMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFatalfCounter) < 1 {
		if m.FatalfMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TesterMock.Fatalf")
		} else {
			m.t.Errorf("Expected call to TesterMock.Fatalf with params: %#v", *m.FatalfMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFatalf != nil && mm_atomic.LoadUint64(&m.afterFatalfCounter) < 1 {
		m.t.Error("Expected call to TesterMock.Fatalf")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TesterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockErrorInspect()

		m.MinimockErrorfInspect()

		m.MinimockFailNowInspect()

		m.MinimockFatalInspect()

		m.MinimockFatalfInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TesterMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TesterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockErrorDone() &&
		m.MinimockErrorfDone() &&
		m.MinimockFailNowDone() &&
		m.MinimockFatalDone() &&
		m.MinimockFatalfDone()
}
