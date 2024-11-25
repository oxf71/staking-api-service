// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dbmodel "github.com/babylonlabs-io/staking-api-service/internal/shared/db/model"
	mock "github.com/stretchr/testify/mock"

	v2dbmodel "github.com/babylonlabs-io/staking-api-service/internal/v2/db/model"
)

// V2DBClient is an autogenerated mock type for the V2DBClient type
type V2DBClient struct {
	mock.Mock
}

// DeleteUnprocessableMessage provides a mock function with given fields: ctx, Receipt
func (_m *V2DBClient) DeleteUnprocessableMessage(ctx context.Context, Receipt interface{}) error {
	ret := _m.Called(ctx, Receipt)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUnprocessableMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		r0 = rf(ctx, Receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindPkMappingsByNativeSegwitAddress provides a mock function with given fields: ctx, nativeSegwitAddresses
func (_m *V2DBClient) FindPkMappingsByNativeSegwitAddress(ctx context.Context, nativeSegwitAddresses []string) ([]*dbmodel.PkAddressMapping, error) {
	ret := _m.Called(ctx, nativeSegwitAddresses)

	if len(ret) == 0 {
		panic("no return value specified for FindPkMappingsByNativeSegwitAddress")
	}

	var r0 []*dbmodel.PkAddressMapping
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*dbmodel.PkAddressMapping, error)); ok {
		return rf(ctx, nativeSegwitAddresses)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*dbmodel.PkAddressMapping); ok {
		r0 = rf(ctx, nativeSegwitAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.PkAddressMapping)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, nativeSegwitAddresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindPkMappingsByTaprootAddress provides a mock function with given fields: ctx, taprootAddresses
func (_m *V2DBClient) FindPkMappingsByTaprootAddress(ctx context.Context, taprootAddresses []string) ([]*dbmodel.PkAddressMapping, error) {
	ret := _m.Called(ctx, taprootAddresses)

	if len(ret) == 0 {
		panic("no return value specified for FindPkMappingsByTaprootAddress")
	}

	var r0 []*dbmodel.PkAddressMapping
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*dbmodel.PkAddressMapping, error)); ok {
		return rf(ctx, taprootAddresses)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*dbmodel.PkAddressMapping); ok {
		r0 = rf(ctx, taprootAddresses)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dbmodel.PkAddressMapping)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, taprootAddresses)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUnprocessableMessages provides a mock function with given fields: ctx
func (_m *V2DBClient) FindUnprocessableMessages(ctx context.Context) ([]dbmodel.UnprocessableMessageDocument, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindUnprocessableMessages")
	}

	var r0 []dbmodel.UnprocessableMessageDocument
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dbmodel.UnprocessableMessageDocument, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dbmodel.UnprocessableMessageDocument); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dbmodel.UnprocessableMessageDocument)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOverallStats provides a mock function with given fields: ctx
func (_m *V2DBClient) GetOverallStats(ctx context.Context) (*v2dbmodel.V2OverallStatsDocument, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetOverallStats")
	}

	var r0 *v2dbmodel.V2OverallStatsDocument
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*v2dbmodel.V2OverallStatsDocument, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *v2dbmodel.V2OverallStatsDocument); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2dbmodel.V2OverallStatsDocument)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerStats provides a mock function with given fields: ctx, stakerPKHex
func (_m *V2DBClient) GetStakerStats(ctx context.Context, stakerPKHex string) (*v2dbmodel.V2StakerStatsDocument, error) {
	ret := _m.Called(ctx, stakerPKHex)

	if len(ret) == 0 {
		panic("no return value specified for GetStakerStats")
	}

	var r0 *v2dbmodel.V2StakerStatsDocument
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*v2dbmodel.V2StakerStatsDocument, error)); ok {
		return rf(ctx, stakerPKHex)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *v2dbmodel.V2StakerStatsDocument); ok {
		r0 = rf(ctx, stakerPKHex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v2dbmodel.V2StakerStatsDocument)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, stakerPKHex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertPkAddressMappings provides a mock function with given fields: ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven
func (_m *V2DBClient) InsertPkAddressMappings(ctx context.Context, stakerPkHex string, taproot string, nativeSigwitOdd string, nativeSigwitEven string) error {
	ret := _m.Called(ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven)

	if len(ret) == 0 {
		panic("no return value specified for InsertPkAddressMappings")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) error); ok {
		r0 = rf(ctx, stakerPkHex, taproot, nativeSigwitOdd, nativeSigwitEven)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ping provides a mock function with given fields: ctx
func (_m *V2DBClient) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveUnprocessableMessage provides a mock function with given fields: ctx, messageBody, receipt
func (_m *V2DBClient) SaveUnprocessableMessage(ctx context.Context, messageBody string, receipt string) error {
	ret := _m.Called(ctx, messageBody, receipt)

	if len(ret) == 0 {
		panic("no return value specified for SaveUnprocessableMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, messageBody, receipt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewV2DBClient creates a new instance of V2DBClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewV2DBClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *V2DBClient {
	mock := &V2DBClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}