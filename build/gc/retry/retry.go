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

func RetryWithData(uri string, data []string, httpCall func(uri string, data string) (string, error)) func(retryConfig *RetryConfiguration) (string, error) {
	return func(retryConfig *RetryConfiguration) (string, error) {
		if retryConfig == nil {
			retryConfig = &RetryConfiguration{
				RetryWaitMin: 5 * time.Second,
				RetryWaitMax: 60 * time.Second,
				RetryMax:     20,
			}
		}
		retryConfiguration = retryConfig

		// if there is only one item in the array (req body), then just return the response object
		if len(data) == 1 {
			res, err := httpCall(uri, data[0])
			return res, err
		}

		// if there is more than one item in the array (req bodies), the response will be an array of response objects
		var response string = "["
		for i, reqBody := range data {
			res, err := httpCall(uri, reqBody)
			retryConfiguration = nil
			if err != nil {
				if i == len(data)-1 {
					response += err.Error() + "]"
				} else {
					response += err.Error() + ","
				}
			} else {
				if i == len(data)-1 {
					response += res + "]"
				} else {
					response += res + ","
				}
			}
		}

		// returning response and nil error as the error objects will be concatenated onto the response (if there are any)
		return response, nil
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
