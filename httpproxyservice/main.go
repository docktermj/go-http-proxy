package httpproxyservice

import (
	"context"
	"net/http"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// The HTTPProxyService interface is...
type HTTPProxyService interface {
	Handler(context.Context) *http.ServeMux
}
