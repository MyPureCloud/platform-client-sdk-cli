package retry

import (
	"net/http"
	"time"
)

type RetryConfiguration struct {
	RetryWaitMin   time.Duration  `json:"retry_wait_min,omitempty"`
	RetryWaitMax   time.Duration  `json:"retry_wait_max,omitempty"`
	RetryMax       int            `json:"retry_max,omitempty"`
	RequestLogHook RequestLogHook `json:"request_log_hook,omitempty"`
}

var retryConfiguration *RetryConfiguration

func GetRetryConfiguration() *RetryConfiguration {
	return retryConfiguration
}

type RequestLogHook func(*http.Request, int)

func RetryWithData(uri string, data string, httpCall func(uri string, data string) (string, error)) func(retryConfig *RetryConfiguration) (string, error) {
	return func(retryConfig *RetryConfiguration) (string, error) {
		if retryConfig == nil {
			retryConfig = &RetryConfiguration{
				RetryWaitMin: 5 * time.Second,
				RetryWaitMax: 60 * time.Second,
				RetryMax:     20,
			}
		}
		retryConfiguration = retryConfig

		response, err := httpCall(uri, data)
		retryConfiguration = nil

		return response, err
	}
}

func Retry(uri string, httpCall func(uri string) (string, error)) func(retryConfig *RetryConfiguration) (string, error) {
	return func(retryConfig *RetryConfiguration) (string, error) {
		if retryConfig == nil {
			retryConfig = &RetryConfiguration{
				RetryWaitMin: 5 * time.Second,
				RetryWaitMax: 60 * time.Second,
				RetryMax:     20,
			}
		}
		retryConfiguration = retryConfig

		response, err := httpCall(uri)
		retryConfiguration = nil

		return response, err
	}
}
