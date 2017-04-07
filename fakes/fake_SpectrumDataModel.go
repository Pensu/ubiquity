// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/IBM/ubiquity/local/spectrumscale"
	_ "github.com/mattn/go-sqlite3"
)

type FakeSpectrumDataModel struct {
	CreateVolumeTableStub        func() error
	createVolumeTableMutex       sync.RWMutex
	createVolumeTableArgsForCall []struct{}
	createVolumeTableReturns     struct {
		result1 error
	}
	createVolumeTableReturnsOnCall map[int]struct {
		result1 error
	}
	SetClusterIdStub        func(string)
	setClusterIdMutex       sync.RWMutex
	setClusterIdArgsForCall []struct {
		arg1 string
	}
	GetClusterIdStub        func() string
	getClusterIdMutex       sync.RWMutex
	getClusterIdArgsForCall []struct{}
	getClusterIdReturns     struct {
		result1 string
	}
	getClusterIdReturnsOnCall map[int]struct {
		result1 string
	}
	DeleteVolumeStub        func(name string) error
	deleteVolumeMutex       sync.RWMutex
	deleteVolumeArgsForCall []struct {
		name string
	}
	deleteVolumeReturns struct {
		result1 error
	}
	deleteVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	InsertFilesetVolumeStub        func(fileset, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error
	insertFilesetVolumeMutex       sync.RWMutex
	insertFilesetVolumeArgsForCall []struct {
		fileset       string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}
	insertFilesetVolumeReturns struct {
		result1 error
	}
	insertFilesetVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	InsertLightweightVolumeStub        func(fileset, directory, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error
	insertLightweightVolumeMutex       sync.RWMutex
	insertLightweightVolumeArgsForCall []struct {
		fileset       string
		directory     string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}
	insertLightweightVolumeReturns struct {
		result1 error
	}
	insertLightweightVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	InsertFilesetQuotaVolumeStub        func(fileset, quota, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error
	insertFilesetQuotaVolumeMutex       sync.RWMutex
	insertFilesetQuotaVolumeArgsForCall []struct {
		fileset       string
		quota         string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}
	insertFilesetQuotaVolumeReturns struct {
		result1 error
	}
	insertFilesetQuotaVolumeReturnsOnCall map[int]struct {
		result1 error
	}
	GetVolumeStub        func(name string) (spectrumscale.SpectrumScaleVolume, bool, error)
	getVolumeMutex       sync.RWMutex
	getVolumeArgsForCall []struct {
		name string
	}
	getVolumeReturns struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}
	getVolumeReturnsOnCall map[int]struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}
	ListVolumesStub        func() ([]spectrumscale.SpectrumScaleVolume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct{}
	listVolumesReturns     struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}
	listVolumesReturnsOnCall map[int]struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpectrumDataModel) CreateVolumeTable() error {
	fake.createVolumeTableMutex.Lock()
	ret, specificReturn := fake.createVolumeTableReturnsOnCall[len(fake.createVolumeTableArgsForCall)]
	fake.createVolumeTableArgsForCall = append(fake.createVolumeTableArgsForCall, struct{}{})
	fake.recordInvocation("CreateVolumeTable", []interface{}{})
	fake.createVolumeTableMutex.Unlock()
	if fake.CreateVolumeTableStub != nil {
		return fake.CreateVolumeTableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createVolumeTableReturns.result1
}

func (fake *FakeSpectrumDataModel) CreateVolumeTableCallCount() int {
	fake.createVolumeTableMutex.RLock()
	defer fake.createVolumeTableMutex.RUnlock()
	return len(fake.createVolumeTableArgsForCall)
}

func (fake *FakeSpectrumDataModel) CreateVolumeTableReturns(result1 error) {
	fake.CreateVolumeTableStub = nil
	fake.createVolumeTableReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) CreateVolumeTableReturnsOnCall(i int, result1 error) {
	fake.CreateVolumeTableStub = nil
	if fake.createVolumeTableReturnsOnCall == nil {
		fake.createVolumeTableReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createVolumeTableReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) SetClusterId(arg1 string) {
	fake.setClusterIdMutex.Lock()
	fake.setClusterIdArgsForCall = append(fake.setClusterIdArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetClusterId", []interface{}{arg1})
	fake.setClusterIdMutex.Unlock()
	if fake.SetClusterIdStub != nil {
		fake.SetClusterIdStub(arg1)
	}
}

func (fake *FakeSpectrumDataModel) SetClusterIdCallCount() int {
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	return len(fake.setClusterIdArgsForCall)
}

func (fake *FakeSpectrumDataModel) SetClusterIdArgsForCall(i int) string {
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	return fake.setClusterIdArgsForCall[i].arg1
}

func (fake *FakeSpectrumDataModel) GetClusterId() string {
	fake.getClusterIdMutex.Lock()
	ret, specificReturn := fake.getClusterIdReturnsOnCall[len(fake.getClusterIdArgsForCall)]
	fake.getClusterIdArgsForCall = append(fake.getClusterIdArgsForCall, struct{}{})
	fake.recordInvocation("GetClusterId", []interface{}{})
	fake.getClusterIdMutex.Unlock()
	if fake.GetClusterIdStub != nil {
		return fake.GetClusterIdStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getClusterIdReturns.result1
}

func (fake *FakeSpectrumDataModel) GetClusterIdCallCount() int {
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	return len(fake.getClusterIdArgsForCall)
}

func (fake *FakeSpectrumDataModel) GetClusterIdReturns(result1 string) {
	fake.GetClusterIdStub = nil
	fake.getClusterIdReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpectrumDataModel) GetClusterIdReturnsOnCall(i int, result1 string) {
	fake.GetClusterIdStub = nil
	if fake.getClusterIdReturnsOnCall == nil {
		fake.getClusterIdReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getClusterIdReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSpectrumDataModel) DeleteVolume(name string) error {
	fake.deleteVolumeMutex.Lock()
	ret, specificReturn := fake.deleteVolumeReturnsOnCall[len(fake.deleteVolumeArgsForCall)]
	fake.deleteVolumeArgsForCall = append(fake.deleteVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteVolume", []interface{}{name})
	fake.deleteVolumeMutex.Unlock()
	if fake.DeleteVolumeStub != nil {
		return fake.DeleteVolumeStub(name)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) DeleteVolumeCallCount() int {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return len(fake.deleteVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) DeleteVolumeArgsForCall(i int) string {
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	return fake.deleteVolumeArgsForCall[i].name
}

func (fake *FakeSpectrumDataModel) DeleteVolumeReturns(result1 error) {
	fake.DeleteVolumeStub = nil
	fake.deleteVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) DeleteVolumeReturnsOnCall(i int, result1 error) {
	fake.DeleteVolumeStub = nil
	if fake.deleteVolumeReturnsOnCall == nil {
		fake.deleteVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolume(fileset string, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error {
	fake.insertFilesetVolumeMutex.Lock()
	ret, specificReturn := fake.insertFilesetVolumeReturnsOnCall[len(fake.insertFilesetVolumeArgsForCall)]
	fake.insertFilesetVolumeArgsForCall = append(fake.insertFilesetVolumeArgsForCall, struct {
		fileset       string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}{fileset, volumeName, filesystem, isPreexisting, opts})
	fake.recordInvocation("InsertFilesetVolume", []interface{}{fileset, volumeName, filesystem, isPreexisting, opts})
	fake.insertFilesetVolumeMutex.Unlock()
	if fake.InsertFilesetVolumeStub != nil {
		return fake.InsertFilesetVolumeStub(fileset, volumeName, filesystem, isPreexisting, opts)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertFilesetVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeCallCount() int {
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	return len(fake.insertFilesetVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeArgsForCall(i int) (string, string, string, bool, map[string]interface{}) {
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	return fake.insertFilesetVolumeArgsForCall[i].fileset, fake.insertFilesetVolumeArgsForCall[i].volumeName, fake.insertFilesetVolumeArgsForCall[i].filesystem, fake.insertFilesetVolumeArgsForCall[i].isPreexisting, fake.insertFilesetVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeReturns(result1 error) {
	fake.InsertFilesetVolumeStub = nil
	fake.insertFilesetVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetVolumeReturnsOnCall(i int, result1 error) {
	fake.InsertFilesetVolumeStub = nil
	if fake.insertFilesetVolumeReturnsOnCall == nil {
		fake.insertFilesetVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertFilesetVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolume(fileset string, directory string, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error {
	fake.insertLightweightVolumeMutex.Lock()
	ret, specificReturn := fake.insertLightweightVolumeReturnsOnCall[len(fake.insertLightweightVolumeArgsForCall)]
	fake.insertLightweightVolumeArgsForCall = append(fake.insertLightweightVolumeArgsForCall, struct {
		fileset       string
		directory     string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}{fileset, directory, volumeName, filesystem, isPreexisting, opts})
	fake.recordInvocation("InsertLightweightVolume", []interface{}{fileset, directory, volumeName, filesystem, isPreexisting, opts})
	fake.insertLightweightVolumeMutex.Unlock()
	if fake.InsertLightweightVolumeStub != nil {
		return fake.InsertLightweightVolumeStub(fileset, directory, volumeName, filesystem, isPreexisting, opts)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertLightweightVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeCallCount() int {
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	return len(fake.insertLightweightVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeArgsForCall(i int) (string, string, string, string, bool, map[string]interface{}) {
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	return fake.insertLightweightVolumeArgsForCall[i].fileset, fake.insertLightweightVolumeArgsForCall[i].directory, fake.insertLightweightVolumeArgsForCall[i].volumeName, fake.insertLightweightVolumeArgsForCall[i].filesystem, fake.insertLightweightVolumeArgsForCall[i].isPreexisting, fake.insertLightweightVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeReturns(result1 error) {
	fake.InsertLightweightVolumeStub = nil
	fake.insertLightweightVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertLightweightVolumeReturnsOnCall(i int, result1 error) {
	fake.InsertLightweightVolumeStub = nil
	if fake.insertLightweightVolumeReturnsOnCall == nil {
		fake.insertLightweightVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertLightweightVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolume(fileset string, quota string, volumeName string, filesystem string, isPreexisting bool, opts map[string]interface{}) error {
	fake.insertFilesetQuotaVolumeMutex.Lock()
	ret, specificReturn := fake.insertFilesetQuotaVolumeReturnsOnCall[len(fake.insertFilesetQuotaVolumeArgsForCall)]
	fake.insertFilesetQuotaVolumeArgsForCall = append(fake.insertFilesetQuotaVolumeArgsForCall, struct {
		fileset       string
		quota         string
		volumeName    string
		filesystem    string
		isPreexisting bool
		opts          map[string]interface{}
	}{fileset, quota, volumeName, filesystem, isPreexisting, opts})
	fake.recordInvocation("InsertFilesetQuotaVolume", []interface{}{fileset, quota, volumeName, filesystem, isPreexisting, opts})
	fake.insertFilesetQuotaVolumeMutex.Unlock()
	if fake.InsertFilesetQuotaVolumeStub != nil {
		return fake.InsertFilesetQuotaVolumeStub(fileset, quota, volumeName, filesystem, isPreexisting, opts)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.insertFilesetQuotaVolumeReturns.result1
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeCallCount() int {
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	return len(fake.insertFilesetQuotaVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeArgsForCall(i int) (string, string, string, string, bool, map[string]interface{}) {
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	return fake.insertFilesetQuotaVolumeArgsForCall[i].fileset, fake.insertFilesetQuotaVolumeArgsForCall[i].quota, fake.insertFilesetQuotaVolumeArgsForCall[i].volumeName, fake.insertFilesetQuotaVolumeArgsForCall[i].filesystem, fake.insertFilesetQuotaVolumeArgsForCall[i].isPreexisting, fake.insertFilesetQuotaVolumeArgsForCall[i].opts
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeReturns(result1 error) {
	fake.InsertFilesetQuotaVolumeStub = nil
	fake.insertFilesetQuotaVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) InsertFilesetQuotaVolumeReturnsOnCall(i int, result1 error) {
	fake.InsertFilesetQuotaVolumeStub = nil
	if fake.insertFilesetQuotaVolumeReturnsOnCall == nil {
		fake.insertFilesetQuotaVolumeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.insertFilesetQuotaVolumeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSpectrumDataModel) GetVolume(name string) (spectrumscale.SpectrumScaleVolume, bool, error) {
	fake.getVolumeMutex.Lock()
	ret, specificReturn := fake.getVolumeReturnsOnCall[len(fake.getVolumeArgsForCall)]
	fake.getVolumeArgsForCall = append(fake.getVolumeArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetVolume", []interface{}{name})
	fake.getVolumeMutex.Unlock()
	if fake.GetVolumeStub != nil {
		return fake.GetVolumeStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getVolumeReturns.result1, fake.getVolumeReturns.result2, fake.getVolumeReturns.result3
}

func (fake *FakeSpectrumDataModel) GetVolumeCallCount() int {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return len(fake.getVolumeArgsForCall)
}

func (fake *FakeSpectrumDataModel) GetVolumeArgsForCall(i int) string {
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	return fake.getVolumeArgsForCall[i].name
}

func (fake *FakeSpectrumDataModel) GetVolumeReturns(result1 spectrumscale.SpectrumScaleVolume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	fake.getVolumeReturns = struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpectrumDataModel) GetVolumeReturnsOnCall(i int, result1 spectrumscale.SpectrumScaleVolume, result2 bool, result3 error) {
	fake.GetVolumeStub = nil
	if fake.getVolumeReturnsOnCall == nil {
		fake.getVolumeReturnsOnCall = make(map[int]struct {
			result1 spectrumscale.SpectrumScaleVolume
			result2 bool
			result3 error
		})
	}
	fake.getVolumeReturnsOnCall[i] = struct {
		result1 spectrumscale.SpectrumScaleVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpectrumDataModel) ListVolumes() ([]spectrumscale.SpectrumScaleVolume, error) {
	fake.listVolumesMutex.Lock()
	ret, specificReturn := fake.listVolumesReturnsOnCall[len(fake.listVolumesArgsForCall)]
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct{}{})
	fake.recordInvocation("ListVolumes", []interface{}{})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
}

func (fake *FakeSpectrumDataModel) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeSpectrumDataModel) ListVolumesReturns(result1 []spectrumscale.SpectrumScaleVolume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumDataModel) ListVolumesReturnsOnCall(i int, result1 []spectrumscale.SpectrumScaleVolume, result2 error) {
	fake.ListVolumesStub = nil
	if fake.listVolumesReturnsOnCall == nil {
		fake.listVolumesReturnsOnCall = make(map[int]struct {
			result1 []spectrumscale.SpectrumScaleVolume
			result2 error
		})
	}
	fake.listVolumesReturnsOnCall[i] = struct {
		result1 []spectrumscale.SpectrumScaleVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeSpectrumDataModel) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createVolumeTableMutex.RLock()
	defer fake.createVolumeTableMutex.RUnlock()
	fake.setClusterIdMutex.RLock()
	defer fake.setClusterIdMutex.RUnlock()
	fake.getClusterIdMutex.RLock()
	defer fake.getClusterIdMutex.RUnlock()
	fake.deleteVolumeMutex.RLock()
	defer fake.deleteVolumeMutex.RUnlock()
	fake.insertFilesetVolumeMutex.RLock()
	defer fake.insertFilesetVolumeMutex.RUnlock()
	fake.insertLightweightVolumeMutex.RLock()
	defer fake.insertLightweightVolumeMutex.RUnlock()
	fake.insertFilesetQuotaVolumeMutex.RLock()
	defer fake.insertFilesetQuotaVolumeMutex.RUnlock()
	fake.getVolumeMutex.RLock()
	defer fake.getVolumeMutex.RUnlock()
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeSpectrumDataModel) recordInvocation(key string, args []interface{}) {
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

var _ spectrumscale.SpectrumDataModel = new(FakeSpectrumDataModel)
