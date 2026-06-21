package interceptor

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/serenite11/kucoin-universal-sdk/sdk/golang/pkg/common/logger"
)

// StartTimeKey Define a constant for the start time key in the context
const startTimeKey = "startTime"

// LoggingInterceptor logs request/response details, including elapsed time and success status
type LoggingInterceptor struct {
	debug bool
	log   logger.Logger
}

// NewLoggingInterceptor Create a new LoggingInterceptor with debug mode option
func NewLoggingInterceptor(debug bool, log logger.Logger) *LoggingInterceptor {
	return &LoggingInterceptor{debug: debug, log: log}
}

// Before sends the request and logs the URI and headers if debug is enabled
func (l *LoggingInterceptor) Before(req *http.Request) (*http.Request, error) {

	ctx := context.WithValue(req.Context(), startTimeKey, time.Now())

	if l.debug {
		l.log.Infof("Request Headers:")
		for key, values := range req.Header {
			for _, value := range values {
				l.log.Infof("%s: %s\n", key, value)
			}
		}
		if req.Body != nil {
			bodyBytes, err := io.ReadAll(req.Body)
			if err == nil {
				l.log.Infof("Request Body: %s", string(bodyBytes))
				// Re-fill the Body after reading
				req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}
	}

	req = req.Clone(ctx)

	return req, nil
}

// After logs the response details, including time taken and success status
func (l *LoggingInterceptor) After(req *http.Request, resp *http.Response, err error) (*http.Response, error) {
	startTime, ok := req.Context().Value(startTimeKey).(time.Time)
	if !ok {
		l.log.Errorf("cannot get startTime from context")
		startTime = time.Now()
	}
	duration := time.Since(startTime)

	errMsg := "nil"
	if err != nil {
		errMsg = err.Error()
	}

	responseStatus := "None"
	if resp != nil {
		responseStatus = resp.Status
	}

	l.log.Infof("[Access] method=%s, url=%s, status=%s, cost=%v, err=%v", req.Method, req.URL.String(), responseStatus, duration, errMsg)

	// If debug is enabled, log response headers and body
	if l.debug {
		fmt.Println("Response Headers:")
		for key, values := range resp.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", key, value)
			}
		}
		bodyBytes, err := io.ReadAll(resp.Body)
		if err == nil {
			fmt.Printf("Response Body: %s\n", string(bodyBytes))
			// Re-fill the response body after reading
			resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}
	}

	return resp, err
}
