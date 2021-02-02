package retry

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/logger"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/models"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"net/http"
	"strconv"
	"time"
)

type RetryConfiguration struct {
	MaxRetryTimeSec          int
	MaxRetriesBeforeQuitting int
}

type retry struct {
	retryAfterMs             int64
	retryCountBeforeQuitting int
	RetryConfiguration
}

func RetryWithData(uri string, data string, httpCall func(uri string, data string) (string, error)) func(retryConfiguration *RetryConfiguration) (string, error) {
	return func(retryConfiguration *RetryConfiguration) (string, error) {
		if retryConfiguration == nil {
			retryConfiguration = &RetryConfiguration{
				MaxRetriesBeforeQuitting: 3,
				MaxRetryTimeSec: 10,
			}
		}
		retry := retry{
			RetryConfiguration: *retryConfiguration,
		}

		response := ""
		var err error
		now := time.Now()
		for ok := true; ok; ok = retry.shouldRetry(now, err) {
			response, err = httpCall(uri, data)
		}
		return response, err
	}
}

func Retry(uri string, httpCall func(uri string) (string, error)) func(retryConfiguration *RetryConfiguration) (string, error) {
	return func(retryConfiguration *RetryConfiguration) (string, error) {
		if retryConfiguration == nil {
			retryConfiguration = &RetryConfiguration{
				MaxRetriesBeforeQuitting: 3,
				MaxRetryTimeSec: 10,
			}
		}
		retry := retry{
			RetryConfiguration: *retryConfiguration,
		}

		response := ""
		var err error
		now := time.Now()
		for ok := true; ok; ok = retry.shouldRetry(now, err) {
			response, err = httpCall(uri)
		}
		return response, err
	}
}

func (r *retry) shouldRetry(startTime time.Time, errorValue error) bool {
	if errorValue == nil {
		return false
	}

	e, ok := errorValue.(models.HttpStatusError)
	if !ok {
		return false
	}

	if time.Since(startTime) < utils.SecondsToNanoSeconds(r.MaxRetryTimeSec) && e.StatusCode == http.StatusTooManyRequests {
		r.retryAfterMs = getRetryAfterValue(e.Headers)
		r.retryCountBeforeQuitting++
		if r.retryCountBeforeQuitting < r.MaxRetriesBeforeQuitting {
			logger.Warnf("Got rate-limited. Sleeping for %v milliseconds before retrying", r.retryAfterMs)
			time.Sleep(utils.MilliSecondsToNanoSeconds(r.retryAfterMs))
			return true
		} else {
			logger.Warnf("Exceeded maximum rate-limit retries (%v). Will not retry again", r.MaxRetriesBeforeQuitting)
		}
	}
	return false
}

func getRetryAfterValue(headers map[string][]string) int64 {
	defaultValue := int64(3000)
	retryAfterValues := headers["Retry-After"]
	if retryAfterValues == nil {
		return defaultValue
	}

	returnValue := int64(0)
	for _, retryAfter := range retryAfterValues {
		if retryAfter != "" {
			returnValue, _ = strconv.ParseInt(retryAfter, 10, 64)
			break
		}
	}

	// Edge case where the retry-after header has no value
	if returnValue == 0 {
		returnValue = defaultValue
	}

	return returnValue
}
