package io

type HttpHeaderOption func(map[string]string)

func HttpHeaders(options ...HttpHeaderOption) map[string]string {
	headers := map[string]string{}
	for _, option := range options {
		option(headers)
	}
	return headers
}

func BearerToken(token string) HttpHeaderOption {
	return func(headers map[string]string) {
		headers["Authentication"] = "Bearer " + token
	}
}

func ClientIp(ip string) HttpHeaderOption {
	return func(headers map[string]string) {
		headers["HTTP_CLIENT_IP"] = ip
	}
}
