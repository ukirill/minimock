// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.Formatter -o ./tests/formatter_mock.go -n FormatterMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// FormatterMock implements Formatter
type FormatterMock struct {
	t minimock.Tester

	funcFormat          func(s1 string, p1 ...interface{}) (s2 string)
	inspectFuncFormat   func(s1 string, p1 ...interface{})
	afterFormatCounter  uint64
	beforeFormatCounter uint64
	FormatMock          mFormatterMockFormat
}

// NewFormatterMock returns a mock for Formatter
func NewFormatterMock(t minimock.Tester) *FormatterMock {
	m := &FormatterMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.FormatMock = mFormatterMockFormat{mock: m}
	m.FormatMock.callArgs = []*FormatterMockFormatParams{}

	return m
}

type mFormatterMockFormat struct {
	mock               *FormatterMock
	defaultExpectation *FormatterMockFormatExpectation
	expectations       []*FormatterMockFormatExpectation

	callArgs []*FormatterMockFormatParams
	mutex    sync.RWMutex
}

// FormatterMockFormatExpectation specifies expectation struct of the Formatter.Format
type FormatterMockFormatExpectation struct {
	mock    *FormatterMock
	params  *FormatterMockFormatParams
	results *FormatterMockFormatResults
	Counter uint64
}

// FormatterMockFormatParams contains parameters of the Formatter.Format
type FormatterMockFormatParams struct {
	s1 string
	p1 []interface{}
}

// FormatterMockFormatResults contains results of the Formatter.Format
type FormatterMockFormatResults struct {
	s2 string
}

// Expect sets up expected params for Formatter.Format
func (mmFormat *mFormatterMockFormat) Expect(s1 string, p1 ...interface{}) *mFormatterMockFormat {
	if mmFormat.mock.funcFormat != nil {
		mmFormat.mock.t.Fatalf("FormatterMock.Format mock is already set by Set")
	}

	if mmFormat.defaultExpectation == nil {
		mmFormat.defaultExpectation = &FormatterMockFormatExpectation{}
	}

	mmFormat.defaultExpectation.params = &FormatterMockFormatParams{s1, p1}
	for _, e := range mmFormat.expectations {
		if minimock.Equal(e.params, mmFormat.defaultExpectation.params) {
			mmFormat.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmFormat.defaultExpectation.params)
		}
	}

	return mmFormat
}

// Inspect accepts an inspector function that has same arguments as the Formatter.Format
func (mmFormat *mFormatterMockFormat) Inspect(f func(s1 string, p1 ...interface{})) *mFormatterMockFormat {
	if mmFormat.mock.inspectFuncFormat != nil {
		mmFormat.mock.t.Fatalf("Inspect function is already set for FormatterMock.Format")
	}

	mmFormat.mock.inspectFuncFormat = f

	return mmFormat
}

// Return sets up results that will be returned by Formatter.Format
func (mmFormat *mFormatterMockFormat) Return(s2 string) *FormatterMock {
	if mmFormat.mock.funcFormat != nil {
		mmFormat.mock.t.Fatalf("FormatterMock.Format mock is already set by Set")
	}

	if mmFormat.defaultExpectation == nil {
		mmFormat.defaultExpectation = &FormatterMockFormatExpectation{mock: mmFormat.mock}
	}
	mmFormat.defaultExpectation.results = &FormatterMockFormatResults{s2}
	return mmFormat.mock
}

// Set uses given function f to mock the Formatter.Format method
func (mmFormat *mFormatterMockFormat) Set(f func(s1 string, p1 ...interface{}) (s2 string)) *FormatterMock {
	if mmFormat.defaultExpectation != nil {
		mmFormat.mock.t.Fatalf("Default expectation is already set for the Formatter.Format method")
	}

	if len(mmFormat.expectations) > 0 {
		mmFormat.mock.t.Fatalf("Some expectations are already set for the Formatter.Format method")
	}

	mmFormat.mock.funcFormat = f
	return mmFormat.mock
}

// When sets expectation for the Formatter.Format which will trigger the result defined by the following
// Then helper
func (mmFormat *mFormatterMockFormat) When(s1 string, p1 ...interface{}) *FormatterMockFormatExpectation {
	if mmFormat.mock.funcFormat != nil {
		mmFormat.mock.t.Fatalf("FormatterMock.Format mock is already set by Set")
	}

	expectation := &FormatterMockFormatExpectation{
		mock:   mmFormat.mock,
		params: &FormatterMockFormatParams{s1, p1},
	}
	mmFormat.expectations = append(mmFormat.expectations, expectation)
	return expectation
}

// Then sets up Formatter.Format return parameters for the expectation previously defined by the When method
func (e *FormatterMockFormatExpectation) Then(s2 string) *FormatterMock {
	e.results = &FormatterMockFormatResults{s2}
	return e.mock
}

// Format implements Formatter
func (mmFormat *FormatterMock) Format(s1 string, p1 ...interface{}) (s2 string) {
	mm_atomic.AddUint64(&mmFormat.beforeFormatCounter, 1)
	defer mm_atomic.AddUint64(&mmFormat.afterFormatCounter, 1)

	if mmFormat.inspectFuncFormat != nil {
		mmFormat.inspectFuncFormat(s1, p1...)
	}

	mm_params := &FormatterMockFormatParams{s1, p1}

	// Record call args
	mmFormat.FormatMock.mutex.Lock()
	mmFormat.FormatMock.callArgs = append(mmFormat.FormatMock.callArgs, mm_params)
	mmFormat.FormatMock.mutex.Unlock()

	for _, e := range mmFormat.FormatMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s2
		}
	}

	if mmFormat.FormatMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmFormat.FormatMock.defaultExpectation.Counter, 1)
		mm_want := mmFormat.FormatMock.defaultExpectation.params
		mm_got := FormatterMockFormatParams{s1, p1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmFormat.t.Errorf("FormatterMock.Format got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmFormat.FormatMock.defaultExpectation.results
		if mm_results == nil {
			mmFormat.t.Fatal("No results are set for the FormatterMock.Format")
		}
		return (*mm_results).s2
	}
	if mmFormat.funcFormat != nil {
		return mmFormat.funcFormat(s1, p1...)
	}
	mmFormat.t.Fatalf("Unexpected call to FormatterMock.Format. %v %v", s1, p1)
	return
}

// FormatAfterCounter returns a count of finished FormatterMock.Format invocations
func (mmFormat *FormatterMock) FormatAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFormat.afterFormatCounter)
}

// FormatBeforeCounter returns a count of FormatterMock.Format invocations
func (mmFormat *FormatterMock) FormatBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmFormat.beforeFormatCounter)
}

// Calls returns a list of arguments used in each call to FormatterMock.Format.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmFormat *mFormatterMockFormat) Calls() []*FormatterMockFormatParams {
	mmFormat.mutex.RLock()

	argCopy := make([]*FormatterMockFormatParams, len(mmFormat.callArgs))
	copy(argCopy, mmFormat.callArgs)

	mmFormat.mutex.RUnlock()

	return argCopy
}

// MinimockFormatDone returns true if the count of the Format invocations corresponds
// the number of defined expectations
func (m *FormatterMock) MinimockFormatDone() bool {
	for _, e := range m.FormatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FormatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFormatCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFormat != nil && mm_atomic.LoadUint64(&m.afterFormatCounter) < 1 {
		return false
	}
	return true
}

// MinimockFormatInspect logs each unmet expectation
func (m *FormatterMock) MinimockFormatInspect() {
	for _, e := range m.FormatMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FormatterMock.Format with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.FormatMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterFormatCounter) < 1 {
		if m.FormatMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FormatterMock.Format")
		} else {
			m.t.Errorf("Expected call to FormatterMock.Format with params: %#v", *m.FormatMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcFormat != nil && mm_atomic.LoadUint64(&m.afterFormatCounter) < 1 {
		m.t.Error("Expected call to FormatterMock.Format")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FormatterMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockFormatInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FormatterMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *FormatterMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockFormatDone()
}
