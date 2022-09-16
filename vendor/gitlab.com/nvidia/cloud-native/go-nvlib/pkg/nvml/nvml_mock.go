// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package nvml

import (
	"sync"
)

// Ensure, that InterfaceMock does implement Interface.
// If this is not the case, regenerate this file with moq.
var _ Interface = &InterfaceMock{}

// InterfaceMock is a mock implementation of Interface.
//
// 	func TestSomethingThatUsesInterface(t *testing.T) {
//
// 		// make and configure a mocked Interface
// 		mockedInterface := &InterfaceMock{
// 			DeviceGetCountFunc: func() (int, Return) {
// 				panic("mock out the DeviceGetCount method")
// 			},
// 			DeviceGetHandleByIndexFunc: func(Index int) (Device, Return) {
// 				panic("mock out the DeviceGetHandleByIndex method")
// 			},
// 			DeviceGetHandleByUUIDFunc: func(UUID string) (Device, Return) {
// 				panic("mock out the DeviceGetHandleByUUID method")
// 			},
// 			ErrorStringFunc: func(r Return) string {
// 				panic("mock out the ErrorString method")
// 			},
// 			InitFunc: func() Return {
// 				panic("mock out the Init method")
// 			},
// 			ShutdownFunc: func() Return {
// 				panic("mock out the Shutdown method")
// 			},
// 			SystemGetCudaDriverVersionFunc: func() (int, Return) {
// 				panic("mock out the SystemGetCudaDriverVersion method")
// 			},
// 			SystemGetDriverVersionFunc: func() (string, Return) {
// 				panic("mock out the SystemGetDriverVersion method")
// 			},
// 		}
//
// 		// use mockedInterface in code that requires Interface
// 		// and then make assertions.
//
// 	}
type InterfaceMock struct {
	// DeviceGetCountFunc mocks the DeviceGetCount method.
	DeviceGetCountFunc func() (int, Return)

	// DeviceGetHandleByIndexFunc mocks the DeviceGetHandleByIndex method.
	DeviceGetHandleByIndexFunc func(Index int) (Device, Return)

	// DeviceGetHandleByUUIDFunc mocks the DeviceGetHandleByUUID method.
	DeviceGetHandleByUUIDFunc func(UUID string) (Device, Return)

	// ErrorStringFunc mocks the ErrorString method.
	ErrorStringFunc func(r Return) string

	// InitFunc mocks the Init method.
	InitFunc func() Return

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func() Return

	// SystemGetCudaDriverVersionFunc mocks the SystemGetCudaDriverVersion method.
	SystemGetCudaDriverVersionFunc func() (int, Return)

	// SystemGetDriverVersionFunc mocks the SystemGetDriverVersion method.
	SystemGetDriverVersionFunc func() (string, Return)

	// calls tracks calls to the methods.
	calls struct {
		// DeviceGetCount holds details about calls to the DeviceGetCount method.
		DeviceGetCount []struct {
		}
		// DeviceGetHandleByIndex holds details about calls to the DeviceGetHandleByIndex method.
		DeviceGetHandleByIndex []struct {
			// Index is the Index argument value.
			Index int
		}
		// DeviceGetHandleByUUID holds details about calls to the DeviceGetHandleByUUID method.
		DeviceGetHandleByUUID []struct {
			// UUID is the UUID argument value.
			UUID string
		}
		// ErrorString holds details about calls to the ErrorString method.
		ErrorString []struct {
			// R is the r argument value.
			R Return
		}
		// Init holds details about calls to the Init method.
		Init []struct {
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
		}
		// SystemGetCudaDriverVersion holds details about calls to the SystemGetCudaDriverVersion method.
		SystemGetCudaDriverVersion []struct {
		}
		// SystemGetDriverVersion holds details about calls to the SystemGetDriverVersion method.
		SystemGetDriverVersion []struct {
		}
	}
	lockDeviceGetCount             sync.RWMutex
	lockDeviceGetHandleByIndex     sync.RWMutex
	lockDeviceGetHandleByUUID      sync.RWMutex
	lockErrorString                sync.RWMutex
	lockInit                       sync.RWMutex
	lockShutdown                   sync.RWMutex
	lockSystemGetCudaDriverVersion sync.RWMutex
	lockSystemGetDriverVersion     sync.RWMutex
}

// DeviceGetCount calls DeviceGetCountFunc.
func (mock *InterfaceMock) DeviceGetCount() (int, Return) {
	if mock.DeviceGetCountFunc == nil {
		panic("InterfaceMock.DeviceGetCountFunc: method is nil but Interface.DeviceGetCount was just called")
	}
	callInfo := struct {
	}{}
	mock.lockDeviceGetCount.Lock()
	mock.calls.DeviceGetCount = append(mock.calls.DeviceGetCount, callInfo)
	mock.lockDeviceGetCount.Unlock()
	return mock.DeviceGetCountFunc()
}

// DeviceGetCountCalls gets all the calls that were made to DeviceGetCount.
// Check the length with:
//     len(mockedInterface.DeviceGetCountCalls())
func (mock *InterfaceMock) DeviceGetCountCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockDeviceGetCount.RLock()
	calls = mock.calls.DeviceGetCount
	mock.lockDeviceGetCount.RUnlock()
	return calls
}

// DeviceGetHandleByIndex calls DeviceGetHandleByIndexFunc.
func (mock *InterfaceMock) DeviceGetHandleByIndex(Index int) (Device, Return) {
	if mock.DeviceGetHandleByIndexFunc == nil {
		panic("InterfaceMock.DeviceGetHandleByIndexFunc: method is nil but Interface.DeviceGetHandleByIndex was just called")
	}
	callInfo := struct {
		Index int
	}{
		Index: Index,
	}
	mock.lockDeviceGetHandleByIndex.Lock()
	mock.calls.DeviceGetHandleByIndex = append(mock.calls.DeviceGetHandleByIndex, callInfo)
	mock.lockDeviceGetHandleByIndex.Unlock()
	return mock.DeviceGetHandleByIndexFunc(Index)
}

// DeviceGetHandleByIndexCalls gets all the calls that were made to DeviceGetHandleByIndex.
// Check the length with:
//     len(mockedInterface.DeviceGetHandleByIndexCalls())
func (mock *InterfaceMock) DeviceGetHandleByIndexCalls() []struct {
	Index int
} {
	var calls []struct {
		Index int
	}
	mock.lockDeviceGetHandleByIndex.RLock()
	calls = mock.calls.DeviceGetHandleByIndex
	mock.lockDeviceGetHandleByIndex.RUnlock()
	return calls
}

// DeviceGetHandleByUUID calls DeviceGetHandleByUUIDFunc.
func (mock *InterfaceMock) DeviceGetHandleByUUID(UUID string) (Device, Return) {
	if mock.DeviceGetHandleByUUIDFunc == nil {
		panic("InterfaceMock.DeviceGetHandleByUUIDFunc: method is nil but Interface.DeviceGetHandleByUUID was just called")
	}
	callInfo := struct {
		UUID string
	}{
		UUID: UUID,
	}
	mock.lockDeviceGetHandleByUUID.Lock()
	mock.calls.DeviceGetHandleByUUID = append(mock.calls.DeviceGetHandleByUUID, callInfo)
	mock.lockDeviceGetHandleByUUID.Unlock()
	return mock.DeviceGetHandleByUUIDFunc(UUID)
}

// DeviceGetHandleByUUIDCalls gets all the calls that were made to DeviceGetHandleByUUID.
// Check the length with:
//     len(mockedInterface.DeviceGetHandleByUUIDCalls())
func (mock *InterfaceMock) DeviceGetHandleByUUIDCalls() []struct {
	UUID string
} {
	var calls []struct {
		UUID string
	}
	mock.lockDeviceGetHandleByUUID.RLock()
	calls = mock.calls.DeviceGetHandleByUUID
	mock.lockDeviceGetHandleByUUID.RUnlock()
	return calls
}

// ErrorString calls ErrorStringFunc.
func (mock *InterfaceMock) ErrorString(r Return) string {
	if mock.ErrorStringFunc == nil {
		panic("InterfaceMock.ErrorStringFunc: method is nil but Interface.ErrorString was just called")
	}
	callInfo := struct {
		R Return
	}{
		R: r,
	}
	mock.lockErrorString.Lock()
	mock.calls.ErrorString = append(mock.calls.ErrorString, callInfo)
	mock.lockErrorString.Unlock()
	return mock.ErrorStringFunc(r)
}

// ErrorStringCalls gets all the calls that were made to ErrorString.
// Check the length with:
//     len(mockedInterface.ErrorStringCalls())
func (mock *InterfaceMock) ErrorStringCalls() []struct {
	R Return
} {
	var calls []struct {
		R Return
	}
	mock.lockErrorString.RLock()
	calls = mock.calls.ErrorString
	mock.lockErrorString.RUnlock()
	return calls
}

// Init calls InitFunc.
func (mock *InterfaceMock) Init() Return {
	if mock.InitFunc == nil {
		panic("InterfaceMock.InitFunc: method is nil but Interface.Init was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInit.Lock()
	mock.calls.Init = append(mock.calls.Init, callInfo)
	mock.lockInit.Unlock()
	return mock.InitFunc()
}

// InitCalls gets all the calls that were made to Init.
// Check the length with:
//     len(mockedInterface.InitCalls())
func (mock *InterfaceMock) InitCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInit.RLock()
	calls = mock.calls.Init
	mock.lockInit.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *InterfaceMock) Shutdown() Return {
	if mock.ShutdownFunc == nil {
		panic("InterfaceMock.ShutdownFunc: method is nil but Interface.Shutdown was just called")
	}
	callInfo := struct {
	}{}
	mock.lockShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	mock.lockShutdown.Unlock()
	return mock.ShutdownFunc()
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedInterface.ShutdownCalls())
func (mock *InterfaceMock) ShutdownCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockShutdown.RLock()
	calls = mock.calls.Shutdown
	mock.lockShutdown.RUnlock()
	return calls
}

// SystemGetCudaDriverVersion calls SystemGetCudaDriverVersionFunc.
func (mock *InterfaceMock) SystemGetCudaDriverVersion() (int, Return) {
	if mock.SystemGetCudaDriverVersionFunc == nil {
		panic("InterfaceMock.SystemGetCudaDriverVersionFunc: method is nil but Interface.SystemGetCudaDriverVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSystemGetCudaDriverVersion.Lock()
	mock.calls.SystemGetCudaDriverVersion = append(mock.calls.SystemGetCudaDriverVersion, callInfo)
	mock.lockSystemGetCudaDriverVersion.Unlock()
	return mock.SystemGetCudaDriverVersionFunc()
}

// SystemGetCudaDriverVersionCalls gets all the calls that were made to SystemGetCudaDriverVersion.
// Check the length with:
//     len(mockedInterface.SystemGetCudaDriverVersionCalls())
func (mock *InterfaceMock) SystemGetCudaDriverVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSystemGetCudaDriverVersion.RLock()
	calls = mock.calls.SystemGetCudaDriverVersion
	mock.lockSystemGetCudaDriverVersion.RUnlock()
	return calls
}

// SystemGetDriverVersion calls SystemGetDriverVersionFunc.
func (mock *InterfaceMock) SystemGetDriverVersion() (string, Return) {
	if mock.SystemGetDriverVersionFunc == nil {
		panic("InterfaceMock.SystemGetDriverVersionFunc: method is nil but Interface.SystemGetDriverVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSystemGetDriverVersion.Lock()
	mock.calls.SystemGetDriverVersion = append(mock.calls.SystemGetDriverVersion, callInfo)
	mock.lockSystemGetDriverVersion.Unlock()
	return mock.SystemGetDriverVersionFunc()
}

// SystemGetDriverVersionCalls gets all the calls that were made to SystemGetDriverVersion.
// Check the length with:
//     len(mockedInterface.SystemGetDriverVersionCalls())
func (mock *InterfaceMock) SystemGetDriverVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSystemGetDriverVersion.RLock()
	calls = mock.calls.SystemGetDriverVersion
	mock.lockSystemGetDriverVersion.RUnlock()
	return calls
}
