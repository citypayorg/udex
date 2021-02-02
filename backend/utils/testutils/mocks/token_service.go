// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	bson "github.com/globalsign/mgo/bson"

	mock "github.com/stretchr/testify/mock"

	types "github.com/citypayorg/udex/tree/udex/backend/types"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

// CheckByAssetOrSymbol provides a mock function with given fields: a
func (_m *TokenService) CheckByAssetOrSymbol(a string) (*types.Token, error) {
	ret := _m.Called(a)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(string) *types.Token); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: token
func (_m *TokenService) Create(token *types.Token) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Token) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *TokenService) GetAll() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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

// GetBaseTokens provides a mock function with given fields:
func (_m *TokenService) GetBaseTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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

// GetByAsset provides a mock function with given fields: a
func (_m *TokenService) GetByAsset(a string) (*types.Token, error) {
	ret := _m.Called(a)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(string) *types.Token); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByAssetOrSymbol provides a mock function with given fields: a
func (_m *TokenService) GetByAssetOrSymbol(a string) (*types.Token, error) {
	ret := _m.Called(a)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(string) *types.Token); ok {
		r0 = rf(a)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(a)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *TokenService) GetByID(id bson.ObjectId) (*types.Token, error) {
	ret := _m.Called(id)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Token); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
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

// GetBySymbol provides a mock function with given fields: s
func (_m *TokenService) GetBySymbol(s string) (*types.Token, error) {
	ret := _m.Called(s)

	var r0 *types.Token
	if rf, ok := ret.Get(0).(func(string) *types.Token); ok {
		r0 = rf(s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListedBaseTokens provides a mock function with given fields:
func (_m *TokenService) GetListedBaseTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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

// GetListedTokens provides a mock function with given fields:
func (_m *TokenService) GetListedTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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

// GetQuoteTokens provides a mock function with given fields:
func (_m *TokenService) GetQuoteTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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

// GetUnlistedTokens provides a mock function with given fields:
func (_m *TokenService) GetUnlistedTokens() ([]types.Token, error) {
	ret := _m.Called()

	var r0 []types.Token
	if rf, ok := ret.Get(0).(func() []types.Token); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Token)
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
