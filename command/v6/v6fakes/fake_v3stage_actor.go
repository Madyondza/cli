// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v6"
)

type FakeV3StageActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	StagePackageStub        func(packageGUID string, appName string) (<-chan v3action.Droplet, <-chan v3action.Warnings, <-chan error)
	stagePackageMutex       sync.RWMutex
	stagePackageArgsForCall []struct {
		packageGUID string
		appName     string
	}
	stagePackageReturns struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}
	stagePackageReturnsOnCall map[int]struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3StageActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct{}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.cloudControllerAPIVersionReturns.result1
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3StageActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpace(appName string, spaceGUID string, client v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
		client    v3action.NOAAClient
	}{appName, spaceGUID, client})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{appName, spaceGUID, client})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(appName, spaceGUID, client)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result1, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result2, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result3, fake.getStreamingLogsForApplicationByNameAndSpaceReturns.result4
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v3action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].appName, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].spaceGUID, fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i].client
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3StageActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v3action.LogMessage
			result2 <-chan error
			result3 v3action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeV3StageActor) StagePackage(packageGUID string, appName string) (<-chan v3action.Droplet, <-chan v3action.Warnings, <-chan error) {
	fake.stagePackageMutex.Lock()
	ret, specificReturn := fake.stagePackageReturnsOnCall[len(fake.stagePackageArgsForCall)]
	fake.stagePackageArgsForCall = append(fake.stagePackageArgsForCall, struct {
		packageGUID string
		appName     string
	}{packageGUID, appName})
	fake.recordInvocation("StagePackage", []interface{}{packageGUID, appName})
	fake.stagePackageMutex.Unlock()
	if fake.StagePackageStub != nil {
		return fake.StagePackageStub(packageGUID, appName)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.stagePackageReturns.result1, fake.stagePackageReturns.result2, fake.stagePackageReturns.result3
}

func (fake *FakeV3StageActor) StagePackageCallCount() int {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return len(fake.stagePackageArgsForCall)
}

func (fake *FakeV3StageActor) StagePackageArgsForCall(i int) (string, string) {
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	return fake.stagePackageArgsForCall[i].packageGUID, fake.stagePackageArgsForCall[i].appName
}

func (fake *FakeV3StageActor) StagePackageReturns(result1 <-chan v3action.Droplet, result2 <-chan v3action.Warnings, result3 <-chan error) {
	fake.StagePackageStub = nil
	fake.stagePackageReturns = struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3StageActor) StagePackageReturnsOnCall(i int, result1 <-chan v3action.Droplet, result2 <-chan v3action.Warnings, result3 <-chan error) {
	fake.StagePackageStub = nil
	if fake.stagePackageReturnsOnCall == nil {
		fake.stagePackageReturnsOnCall = make(map[int]struct {
			result1 <-chan v3action.Droplet
			result2 <-chan v3action.Warnings
			result3 <-chan error
		})
	}
	fake.stagePackageReturnsOnCall[i] = struct {
		result1 <-chan v3action.Droplet
		result2 <-chan v3action.Warnings
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeV3StageActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.stagePackageMutex.RLock()
	defer fake.stagePackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3StageActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3StageActor = new(FakeV3StageActor)