// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ports

import (
	"sync"
)

// Ensure, that CipherMock does implement Cipher.
// If this is not the case, regenerate this file with moq.
var _ Cipher = &CipherMock{}

// CipherMock is a mock implementation of Cipher.
//
// 	func TestSomethingThatUsesCipher(t *testing.T) {
//
// 		// make and configure a mocked Cipher
// 		mockedCipher := &CipherMock{
// 			DecryptFunc: func(val string) (string, error) {
// 				panic("mock out the Decrypt method")
// 			},
// 			EncryptFunc: func(id string) (string, error) {
// 				panic("mock out the Encrypt method")
// 			},
// 		}
//
// 		// use mockedCipher in code that requires Cipher
// 		// and then make assertions.
//
// 	}
type CipherMock struct {
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
func (mock *CipherMock) Decrypt(val string) (string, error) {
	if mock.DecryptFunc == nil {
		panic("CipherMock.DecryptFunc: method is nil but Cipher.Decrypt was just called")
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
//     len(mockedCipher.DecryptCalls())
func (mock *CipherMock) DecryptCalls() []struct {
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
func (mock *CipherMock) Encrypt(id string) (string, error) {
	if mock.EncryptFunc == nil {
		panic("CipherMock.EncryptFunc: method is nil but Cipher.Encrypt was just called")
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
//     len(mockedCipher.EncryptCalls())
func (mock *CipherMock) EncryptCalls() []struct {
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
