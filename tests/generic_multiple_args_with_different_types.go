// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

package tests

//go:generate minimock -i github.com/gojuno/minimock/v3/tests.genericMultipleTypes -o ./tests/generic_multiple_args_with_different_types.go -n GenericMultipleTypesMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"google.golang.org/protobuf/proto"
)

// GenericMultipleTypesMock implements genericMultipleTypes
type GenericMultipleTypesMock[T proto.Message, K any] struct {
	t minimock.Tester

	funcName          func(t1 T, k1 K)
	inspectFuncName   func(t1 T, k1 K)
	afterNameCounter  uint64
	beforeNameCounter uint64
	NameMock          mGenericMultipleTypesMockName[T, K]
}

// NewGenericMultipleTypesMock returns a mock for genericMultipleTypes
func NewGenericMultipleTypesMock[T proto.Message, K any](t minimock.Tester) *GenericMultipleTypesMock[T, K] {
	m := &GenericMultipleTypesMock[T, K]{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NameMock = mGenericMultipleTypesMockName[T, K]{mock: m}
	m.NameMock.callArgs = []*GenericMultipleTypesMockNameParams[T, K]{}

	return m
}

type mGenericMultipleTypesMockName[T proto.Message, K any] struct {
	mock               *GenericMultipleTypesMock[T, K]
	defaultExpectation *GenericMultipleTypesMockNameExpectation[T, K]
	expectations       []*GenericMultipleTypesMockNameExpectation[T, K]

	callArgs []*GenericMultipleTypesMockNameParams[T, K]
	mutex    sync.RWMutex
}

// GenericMultipleTypesMockNameExpectation specifies expectation struct of the genericMultipleTypes.Name
type GenericMultipleTypesMockNameExpectation[T proto.Message, K any] struct {
	mock   *GenericMultipleTypesMock[T, K]
	params *GenericMultipleTypesMockNameParams[T, K]

	Counter uint64
}

// GenericMultipleTypesMockNameParams contains parameters of the genericMultipleTypes.Name
type GenericMultipleTypesMockNameParams[T proto.Message, K any] struct {
	t1 T
	k1 K
}

// Expect sets up expected params for genericMultipleTypes.Name
func (mmName *mGenericMultipleTypesMockName[T, K]) Expect(t1 T, k1 K) *mGenericMultipleTypesMockName[T, K] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericMultipleTypesMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericMultipleTypesMockNameExpectation[T, K]{}
	}

	mmName.defaultExpectation.params = &GenericMultipleTypesMockNameParams[T, K]{t1, k1}
	for _, e := range mmName.expectations {
		if minimock.Equal(e.params, mmName.defaultExpectation.params) {
			mmName.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmName.defaultExpectation.params)
		}
	}

	return mmName
}

// Inspect accepts an inspector function that has same arguments as the genericMultipleTypes.Name
func (mmName *mGenericMultipleTypesMockName[T, K]) Inspect(f func(t1 T, k1 K)) *mGenericMultipleTypesMockName[T, K] {
	if mmName.mock.inspectFuncName != nil {
		mmName.mock.t.Fatalf("Inspect function is already set for GenericMultipleTypesMock.Name")
	}

	mmName.mock.inspectFuncName = f

	return mmName
}

// Return sets up results that will be returned by genericMultipleTypes.Name
func (mmName *mGenericMultipleTypesMockName[T, K]) Return() *GenericMultipleTypesMock[T, K] {
	if mmName.mock.funcName != nil {
		mmName.mock.t.Fatalf("GenericMultipleTypesMock.Name mock is already set by Set")
	}

	if mmName.defaultExpectation == nil {
		mmName.defaultExpectation = &GenericMultipleTypesMockNameExpectation[T, K]{mock: mmName.mock}
	}

	return mmName.mock
}

// Set uses given function f to mock the genericMultipleTypes.Name method
func (mmName *mGenericMultipleTypesMockName[T, K]) Set(f func(t1 T, k1 K)) *GenericMultipleTypesMock[T, K] {
	if mmName.defaultExpectation != nil {
		mmName.mock.t.Fatalf("Default expectation is already set for the genericMultipleTypes.Name method")
	}

	if len(mmName.expectations) > 0 {
		mmName.mock.t.Fatalf("Some expectations are already set for the genericMultipleTypes.Name method")
	}

	mmName.mock.funcName = f
	return mmName.mock
}

// Name implements genericMultipleTypes
func (mmName *GenericMultipleTypesMock[T, K]) Name(t1 T, k1 K) {
	mm_atomic.AddUint64(&mmName.beforeNameCounter, 1)
	defer mm_atomic.AddUint64(&mmName.afterNameCounter, 1)

	if mmName.inspectFuncName != nil {
		mmName.inspectFuncName(t1, k1)
	}

	mm_params := &GenericMultipleTypesMockNameParams[T, K]{t1, k1}

	// Record call args
	mmName.NameMock.mutex.Lock()
	mmName.NameMock.callArgs = append(mmName.NameMock.callArgs, mm_params)
	mmName.NameMock.mutex.Unlock()

	for _, e := range mmName.NameMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmName.NameMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmName.NameMock.defaultExpectation.Counter, 1)
		mm_want := mmName.NameMock.defaultExpectation.params
		mm_got := GenericMultipleTypesMockNameParams[T, K]{t1, k1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmName.t.Errorf("GenericMultipleTypesMock.Name got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmName.funcName != nil {
		mmName.funcName(t1, k1)
		return
	}
	mmName.t.Fatalf("Unexpected call to GenericMultipleTypesMock.Name. %v %v", t1, k1)

}

// NameAfterCounter returns a count of finished GenericMultipleTypesMock.Name invocations
func (mmName *GenericMultipleTypesMock[T, K]) NameAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.afterNameCounter)
}

// NameBeforeCounter returns a count of GenericMultipleTypesMock.Name invocations
func (mmName *GenericMultipleTypesMock[T, K]) NameBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmName.beforeNameCounter)
}

// Calls returns a list of arguments used in each call to GenericMultipleTypesMock.Name.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmName *mGenericMultipleTypesMockName[T, K]) Calls() []*GenericMultipleTypesMockNameParams[T, K] {
	mmName.mutex.RLock()

	argCopy := make([]*GenericMultipleTypesMockNameParams[T, K], len(mmName.callArgs))
	copy(argCopy, mmName.callArgs)

	mmName.mutex.RUnlock()

	return argCopy
}

// MinimockNameDone returns true if the count of the Name invocations corresponds
// the number of defined expectations
func (m *GenericMultipleTypesMock[T, K]) MinimockNameDone() bool {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		return false
	}
	return true
}

// MinimockNameInspect logs each unmet expectation
func (m *GenericMultipleTypesMock[T, K]) MinimockNameInspect() {
	for _, e := range m.NameMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GenericMultipleTypesMock.Name with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NameMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		if m.NameMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GenericMultipleTypesMock.Name")
		} else {
			m.t.Errorf("Expected call to GenericMultipleTypesMock.Name with params: %#v", *m.NameMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcName != nil && mm_atomic.LoadUint64(&m.afterNameCounter) < 1 {
		m.t.Error("Expected call to GenericMultipleTypesMock.Name")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GenericMultipleTypesMock[T, K]) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockNameInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GenericMultipleTypesMock[T, K]) MinimockWait(timeout mm_time.Duration) {
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

func (m *GenericMultipleTypesMock[T, K]) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNameDone()
}
