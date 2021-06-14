// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package accounts

import (
	"context"
	"github.com/sptGabriel/banking/app/domain/entities/accounts"
	"github.com/sptGabriel/banking/app/domain/vos"
	"sync"
)

// Ensure, that UseCaseMock does implement UseCase.
// If this is not the case, regenerate this file with moq.
var _ UseCase = &UseCaseMock{}

// UseCaseMock is a mock implementation of UseCase.
//
// 	func TestSomethingThatUsesUseCase(t *testing.T) {
//
// 		// make and configure a mocked UseCase
// 		mockedUseCase := &UseCaseMock{
// 			CreateAccountFunc: func(ctx context.Context, account accounts.Account) error {
// 				panic("mock out the CreateAccount method")
// 			},
// 			GetAllFunc: func(ctx context.Context) ([]accounts.Account, error) {
// 				panic("mock out the GetAll method")
// 			},
// 			GetBalanceFunc: func(ctx context.Context, id vos.AccountId) (*accounts.Account, error) {
// 				panic("mock out the GetBalance method")
// 			},
// 		}
//
// 		// use mockedUseCase in code that requires UseCase
// 		// and then make assertions.
//
// 	}
type UseCaseMock struct {
	// CreateAccountFunc mocks the CreateAccount method.
	CreateAccountFunc func(ctx context.Context, account accounts.Account) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) ([]accounts.Account, error)

	// GetBalanceFunc mocks the GetBalance method.
	GetBalanceFunc func(ctx context.Context, id vos.AccountId) (*accounts.Account, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateAccount holds details about calls to the CreateAccount method.
		CreateAccount []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Account is the account argument value.
			Account accounts.Account
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetBalance holds details about calls to the GetBalance method.
		GetBalance []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ID is the id argument value.
			ID vos.AccountId
		}
	}
	lockCreateAccount sync.RWMutex
	lockGetAll        sync.RWMutex
	lockGetBalance    sync.RWMutex
}

// CreateAccount calls CreateAccountFunc.
func (mock *UseCaseMock) CreateAccount(ctx context.Context, account accounts.Account) error {
	if mock.CreateAccountFunc == nil {
		panic("UseCaseMock.CreateAccountFunc: method is nil but UseCase.CreateAccount was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Account accounts.Account
	}{
		Ctx:     ctx,
		Account: account,
	}
	mock.lockCreateAccount.Lock()
	mock.calls.CreateAccount = append(mock.calls.CreateAccount, callInfo)
	mock.lockCreateAccount.Unlock()
	return mock.CreateAccountFunc(ctx, account)
}

// CreateAccountCalls gets all the calls that were made to CreateAccount.
// Check the length with:
//     len(mockedUseCase.CreateAccountCalls())
func (mock *UseCaseMock) CreateAccountCalls() []struct {
	Ctx     context.Context
	Account accounts.Account
} {
	var calls []struct {
		Ctx     context.Context
		Account accounts.Account
	}
	mock.lockCreateAccount.RLock()
	calls = mock.calls.CreateAccount
	mock.lockCreateAccount.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *UseCaseMock) GetAll(ctx context.Context) ([]accounts.Account, error) {
	if mock.GetAllFunc == nil {
		panic("UseCaseMock.GetAllFunc: method is nil but UseCase.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//     len(mockedUseCase.GetAllCalls())
func (mock *UseCaseMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetBalance calls GetBalanceFunc.
func (mock *UseCaseMock) GetBalance(ctx context.Context, id vos.AccountId) (*accounts.Account, error) {
	if mock.GetBalanceFunc == nil {
		panic("UseCaseMock.GetBalanceFunc: method is nil but UseCase.GetBalance was just called")
	}
	callInfo := struct {
		Ctx context.Context
		ID  vos.AccountId
	}{
		Ctx: ctx,
		ID:  id,
	}
	mock.lockGetBalance.Lock()
	mock.calls.GetBalance = append(mock.calls.GetBalance, callInfo)
	mock.lockGetBalance.Unlock()
	return mock.GetBalanceFunc(ctx, id)
}

// GetBalanceCalls gets all the calls that were made to GetBalance.
// Check the length with:
//     len(mockedUseCase.GetBalanceCalls())
func (mock *UseCaseMock) GetBalanceCalls() []struct {
	Ctx context.Context
	ID  vos.AccountId
} {
	var calls []struct {
		Ctx context.Context
		ID  vos.AccountId
	}
	mock.lockGetBalance.RLock()
	calls = mock.calls.GetBalance
	mock.lockGetBalance.RUnlock()
	return calls
}
