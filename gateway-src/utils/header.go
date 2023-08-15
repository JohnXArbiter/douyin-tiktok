package utils

import "net/http"

func CopyHeader(dest *http.Header, src http.Header) {
	for k, v := range src {
		dest.Set(k, v[0])
	}
}
