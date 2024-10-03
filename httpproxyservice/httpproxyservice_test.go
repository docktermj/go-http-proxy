package httpproxyservice

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	httpProxyServiceSingleton HttpProxyService
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) HttpProxyService {
	if httpProxyServiceSingleton == nil {
		httpProxyServiceSingleton = &BasicHttpProxyService{}
	}
	return httpProxyServiceSingleton
}

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------
