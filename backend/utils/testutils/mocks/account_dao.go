// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bson "github.com/globalsign/mgo/bson"

	mock "github.com/stretchr/testify/mock"

	types "github.com/citypayorg/udex/tree/udex/backend/types"
)

// AccountDao is an autogenerated mock type for the AccountDao type
type AccountDao struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountDao) Create(account *types.Account) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Account) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Drop provides a mock function with given fields:
func (_m *AccountDao) Drop() {
	_m.Called()
}

// FindOrCreate provides a mock function with given fields: addr
func (_m *AccountDao) FindOrCreate(addr string) (*types.Account, error) {
	ret := _m.Called(addr)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(string) *types.Account); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *AccountDao) GetAll() ([]types.Account, error) {
	ret := _m.Called()

	var r0 []types.Account
	if rf, ok := ret.Get(0).(func() []types.Account); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAddress provides a mock function with given fields: owner
func (_m *AccountDao) GetByAddress(owner string) (*types.Account, error) {
	ret := _m.Called(owner)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(string) *types.Account); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *AccountDao) GetByID(id bson.ObjectId) (*types.Account, error) {
	ret := _m.Called(id)

	var r0 *types.Account
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Account); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalance provides a mock function with given fields: owner, token
func (_m *AccountDao) GetTokenBalance(owner string, token string) (*types.TokenBalance, error) {
	ret := _m.Called(owner, token)

	var r0 *types.TokenBalance
	if rf, ok := ret.Get(0).(func(string, string) *types.TokenBalance); ok {
		r0 = rf(owner, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(owner, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenBalances provides a mock function with given fields: owner
func (_m *AccountDao) GetTokenBalances(owner string) (map[string]*types.TokenBalance, error) {
	ret := _m.Called(owner)

	var r0 map[string]*types.TokenBalance
	if rf, ok := ret.Get(0).(func(string) map[string]*types.TokenBalance); ok {
		r0 = rf(owner)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*types.TokenBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(owner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBalance provides a mock function with given fields: owner, token, balance
func (_m *AccountDao) UpdateBalance(owner string, token string, balance int64) error {
	ret := _m.Called(owner, token, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(owner, token, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTokenBalance provides a mock function with given fields: owner, token, tokenBalance
func (_m *AccountDao) UpdateTokenBalance(owner string, token string, tokenBalance *types.TokenBalance) error {
	ret := _m.Called(owner, token, tokenBalance)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *types.TokenBalance) error); ok {
		r0 = rf(owner, token, tokenBalance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
