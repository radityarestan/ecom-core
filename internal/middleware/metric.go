package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/radityarestan/ecom-core/internal/shared/dto"
	"strconv"
)

var (
	totalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Number of get requests.",
		},
		[]string{"path", "method"})

	responseStatus = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_response_status",
			Help: "Status of HTTP response",
		},
		[]string{"path", "method", "status"})

	httpLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_response_time",
			Help: "Duration of HTTP requests.",
		}, []string{"path", "method"})

	httpError = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Number of errors.",
		}, []string{"path", "method", "exception_name"})
)

func init() {
	prometheus.Register(httpLatency)

	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
	prometheus.MustRegister(httpError)
}

func MetricsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		method := c.Request().Method

		timer := prometheus.NewTimer(httpLatency.WithLabelValues(path, method))
		defer timer.ObserveDuration()

		totalRequests.WithLabelValues(path, method).Inc()

		err := next(c)
		if err != nil {
			return err
		}

		statusCode := c.Response().Status
		responseStatus.WithLabelValues(path, method, strconv.Itoa(statusCode)).Inc()

		if statusCode >= 400 {
			exceptionName := c.Get(dto.StatusError).(string)
			httpError.WithLabelValues(path, method, exceptionName).Inc()
		}

		return nil
	}
}
