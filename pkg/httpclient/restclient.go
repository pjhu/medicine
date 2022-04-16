package httpclient

import (
	"time"

	"pjhu/medicine/pkg/errors"

	"github.com/go-resty/resty/v2"
)

func BuildRestClient() *resty.Client {
	client := resty.New()
	// Retries are configured per client
	client.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(1 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(1 * time.Second).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			errWithCode := errors.NewErrorWithCode(errors.SystemInternalError, "quota exceeded")
			return 0, errWithCode
		})
	return client
}
