//+build !swipe

// Code generated by Swipe v1.22.3. DO NOT EDIT.

//go:generate swipe
package rest

import (
	"context"
	"github.com/go-kit/kit/metrics"
	prometheus2 "github.com/go-kit/kit/metrics/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/swipe-io/swipe/fixtures/service"
	"github.com/swipe-io/swipe/fixtures/user"
	"time"
)

type instrumentingMiddlewareServiceInterface struct {
	next           service.Interface
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
}

func (s *instrumentingMiddlewareServiceInterface) Create(ctx context.Context, name string, data []byte) (_ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Create").Add(1)
		s.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Create(ctx, name, data)
}

func (s *instrumentingMiddlewareServiceInterface) Delete(ctx context.Context, id uint) (a string, b string, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Delete").Add(1)
		s.requestLatency.With("method", "Delete").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Delete(ctx, id)
}

func (s *instrumentingMiddlewareServiceInterface) Get(ctx context.Context, id int, name string, fname string, price float32, n int, b int, c int) (data user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "Get").Add(1)
		s.requestLatency.With("method", "Get").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.Get(ctx, id, name, fname, price, n, b, c)
}

func (s *instrumentingMiddlewareServiceInterface) GetAll(ctx context.Context) (_ []*user.User, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetAll").Add(1)
		s.requestLatency.With("method", "GetAll").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.GetAll(ctx)
}

func (s *instrumentingMiddlewareServiceInterface) TestMethod(data map[string]interface{}, ss interface{}) (states map[string]map[int][]string, _ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TestMethod").Add(1)
		s.requestLatency.With("method", "TestMethod").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.TestMethod(data, ss)
}

func (s *instrumentingMiddlewareServiceInterface) TestMethod2(ctx context.Context, ns string, utype string, user string, restype string, resource string, permission string) (_ error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "TestMethod2").Add(1)
		s.requestLatency.With("method", "TestMethod2").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.next.TestMethod2(ctx, ns, utype, user, restype, resource, permission)
}

func NewInstrumentingMiddlewareServiceInterface(s service.Interface, requestCount metrics.Counter, requestLatency metrics.Histogram) service.Interface {
	if requestCount == nil {
		requestCount = prometheus2.NewCounterFrom(prometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "api",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, []string{"method"})

	}
	if requestLatency == nil {
		requestLatency = prometheus2.NewSummaryFrom(prometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "api",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, []string{"method"})

	}
	return &instrumentingMiddlewareServiceInterface{next: s, requestCount: requestCount, requestLatency: requestLatency}
}
