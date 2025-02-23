// Code generated by counterfeiter. DO NOT EDIT.
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/pushaction"
	"code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/util/manifest"
)

type FakeV2PushActor struct {
	ApplyStub        func(config pushaction.ApplicationConfig, progressBar pushaction.ProgressBar) (<-chan pushaction.ApplicationConfig, <-chan pushaction.Event, <-chan pushaction.Warnings, <-chan error)
	applyMutex       sync.RWMutex
	applyArgsForCall []struct {
		config      pushaction.ApplicationConfig
		progressBar pushaction.ProgressBar
	}
	applyReturns struct {
		result1 <-chan pushaction.ApplicationConfig
		result2 <-chan pushaction.Event
		result3 <-chan pushaction.Warnings
		result4 <-chan error
	}
	applyReturnsOnCall map[int]struct {
		result1 <-chan pushaction.ApplicationConfig
		result2 <-chan pushaction.Event
		result3 <-chan pushaction.Warnings
		result4 <-chan error
	}
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct{}
	cloudControllerAPIVersionReturns     struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	ConvertToApplicationConfigsStub        func(orgGUID string, spaceGUID string, noStart bool, apps []manifest.Application) ([]pushaction.ApplicationConfig, pushaction.Warnings, error)
	convertToApplicationConfigsMutex       sync.RWMutex
	convertToApplicationConfigsArgsForCall []struct {
		orgGUID   string
		spaceGUID string
		noStart   bool
		apps      []manifest.Application
	}
	convertToApplicationConfigsReturns struct {
		result1 []pushaction.ApplicationConfig
		result2 pushaction.Warnings
		result3 error
	}
	convertToApplicationConfigsReturnsOnCall map[int]struct {
		result1 []pushaction.ApplicationConfig
		result2 pushaction.Warnings
		result3 error
	}
	MergeAndValidateSettingsAndManifestsStub        func(cmdSettings pushaction.CommandLineSettings, apps []manifest.Application) ([]manifest.Application, error)
	mergeAndValidateSettingsAndManifestsMutex       sync.RWMutex
	mergeAndValidateSettingsAndManifestsArgsForCall []struct {
		cmdSettings pushaction.CommandLineSettings
		apps        []manifest.Application
	}
	mergeAndValidateSettingsAndManifestsReturns struct {
		result1 []manifest.Application
		result2 error
	}
	mergeAndValidateSettingsAndManifestsReturnsOnCall map[int]struct {
		result1 []manifest.Application
		result2 error
	}
	ReadManifestStub        func(pathToManifest string) ([]manifest.Application, error)
	readManifestMutex       sync.RWMutex
	readManifestArgsForCall []struct {
		pathToManifest string
	}
	readManifestReturns struct {
		result1 []manifest.Application
		result2 error
	}
	readManifestReturnsOnCall map[int]struct {
		result1 []manifest.Application
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV2PushActor) Apply(config pushaction.ApplicationConfig, progressBar pushaction.ProgressBar) (<-chan pushaction.ApplicationConfig, <-chan pushaction.Event, <-chan pushaction.Warnings, <-chan error) {
	fake.applyMutex.Lock()
	ret, specificReturn := fake.applyReturnsOnCall[len(fake.applyArgsForCall)]
	fake.applyArgsForCall = append(fake.applyArgsForCall, struct {
		config      pushaction.ApplicationConfig
		progressBar pushaction.ProgressBar
	}{config, progressBar})
	fake.recordInvocation("Apply", []interface{}{config, progressBar})
	fake.applyMutex.Unlock()
	if fake.ApplyStub != nil {
		return fake.ApplyStub(config, progressBar)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.applyReturns.result1, fake.applyReturns.result2, fake.applyReturns.result3, fake.applyReturns.result4
}

func (fake *FakeV2PushActor) ApplyCallCount() int {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return len(fake.applyArgsForCall)
}

func (fake *FakeV2PushActor) ApplyArgsForCall(i int) (pushaction.ApplicationConfig, pushaction.ProgressBar) {
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	return fake.applyArgsForCall[i].config, fake.applyArgsForCall[i].progressBar
}

func (fake *FakeV2PushActor) ApplyReturns(result1 <-chan pushaction.ApplicationConfig, result2 <-chan pushaction.Event, result3 <-chan pushaction.Warnings, result4 <-chan error) {
	fake.ApplyStub = nil
	fake.applyReturns = struct {
		result1 <-chan pushaction.ApplicationConfig
		result2 <-chan pushaction.Event
		result3 <-chan pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakeV2PushActor) ApplyReturnsOnCall(i int, result1 <-chan pushaction.ApplicationConfig, result2 <-chan pushaction.Event, result3 <-chan pushaction.Warnings, result4 <-chan error) {
	fake.ApplyStub = nil
	if fake.applyReturnsOnCall == nil {
		fake.applyReturnsOnCall = make(map[int]struct {
			result1 <-chan pushaction.ApplicationConfig
			result2 <-chan pushaction.Event
			result3 <-chan pushaction.Warnings
			result4 <-chan error
		})
	}
	fake.applyReturnsOnCall[i] = struct {
		result1 <-chan pushaction.ApplicationConfig
		result2 <-chan pushaction.Event
		result3 <-chan pushaction.Warnings
		result4 <-chan error
	}{result1, result2, result3, result4}
}

func (fake *FakeV2PushActor) CloudControllerAPIVersion() string {
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

func (fake *FakeV2PushActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakeV2PushActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeV2PushActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
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

func (fake *FakeV2PushActor) ConvertToApplicationConfigs(orgGUID string, spaceGUID string, noStart bool, apps []manifest.Application) ([]pushaction.ApplicationConfig, pushaction.Warnings, error) {
	var appsCopy []manifest.Application
	if apps != nil {
		appsCopy = make([]manifest.Application, len(apps))
		copy(appsCopy, apps)
	}
	fake.convertToApplicationConfigsMutex.Lock()
	ret, specificReturn := fake.convertToApplicationConfigsReturnsOnCall[len(fake.convertToApplicationConfigsArgsForCall)]
	fake.convertToApplicationConfigsArgsForCall = append(fake.convertToApplicationConfigsArgsForCall, struct {
		orgGUID   string
		spaceGUID string
		noStart   bool
		apps      []manifest.Application
	}{orgGUID, spaceGUID, noStart, appsCopy})
	fake.recordInvocation("ConvertToApplicationConfigs", []interface{}{orgGUID, spaceGUID, noStart, appsCopy})
	fake.convertToApplicationConfigsMutex.Unlock()
	if fake.ConvertToApplicationConfigsStub != nil {
		return fake.ConvertToApplicationConfigsStub(orgGUID, spaceGUID, noStart, apps)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.convertToApplicationConfigsReturns.result1, fake.convertToApplicationConfigsReturns.result2, fake.convertToApplicationConfigsReturns.result3
}

func (fake *FakeV2PushActor) ConvertToApplicationConfigsCallCount() int {
	fake.convertToApplicationConfigsMutex.RLock()
	defer fake.convertToApplicationConfigsMutex.RUnlock()
	return len(fake.convertToApplicationConfigsArgsForCall)
}

func (fake *FakeV2PushActor) ConvertToApplicationConfigsArgsForCall(i int) (string, string, bool, []manifest.Application) {
	fake.convertToApplicationConfigsMutex.RLock()
	defer fake.convertToApplicationConfigsMutex.RUnlock()
	return fake.convertToApplicationConfigsArgsForCall[i].orgGUID, fake.convertToApplicationConfigsArgsForCall[i].spaceGUID, fake.convertToApplicationConfigsArgsForCall[i].noStart, fake.convertToApplicationConfigsArgsForCall[i].apps
}

func (fake *FakeV2PushActor) ConvertToApplicationConfigsReturns(result1 []pushaction.ApplicationConfig, result2 pushaction.Warnings, result3 error) {
	fake.ConvertToApplicationConfigsStub = nil
	fake.convertToApplicationConfigsReturns = struct {
		result1 []pushaction.ApplicationConfig
		result2 pushaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2PushActor) ConvertToApplicationConfigsReturnsOnCall(i int, result1 []pushaction.ApplicationConfig, result2 pushaction.Warnings, result3 error) {
	fake.ConvertToApplicationConfigsStub = nil
	if fake.convertToApplicationConfigsReturnsOnCall == nil {
		fake.convertToApplicationConfigsReturnsOnCall = make(map[int]struct {
			result1 []pushaction.ApplicationConfig
			result2 pushaction.Warnings
			result3 error
		})
	}
	fake.convertToApplicationConfigsReturnsOnCall[i] = struct {
		result1 []pushaction.ApplicationConfig
		result2 pushaction.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV2PushActor) MergeAndValidateSettingsAndManifests(cmdSettings pushaction.CommandLineSettings, apps []manifest.Application) ([]manifest.Application, error) {
	var appsCopy []manifest.Application
	if apps != nil {
		appsCopy = make([]manifest.Application, len(apps))
		copy(appsCopy, apps)
	}
	fake.mergeAndValidateSettingsAndManifestsMutex.Lock()
	ret, specificReturn := fake.mergeAndValidateSettingsAndManifestsReturnsOnCall[len(fake.mergeAndValidateSettingsAndManifestsArgsForCall)]
	fake.mergeAndValidateSettingsAndManifestsArgsForCall = append(fake.mergeAndValidateSettingsAndManifestsArgsForCall, struct {
		cmdSettings pushaction.CommandLineSettings
		apps        []manifest.Application
	}{cmdSettings, appsCopy})
	fake.recordInvocation("MergeAndValidateSettingsAndManifests", []interface{}{cmdSettings, appsCopy})
	fake.mergeAndValidateSettingsAndManifestsMutex.Unlock()
	if fake.MergeAndValidateSettingsAndManifestsStub != nil {
		return fake.MergeAndValidateSettingsAndManifestsStub(cmdSettings, apps)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.mergeAndValidateSettingsAndManifestsReturns.result1, fake.mergeAndValidateSettingsAndManifestsReturns.result2
}

func (fake *FakeV2PushActor) MergeAndValidateSettingsAndManifestsCallCount() int {
	fake.mergeAndValidateSettingsAndManifestsMutex.RLock()
	defer fake.mergeAndValidateSettingsAndManifestsMutex.RUnlock()
	return len(fake.mergeAndValidateSettingsAndManifestsArgsForCall)
}

func (fake *FakeV2PushActor) MergeAndValidateSettingsAndManifestsArgsForCall(i int) (pushaction.CommandLineSettings, []manifest.Application) {
	fake.mergeAndValidateSettingsAndManifestsMutex.RLock()
	defer fake.mergeAndValidateSettingsAndManifestsMutex.RUnlock()
	return fake.mergeAndValidateSettingsAndManifestsArgsForCall[i].cmdSettings, fake.mergeAndValidateSettingsAndManifestsArgsForCall[i].apps
}

func (fake *FakeV2PushActor) MergeAndValidateSettingsAndManifestsReturns(result1 []manifest.Application, result2 error) {
	fake.MergeAndValidateSettingsAndManifestsStub = nil
	fake.mergeAndValidateSettingsAndManifestsReturns = struct {
		result1 []manifest.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) MergeAndValidateSettingsAndManifestsReturnsOnCall(i int, result1 []manifest.Application, result2 error) {
	fake.MergeAndValidateSettingsAndManifestsStub = nil
	if fake.mergeAndValidateSettingsAndManifestsReturnsOnCall == nil {
		fake.mergeAndValidateSettingsAndManifestsReturnsOnCall = make(map[int]struct {
			result1 []manifest.Application
			result2 error
		})
	}
	fake.mergeAndValidateSettingsAndManifestsReturnsOnCall[i] = struct {
		result1 []manifest.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) ReadManifest(pathToManifest string) ([]manifest.Application, error) {
	fake.readManifestMutex.Lock()
	ret, specificReturn := fake.readManifestReturnsOnCall[len(fake.readManifestArgsForCall)]
	fake.readManifestArgsForCall = append(fake.readManifestArgsForCall, struct {
		pathToManifest string
	}{pathToManifest})
	fake.recordInvocation("ReadManifest", []interface{}{pathToManifest})
	fake.readManifestMutex.Unlock()
	if fake.ReadManifestStub != nil {
		return fake.ReadManifestStub(pathToManifest)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readManifestReturns.result1, fake.readManifestReturns.result2
}

func (fake *FakeV2PushActor) ReadManifestCallCount() int {
	fake.readManifestMutex.RLock()
	defer fake.readManifestMutex.RUnlock()
	return len(fake.readManifestArgsForCall)
}

func (fake *FakeV2PushActor) ReadManifestArgsForCall(i int) string {
	fake.readManifestMutex.RLock()
	defer fake.readManifestMutex.RUnlock()
	return fake.readManifestArgsForCall[i].pathToManifest
}

func (fake *FakeV2PushActor) ReadManifestReturns(result1 []manifest.Application, result2 error) {
	fake.ReadManifestStub = nil
	fake.readManifestReturns = struct {
		result1 []manifest.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) ReadManifestReturnsOnCall(i int, result1 []manifest.Application, result2 error) {
	fake.ReadManifestStub = nil
	if fake.readManifestReturnsOnCall == nil {
		fake.readManifestReturnsOnCall = make(map[int]struct {
			result1 []manifest.Application
			result2 error
		})
	}
	fake.readManifestReturnsOnCall[i] = struct {
		result1 []manifest.Application
		result2 error
	}{result1, result2}
}

func (fake *FakeV2PushActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.applyMutex.RLock()
	defer fake.applyMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.convertToApplicationConfigsMutex.RLock()
	defer fake.convertToApplicationConfigsMutex.RUnlock()
	fake.mergeAndValidateSettingsAndManifestsMutex.RLock()
	defer fake.mergeAndValidateSettingsAndManifestsMutex.RUnlock()
	fake.readManifestMutex.RLock()
	defer fake.readManifestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV2PushActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.V2PushActor = new(FakeV2PushActor)
