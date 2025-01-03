// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/crhntr/muxt-template-module-htmx/internal/hypertext"
)

type Server struct {
	IndexStub        func(context.Context) hypertext.IndexData
	indexMutex       sync.RWMutex
	indexArgsForCall []struct {
		arg1 context.Context
	}
	indexReturns struct {
		result1 hypertext.IndexData
	}
	indexReturnsOnCall map[int]struct {
		result1 hypertext.IndexData
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Server) Index(arg1 context.Context) hypertext.IndexData {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.IndexStub
	fakeReturns := fake.indexReturns
	fake.recordInvocation("Index", []interface{}{arg1})
	fake.indexMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *Server) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *Server) IndexCalls(stub func(context.Context) hypertext.IndexData) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = stub
}

func (fake *Server) IndexArgsForCall(i int) context.Context {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	argsForCall := fake.indexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Server) IndexReturns(result1 hypertext.IndexData) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 hypertext.IndexData
	}{result1}
}

func (fake *Server) IndexReturnsOnCall(i int, result1 hypertext.IndexData) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 hypertext.IndexData
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 hypertext.IndexData
	}{result1}
}

func (fake *Server) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Server) recordInvocation(key string, args []interface{}) {
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

var _ hypertext.RoutesReceiver = new(Server)
