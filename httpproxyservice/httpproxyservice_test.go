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
// Test interface functions
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) HttpProxyService {
	_ = ctx
	_ = test
	if httpProxyServiceSingleton == nil {
		httpProxyServiceSingleton = &BasicHttpProxyService{}
	}
	return httpProxyServiceSingleton
}

func testError(test *testing.T, ctx context.Context, err error) {
	_ = ctx
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}
