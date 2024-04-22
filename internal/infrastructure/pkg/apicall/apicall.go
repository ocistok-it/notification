package apicall

import (
	"github.com/gojek/heimdall/v7"
	"github.com/gojek/heimdall/v7/httpclient"
	"time"
)

type Apicall struct {
	client *httpclient.Client
}

func New() *Apicall {
	initialTimeout := 2 * time.Millisecond        // Inital timeout
	maxTimeout := 9 * time.Millisecond            // Max time out
	exponentFactor := 2.0                         // Multiplier
	maximumJitterInterval := 2 * time.Millisecond // Max jitter interval. It must be more than 1*time.Millisecond

	backoff := heimdall.NewExponentialBackoff(initialTimeout, maxTimeout, exponentFactor, maximumJitterInterval)

	retrier := heimdall.NewRetrier(backoff)

	timeout := 1000 * time.Millisecond

	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(timeout),
		httpclient.WithRetrier(retrier),
		httpclient.WithRetryCount(4),
	)

	return &Apicall{
		client: client,
	}
}
