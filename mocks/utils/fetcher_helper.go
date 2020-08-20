// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package utils

import (
	context "context"

	types "github.com/coinbase/rosetta-sdk-go/types"
	mock "github.com/stretchr/testify/mock"
)

// FetcherHelper is an autogenerated mock type for the FetcherHelper type
type FetcherHelper struct {
	mock.Mock
}

// AccountBalanceRetry provides a mock function with given fields: ctx, network, account, block
func (_m *FetcherHelper) AccountBalanceRetry(ctx context.Context, network *types.NetworkIdentifier, account *types.AccountIdentifier, block *types.PartialBlockIdentifier) (*types.BlockIdentifier, []*types.Amount, []*types.Coin, map[string]interface{}, *types.Error) {
	ret := _m.Called(ctx, network, account, block)

	var r0 *types.BlockIdentifier
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier) *types.BlockIdentifier); ok {
		r0 = rf(ctx, network, account, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.BlockIdentifier)
		}
	}

	var r1 []*types.Amount
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier) []*types.Amount); ok {
		r1 = rf(ctx, network, account, block)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*types.Amount)
		}
	}

	var r2 []*types.Coin
	if rf, ok := ret.Get(2).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier) []*types.Coin); ok {
		r2 = rf(ctx, network, account, block)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]*types.Coin)
		}
	}

	var r3 map[string]interface{}
	if rf, ok := ret.Get(3).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier) map[string]interface{}); ok {
		r3 = rf(ctx, network, account, block)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(map[string]interface{})
		}
	}

	var r4 *types.Error
	if rf, ok := ret.Get(4).(func(context.Context, *types.NetworkIdentifier, *types.AccountIdentifier, *types.PartialBlockIdentifier) *types.Error); ok {
		r4 = rf(ctx, network, account, block)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).(*types.Error)
		}
	}

	return r0, r1, r2, r3, r4
}

// NetworkList provides a mock function with given fields: ctx, metadata
func (_m *FetcherHelper) NetworkList(ctx context.Context, metadata map[string]interface{}) (*types.NetworkListResponse, *types.Error) {
	ret := _m.Called(ctx, metadata)

	var r0 *types.NetworkListResponse
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) *types.NetworkListResponse); ok {
		r0 = rf(ctx, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkListResponse)
		}
	}

	var r1 *types.Error
	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) *types.Error); ok {
		r1 = rf(ctx, metadata)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}

// NetworkStatusRetry provides a mock function with given fields: ctx, network, metadata
func (_m *FetcherHelper) NetworkStatusRetry(ctx context.Context, network *types.NetworkIdentifier, metadata map[string]interface{}) (*types.NetworkStatusResponse, *types.Error) {
	ret := _m.Called(ctx, network, metadata)

	var r0 *types.NetworkStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) *types.NetworkStatusResponse); ok {
		r0 = rf(ctx, network, metadata)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	var r1 *types.Error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, map[string]interface{}) *types.Error); ok {
		r1 = rf(ctx, network, metadata)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*types.Error)
		}
	}

	return r0, r1
}
