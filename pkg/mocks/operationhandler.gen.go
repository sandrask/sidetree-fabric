// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	batchapi "github.com/trustbloc/sidetree-core-go/pkg/api/batch"
	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
)

type OperationHandler struct {
	PrepareTxnFilesStub        func(ops []*batchapi.Operation) (string, error)
	prepareTxnFilesMutex       sync.RWMutex
	prepareTxnFilesArgsForCall []struct {
		ops []*batchapi.Operation
	}
	prepareTxnFilesReturns struct {
		result1 string
		result2 error
	}
	prepareTxnFilesReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OperationHandler) PrepareTxnFiles(ops []*batchapi.Operation) (string, error) {
	var opsCopy []*batchapi.Operation
	if ops != nil {
		opsCopy = make([]*batchapi.Operation, len(ops))
		copy(opsCopy, ops)
	}
	fake.prepareTxnFilesMutex.Lock()
	ret, specificReturn := fake.prepareTxnFilesReturnsOnCall[len(fake.prepareTxnFilesArgsForCall)]
	fake.prepareTxnFilesArgsForCall = append(fake.prepareTxnFilesArgsForCall, struct {
		ops []*batchapi.Operation
	}{opsCopy})
	fake.recordInvocation("PrepareTxnFiles", []interface{}{opsCopy})
	fake.prepareTxnFilesMutex.Unlock()
	if fake.PrepareTxnFilesStub != nil {
		return fake.PrepareTxnFilesStub(ops)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.prepareTxnFilesReturns.result1, fake.prepareTxnFilesReturns.result2
}

func (fake *OperationHandler) PrepareTxnFilesCallCount() int {
	fake.prepareTxnFilesMutex.RLock()
	defer fake.prepareTxnFilesMutex.RUnlock()
	return len(fake.prepareTxnFilesArgsForCall)
}

func (fake *OperationHandler) PrepareTxnFilesArgsForCall(i int) []*batchapi.Operation {
	fake.prepareTxnFilesMutex.RLock()
	defer fake.prepareTxnFilesMutex.RUnlock()
	return fake.prepareTxnFilesArgsForCall[i].ops
}

func (fake *OperationHandler) PrepareTxnFilesReturns(result1 string, result2 error) {
	fake.PrepareTxnFilesStub = nil
	fake.prepareTxnFilesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationHandler) PrepareTxnFilesReturnsOnCall(i int, result1 string, result2 error) {
	fake.PrepareTxnFilesStub = nil
	if fake.prepareTxnFilesReturnsOnCall == nil {
		fake.prepareTxnFilesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.prepareTxnFilesReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *OperationHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.prepareTxnFilesMutex.RLock()
	defer fake.prepareTxnFilesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OperationHandler) recordInvocation(key string, args []interface{}) {
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

var _ protocol.OperationHandler = new(OperationHandler)
