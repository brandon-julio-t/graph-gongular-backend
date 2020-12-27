package middlewares

import (
	"context"
	"net/http"
)

type cookieWriterProviderKeyStruct struct{ name string }

var cookieWriterProviderKey = cookieWriterProviderKeyStruct{name: "cookie-provider"}

func CookieWriterProviderMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), cookieWriterProviderKey, &w)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UseCookieWriter(ctx context.Context) *http.ResponseWriter {
	if w, ok := ctx.Value(cookieWriterProviderKey).(*http.ResponseWriter); ok {
		return w
	}
	return nil
}
