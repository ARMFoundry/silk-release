// This file was generated by counterfeiter
package fakes

import (
	"sync"
)

type CIDRPool struct {
	GetAvailableStub        func([]string) string
	getAvailableMutex       sync.RWMutex
	getAvailableArgsForCall []struct {
		arg1 []string
	}
	getAvailableReturns struct {
		result1 string
	}
	getAvailableReturnsOnCall map[int]struct {
		result1 string
	}
	IsMemberStub        func(string) bool
	isMemberMutex       sync.RWMutex
	isMemberArgsForCall []struct {
		arg1 string
	}
	isMemberReturns struct {
		result1 bool
	}
	isMemberReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CIDRPool) GetAvailable(arg1 []string) string {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.getAvailableMutex.Lock()
	ret, specificReturn := fake.getAvailableReturnsOnCall[len(fake.getAvailableArgsForCall)]
	fake.getAvailableArgsForCall = append(fake.getAvailableArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	fake.recordInvocation("GetAvailable", []interface{}{arg1Copy})
	fake.getAvailableMutex.Unlock()
	if fake.GetAvailableStub != nil {
		return fake.GetAvailableStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getAvailableReturns.result1
}

func (fake *CIDRPool) GetAvailableCallCount() int {
	fake.getAvailableMutex.RLock()
	defer fake.getAvailableMutex.RUnlock()
	return len(fake.getAvailableArgsForCall)
}

func (fake *CIDRPool) GetAvailableArgsForCall(i int) []string {
	fake.getAvailableMutex.RLock()
	defer fake.getAvailableMutex.RUnlock()
	return fake.getAvailableArgsForCall[i].arg1
}

func (fake *CIDRPool) GetAvailableReturns(result1 string) {
	fake.GetAvailableStub = nil
	fake.getAvailableReturns = struct {
		result1 string
	}{result1}
}

func (fake *CIDRPool) GetAvailableReturnsOnCall(i int, result1 string) {
	fake.GetAvailableStub = nil
	if fake.getAvailableReturnsOnCall == nil {
		fake.getAvailableReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getAvailableReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *CIDRPool) IsMember(arg1 string) bool {
	fake.isMemberMutex.Lock()
	ret, specificReturn := fake.isMemberReturnsOnCall[len(fake.isMemberArgsForCall)]
	fake.isMemberArgsForCall = append(fake.isMemberArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsMember", []interface{}{arg1})
	fake.isMemberMutex.Unlock()
	if fake.IsMemberStub != nil {
		return fake.IsMemberStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isMemberReturns.result1
}

func (fake *CIDRPool) IsMemberCallCount() int {
	fake.isMemberMutex.RLock()
	defer fake.isMemberMutex.RUnlock()
	return len(fake.isMemberArgsForCall)
}

func (fake *CIDRPool) IsMemberArgsForCall(i int) string {
	fake.isMemberMutex.RLock()
	defer fake.isMemberMutex.RUnlock()
	return fake.isMemberArgsForCall[i].arg1
}

func (fake *CIDRPool) IsMemberReturns(result1 bool) {
	fake.IsMemberStub = nil
	fake.isMemberReturns = struct {
		result1 bool
	}{result1}
}

func (fake *CIDRPool) IsMemberReturnsOnCall(i int, result1 bool) {
	fake.IsMemberStub = nil
	if fake.isMemberReturnsOnCall == nil {
		fake.isMemberReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMemberReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *CIDRPool) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAvailableMutex.RLock()
	defer fake.getAvailableMutex.RUnlock()
	fake.isMemberMutex.RLock()
	defer fake.isMemberMutex.RUnlock()
	return fake.invocations
}

func (fake *CIDRPool) recordInvocation(key string, args []interface{}) {
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
