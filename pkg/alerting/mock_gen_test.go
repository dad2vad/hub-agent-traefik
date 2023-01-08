// Code generated by mocktail; DO NOT EDIT.

package alerting

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/traefik/hub-agent-traefik/pkg/metrics"
)

// dataPointsFinderMock mock of DataPointsFinder.
type dataPointsFinderMock struct{ mock.Mock }

// newDataPointsFinderMock creates a new dataPointsFinderMock.
func newDataPointsFinderMock(tb testing.TB) *dataPointsFinderMock {
	tb.Helper()

	m := &dataPointsFinderMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *dataPointsFinderMock) FindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) metrics.DataPoints {
	_ret := _m.Called(table, edgeIngress, from, to)

	if _rf, ok := _ret.Get(0).(func(string, string, time.Time, time.Time) metrics.DataPoints); ok {
		return _rf(table, edgeIngress, from, to)
	}

	_ra0, _ := _ret.Get(0).(metrics.DataPoints)

	return _ra0
}

func (_m *dataPointsFinderMock) OnFindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) *dataPointsFinderFindByEdgeIngressCall {
	return &dataPointsFinderFindByEdgeIngressCall{Call: _m.Mock.On("FindByEdgeIngress", table, edgeIngress, from, to), Parent: _m}
}

func (_m *dataPointsFinderMock) OnFindByEdgeIngressRaw(table interface{}, edgeIngress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByEdgeIngressCall {
	return &dataPointsFinderFindByEdgeIngressCall{Call: _m.Mock.On("FindByEdgeIngress", table, edgeIngress, from, to), Parent: _m}
}

type dataPointsFinderFindByEdgeIngressCall struct {
	*mock.Call
	Parent *dataPointsFinderMock
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Panic(msg string) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Once() *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Twice() *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Times(i int) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) WaitUntil(w <-chan time.Time) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) After(d time.Duration) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Run(fn func(args mock.Arguments)) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) Maybe() *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) TypedReturns(a metrics.DataPoints) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) ReturnsFn(fn func(string, string, time.Time, time.Time) metrics.DataPoints) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) TypedRun(fn func(string, string, time.Time, time.Time)) *dataPointsFinderFindByEdgeIngressCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_table := args.String(0)
		_edgeIngress := args.String(1)
		_from, _ := args.Get(2).(time.Time)
		_to, _ := args.Get(3).(time.Time)
		fn(_table, _edgeIngress, _from, _to)
	})
	return _c
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngress(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByIngress(table string, ingress string, from time.Time, to time.Time) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngress(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndService(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByService(table string, service string, from time.Time, to time.Time) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByService(table, service, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByEdgeIngressRaw(table interface{}, edgeIngress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngressRaw(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByIngressRaw(table interface{}, ingress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngressRaw(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByIngressAndServiceRaw(table interface{}, ingress interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndServiceRaw(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByEdgeIngressCall) OnFindByServiceRaw(table interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByServiceRaw(table, service, from, to)
}

func (_m *dataPointsFinderMock) FindByIngress(table string, ingress string, from time.Time, to time.Time) metrics.DataPoints {
	_ret := _m.Called(table, ingress, from, to)

	if _rf, ok := _ret.Get(0).(func(string, string, time.Time, time.Time) metrics.DataPoints); ok {
		return _rf(table, ingress, from, to)
	}

	_ra0, _ := _ret.Get(0).(metrics.DataPoints)

	return _ra0
}

func (_m *dataPointsFinderMock) OnFindByIngress(table string, ingress string, from time.Time, to time.Time) *dataPointsFinderFindByIngressCall {
	return &dataPointsFinderFindByIngressCall{Call: _m.Mock.On("FindByIngress", table, ingress, from, to), Parent: _m}
}

func (_m *dataPointsFinderMock) OnFindByIngressRaw(table interface{}, ingress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressCall {
	return &dataPointsFinderFindByIngressCall{Call: _m.Mock.On("FindByIngress", table, ingress, from, to), Parent: _m}
}

type dataPointsFinderFindByIngressCall struct {
	*mock.Call
	Parent *dataPointsFinderMock
}

func (_c *dataPointsFinderFindByIngressCall) Panic(msg string) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) Once() *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) Twice() *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) Times(i int) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) WaitUntil(w <-chan time.Time) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) After(d time.Duration) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) Run(fn func(args mock.Arguments)) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) Maybe() *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) TypedReturns(a metrics.DataPoints) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) ReturnsFn(fn func(string, string, time.Time, time.Time) metrics.DataPoints) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) TypedRun(fn func(string, string, time.Time, time.Time)) *dataPointsFinderFindByIngressCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_table := args.String(0)
		_ingress := args.String(1)
		_from, _ := args.Get(2).(time.Time)
		_to, _ := args.Get(3).(time.Time)
		fn(_table, _ingress, _from, _to)
	})
	return _c
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngress(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByIngress(table string, ingress string, from time.Time, to time.Time) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngress(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndService(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByService(table string, service string, from time.Time, to time.Time) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByService(table, service, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByEdgeIngressRaw(table interface{}, edgeIngress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngressRaw(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByIngressRaw(table interface{}, ingress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngressRaw(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByIngressAndServiceRaw(table interface{}, ingress interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndServiceRaw(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByIngressCall) OnFindByServiceRaw(table interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByServiceRaw(table, service, from, to)
}

func (_m *dataPointsFinderMock) FindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) (metrics.DataPoints, error) {
	_ret := _m.Called(table, ingress, service, from, to)

	if _rf, ok := _ret.Get(0).(func(string, string, string, time.Time, time.Time) (metrics.DataPoints, error)); ok {
		return _rf(table, ingress, service, from, to)
	}

	_ra0, _ := _ret.Get(0).(metrics.DataPoints)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *dataPointsFinderMock) OnFindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	return &dataPointsFinderFindByIngressAndServiceCall{Call: _m.Mock.On("FindByIngressAndService", table, ingress, service, from, to), Parent: _m}
}

func (_m *dataPointsFinderMock) OnFindByIngressAndServiceRaw(table interface{}, ingress interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressAndServiceCall {
	return &dataPointsFinderFindByIngressAndServiceCall{Call: _m.Mock.On("FindByIngressAndService", table, ingress, service, from, to), Parent: _m}
}

type dataPointsFinderFindByIngressAndServiceCall struct {
	*mock.Call
	Parent *dataPointsFinderMock
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Panic(msg string) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Once() *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Twice() *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Times(i int) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) WaitUntil(w <-chan time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) After(d time.Duration) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Run(fn func(args mock.Arguments)) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) Maybe() *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) TypedReturns(a metrics.DataPoints, b error) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) ReturnsFn(fn func(string, string, string, time.Time, time.Time) (metrics.DataPoints, error)) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) TypedRun(fn func(string, string, string, time.Time, time.Time)) *dataPointsFinderFindByIngressAndServiceCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_table := args.String(0)
		_ingress := args.String(1)
		_service := args.String(2)
		_from, _ := args.Get(3).(time.Time)
		_to, _ := args.Get(4).(time.Time)
		fn(_table, _ingress, _service, _from, _to)
	})
	return _c
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngress(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByIngress(table string, ingress string, from time.Time, to time.Time) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngress(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndService(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByService(table string, service string, from time.Time, to time.Time) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByService(table, service, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByEdgeIngressRaw(table interface{}, edgeIngress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngressRaw(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByIngressRaw(table interface{}, ingress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngressRaw(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByIngressAndServiceRaw(table interface{}, ingress interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndServiceRaw(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByIngressAndServiceCall) OnFindByServiceRaw(table interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByServiceRaw(table, service, from, to)
}

func (_m *dataPointsFinderMock) FindByService(table string, service string, from time.Time, to time.Time) metrics.DataPoints {
	_ret := _m.Called(table, service, from, to)

	if _rf, ok := _ret.Get(0).(func(string, string, time.Time, time.Time) metrics.DataPoints); ok {
		return _rf(table, service, from, to)
	}

	_ra0, _ := _ret.Get(0).(metrics.DataPoints)

	return _ra0
}

func (_m *dataPointsFinderMock) OnFindByService(table string, service string, from time.Time, to time.Time) *dataPointsFinderFindByServiceCall {
	return &dataPointsFinderFindByServiceCall{Call: _m.Mock.On("FindByService", table, service, from, to), Parent: _m}
}

func (_m *dataPointsFinderMock) OnFindByServiceRaw(table interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByServiceCall {
	return &dataPointsFinderFindByServiceCall{Call: _m.Mock.On("FindByService", table, service, from, to), Parent: _m}
}

type dataPointsFinderFindByServiceCall struct {
	*mock.Call
	Parent *dataPointsFinderMock
}

func (_c *dataPointsFinderFindByServiceCall) Panic(msg string) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) Once() *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) Twice() *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) Times(i int) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) WaitUntil(w <-chan time.Time) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) After(d time.Duration) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) Run(fn func(args mock.Arguments)) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) Maybe() *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) TypedReturns(a metrics.DataPoints) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) ReturnsFn(fn func(string, string, time.Time, time.Time) metrics.DataPoints) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) TypedRun(fn func(string, string, time.Time, time.Time)) *dataPointsFinderFindByServiceCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_table := args.String(0)
		_service := args.String(1)
		_from, _ := args.Get(2).(time.Time)
		_to, _ := args.Get(3).(time.Time)
		fn(_table, _service, _from, _to)
	})
	return _c
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByEdgeIngress(table string, edgeIngress string, from time.Time, to time.Time) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngress(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByIngress(table string, ingress string, from time.Time, to time.Time) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngress(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByIngressAndService(table string, ingress string, service string, from time.Time, to time.Time) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndService(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByService(table string, service string, from time.Time, to time.Time) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByService(table, service, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByEdgeIngressRaw(table interface{}, edgeIngress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByEdgeIngressCall {
	return _c.Parent.OnFindByEdgeIngressRaw(table, edgeIngress, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByIngressRaw(table interface{}, ingress interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressCall {
	return _c.Parent.OnFindByIngressRaw(table, ingress, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByIngressAndServiceRaw(table interface{}, ingress interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByIngressAndServiceCall {
	return _c.Parent.OnFindByIngressAndServiceRaw(table, ingress, service, from, to)
}

func (_c *dataPointsFinderFindByServiceCall) OnFindByServiceRaw(table interface{}, service interface{}, from interface{}, to interface{}) *dataPointsFinderFindByServiceCall {
	return _c.Parent.OnFindByServiceRaw(table, service, from, to)
}

// backendMock mock of Backend.
type backendMock struct{ mock.Mock }

// newBackendMock creates a new backendMock.
func newBackendMock(tb testing.TB) *backendMock {
	tb.Helper()

	m := &backendMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *backendMock) GetRules(_ context.Context) ([]Rule, error) {
	_ret := _m.Called()

	_ra0, _ := _ret.Get(0).([]Rule)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *backendMock) OnGetRules() *backendGetRulesCall {
	return &backendGetRulesCall{Call: _m.Mock.On("GetRules"), Parent: _m}
}

func (_m *backendMock) OnGetRulesRaw() *backendGetRulesCall {
	return &backendGetRulesCall{Call: _m.Mock.On("GetRules"), Parent: _m}
}

type backendGetRulesCall struct {
	*mock.Call
	Parent *backendMock
}

func (_c *backendGetRulesCall) Panic(msg string) *backendGetRulesCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *backendGetRulesCall) Once() *backendGetRulesCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *backendGetRulesCall) Twice() *backendGetRulesCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *backendGetRulesCall) Times(i int) *backendGetRulesCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *backendGetRulesCall) WaitUntil(w <-chan time.Time) *backendGetRulesCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *backendGetRulesCall) After(d time.Duration) *backendGetRulesCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *backendGetRulesCall) Run(fn func(args mock.Arguments)) *backendGetRulesCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *backendGetRulesCall) Maybe() *backendGetRulesCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *backendGetRulesCall) TypedReturns(a []Rule, b error) *backendGetRulesCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *backendGetRulesCall) ReturnsFn(fn func() ([]Rule, error)) *backendGetRulesCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *backendGetRulesCall) TypedRun(fn func()) *backendGetRulesCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		fn()
	})
	return _c
}

func (_c *backendGetRulesCall) OnGetRules() *backendGetRulesCall {
	return _c.Parent.OnGetRules()
}

func (_c *backendGetRulesCall) OnPreflightAlerts(alerts []Alert) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlerts(alerts)
}

func (_c *backendGetRulesCall) OnSendAlerts(alerts []Alert) *backendSendAlertsCall {
	return _c.Parent.OnSendAlerts(alerts)
}

func (_c *backendGetRulesCall) OnGetRulesRaw() *backendGetRulesCall {
	return _c.Parent.OnGetRulesRaw()
}

func (_c *backendGetRulesCall) OnPreflightAlertsRaw(alerts interface{}) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlertsRaw(alerts)
}

func (_c *backendGetRulesCall) OnSendAlertsRaw(alerts interface{}) *backendSendAlertsCall {
	return _c.Parent.OnSendAlertsRaw(alerts)
}

func (_m *backendMock) PreflightAlerts(_ context.Context, alerts []Alert) ([]Alert, error) {
	_ret := _m.Called(alerts)

	if _rf, ok := _ret.Get(0).(func([]Alert) ([]Alert, error)); ok {
		return _rf(alerts)
	}

	_ra0, _ := _ret.Get(0).([]Alert)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *backendMock) OnPreflightAlerts(alerts []Alert) *backendPreflightAlertsCall {
	return &backendPreflightAlertsCall{Call: _m.Mock.On("PreflightAlerts", alerts), Parent: _m}
}

func (_m *backendMock) OnPreflightAlertsRaw(alerts interface{}) *backendPreflightAlertsCall {
	return &backendPreflightAlertsCall{Call: _m.Mock.On("PreflightAlerts", alerts), Parent: _m}
}

type backendPreflightAlertsCall struct {
	*mock.Call
	Parent *backendMock
}

func (_c *backendPreflightAlertsCall) Panic(msg string) *backendPreflightAlertsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *backendPreflightAlertsCall) Once() *backendPreflightAlertsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *backendPreflightAlertsCall) Twice() *backendPreflightAlertsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *backendPreflightAlertsCall) Times(i int) *backendPreflightAlertsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *backendPreflightAlertsCall) WaitUntil(w <-chan time.Time) *backendPreflightAlertsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *backendPreflightAlertsCall) After(d time.Duration) *backendPreflightAlertsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *backendPreflightAlertsCall) Run(fn func(args mock.Arguments)) *backendPreflightAlertsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *backendPreflightAlertsCall) Maybe() *backendPreflightAlertsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *backendPreflightAlertsCall) TypedReturns(a []Alert, b error) *backendPreflightAlertsCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *backendPreflightAlertsCall) ReturnsFn(fn func([]Alert) ([]Alert, error)) *backendPreflightAlertsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *backendPreflightAlertsCall) TypedRun(fn func([]Alert)) *backendPreflightAlertsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_alerts, _ := args.Get(0).([]Alert)
		fn(_alerts)
	})
	return _c
}

func (_c *backendPreflightAlertsCall) OnGetRules() *backendGetRulesCall {
	return _c.Parent.OnGetRules()
}

func (_c *backendPreflightAlertsCall) OnPreflightAlerts(alerts []Alert) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlerts(alerts)
}

func (_c *backendPreflightAlertsCall) OnSendAlerts(alerts []Alert) *backendSendAlertsCall {
	return _c.Parent.OnSendAlerts(alerts)
}

func (_c *backendPreflightAlertsCall) OnGetRulesRaw() *backendGetRulesCall {
	return _c.Parent.OnGetRulesRaw()
}

func (_c *backendPreflightAlertsCall) OnPreflightAlertsRaw(alerts interface{}) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlertsRaw(alerts)
}

func (_c *backendPreflightAlertsCall) OnSendAlertsRaw(alerts interface{}) *backendSendAlertsCall {
	return _c.Parent.OnSendAlertsRaw(alerts)
}

func (_m *backendMock) SendAlerts(_ context.Context, alerts []Alert) error {
	_ret := _m.Called(alerts)

	if _rf, ok := _ret.Get(0).(func([]Alert) error); ok {
		return _rf(alerts)
	}

	_ra0 := _ret.Error(0)

	return _ra0
}

func (_m *backendMock) OnSendAlerts(alerts []Alert) *backendSendAlertsCall {
	return &backendSendAlertsCall{Call: _m.Mock.On("SendAlerts", alerts), Parent: _m}
}

func (_m *backendMock) OnSendAlertsRaw(alerts interface{}) *backendSendAlertsCall {
	return &backendSendAlertsCall{Call: _m.Mock.On("SendAlerts", alerts), Parent: _m}
}

type backendSendAlertsCall struct {
	*mock.Call
	Parent *backendMock
}

func (_c *backendSendAlertsCall) Panic(msg string) *backendSendAlertsCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *backendSendAlertsCall) Once() *backendSendAlertsCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *backendSendAlertsCall) Twice() *backendSendAlertsCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *backendSendAlertsCall) Times(i int) *backendSendAlertsCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *backendSendAlertsCall) WaitUntil(w <-chan time.Time) *backendSendAlertsCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *backendSendAlertsCall) After(d time.Duration) *backendSendAlertsCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *backendSendAlertsCall) Run(fn func(args mock.Arguments)) *backendSendAlertsCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *backendSendAlertsCall) Maybe() *backendSendAlertsCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *backendSendAlertsCall) TypedReturns(a error) *backendSendAlertsCall {
	_c.Call = _c.Return(a)
	return _c
}

func (_c *backendSendAlertsCall) ReturnsFn(fn func([]Alert) error) *backendSendAlertsCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *backendSendAlertsCall) TypedRun(fn func([]Alert)) *backendSendAlertsCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_alerts, _ := args.Get(0).([]Alert)
		fn(_alerts)
	})
	return _c
}

func (_c *backendSendAlertsCall) OnGetRules() *backendGetRulesCall {
	return _c.Parent.OnGetRules()
}

func (_c *backendSendAlertsCall) OnPreflightAlerts(alerts []Alert) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlerts(alerts)
}

func (_c *backendSendAlertsCall) OnSendAlerts(alerts []Alert) *backendSendAlertsCall {
	return _c.Parent.OnSendAlerts(alerts)
}

func (_c *backendSendAlertsCall) OnGetRulesRaw() *backendGetRulesCall {
	return _c.Parent.OnGetRulesRaw()
}

func (_c *backendSendAlertsCall) OnPreflightAlertsRaw(alerts interface{}) *backendPreflightAlertsCall {
	return _c.Parent.OnPreflightAlertsRaw(alerts)
}

func (_c *backendSendAlertsCall) OnSendAlertsRaw(alerts interface{}) *backendSendAlertsCall {
	return _c.Parent.OnSendAlertsRaw(alerts)
}

// processorMock mock of Processor.
type processorMock struct{ mock.Mock }

// newProcessorMock creates a new processorMock.
func newProcessorMock(tb testing.TB) *processorMock {
	tb.Helper()

	m := &processorMock{}
	m.Mock.Test(tb)

	tb.Cleanup(func() { m.AssertExpectations(tb) })

	return m
}

func (_m *processorMock) Process(rule *Rule) (*Alert, error) {
	_ret := _m.Called(rule)

	if _rf, ok := _ret.Get(0).(func(*Rule) (*Alert, error)); ok {
		return _rf(rule)
	}

	_ra0, _ := _ret.Get(0).(*Alert)
	_rb1 := _ret.Error(1)

	return _ra0, _rb1
}

func (_m *processorMock) OnProcess(rule *Rule) *processorProcessCall {
	return &processorProcessCall{Call: _m.Mock.On("Process", rule), Parent: _m}
}

func (_m *processorMock) OnProcessRaw(rule interface{}) *processorProcessCall {
	return &processorProcessCall{Call: _m.Mock.On("Process", rule), Parent: _m}
}

type processorProcessCall struct {
	*mock.Call
	Parent *processorMock
}

func (_c *processorProcessCall) Panic(msg string) *processorProcessCall {
	_c.Call = _c.Call.Panic(msg)
	return _c
}

func (_c *processorProcessCall) Once() *processorProcessCall {
	_c.Call = _c.Call.Once()
	return _c
}

func (_c *processorProcessCall) Twice() *processorProcessCall {
	_c.Call = _c.Call.Twice()
	return _c
}

func (_c *processorProcessCall) Times(i int) *processorProcessCall {
	_c.Call = _c.Call.Times(i)
	return _c
}

func (_c *processorProcessCall) WaitUntil(w <-chan time.Time) *processorProcessCall {
	_c.Call = _c.Call.WaitUntil(w)
	return _c
}

func (_c *processorProcessCall) After(d time.Duration) *processorProcessCall {
	_c.Call = _c.Call.After(d)
	return _c
}

func (_c *processorProcessCall) Run(fn func(args mock.Arguments)) *processorProcessCall {
	_c.Call = _c.Call.Run(fn)
	return _c
}

func (_c *processorProcessCall) Maybe() *processorProcessCall {
	_c.Call = _c.Call.Maybe()
	return _c
}

func (_c *processorProcessCall) TypedReturns(a *Alert, b error) *processorProcessCall {
	_c.Call = _c.Return(a, b)
	return _c
}

func (_c *processorProcessCall) ReturnsFn(fn func(*Rule) (*Alert, error)) *processorProcessCall {
	_c.Call = _c.Return(fn)
	return _c
}

func (_c *processorProcessCall) TypedRun(fn func(*Rule)) *processorProcessCall {
	_c.Call = _c.Call.Run(func(args mock.Arguments) {
		_rule, _ := args.Get(0).(*Rule)
		fn(_rule)
	})
	return _c
}

func (_c *processorProcessCall) OnProcess(rule *Rule) *processorProcessCall {
	return _c.Parent.OnProcess(rule)
}

func (_c *processorProcessCall) OnProcessRaw(rule interface{}) *processorProcessCall {
	return _c.Parent.OnProcessRaw(rule)
}