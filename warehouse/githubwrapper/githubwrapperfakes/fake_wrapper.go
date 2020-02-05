// Code generated by counterfeiter. DO NOT EDIT.
package githubwrapperfakes

import (
	"sync"

	"github.com/concourse/dutyfree/githubwrapper"
)

type FakeWrapper struct {
	GetStarsStub        func(map[string]int) (map[string]int, error)
	getStarsMutex       sync.RWMutex
	getStarsArgsForCall []struct {
		arg1 map[string]int
	}
	getStarsReturns struct {
		result1 map[string]int
		result2 error
	}
	getStarsReturnsOnCall map[int]struct {
		result1 map[string]int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWrapper) GetStars(arg1 map[string]int) (map[string]int, error) {
	fake.getStarsMutex.Lock()
	ret, specificReturn := fake.getStarsReturnsOnCall[len(fake.getStarsArgsForCall)]
	fake.getStarsArgsForCall = append(fake.getStarsArgsForCall, struct {
		arg1 map[string]int
	}{arg1})
	fake.recordInvocation("GetStars", []interface{}{arg1})
	fake.getStarsMutex.Unlock()
	if fake.GetStarsStub != nil {
		return fake.GetStarsStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getStarsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeWrapper) GetStarsCallCount() int {
	fake.getStarsMutex.RLock()
	defer fake.getStarsMutex.RUnlock()
	return len(fake.getStarsArgsForCall)
}

func (fake *FakeWrapper) GetStarsCalls(stub func(map[string]int) (map[string]int, error)) {
	fake.getStarsMutex.Lock()
	defer fake.getStarsMutex.Unlock()
	fake.GetStarsStub = stub
}

func (fake *FakeWrapper) GetStarsArgsForCall(i int) map[string]int {
	fake.getStarsMutex.RLock()
	defer fake.getStarsMutex.RUnlock()
	argsForCall := fake.getStarsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWrapper) GetStarsReturns(result1 map[string]int, result2 error) {
	fake.getStarsMutex.Lock()
	defer fake.getStarsMutex.Unlock()
	fake.GetStarsStub = nil
	fake.getStarsReturns = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

func (fake *FakeWrapper) GetStarsReturnsOnCall(i int, result1 map[string]int, result2 error) {
	fake.getStarsMutex.Lock()
	defer fake.getStarsMutex.Unlock()
	fake.GetStarsStub = nil
	if fake.getStarsReturnsOnCall == nil {
		fake.getStarsReturnsOnCall = make(map[int]struct {
			result1 map[string]int
			result2 error
		})
	}
	fake.getStarsReturnsOnCall[i] = struct {
		result1 map[string]int
		result2 error
	}{result1, result2}
}

func (fake *FakeWrapper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStarsMutex.RLock()
	defer fake.getStarsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWrapper) recordInvocation(key string, args []interface{}) {
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

var _ githubwrapper.Wrapper = new(FakeWrapper)
