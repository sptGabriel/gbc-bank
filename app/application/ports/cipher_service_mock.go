// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ports

import (
	"sync"
)

// Ensure, that CipherServiceMock does implement CipherService.
// If this is not the case, regenerate this file with moq.
var _ CipherService = &CipherServiceMock{}

// CipherServiceMock is a mock implementation of CipherService.
//
// 	func TestSomethingThatUsesCipherService(t *testing.T) {
//
// 		// make and configure a mocked CipherService
// 		mockedCipherService := &CipherServiceMock{
// 			DecryptFunc: func(val string) (string, error) {
// 				panic("mock out the Decrypt method")
// 			},
// 			EncryptFunc: func(id string) (string, error) {
// 				panic("mock out the Encrypt method")
// 			},
// 		}
//
// 		// use mockedCipherService in code that requires CipherService
// 		// and then make assertions.
//
// 	}
type CipherServiceMock struct {
	// DecryptFunc mocks the Decrypt method.
	DecryptFunc func(val string) (string, error)

	// EncryptFunc mocks the Encrypt method.
	EncryptFunc func(id string) (string, error)

	// calls tracks calls to the methods.
	calls struct {
		// Decrypt holds details about calls to the Decrypt method.
		Decrypt []struct {
			// Val is the val argument value.
			Val string
		}
		// Encrypt holds details about calls to the Encrypt method.
		Encrypt []struct {
			// ID is the id argument value.
			ID string
		}
	}
	lockDecrypt sync.RWMutex
	lockEncrypt sync.RWMutex
}

// Decrypt calls DecryptFunc.
func (mock *CipherServiceMock) Decrypt(val string) (string, error) {
	if mock.DecryptFunc == nil {
		panic("CipherServiceMock.DecryptFunc: method is nil but CipherService.Decrypt was just called")
	}
	callInfo := struct {
		Val string
	}{
		Val: val,
	}
	mock.lockDecrypt.Lock()
	mock.calls.Decrypt = append(mock.calls.Decrypt, callInfo)
	mock.lockDecrypt.Unlock()
	return mock.DecryptFunc(val)
}

// DecryptCalls gets all the calls that were made to Decrypt.
// Check the length with:
//     len(mockedCipherService.DecryptCalls())
func (mock *CipherServiceMock) DecryptCalls() []struct {
	Val string
} {
	var calls []struct {
		Val string
	}
	mock.lockDecrypt.RLock()
	calls = mock.calls.Decrypt
	mock.lockDecrypt.RUnlock()
	return calls
}

// Encrypt calls EncryptFunc.
func (mock *CipherServiceMock) Encrypt(id string) (string, error) {
	if mock.EncryptFunc == nil {
		panic("CipherServiceMock.EncryptFunc: method is nil but CipherService.Encrypt was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	mock.lockEncrypt.Lock()
	mock.calls.Encrypt = append(mock.calls.Encrypt, callInfo)
	mock.lockEncrypt.Unlock()
	return mock.EncryptFunc(id)
}

// EncryptCalls gets all the calls that were made to Encrypt.
// Check the length with:
//     len(mockedCipherService.EncryptCalls())
func (mock *CipherServiceMock) EncryptCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	mock.lockEncrypt.RLock()
	calls = mock.calls.Encrypt
	mock.lockEncrypt.RUnlock()
	return calls
}
