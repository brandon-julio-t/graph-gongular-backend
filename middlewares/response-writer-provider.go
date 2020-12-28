package middlewares

import (
	"context"
	"net/http"
)

type responseWriterProviderKeyStruct struct{ name string }

var responseWriterProviderKey = responseWriterProviderKeyStruct{name: "response-writer-provider"}

func CookieWriterProviderMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), responseWriterProviderKey, &w)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UseResponseWriter(ctx context.Context) *http.ResponseWriter {
	if w, ok := ctx.Value(responseWriterProviderKey).(*http.ResponseWriter); ok {
		return w
	}
	return nil
}
