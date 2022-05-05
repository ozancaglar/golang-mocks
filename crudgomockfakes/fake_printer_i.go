// Code generated by counterfeiter. DO NOT EDIT.
package crudgomockfakes

import (
	"printergomocks"
	"sync"
)

type FakePrinterI struct {
	PrintStub        func() error
	printMutex       sync.RWMutex
	printArgsForCall []struct {
	}
	printReturns struct {
		result1 error
	}
	printReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePrinterI) Print() error {
	fake.printMutex.Lock()
	ret, specificReturn := fake.printReturnsOnCall[len(fake.printArgsForCall)]
	fake.printArgsForCall = append(fake.printArgsForCall, struct {
	}{})
	stub := fake.PrintStub
	fakeReturns := fake.printReturns
	fake.recordInvocation("Print", []interface{}{})
	fake.printMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakePrinterI) PrintCallCount() int {
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	return len(fake.printArgsForCall)
}

func (fake *FakePrinterI) PrintCalls(stub func() error) {
	fake.printMutex.Lock()
	defer fake.printMutex.Unlock()
	fake.PrintStub = stub
}

func (fake *FakePrinterI) PrintReturns(result1 error) {
	fake.printMutex.Lock()
	defer fake.printMutex.Unlock()
	fake.PrintStub = nil
	fake.printReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePrinterI) PrintReturnsOnCall(i int, result1 error) {
	fake.printMutex.Lock()
	defer fake.printMutex.Unlock()
	fake.PrintStub = nil
	if fake.printReturnsOnCall == nil {
		fake.printReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.printReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePrinterI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.printMutex.RLock()
	defer fake.printMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePrinterI) recordInvocation(key string, args []interface{}) {
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

var _ printergomocks.PrinterI = new(FakePrinterI)