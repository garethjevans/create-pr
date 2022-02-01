// Code generated by counterfeiter. DO NOT EDIT.
package pkgfakes

import (
	"sync"

	"github.com/garethjevans/create-pr/pkg"
)

type FakeGitter struct {
	GetOriginStub        func() (string, error)
	getOriginMutex       sync.RWMutex
	getOriginArgsForCall []struct {
	}
	getOriginReturns struct {
		result1 string
		result2 error
	}
	getOriginReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	HasLocalChangesStub        func() (bool, error)
	hasLocalChangesMutex       sync.RWMutex
	hasLocalChangesArgsForCall []struct {
	}
	hasLocalChangesReturns struct {
		result1 bool
		result2 error
	}
	hasLocalChangesReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeGitter) GetOrigin() (string, error) {
	fake.getOriginMutex.Lock()
	ret, specificReturn := fake.getOriginReturnsOnCall[len(fake.getOriginArgsForCall)]
	fake.getOriginArgsForCall = append(fake.getOriginArgsForCall, struct {
	}{})
	stub := fake.GetOriginStub
	fakeReturns := fake.getOriginReturns
	fake.recordInvocation("GetOrigin", []interface{}{})
	fake.getOriginMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitter) GetOriginCallCount() int {
	fake.getOriginMutex.RLock()
	defer fake.getOriginMutex.RUnlock()
	return len(fake.getOriginArgsForCall)
}

func (fake *FakeGitter) GetOriginCalls(stub func() (string, error)) {
	fake.getOriginMutex.Lock()
	defer fake.getOriginMutex.Unlock()
	fake.GetOriginStub = stub
}

func (fake *FakeGitter) GetOriginReturns(result1 string, result2 error) {
	fake.getOriginMutex.Lock()
	defer fake.getOriginMutex.Unlock()
	fake.GetOriginStub = nil
	fake.getOriginReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitter) GetOriginReturnsOnCall(i int, result1 string, result2 error) {
	fake.getOriginMutex.Lock()
	defer fake.getOriginMutex.Unlock()
	fake.GetOriginStub = nil
	if fake.getOriginReturnsOnCall == nil {
		fake.getOriginReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getOriginReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeGitter) HasLocalChanges() (bool, error) {
	fake.hasLocalChangesMutex.Lock()
	ret, specificReturn := fake.hasLocalChangesReturnsOnCall[len(fake.hasLocalChangesArgsForCall)]
	fake.hasLocalChangesArgsForCall = append(fake.hasLocalChangesArgsForCall, struct {
	}{})
	stub := fake.HasLocalChangesStub
	fakeReturns := fake.hasLocalChangesReturns
	fake.recordInvocation("HasLocalChanges", []interface{}{})
	fake.hasLocalChangesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeGitter) HasLocalChangesCallCount() int {
	fake.hasLocalChangesMutex.RLock()
	defer fake.hasLocalChangesMutex.RUnlock()
	return len(fake.hasLocalChangesArgsForCall)
}

func (fake *FakeGitter) HasLocalChangesCalls(stub func() (bool, error)) {
	fake.hasLocalChangesMutex.Lock()
	defer fake.hasLocalChangesMutex.Unlock()
	fake.HasLocalChangesStub = stub
}

func (fake *FakeGitter) HasLocalChangesReturns(result1 bool, result2 error) {
	fake.hasLocalChangesMutex.Lock()
	defer fake.hasLocalChangesMutex.Unlock()
	fake.HasLocalChangesStub = nil
	fake.hasLocalChangesReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitter) HasLocalChangesReturnsOnCall(i int, result1 bool, result2 error) {
	fake.hasLocalChangesMutex.Lock()
	defer fake.hasLocalChangesMutex.Unlock()
	fake.HasLocalChangesStub = nil
	if fake.hasLocalChangesReturnsOnCall == nil {
		fake.hasLocalChangesReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasLocalChangesReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeGitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getOriginMutex.RLock()
	defer fake.getOriginMutex.RUnlock()
	fake.hasLocalChangesMutex.RLock()
	defer fake.hasLocalChangesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeGitter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pkg.Gitter = new(FakeGitter)
