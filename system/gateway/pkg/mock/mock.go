package mock

import "gateway/internal/router/middleware"

func SetRequestedAt(path string, requestedAt int64) {
	middleware.RateLimitMap[path] = requestedAt
}
